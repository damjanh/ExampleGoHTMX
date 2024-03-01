package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Pet struct {
  Name string
  Animal string
}

func main() {

  data := map[string][]Pet{
    "Pets": {
      {Name: "Lady", Animal: "Bern Shepherd"},
      {Name: "Loki", Animal: "British Shorthair"},
      {Name: "Stella", Animal: "Labrador"},
    },
  }

  h1 := func(w http.ResponseWriter, _ *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, data)
  }

  h2 := func(w http.ResponseWriter, r *http.Request) {
    // Simulate loading
    time.Sleep(1 * time.Second)
    name := r.PostFormValue("name")
    animal := r.PostFormValue("animal")
    data["Pets"] = append(data["Pets"], Pet{Name: name, Animal: animal})
    fmt.Println(data)
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "pet-list-element", Pet{Name: name, Animal: animal})
  }

  http.HandleFunc("/", h1)

  http.HandleFunc("/add-pet/", h2)

  log.Fatal( http.ListenAndServe(":8888", nil))
}
