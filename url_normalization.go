package apify

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/idna"
)

// The fields the OpenAPI specification declares as `format: uri` are returned in their
// normalized form, the same string for the same field in every Apify client (the JavaScript
// client returns the WHATWG URL serialization, and the Python one an equivalent): the scheme and
// host are lowercased, a non-ASCII host is converted to punycode, the default port is dropped, an
// empty path becomes "/", dot segments are resolved and characters that are unsafe in a URL are
// percent-encoded.

// defaultPorts maps a special URL scheme to its default port, which normalization drops.
var defaultPorts = map[string]string{
	"http":  "80",
	"https": "443",
	"ws":    "80",
	"wss":   "443",
	"ftp":   "21",
}

// urlSchemePattern matches the scheme at the start of an absolute URL.
var urlSchemePattern = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9+.-]*):`)

// maxPort is the largest valid TCP/UDP port number.
const maxPort = 65535

// Bytes that are percent-encoded in each URL component, in addition to every control character,
// space and non-ASCII byte (the WHATWG percent-encode sets).
const (
	unsafePathBytes     = `"#<>?` + "`{}"
	unsafeQueryBytes    = `"#<>'`
	unsafeFragmentBytes = `"<>` + "`"
	unsafeUserinfoBytes = `"#<>?` + "`{}/:;=@[\\]^|"
)

// idnaProfile converts an internationalized host name to its ASCII (punycode) form. The checks
// that only apply to registrable names (hyphen rules, bidi, joiners, label length) are off,
// because the host of a URL is not necessarily a registrable one.
var idnaProfile = idna.New(
	idna.MapForLookup(),
	idna.StrictDomainName(false),
	idna.VerifyDNSLength(false),
	idna.CheckHyphens(false),
	idna.CheckJoiners(false),
	idna.BidiRule(),
)

// normalizeURL returns the normalized form of an absolute URL, or an error if raw is not one.
func normalizeURL(raw string) (string, error) {
	rest := strings.TrimSpace(raw)
	match := urlSchemePattern.FindStringSubmatch(rest)
	if match == nil {
		return "", fmt.Errorf("invalid URL %q: no scheme", raw)
	}
	scheme := strings.ToLower(match[1])
	rest = rest[len(match[0]):]
	defaultPort, special := defaultPorts[scheme]
	special = special || scheme == "file"

	var fragment string
	hasFragment := false
	if i := strings.IndexByte(rest, '#'); i >= 0 {
		rest, fragment, hasFragment = rest[:i], rest[i+1:], true
	}
	var query string
	hasQuery := false
	if i := strings.IndexByte(rest, '?'); i >= 0 {
		rest, query, hasQuery = rest[:i], rest[i+1:], true
	}
	if special {
		rest = strings.ReplaceAll(rest, `\`, "/")
	}

	var out strings.Builder
	out.WriteString(scheme)
	out.WriteByte(':')
	path := rest
	if authority, ok := strings.CutPrefix(rest, "//"); ok {
		end := strings.IndexByte(authority, '/')
		if end < 0 {
			end = len(authority)
		}
		path = authority[end:]
		host, err := normalizeAuthority(authority[:end], defaultPort)
		if err != nil {
			return "", fmt.Errorf("invalid URL %q: %w", raw, err)
		}
		out.WriteString("//")
		out.WriteString(host)
	} else if special {
		return "", fmt.Errorf("invalid URL %q: %q URLs need a host", raw, scheme)
	}

	if special && path == "" {
		path = "/"
	}
	if special || strings.HasPrefix(path, "/") {
		path = removeDotSegments(path)
	}
	out.WriteString(percentEncode(path, unsafePathBytes))
	if hasQuery {
		unsafe := unsafeQueryBytes
		if !special {
			unsafe = `"#<>`
		}
		out.WriteByte('?')
		out.WriteString(percentEncode(query, unsafe))
	}
	if hasFragment {
		out.WriteByte('#')
		out.WriteString(percentEncode(fragment, unsafeFragmentBytes))
	}
	return out.String(), nil
}

