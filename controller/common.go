package controller

import(
	"net/http"
	"html/template"
)

func CommonShowUnauthorised(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("views/unauthorised.tmpl", "views/templates/success_template.tmpl")
    t.ExecuteTemplate(w, "unauthorised", nil)
}