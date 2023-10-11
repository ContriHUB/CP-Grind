package main

import (
    "html/template"
    "net/http"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type UserProfile struct {
    gorm.Model
    Username string
    Email    string
    Bio      string
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    var user UserProfile
    db.First(&user)
    renderTemplate(w, "profile.html", user)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    var user UserProfile
    db.First(&user)
    renderTemplate(w, "edit.html", user)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    var user UserProfile
    db.First(&user)

    user.Username = r.FormValue("username")
    user.Email = r.FormValue("email")
    user.Bio = r.FormValue("bio")

    db.Save(&user)

    http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles("templates/" + tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    var err error
    db, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()
    db.AutoMigrate(&UserProfile{})

    http.HandleFunc("/profile", profileHandler)
    http.HandleFunc("/edit", editHandler)
    http.HandleFunc("/update", updateHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8080", nil)
}
