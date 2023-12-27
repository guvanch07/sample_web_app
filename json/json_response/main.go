package main

import (
	"encoding/json"
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
	user1 := User{Name: "Jakob", Id: 1}
	err := WriteJson(w, http.StatusOK, user1)

	if err != nil {
		errorMap := map[string]any{
			"ok":    false,
			"error": err,
		}
		WriteJson(w, http.StatusInternalServerError, errorMap)
		return
	}
}
