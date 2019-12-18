package main

import (
    "html/template"
    "net/http"
)

type Todo struct {
    Title string
    Done  bool
    Num int
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "This is TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false, Num: 100},
                {Title: "Task 2", Done: true, Num: 200},
                {Title: "Task 3", Done: true, Num: 300},
                {Title: "Task 4", Done: true, Num: 400},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":80", nil)
}
