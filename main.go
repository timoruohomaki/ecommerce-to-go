package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	// router

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home page"))
	})

	// auth

	mux.HandleFunc("GET /auth/login", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /auth/register", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /auth/logout", func(w http.ResponseWriter, r *http.Request) {})

	// dashboard

	mux.HandleFunc("GET /dashboard", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /dashboard/products", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /dashboard/product/add", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /dashboard/product/update/{id}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /dashboard/product/delete/{id}", func(w http.ResponseWriter, r *http.Request) {})

	// listener

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(srv.ListenAndServe())

}