// normalizeAuthority normalizes the "userinfo@host:port" part of a URL.
func normalizeAuthority(authority, defaultPort string) (string, error) {
	var userinfo string
	if i := strings.LastIndexByte(authority, '@'); i >= 0 {
		userinfo, authority = authority[:i], authority[i+1:]
	}
	host, port := authority, ""
	if strings.HasPrefix(authority, "[") {
		end := strings.IndexByte(authority, ']')
		if end < 0 {
			return "", errors.New("unterminated IPv6 address")
		}
		host = strings.ToLower(authority[:end+1])
		port = strings.TrimPrefix(authority[end+1:], ":")
		if tail := authority[end+1:]; tail != "" && !strings.HasPrefix(tail, ":") {
			return "", errors.New("unexpected text after the IPv6 address")
		}
	} else {
		if i := strings.LastIndexByte(authority, ':'); i >= 0 {
			host, port = authority[:i], authority[i+1:]
		}
		if strings.ContainsFunc(host, isForbiddenHostRune) {
			return "", fmt.Errorf("invalid host %q", host)
		}
		ascii, err := idnaProfile.ToASCII(strings.ToLower(host))
		if err != nil {
			return "", fmt.Errorf("invalid host %q: %w", host, err)
		}
		host = ascii
	}
	if host == "" {
		return "", errors.New("empty host")
	}
	if port != "" {
		number, err := strconv.Atoi(port)
		if err != nil || number < 0 || number > maxPort {
			return "", fmt.Errorf("invalid port %q", port)
		}
		port = strconv.Itoa(number)
		if port == defaultPort {
			port = ""
		}
	}

	var out strings.Builder
	if userinfo != "" {
		username, password, hasPassword := strings.Cut(userinfo, ":")
		out.WriteString(percentEncode(username, unsafeUserinfoBytes))
		if hasPassword {
			out.WriteByte(':')
			out.WriteString(percentEncode(password, unsafeUserinfoBytes))
		}
		out.WriteByte('@')
	}
	out.WriteString(host)
	if port != "" {
		out.WriteByte(':')
		out.WriteString(port)
	}
	return out.String(), nil
}

// forbiddenHostBytes are the ASCII characters a host may not contain (besides control characters
// and the space).
const forbiddenHostBytes = "#%/<>?@[\\]^|"

// isForbiddenHostRune reports whether r may not appear in a host.
func isForbiddenHostRune(r rune) bool {
	return r <= ' ' || r == 0x7F || strings.ContainsRune(forbiddenHostBytes, r)
}

// removeDotSegments resolves the "." and ".." segments of an absolute path.
func removeDotSegments(path string) string {
	if path == "" {
		return path
	}
	segments := strings.Split(path, "/")
	kept := make([]string, 0, len(segments))
	for i, segment := range segments {
		last := i == len(segments)-1
		switch strings.ToLower(segment) {
		case ".", "%2e":
			if last {
				kept = append(kept, "")
			}
		case "..", ".%2e", "%2e.", "%2e%2e":
			if len(kept) > 1 {
				kept = kept[:len(kept)-1]
			}
			if last {
				kept = append(kept, "")
			}
		default:
			kept = append(kept, segment)
		}
	}
	return strings.Join(kept, "/")
}

// percentEncode percent-encodes every byte of s that is a control character, a space, a
// non-ASCII byte or listed in unsafe. An existing percent escape is kept as it is.
func percentEncode(s, unsafe string) string {
	const upperHex = "0123456789ABCDEF"
	var out strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b <= ' ' || b >= 0x7F || strings.IndexByte(unsafe, b) >= 0 {
			out.WriteByte('%')
			out.WriteByte(upperHex[b>>4])
			out.WriteByte(upperHex[b&0x0F])
			continue
		}
		out.WriteByte(b)
	}
	return out.String()
}

// normalizeURLField normalizes the URL held in *field in place. An empty value (a missing or
// null field) is left as it is; any other value that is not an absolute URL is an error naming
// the field.
func normalizeURLField(name string, field *string) error {
	if *field == "" {
		return nil
	}
	normalized, err := normalizeURL(*field)
	if err != nil {
		return fmt.Errorf("field %q: %w", name, err)
	}
	*field = normalized
	return nil
}

// normalizeOptionalURLField is [normalizeURLField] for a nullable field held as a *string.
func normalizeOptionalURLField(name string, field **string) error {
	if *field == nil {
		return nil
	}
	return normalizeURLField(name, *field)
}

// normalizeURLMap normalizes every value of a map of URLs in place.
func normalizeURLMap(name string, urls map[string]string) error {
	for key, value := range urls {
		if err := normalizeURLField(name+"."+key, &value); err != nil {
			return err
		}
		urls[key] = value
	}
	return nil
}
