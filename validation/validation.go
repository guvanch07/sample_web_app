package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u User) validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 15)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Phone, is.E164),
	)
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

	err = user.validate()
	errors, ok := err.(validation.Errors)
	if ok {
		for key, err := range errors {
			verr, ok := err.(validation.Error)
			if ok {
				fmt.Printf("%#v %#v %#v\n", key, verr.Code(), verr.Error())
			}
		}
	}

	// if err != nil {
	// 	errorMap := map[string]any{
	// 		"ok":    false,
	// 		"error": err,
	// 	}
	// 	WriteJson(w, http.StatusBadRequest, errorMap)
	// 	return
	// }

	fmt.Printf("user %v", user)
	WriteJson(w, http.StatusOK, map[string]any{
		"ok": true,
	})
}
