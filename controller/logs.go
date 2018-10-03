package controller

import(
    "net/http"
    "html/template"
    "strconv"
    "log"

    "../model"
    "../library"
    "../middleware"
)

func Logs(w http.ResponseWriter, r *http.Request){
	middleware.RedirectIfNotLoggedIn(w, r)

	role := model.GetUserRole(r)
	user_id := model.GetUserId(r)

	if role == library.Role_anonymous {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    var logs interface{}

    if role == library.Role_admin {
		logs = model.GetAllLogUsers()
	} else if role == library.Role_user_type_1 {
		http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
		return
	} else if role == library.Role_user_type_2 {
		http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
		return
	} else if role == library.Role_user_type_3 {
		logs = model.GetAllLogUsersForType3(user_id)
	}

	data := struct{
                Header Header
                Logs interface{}
            }{getHeader(r, "Logs"), logs }

    t, _ := template.ParseFiles("views/log_users.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "logs_users", data)
}

func UserLogView(w http.ResponseWriter, r *http.Request){
	middleware.RedirectIfNotLoggedIn(w, r)
	
	ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    sId := ids[0]
    user_id, _ := strconv.Atoi(sId)

    logs := model.GetUserLogs(user_id)

    data := struct{
                Header Header
                Logs interface{}
            }{getHeader(r, "Logs"), logs }

    t, _ := template.ParseFiles("views/logs.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "logs", data)
}