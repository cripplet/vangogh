package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

type Widget struct {
  Name  string
  Price int
}

type ViewData struct {
  Name    string
  Widgets []Widget
}

func main() {
  var err error
  testTemplate, err = template.ParseFiles("lib/core/template/hello.gohtml")
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe("0.0.0.0:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  err := testTemplate.Execute(w, ViewData{
    Name: "John Smith",
    Widgets: []Widget{
      Widget{"Blue Widget", 12},
      Widget{"Red Widget", 12},
      Widget{"Green Widget", 12},
    },
  })
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
