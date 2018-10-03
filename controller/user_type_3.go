package controller

import(
	"log"
	"net/http"
    "html/template"
    "strconv"

    "../model"
    "../library"
    "../middleware"
)

func type3Dashboard(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    if model.GetUserRole(r) != library.Role_user_type_3 {
        http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
        return;
    }

	users := model.GetUsersListForUserType3(); // GetUsersWithStatus("5");

    data := struct{
                Header Header
                List interface{}
            }{getHeader(r, "Dashboard | User Type 3"), users }

	t, _ := template.ParseFiles("views/type_3_dashboard.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "dashboard", data)
}

func Type3UserFiles(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)
    
    if model.GetUserRole(r) != library.Role_user_type_3 {
        http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
        return;
    }

	ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    user_id, _ := strconv.Atoi(ids[0])
    user := model.GetUserDetails(user_id)

    if r.Method == "GET" {
        files := model.GetUserFiles(user_id)

        message:= library.GetCookie(r, "message")
        user.Message = message;
        library.PutCookie(w, "message", "")

        data := struct{
                    Header Header
                    Files interface{} 
                    User model.User
                }{getHeader(r, "View User Files | User Type 3"), files, user}
        t, _ := template.ParseFiles("views/type_3_user_files.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
        t.ExecuteTemplate(w, "user_files", data)
    }   else {
        status, _ := strconv.Atoi(r.PostFormValue("status"))

        message := "";
        if user.Status != status {
            model.UpdateUserStatus(user_id, status)
            message = "User status updated successfully!"

            SendStatusChangedEmail(user, user.Status, status)

        }else{
            message = "No change"
        }

        library.PutCookie(w, "message", message)
        http.Redirect(w, r, "/user_files?id=" + r.PostFormValue("id"), http.StatusSeeOther)
    }
        
}