package parser

import (
	"bytes"

	"github.com/yuin/goldmark"
)

// ParseMD is used to parse markdown
func ParseMD(source string) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(source), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
