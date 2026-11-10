package apify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

// Constants of the storage-content signing scheme of @apify/utilities.
const (
	// storageContentSignatureVersion is the version tag embedded in signatures.
	storageContentSignatureVersion = "0"
	// hmacSignatureHexLen is the number of leading hex characters of the HMAC digest used.
	hmacSignatureHexLen = 30
	// base62Alphabet is the alphabet (lowercase first) used to encode the truncated HMAC.
	base62Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// urlSigningSecretExtraKey is the Extra field name the API uses to expose a private
	// storage's (dataset or key-value store) HMAC signing secret. Consulted by the
	// dataset/key-value-store public-URL builders to decide whether a URL needs a signature.
	urlSigningSecretExtraKey = "urlSigningSecretKey"
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

// signStorageContent signs the public URL of a storage resource, compatibly with
// createStorageContentSignature of @apify/utilities. expiresInSecs is nil for no expiry.
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

// validateExpiresInSecs rejects an expiresInSecs that is not positive. A zero or negative expiry
// gives a URL that is already expired when it is made, which is never what the caller meant;
// leave expiresInSecs nil for a URL that never expires.
func validateExpiresInSecs(expiresInSecs *int64) error {
	if expiresInSecs != nil && *expiresInSecs <= 0 {
		return fmt.Errorf("expiresInSecs must be positive, got %d (leave it nil for a URL that never expires)", *expiresInSecs)
	}
	return nil
}

// itoa is a tiny int64-to-string helper to keep signStorageContent readable.
func itoa(v int64) string {
	return big.NewInt(v).String()
}
