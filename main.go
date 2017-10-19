//main.go

package main

import (
  "log"
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

  fs := http.FileServer(http.Dir("views"))
  http.Handle("/views/", http.StripPrefix("/views/", fs))

  http.HandleFunc("/", home)

  http.ListenAndServe(":"+port, nil)
  log.Println("Listening...")
  //http.ListenAndServe(":8080", nil)
}
