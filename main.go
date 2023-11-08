package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/abhinandpn/TodoApp/Todo"
)

var Tmp *template.Template

func main() {
	mux := http.NewServeMux()

	// Check if the templates directory exists
	if _, err := os.Stat("templates"); os.IsNotExist(err) {
		log.Fatal("Templates directory does not exist")
	}

	// List files in the templates directory
	err := Todo.FilePass()
	if err != nil {
		log.Fatalf("Error occurred while parsing templates: %v", err)
	}

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	mux.HandleFunc("/todo", Todo.Todo)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
