package main

import (
	"sort"
	"strconv"
	"strings"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func Prettify(m map[string]int) string {
	var builder strings.Builder
	if len(m) == 0 {
		builder.WriteString("{}")
	} else if len(m) == 1 {
		for k, v := range m {
			builder.WriteString("{ ")
			builder.WriteString(k)
			builder.WriteString(": ")
			builder.WriteString(strconv.Itoa(v))
			builder.WriteString(" }")
		}
	} else {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		builder.WriteString("{\n")
		for _, k := range keys {
			builder.WriteString("    ")
			builder.WriteString(k)
			builder.WriteString(": ")
			builder.WriteString(strconv.Itoa(m[k]))
			builder.WriteString(",\n")
		}
		builder.WriteString("}")
	}
	return builder.String()
}

// конец решения

// func Test(t *testing.T) {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
// 	got := prettify(m)
// 	if got != want {
// 		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
// 	}
// }
