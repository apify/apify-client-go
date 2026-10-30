package apify

import (
	"errors"
	"fmt"
	"net/netip"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// The fields the OpenAPI specification declares as `format: uri` are returned in a normalized
// form, close to the WHATWG URL serialization that the JavaScript client returns: the scheme and
// host are lowercased, a non-ASCII host is converted to punycode, a numeric (IPv4) or IPv6 host is
// written in its canonical form, the default port is dropped, an empty path becomes "/", dot
// segments are resolved and characters that are unsafe in a URL are percent-encoded. It is built
// for the http(s), ws(s), ftp and file URLs the API returns. The host of any other scheme is kept
// as given. It does not apply the full Unicode (UTS #46) mapping to a non-ASCII host.

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
	unsafePathBytes       = `"#<>?` + "`{}"
	unsafeQueryBytes      = `"#<>'`
	unsafeOpaqueQueryByte = `"#<>`
	unsafeFragmentBytes   = `"<>` + "`"
	unsafeUserinfoBytes   = `"#<>?` + "`{}/:;=@[\\]^|"
)

// forbiddenHostBytes are the ASCII characters a host may not contain (besides control characters
// and the space).
const forbiddenHostBytes = "#%/:<>?@[\\]^|"

// ipv4PartCount is the number of parts of a dotted IPv4 address, and ipv4PartBits the bits of
// each; ipv6PieceCount is the number of 16-bit pieces of an IPv6 address.
const (
	ipv4PartCount  = 4
	ipv4PartBits   = 8
	ipv6PieceCount = 8
	maxIPv4Part    = 1<<ipv4PartBits - 1
)

