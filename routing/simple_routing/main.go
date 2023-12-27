package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/", index)
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handle404(w, r)
		return
	}
	w.Write([]byte("Home"))
}
func hello(w http.ResponseWriter, r *http.Request) {
	pathRegexp := regexp.MustCompile(`^/hello/\w+$`)
	if !pathRegexp.Match([]byte(r.URL.Path)) {
		handle404(w, r)
		return
	}
	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Page No"))
}
