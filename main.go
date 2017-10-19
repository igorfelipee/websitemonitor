//main.go

package main

import (

  "net/http"
  "io"
  "os"
)

func home(w http.ResponseWriter, r *http.Request){
  io.WriteString(w, "Hello, World!")
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
