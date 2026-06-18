package apify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"time"
)

// Constants for the Apify storage-content signing scheme, matching the platform's
// @apify/utilities implementation that the reference clients rely on.
const (
	// storageContentSignatureVersion is the version tag embedded in signatures. The
	// reference clients rely on the upstream default of "0".
	storageContentSignatureVersion = "0"
	// hmacSignatureHexLen is the number of leading hex characters of the HMAC digest used.
	hmacSignatureHexLen = 30
	// base62Alphabet is the alphabet (lowercase first) used to encode the truncated HMAC.
	base62Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// createHmacSignature computes an Apify URL-signing signature, byte-for-byte compatible
// with the platform's @apify/utilities createHmacSignature.
//
// The algorithm is HMAC-SHA256(secretKey, message) as lowercase hex, take the first 30
// hex characters, interpret them as a big integer, then base62-encode (alphabet 0-9a-zA-Z).
func createHmacSignature(secretKey, message string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	digest := mac.Sum(nil)

	full := hex.EncodeToString(digest)
	truncated := full[:hmacSignatureHexLen]
	value := new(big.Int)
	value.SetString(truncated, 16)
	return toBase62(value)
}

// toBase62 encodes a non-negative big integer in base62 using the 0-9a-zA-Z alphabet.
func toBase62(value *big.Int) string {
	if value.Sign() == 0 {
		return "0"
	}
	base := big.NewInt(int64(len(base62Alphabet)))
	mod := new(big.Int)
	v := new(big.Int).Set(value)
	var digits []byte
	for v.Sign() > 0 {
		v.DivMod(v, base, mod)
		digits = append(digits, base62Alphabet[mod.Int64()])
	}
	// reverse
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

// signStorageContent builds a storage-content signature for a resource's public URL,
// byte-for-byte compatible with @apify/utilities createStorageContentSignature.
//
// It signs the message "{version}.{expiresAtMillis}.{resourceId}" (expiresAtMillis is the
// absolute expiry in ms, or 0 for a non-expiring URL) with createHmacSignature, then
// returns the base64url (no padding) encoding of "{version}.{expiresAtMillis}.{hmac}".
//
// expiresInSecs is optional (nil for a non-expiring URL).
func signStorageContent(secretKey, resourceID string, expiresInSecs *int64) string {
	var expiresAtMillis int64
	if expiresInSecs != nil {
		expiresAtMillis = time.Now().UnixMilli() + *expiresInSecs*1000
	}
	version := storageContentSignatureVersion
	message := version + "." + itoa(expiresAtMillis) + "." + resourceID
	mac := createHmacSignature(secretKey, message)
	envelope := version + "." + itoa(expiresAtMillis) + "." + mac
	return base64.RawURLEncoding.EncodeToString([]byte(envelope))
}

// itoa is a tiny int64-to-string helper to keep signStorageContent readable.
func itoa(v int64) string {
	return big.NewInt(v).String()
}
