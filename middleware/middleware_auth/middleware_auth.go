package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/getMe", GetMeHandler)

	middlewares := []func(http.Handler) http.Handler{
		LoggingMiddleware,
		AuthMiddleware,
	}
	handler := http.Handler(mux)
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, map[string]any{
		"ok": true,
	})
}

func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("user").(User)
	if user.Id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		WriteJSON(w, map[string]any{
			"ok":    false,
			"error": "unauthoried",
		})
		return
	}
	WriteJSON(w, map[string]any{
		"ok":   true,
		"user": user,
	})
}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session")
		if cookie != nil {
			sessionId := cookie.Value
			user, _ := GetUserBySessionId(sessionId)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)
			r = r.WithContext(ctx)
		}
		handler.ServeHTTP(w, r)
	})
}

func WriteJSON(w io.Writer, v any) {
	bytes, _ := json.Marshal(v)
	w.Write(bytes)
}

type User struct {
	Name string
	Id   int
}

func GetUserBySessionId(sessionId string) (User, error) {
	if sessionId == "123" {
		return User{Id: 1, Name: "admin"}, nil
	}
	return User{}, errors.New("sesseion not found")
}
