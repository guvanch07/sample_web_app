package main

import (
	"crypto/rand"
	"io"
)

func RandomReader(max int) io.Reader {
	return &randomReader{max: max}
}

type randomReader struct {
	max   int
	count int
}

func (r *randomReader) Read(p []byte) (n int, err error) {
	if r.count >= r.max {
		return 0, io.EOF
	}

	lenP := len(p)
	if lenP > r.max-r.count {
		lenP = r.max - r.count
	}

	// Генерируем случайные байты с использованием rand.Read()
	randomBytes := make([]byte, lenP)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return 0, err
	}

	// Копируем случайные байты в буфер p
	copy(p, randomBytes)
	r.count += lenP

	return lenP, nil
}

// using io.LimitReader()
type randomReader1 struct{}

func (r *randomReader) Read1(p []byte) (n int, err error) {
	return rand.Read(p)
}

func RandomReader1(max int) io.Reader {
	rd := &randomReader{}
	return io.LimitReader(rd, int64(max))
}

// not using LimitReader()
type randomReader2 struct {
	n int
}

func (r *randomReader2) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if len(p) > r.n {
		p = p[:r.n]
	}
	n, err = rand.Read(p)
	r.n -= n
	return
}

func RandomReader2(max int) io.Reader {
	return &randomReader2{n: max}
}
