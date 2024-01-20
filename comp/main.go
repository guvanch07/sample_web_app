package main

import (
	"bytes"
	"fmt"
	"strings"
)

func slugify(src string) string {
	var result bytes.Buffer
	var currentWord bytes.Buffer
	src = strings.ToLower(src)

	isSafe := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-'
	}
	result.Grow(1)

	for i := 0; i < len(src); i++ {
		char := src[i]

		if isSafe(char) {
			currentWord.WriteByte(char)
		} else if currentWord.Len() > 0 {
			if result.Len() > 0 {
				result.WriteByte('-')
			}
			result.Write(currentWord.Bytes())
			currentWord.Reset()
		}
	}

	if currentWord.Len() > 0 {
		if result.Len() > 0 {
			result.WriteByte('-')
		}
		result.Write(currentWord.Bytes())
	}

	return result.String()
}
func slugify1(src string) string {
	var result bytes.Buffer
	src = strings.ToLower(src)

	words := strings.FieldsFunc(src, func(r rune) bool {
		return (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-'
	})
	for i := 0; i < len(words); i++ {
		result.Grow(i)
		b := words[i]
		result.WriteString(b)
		if len(words)-1 != i {
			result.WriteByte('-')
		}
	}
	return result.String()
}

func main() {
	const phrase = "A 100x Investment (2019)"
	slug := slugify1(phrase)
	fmt.Println(slug)
	// Output: attention-attention-
}
