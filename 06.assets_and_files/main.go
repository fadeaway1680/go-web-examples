package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Go server is running on :8011")
	http.ListenAndServe(":8011", nil)
}