package Todo

import (
	"html/template"
	"net/http"

	"github.com/abhinandpn/TodoApp/util"
)

var Tmpl *template.Template

func Todo(w http.ResponseWriter, r *http.Request) {
	if Tmpl != nil {
		data := util.PageData{
			Title: "Todo List",
			Todos: []util.Todo{
				{Item: "install Go", Done: true},
				{Item: "Run Go", Done: true},
				{Item: "Create Welcome message", Done: true},
				{Item: "Create Package and importing", Done: true},
			},
		}
		Tmpl.Execute(w, data)
	} else {
		http.Error(w, "Template not initialized", http.StatusInternalServerError)
	}
}
