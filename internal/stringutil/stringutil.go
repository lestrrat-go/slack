package stringutil

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/iancoleman/strcase"
)

// Copied from golint
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

func Snake(s string) string {
	for wd, _ := range commonInitialisms {
		tmp := s
		i := strings.Index(tmp, wd)
		if i == -1 {
			continue
		}

		var buf bytes.Buffer
		for i > -1 {
			if i > 0 {
				buf.WriteString(tmp[:i])
				if tmp[i-1] != '_' {
					buf.WriteByte('_')
				}
			}

			buf.WriteString(strings.ToLower(wd))
			tmp = tmp[i+len(wd):]
			i = strings.Index(tmp, wd)
		}

		if len(tmp) > 0 {
			buf.WriteString(tmp)
		}

		s = buf.String()
	}

	// grr, apparently strcase doesn't handle foo-bar-baz (the "-"s)
	s = strings.Replace(s, "-", " ", -1)
	s = strcase.ToSnake(s)

	// Except...
	s = strings.Replace(s, "o_auth", "oauth", -1)
	s = strings.Replace(s, "open_api", "openapi", -1)
	return s
}

func Camel(s string) string {
	return strcase.ToCamel(s)
}

func LowerCamel(s string) string {
	return LcFirst(strcase.ToCamel(s))
}

func LcFirst(s string) string {
	if len(s) <= 0 {
		return s
	}

	r, w := utf8.DecodeRuneInString(s)
	var buf bytes.Buffer
	buf.WriteRune(unicode.ToLower(r))
	buf.WriteString(s[w:])
	return buf.String()
}

func UcFirst(s string) string {
	if len(s) <= 0 {
		return s
	}

	r, w := utf8.DecodeRuneInString(s)
	var buf bytes.Buffer
	buf.WriteRune(unicode.ToUpper(r))
	buf.WriteString(s[w:])
	return buf.String()
}
