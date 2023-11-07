package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/abhinandpn/TodoApp/Todo"
)

var Tmpl *template.Template

func main() {

	mux := http.NewServeMux()

	// Check if the templates directory exists
	if _, err := os.Stat("templetes"); os.IsNotExist(err) {
		log.Fatal("Templates directory does not exist")
	}

	// List files in the templates directory
	files, err := filepath.Glob("templetes/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List of HTML files in templates directory:")
	for _, file := range files {
		fmt.Println(file)
	}

	Tmpl = template.Must(template.ParseGlob("templetes/*.html"))
	mux.HandleFunc("/todo", Todo.Todo)
	log.Fatal(http.ListenAndServe(":9090", mux))
}
