package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/foo", FooHandler)

	r.Use(MyMiddleware)
	r.Use(SecondMiddleware)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	w.Write([]byte("Home"))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("foo")
	w.Write([]byte("Foo"))
}

func MyMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before")
		handler.ServeHTTP(w, r)
		fmt.Println("after")
	})
}

func SecondMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before 2")
		handler.ServeHTTP(w, r)
		fmt.Println("after 2")
	})
}
