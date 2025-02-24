package main

import (
  "net/http"
  "html/template"
)

func main() {
  http.HandleFunc("/", helloHandler)
  http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("static/index.html"))
  tmpl.Execute(w, nil)
}
