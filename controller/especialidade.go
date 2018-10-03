package controller

import(
	"net/http"
    "html/template"
    "strconv"
    "log"

    "../middleware"
    "../model"
)

func CreateEspecialidade(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    message := ""
    has_error := false
    especialidade := model.Especialidade{Name : "", Expiry_date : "", Description : ""}

    if r.Method == "POST" {
		r.ParseForm()

		name := r.FormValue("name")
        expiry_date := r.FormValue("expiry_date")
        description := r.FormValue("description")

    	model.CreateEspecialidade(name, expiry_date, description)

    	message = "Especialidade added succesfully!"

    	log_message := "New Especialidade [Name: " + name + "] created"
    	model.AddLog(model.GetUserId(r), log_message)
    	log.Print(log_message)
	}

	data := struct{
                Header Header
                Message string
                Has_error bool
                Especialidade model.Especialidade
            }{getHeader(r, "Create New Especialidade"), message, has_error, especialidade}

    t, _ := template.ParseFiles("views/especialidade_add.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "create_especialidade", data)
}

func ListEspecialidade(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    especialidadeList := model.GetEspecialidadeList();

    data := struct{
                Header Header
                List interface{}
            }{getHeader(r, "Especialidade List"), especialidadeList }

    t, _ := template.ParseFiles("views/especialidade_list.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "especialidade_list", data)
}

func UpdateEspecialidade(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    user_id, _ := strconv.Atoi(ids[0])

    message := ""
    has_error := false

    if r.Method == "POST" {
        r.ParseForm()

        id, _ := strconv.Atoi(r.FormValue("id"))
        name := r.FormValue("name")
        expiry_date := r.FormValue("expiry_date")
        description := r.FormValue("description")

        model.UpdateEspecialidade(id, name, expiry_date, description)

        message = "User Especialidade succesfully!"

        log_message := "Especialidade [Especialidade id: " + strconv.Itoa(id) + "] updated successfully"
        model.AddLog(model.GetUserId(r), log_message)
        log.Print(log_message)
    }

    especialidade := model.GetEspecialidadeDetails(user_id)

    data := struct{
                Header Header
                Message string
                Has_error bool
                Especialidade model.Especialidade
            }{getHeader(r, "Update Especialidade Details"), message, has_error, especialidade}

    t, _ := template.ParseFiles("views/especialidade_edit.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "especialidade_edit", data)
}

func DeleteEspecialidade(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    id, _ := strconv.Atoi(ids[0])

    model.DeleteEspecialidade(id)

    log_message := "Especialidade [Especialidade id: " + strconv.Itoa(id) + "] deleted successfully"
    model.AddLog(model.GetUserId(r), log_message)
    log.Print(log_message)

    http.Redirect(w, r, "/list_especialidade", http.StatusSeeOther)
}