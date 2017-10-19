//main.go

package main

import (
  "html/template"
  "net/http"
  "os"
)

func home(w http.ResponseWriter, r *http.Request ){
  t, _ := template.ParseFiles("./views/index.html")
  token:= 1234
  t.Execute(w, token)
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
