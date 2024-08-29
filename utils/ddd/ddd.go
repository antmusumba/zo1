package main

import (
    "html/template"
    "net/http"
)

type PageData struct {
    Title   string
    Content string
}

func homePage(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "Home Page",
        Content: "Welcome to the Home Page!",
    }
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, data)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "About Page",
        Content: "Welcome to the About Page!",
    }
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, data)
}

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/about", aboutPage)
    http.ListenAndServe(":8080", nil)
}