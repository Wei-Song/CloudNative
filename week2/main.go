package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func response(w http.ResponseWriter, r *http.Request) {

	//get VERSION
	ver := os.Getenv("VERSION")
	h := r.Header
	h.Add("VERSION", ver)

	//write to response header
	for key, values := range h {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}

	fmt.Println("response header :", h)

	//log
	fmt.Printf("visit from ip: %v status code: %d", r.Host, 200)

	fmt.Fprintln(w, h)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
