package main

import (
  "fmt"
  "log"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  log.Printf("got a request! headers: %v", r.Header)
  for _, i := range("hello") {
    fmt.Fprintf(w, "%v :)", i)
  }
}
