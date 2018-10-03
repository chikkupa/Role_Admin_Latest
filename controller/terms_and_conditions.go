package controller

import(
	"net/http"
    "html/template"
    "log"

    "../middleware"
    "../model"
)

func TermsAndConditions(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    message := ""
    has_error := false

    if r.Method == "POST" {
		r.ParseForm()

		terms := r.FormValue("terms")

    	model.SetDetails("TERMS_AND_CONDITIONS", terms)

    	message = "Terms and condtions updated succesfully!"

    	log_message := "Terms and condtions updated"
    	model.AddLog(model.GetUserId(r), log_message)
    	log.Print(log_message)
	}

	terms_and_conditions := model.GetDetails("TERMS_AND_CONDITIONS")

	data := struct{
                Header Header
                Message string
                Has_error bool
                Terms string
            }{getHeader(r, "Terms and Conditions"), message, has_error, terms_and_conditions}

    t, _ := template.ParseFiles("views/terms_and_condtions.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "terms", data)
}