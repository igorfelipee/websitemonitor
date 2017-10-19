//main.go

package main

import (
  "html/template"
  "net/http"
  "os"
)


func home(w http.ResponseWriter, r *http.Request ){
  t := template.New("/views/index.html")

  t, _ := t.ParseFiles("/views/index.html")
  t.Execute(w)
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
