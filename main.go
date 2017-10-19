//main.go

package main

import (
  "html/template"
  "net/http"
  "os"
)

func home(w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, p)
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", home)
  http.ListenAndServe(":"+port, nil)

}
