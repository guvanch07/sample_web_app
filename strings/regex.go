package main

import (
	"regexp"
	"strings"
)

func Slugify(src string) string {
	src = strings.ToLower(src)

	// Заменяем все символы, не являющиеся латинскими буквами, цифрами или дефисом, на пробел
	re := regexp.MustCompile("[^a-z0-9-]+")
	safeStr := re.ReplaceAllString(src, " ")

	// Разделяем на "слова" и объединяем через дефис
	words := strings.Fields(safeStr)
	result := strings.Join(words, "-")
	return result
}
