//main.go

package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
)

func home(w http.ResponseWriter, r *http.Request){
  io.WriteString(w, "Hello, World!")
}

func main() {
  port := os.Getenv("PORT")
  http.HandlerFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
