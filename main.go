//main.go

package main

import (
  "html/template"
  "net/http"
  "os"
)

type Page struct {
    Title string
}

func home(w http.ResponseWriter, r *http.Request ){
  page := Page{"Index"}
  tmpl, err := template.ParseFiles("views/index.html")
  if err != nil {
      panic(err)
  }
  err = tmpl.ExecuteTemplate(os.Stdout, "views/index.html", page)
  if err != nil {
      panic(err)
  }
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
