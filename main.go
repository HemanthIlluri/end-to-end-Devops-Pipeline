package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from DevOps Project 🚀 END to END Pipeline Project And Sucessfully Deployed App with ARGOCD")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

