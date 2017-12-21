package main

import (
	"fmt"
	"net/http"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before")
		defer fmt.Println("after")
		h.ServeHTTP(w, r)
	})
}

func MustParams(h http.Handler, params ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		q := r.URL.Query()
		for _, param := range params {
			if len(q.Get(param)) == 0 {
				http.Error(w, "missing "+param, http.StatusBadRequest)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