// normalizeURL returns the normalized form of an absolute URL, or an error if raw is not one.
func normalizeURL(raw string) (string, error) {
	// Leading and trailing control characters and spaces are dropped, tabs and line breaks
	// anywhere are removed.
	cleaned := strings.Map(func(r rune) rune {
		if r == '\t' || r == '\n' || r == '\r' {
			return -1
		}
		return r
	}, strings.TrimFunc(raw, func(r rune) bool { return r <= ' ' }))

	match := urlSchemePattern.FindStringSubmatch(cleaned)
	if match == nil {
		return "", fmt.Errorf("invalid URL %q: no scheme", raw)
	}
	scheme := strings.ToLower(match[1])
	rest := cleaned[len(match[0]):]
	_, special := defaultPorts[scheme]
	isFile := scheme == "file"
	special = special || isFile

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
	hasAuthority := special
	switch {
	case special && !isFile:
		// A special URL always has an authority, whatever number of slashes (even none) follows
		// its scheme.
		authority := strings.TrimLeft(rest, "/")
		end := strings.IndexByte(authority, '/')
		if end < 0 {
			end = len(authority)
		}
		path = authority[end:]
		host, err := normalizeAuthority(authority[:end], authorityRules{defaultPort: defaultPorts[scheme], special: true, allowPort: true, allowUserinfo: true})
		if err != nil {
			return "", fmt.Errorf("invalid URL %q: %w", raw, err)
		}
		out.WriteString("//")
		out.WriteString(host)
	case isFile:
		host := ""
		path = rest
		if authority, ok := strings.CutPrefix(rest, "//"); ok {
			end := strings.IndexByte(authority, '/')
			if end < 0 {
				end = len(authority)
			}
			path = authority[end:]
			if authority[:end] != "" {
				var err error
				if host, err = normalizeAuthority(authority[:end], authorityRules{special: true}); err != nil {
					return "", fmt.Errorf("invalid URL %q: %w", raw, err)
				}
			}
			if host == "localhost" {
				host = ""
			}
		}
		out.WriteString("//")
		out.WriteString(host)
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
	default:
		if authority, ok := strings.CutPrefix(rest, "//"); ok {
			hasAuthority = true
			end := strings.IndexByte(authority, '/')
			if end < 0 {
				end = len(authority)
			}
			path = authority[end:]
			host, err := normalizeAuthority(authority[:end], authorityRules{allowPort: true, allowUserinfo: true})
			if err != nil {
				return "", fmt.Errorf("invalid URL %q: %w", raw, err)
			}
			out.WriteString("//")
			out.WriteString(host)
		}
	}

	if special && path == "" {
		path = "/"
	}
	opaquePath := !special && !strings.HasPrefix(path, "/")
	if !opaquePath {
		path = removeDotSegments(path)
		if !hasAuthority && strings.HasPrefix(path, "//") {
			// Without this marker the path would read as the start of an authority.
			out.WriteString("/.")
		}
		out.WriteString(percentEncode(path, unsafePathBytes))
	} else {
		out.WriteString(percentEncodeControls(path))
	}
	if hasQuery {
		unsafe := unsafeQueryBytes
		if !special {
			unsafe = unsafeOpaqueQueryByte
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

// authorityRules says how the authority of a URL is normalized.
type authorityRules struct {
	// defaultPort is the port that is dropped as the default one of the scheme.
	defaultPort string
	// special is true for the special schemes, whose host is a domain, IPv4 or IPv6 address. The
	// host of any other scheme is opaque and kept as given.
	special bool
	// allowPort is false for a scheme whose URLs have no port (file), and allowUserinfo for one
	// whose URLs have no credentials (file).
	allowPort     bool
	allowUserinfo bool
}

// normalizeAuthority normalizes the "userinfo@host:port" part of a URL.
func normalizeAuthority(authority string, rules authorityRules) (string, error) {
	var userinfo string
	if i := strings.LastIndexByte(authority, '@'); i >= 0 {
		if !rules.allowUserinfo {
			return "", errors.New("credentials are not allowed")
		}
		userinfo, authority = authority[:i], authority[i+1:]
	}
	host, port := authority, ""
	hasPortSeparator := false
	if strings.HasPrefix(authority, "[") {
		end := strings.IndexByte(authority, ']')
		if end < 0 {
			return "", errors.New("unterminated IPv6 address")
		}
		tail := authority[end+1:]
		if tail != "" && !strings.HasPrefix(tail, ":") {
			return "", errors.New("unexpected text after the IPv6 address")
		}
		address, err := netip.ParseAddr(authority[1:end])
		if err != nil || !address.Is6() || address.Zone() != "" {
			return "", fmt.Errorf("invalid IPv6 address %q", authority[1:end])
		}
		host, port = serializeIPv6(address), strings.TrimPrefix(tail, ":")
		hasPortSeparator = tail != ""
	} else {
		if i := strings.LastIndexByte(authority, ':'); i >= 0 {
			host, port, hasPortSeparator = authority[:i], authority[i+1:], true
		}
		var err error
		if host, err = normalizeHost(host, rules.special); err != nil {
			return "", err
		}
	}
	if hasPortSeparator && !rules.allowPort {
		return "", errors.New("a port is not allowed")
	}
	if port != "" {
		number, err := parsePort(port)
		if err != nil || !rules.allowPort {
			return "", fmt.Errorf("invalid port %q", port)
		}
		port = strconv.Itoa(number)
		if port == rules.defaultPort {
			port = ""
		}
	}

	var out strings.Builder
	username, password, _ := strings.Cut(userinfo, ":")
	if username != "" || password != "" {
		out.WriteString(percentEncode(username, unsafeUserinfoBytes))
		if password != "" {
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

// parsePort parses a port: ASCII digits only, at most maxPort.
func parsePort(port string) (int, error) {
	for i := 0; i < len(port); i++ {
		if port[i] < '0' || port[i] > '9' {
			return 0, errors.New("not a number")
		}
	}
	number, err := strconv.Atoi(port)
	if err != nil || number > maxPort {
		return 0, errors.New("out of range")
	}
	return number, nil
}

// normalizeHost normalizes a host name: percent-decoded, lowercased, converted to punycode, and
// an IPv4 address (in any of its numeric notations) written as four decimal parts. An empty host
// is an error for a special scheme. With special false the host is opaque and kept as given.
func normalizeHost(host string, special bool) (string, error) {
	if !special {
		return percentEncode(host, ""), nil
	}
	decoded, err := percentDecode(host)
	if err != nil {
		return "", fmt.Errorf("invalid host %q: %w", host, err)
	}
	if decoded == "" {
		return "", errors.New("empty host")
	}
	if strings.ContainsFunc(decoded, isForbiddenHostRune) {
		return "", fmt.Errorf("invalid host %q", host)
	}
	ascii := hostToASCII(strings.ToLower(decoded))
	if endsInANumber(ascii) {
		return parseIPv4(ascii)
	}
	return ascii, nil
}

// isForbiddenHostRune reports whether r may not appear in a host.
func isForbiddenHostRune(r rune) bool {
	return r <= ' ' || r == 0x7F || strings.ContainsRune(forbiddenHostBytes, r)
}

// percentDecode decodes the percent escapes of s, which must make valid UTF-8.
func percentDecode(s string) (string, error) {
	if !strings.Contains(s, "%") {
		return s, nil
	}
	decoded := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '%' {
			decoded = append(decoded, s[i])
			continue
		}
		if i+3 > len(s) {
			decoded = append(decoded, s[i])
			continue
		}
		value, err := strconv.ParseUint(s[i+1:i+3], 16, 8)
		if err != nil {
			decoded = append(decoded, s[i])
			continue
		}
		decoded = append(decoded, byte(value))
		i += 2
	}
	if !utf8.Valid(decoded) {
		return "", errors.New("not valid UTF-8")
	}
	return string(decoded), nil
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

// percentEncodeControls percent-encodes only the control characters and the non-ASCII bytes of s,
// as is done for the opaque path of a URL such as mailto:.
func percentEncodeControls(s string) string {
	const upperHex = "0123456789ABCDEF"
	var out strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b < ' ' || b >= 0x7F {
			out.WriteByte('%')
			out.WriteByte(upperHex[b>>4])
			out.WriteByte(upperHex[b&0x0F])
			continue
		}
		out.WriteByte(b)
	}
	return out.String()
}

// serializeIPv6 writes an IPv6 address as WHATWG does: lowercase hexadecimal pieces without
// leading zeros, the longest run of zero pieces (of at least two) compressed to "::", in brackets.
func serializeIPv6(address netip.Addr) string {
	bytes := address.As16()
	var pieces [ipv6PieceCount]uint16
	for i := range pieces {
		pieces[i] = uint16(bytes[2*i])<<8 | uint16(bytes[2*i+1])
	}
	compressStart, compressLength := -1, 0
	for i := 0; i < ipv6PieceCount; {
		if pieces[i] != 0 {
			i++
			continue
		}
		j := i
		for j < ipv6PieceCount && pieces[j] == 0 {
			j++
		}
		if j-i > compressLength && j-i >= 2 {
			compressStart, compressLength = i, j-i
		}
		i = j
	}
	var out strings.Builder
	out.WriteByte('[')
	for i := 0; i < ipv6PieceCount; i++ {
		if i == compressStart {
			if i == 0 {
				out.WriteString("::")
			} else {
				out.WriteByte(':')
			}
			i += compressLength - 1
			continue
		}
		out.WriteString(strconv.FormatUint(uint64(pieces[i]), 16))
		if i != ipv6PieceCount-1 {
			out.WriteByte(':')
		}
	}
	out.WriteByte(']')
	return out.String()
}

// endsInANumber reports whether the last label of host is a number (decimal, 0x hexadecimal or
// 0 octal), which makes the whole host an IPv4 address.
func endsInANumber(host string) bool {
	labels := strings.Split(host, ".")
	if labels[len(labels)-1] == "" {
		if len(labels) == 1 {
			return false
		}
		labels = labels[:len(labels)-1]
	}
	last := labels[len(labels)-1]
	if last != "" && strings.Trim(last, "0123456789") == "" {
		return true
	}
	_, ok := parseIPv4Part(last)
	return ok
}

// parseIPv4Part parses one part of an IPv4 address in decimal, 0x hexadecimal or 0 octal.
func parseIPv4Part(part string) (uint64, bool) {
	if part == "" {
		return 0, false
	}
	base := 10
	digits := part
	switch {
	case len(part) >= 2 && (part[:2] == "0x" || part[:2] == "0X"):
		base, digits = 16, part[2:]
	case len(part) >= 2 && part[0] == '0':
		base, digits = 8, part[1:]
	}
	if digits == "" {
		return 0, true // "0x" is zero
	}
	value, err := strconv.ParseUint(digits, base, 64)
	return value, err == nil
}

// parseIPv4 parses an IPv4 host in any of its numeric notations and writes it as four decimal
// parts. A host that ends in a number but is not a valid address is an error.
func parseIPv4(host string) (string, error) {
	parts := strings.Split(host, ".")
	if parts[len(parts)-1] == "" && len(parts) > 1 {
		parts = parts[:len(parts)-1]
	}
	if len(parts) > ipv4PartCount {
		return "", fmt.Errorf("invalid IPv4 address %q", host)
	}
	numbers := make([]uint64, len(parts))
	for i, part := range parts {
		value, ok := parseIPv4Part(part)
		if !ok {
			return "", fmt.Errorf("invalid IPv4 address %q", host)
		}
		numbers[i] = value
	}
	for _, number := range numbers[:len(numbers)-1] {
		if number > maxIPv4Part {
			return "", fmt.Errorf("invalid IPv4 address %q", host)
		}
	}
	last := numbers[len(numbers)-1]
	if last >= 1<<(ipv4PartBits*uint(ipv4PartCount-len(numbers)+1)) {
		return "", fmt.Errorf("invalid IPv4 address %q", host)
	}
	address := last
	for i, number := range numbers[:len(numbers)-1] {
		address += number << (ipv4PartBits * uint(ipv4PartCount-1-i))
	}
	octets := make([]string, ipv4PartCount)
	for i := range octets {
		octets[i] = strconv.FormatUint(address>>(ipv4PartBits*uint(ipv4PartCount-1-i))&maxIPv4Part, 10)
	}
	return strings.Join(octets, "."), nil
}

// Punycode (RFC 3492) parameters.
const (
	punycodeBase        = 36
	punycodeTMin        = 1
	punycodeTMax        = 26
	punycodeSkew        = 38
	punycodeDamp        = 700
	punycodeInitialBias = 72
	punycodeInitialN    = 128
	punycodePrefix      = "xn--"
)

// hostToASCII converts every non-ASCII label of a lowercased host to its punycode form.
func hostToASCII(host string) string {
	labels := strings.Split(host, ".")
	for i, label := range labels {
		if !isASCII(label) {
			labels[i] = punycodePrefix + punycodeEncode(label)
		}
	}
	return strings.Join(labels, ".")
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// punycodeEncode encodes a label as punycode (RFC 3492), without the "xn--" prefix.
func punycodeEncode(label string) string {
	runes := []rune(label)
	var out []byte
	for _, r := range runes {
		if r < punycodeInitialN {
			out = append(out, byte(r))
		}
	}
	basicCount := len(out)
	handled := basicCount
	if basicCount > 0 {
		out = append(out, '-')
	}
	n, delta, bias := rune(punycodeInitialN), 0, punycodeInitialBias
	for handled < len(runes) {
		next := rune(unicode.MaxRune + 1)
		for _, r := range runes {
			if r >= n && r < next {
				next = r
			}
		}
		delta += int(next-n) * (handled + 1)
		n = next
		for _, r := range runes {
			if r < n {
				delta++
			}
			if r != n {
				continue
			}
			q := delta
			for k := punycodeBase; ; k += punycodeBase {
				threshold := k - bias
				if threshold < punycodeTMin {
					threshold = punycodeTMin
				} else if threshold > punycodeTMax {
					threshold = punycodeTMax
				}
				if q < threshold {
					break
				}
				out = append(out, punycodeDigit(threshold+(q-threshold)%(punycodeBase-threshold)))
				q = (q - threshold) / (punycodeBase - threshold)
			}
			out = append(out, punycodeDigit(q))
			bias = adaptPunycodeBias(delta, handled+1, handled == basicCount)
			delta = 0
			handled++
		}
		delta++
		n++
	}
	return string(out)
}

// punycodeDigit maps 0..35 to "a".."z", "0".."9".
func punycodeDigit(d int) byte {
	const letters = 26
	if d < letters {
		return byte('a' + d)
	}
	return byte('0' + d - letters)
}

// adaptPunycodeBias is the bias adaptation function of RFC 3492, section 6.1.
func adaptPunycodeBias(delta, points int, first bool) int {
	if first {
		delta /= punycodeDamp
	} else {
		delta /= 2
	}
	delta += delta / points
	k := 0
	for delta > ((punycodeBase-punycodeTMin)*punycodeTMax)/2 {
		delta /= punycodeBase - punycodeTMin
		k += punycodeBase
	}
	return k + (punycodeBase-punycodeTMin+1)*delta/(delta+punycodeSkew)
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
