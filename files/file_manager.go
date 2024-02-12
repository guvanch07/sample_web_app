package main

import (
	"fmt"
	"math/rand"
	"os"
)

const alphabet = "aeiourtnsl"

type Census struct {
	files map[rune]*os.File
}

func (c *Census) Count() int {
	total := 0
	for _, file := range c.files {
		info, err := file.Stat()
		if err != nil {
			continue
		}
		total += int(info.Size()) / 6 // Размер строки имени включая символ новой строки.
	}
	return total
}

func (c *Census) Add(name string) {
	firstLetter := rune(name[0])
	file, ok := c.files[firstLetter]
	if !ok {
		return
	}
	_, _ = fmt.Fprintf(file, "%s\n", name)
}

func (c *Census) Close() {
	for _, file := range c.files {
		_ = file.Close()
	}
}

func NewCensus() *Census {
	c := &Census{
		files: make(map[rune]*os.File),
	}
	for _, letter := range alphabet {
		filePath := fmt.Sprintf("%c.txt", letter)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", filePath, err)
			continue
		}
		c.files[letter] = file
	}
	return c
}

func mainFile() {
	rand.Seed(0)
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}

func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}
