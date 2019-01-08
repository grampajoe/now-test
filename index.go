package main

import (
  "fmt"
  "log"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  log.Printf("got a request! headers: %v", r.Header)
  fmt.Fprintf(w, "hello")
}
