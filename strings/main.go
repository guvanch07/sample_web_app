package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// начало решения

func main() {
	// directions := []string{
	// 	"100m to intersection",
	// 	"turn right",
	// 	"straight 300m",
	// 	"enter motorway",
	// 	"straight 1.6km",
	// 	"exit motorway",
	// 	"500m straight",
	// 	"turn sharp left",
	// 	"continue 100m to destination",
	// }

	// calcDistance(directions)
	// m := map[string]int{"one": 1, "two": 2, "three": 3}
	// str := Prettify(m)
	// fmt.Println(str)
	const phrase = "Go Is Awesome!"
	s := Slugify(phrase)
	fmt.Println(s)
}

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	//var words []string
	var result int
	for _, val := range directions {
		words := strings.FieldsFunc(val, func(r rune) bool {
			return !unicode.IsDigit(r) && r != 'k' && r != 'm' && r != '.'
		})

		for _, word := range words {
			if len(word) > 1 {
				if strings.Contains(word, "km") {
					splitedKm := strings.Split(word, "km")
					if valueKm, err := strconv.ParseFloat(splitedKm[0], 64); err == nil {
						result += int(valueKm * 1000)
					}
				} else if strings.Contains(word, "m") {
					splitedM := strings.Split(word, "m")
					if valueM, err := strconv.ParseFloat(splitedM[0], 64); err == nil {
						result += int(valueM)
					}
				}
			}
		}
	}
	return result
}

// конец решения

// func Test(t *testing.T) {
// 	directions := []string{
// 		"100m to intersection",
// 		"turn right",
// 		"straight 300m",
// 		"enter motorway",
// 		"straight 5km",
// 		"exit motorway",
// 		"500m straight",
// 		"turn sharp left",
// 		"continue 100m to destination",
// 	}
// 	const want = 6000
// 	got := calcDistance(directions)
// 	if got != want {
// 		t.Errorf("%v: got %v, want %v", directions, got, want)
// 	}
// }
