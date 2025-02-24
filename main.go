package main

import (
  "database/sql"
  "fmt"
  "log"
  "time"
  "net/http"
  "html/template"
  "golang.org/x/crypto/bcrypt"
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
  
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/register", registerHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/dashboard", dashboardHandler)
  http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("static/index.html"))
  tmpl.Execute(w, nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("static/dashboard.html"))
  tmpl.Execute(w, nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    err := r.ParseForm()
    if err != nil {
      http.Error(w, "Unable to parse form", http.StatusBadRequest)
      return
    }

    name := r.FormValue("name")
    email := r.FormValue("email")
    password := r.FormValue("password")

    var existingEmail string
    err = db.QueryRow("SELECT email FROM users WHERE email = ?", email).Scan(&existingEmail)
    if err != nil && err != sql.ErrNoRows {
      http.Error(w, "Unable to query database", http.StatusInternalServerError)
      return
    }
    
    if existingEmail != "" {
      http.Error(w, "Email already registered", http.StatusBadRequest)
      return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
      http.Error(w, "Unable to hash password", http.StatusInternalServerError)
    }

    stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
    if err != nil {
      http.Error(w, "Unable to prepare statement", http.StatusInternalServerError)
      return
    }
    defer stmt.Close()

    _, err = stmt.Exec(name, email, hashedPassword)
    if err != nil {
      http.Error(w, "Unable to execute statement", http.StatusInternalServerError)
      return
    }
    fmt.Fprintf(w, "User %s registered successfully!", name)
  } else {
      tmpl := template.Must(template.ParseFiles("static/register.html"))
      tmpl.Execute(w, nil)
  }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    err := r.ParseForm()
    if err != nil {
      http.Error(w, "Unable to parse form", http.StatusBadRequest)
      return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")

    var dbPassword string
    err = db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)
    if err != nil {
      http.Error(w, "Invalid email or password", http.StatusUnauthorized)
      return
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
    if err != nil {
      http.Error(w, "Invalid email or password", http.StatusUnauthorized)
      return
    }

    http.SetCookie(w, &http.Cookie{
      Name: "session_token",
      Value: "hhbk2012740**0237hhbk2012740**0237hhbk2012740**0237hhbk2012740**0237hhbk2012740**0237",
      Expires: time.Now().Add(24 * time.Hour),
      HttpOnly: true,
    })

    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)    
  } else {
    tmpl := template.Must(template.ParseFiles("static/login.html"))
    tmpl.Execute(w, nil)
  }
}
