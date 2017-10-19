//main.go

package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
  "os"
)


func home(w http.ResponseWriter, r *http.Requestm ){
  t := template.New("/views/index.html")

  t, err := t.ParseFiles("/views/index.html")

  if err!=nil{
      log.Println(err)
  }
  t.Execute(w)
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
