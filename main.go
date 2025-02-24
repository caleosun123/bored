package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "html/template"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

  var err error
  db, err = sql.Open("mysql", "sql5764415:nEIpTAV8Hj@tcp(sql5.freesqldatabase.com:3306)/sql5764415")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Successfully connected to the database")
  
  http.HandleFunc("/", helloHandler)
  http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("static/index.html"))
  tmpl.Execute(w, nil)
}
