package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func main() {
	const phrase = "Understanding the x64 code models (2012)"
	slug := slugify(phrase)
	fmt.Println(slug)
	// a-100x-investment-2019
}

func slugify1(src string) string {
	src = strings.ToLower(src)
	var builder strings.Builder
	words := strings.FieldsFunc(src, func(r rune) bool {
		return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && r != '-'
	})
	builder.Grow(1)
	for i, v := range words {
		builder.WriteString(v)
		if len(words)-1 != i {
			builder.WriteString("-")
		}
	}
	return builder.String()
}

func slugify2(src string) string {
	res := strings.ToLower(src)
	res = strings.Map(purifyChar, res)
	words := strings.Fields(res)
	return strings.Join(words, "-")
}

func purifyChar(r rune) rune {
	const validChars string = "abcdefghijklmnopqrstuvwxyz01234567890- "
	if strings.IndexRune(validChars, r) == -1 {
		return ' '
	}
	return r
}

func slugify(src string) string {
	var result bytes.Buffer
	src = strings.ToLower(src)

	isSafe := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-'
	}
	result.Grow(len(src))

	for i := 0; i < len(src); i++ {
		char := src[i]
		if char == ')' || char == '(' {
			continue
		} else if isSafe(char) {
			result.WriteByte(char)
		} else if result.Len() > 0 && char != '-' {
			result.WriteByte('-')
		}
	}
	return result.String()
}

func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
