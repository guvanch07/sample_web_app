package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", UsersHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method != http.MethodPost {
		errorMap := map[string]any{
			"ok":    false,
			"error": "method is not allowed",
		}
		WriteJson(w, http.StatusMethodNotAllowed, errorMap)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		errorMap := map[string]any{
			"ok":    false,
			"error": err,
		}
		WriteJson(w, http.StatusInternalServerError, errorMap)
		return
	}
	fmt.Printf("user %v", user)
	WriteJson(w, http.StatusOK, map[string]any{
		"ok": true,
	})
}
