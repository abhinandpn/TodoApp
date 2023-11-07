package Todo

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
)

func FilePass() error {

	// List files in the templates directory
	files, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List of HTML files in templates directory:")
	for _, file := range files {
		fmt.Println(file)
	}

	Tmpl = template.Must(template.ParseGlob("templates/*.html"))
	return nil
}
