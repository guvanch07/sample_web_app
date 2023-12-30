package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/form", FormHandler)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
	file, header, err := r.FormFile("myfile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	fmt.Println(header.Filename)
	fmt.Println(header.Header)
	fmt.Println(header.Size)

	ext := path.Ext(header.Filename)
	tmpFile, err := os.CreateTemp("tmp", "*"+ext)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}
	defer tmpFile.Close()
	fmt.Println(tmpFile.Name())

	bytes, err := io.ReadAll(tmpFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	n, err := tmpFile.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	fmt.Println(n)
	fmt.Fprintf(w, "File uploaded!")
}
