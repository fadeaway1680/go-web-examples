package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    f(w, r)
  }
}

func foo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "bar")
}

func main() {
  http.HandleFunc("/foo", logging(foo))
  http.HandleFunc("/bar", logging(bar))

  fmt.Println("Go server is running on :8011")
  http.ListenAndServe(":8011", nil)
}