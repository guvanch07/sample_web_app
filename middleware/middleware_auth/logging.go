package main

import (
	"log"
	"net/http"
)

type MyResponseWritter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *MyResponseWritter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w2 := &MyResponseWritter{ResponseWriter: w, StatusCode: http.StatusOK}
		handler.ServeHTTP(w2, r)
		log.Printf("%s\n: [%d]", r.RequestURI, w2.StatusCode)
	})
}
