package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
	"unicode"
)

// информация о количестве цифр в каждом слове
type counter map[string]int

// слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// начало решения

// считает количество цифр в словах
// считает количество цифр в словах
func countDigitsInWords(ctx context.Context, words []string) counter {
	pending := submitWords(ctx, words)
	counted := countWords(ctx, pending)
	return fillStats(ctx, counted)
}

// отправляет слова на подсчет
func submitWords(ctx context.Context, words []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, word := range words {
			select {
			case <-ctx.Done():
				return
			case out <- word:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	fmt.Printf("осталось горутин submitWords: %d\n", runtime.NumGoroutine())
	return out
}

// считает цифры в словах
func countWords(ctx context.Context, in <-chan string) <-chan pair {
	out := make(chan pair)

	go func() {
		defer close(out)
		for word := range in {
			select {
			case <-ctx.Done():
				return
			default:

				count := countDigits(word)
				out <- pair{word, count}
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	fmt.Printf("осталось горутин countWords: %d\n", runtime.NumGoroutine())
	return out
}

// готовит итоговую статистику
func fillStats(ctx context.Context, in <-chan pair) counter {
	stats := counter{}
	for p := range in {
		select {
		case <-ctx.Done():
			return stats
		default:
			stats[p.word] = p.count
		}
	}

	fmt.Printf("осталось горутин fillStats: %d\n", runtime.NumGoroutine())
	return stats
}

// конец решения

// считает количество цифр в слове
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	words := strings.Fields(phrase)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stats := countDigitsInWords(ctx, words)

	fmt.Printf("осталось горутин: %d\n", runtime.NumGoroutine()-1)
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	fmt.Println(stats)
}
