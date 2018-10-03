package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../library"
	"../middleware"
	"../model"
)

func CreateAdminUser(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	message := ""
	has_error := false
	user := model.InitializeAdminUser()

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		password := library.ConvertToSha1(r.FormValue("password"))
		name := r.FormValue("name")
		lastname := r.FormValue("lastname")
		age, _ := strconv.Atoi(r.FormValue("age"))
		gender := r.FormValue("gender")
		status, _ := strconv.Atoi(r.FormValue("status"))
		task := r.FormValue("task")
		role, _ := strconv.Atoi(r.FormValue("role"))

		if !model.IsUsernameAvailable(username) {
			message = "Username (" + username + ") not available"
			log.Print("Username (", username, ") not available")

			user.Username = username
			user.Name = name
			user.Lastname = lastname
			user.Age = age
			user.Gender = gender
			user.Status = status
			user.Task = task
			user.Role = role

			has_error = true
		} else {
			model.AdminUserRegister(username, password, name, lastname, age, gender, 1, status, task, role)

			message = "User added succesfully!"

			log_message := "New user [Username: " + username + "] created"
			model.AddLog(model.GetUserId(r), log_message)
			log.Print(log_message)
		}
	}

	data := struct {
		Header    Header
		Message   string
		Has_error bool
		User      model.Admin_User
	}{getHeader(r, "Create New Admin"), message, has_error, user}

	t, _ := template.ParseFiles("views/register_admin_user.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "register_admin", data)
}

func ListAdminUsers(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	users := model.GetAllAdminUsers()

	data := struct {
		Header Header
		List   interface{}
	}{getHeader(r, "Admin User List | Administrator"), users}

	t, _ := template.ParseFiles("views/admin_users_list.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "admin_user_list", data)
}

func UpdateAdminUser(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	user_id, _ := strconv.Atoi(ids[0])

	message := ""
	has_error := false
	user := model.GetAdminUserDetails(user_id)

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		username := r.FormValue("username")
		password := library.ConvertToSha1(r.FormValue("password"))
		name := r.FormValue("name")
		lastname := r.FormValue("lastname")
		age, _ := strconv.Atoi(r.FormValue("age"))
		gender := r.FormValue("gender")
		status, _ := strconv.Atoi(r.FormValue("status"))
		task := r.FormValue("task")
		role, _ := strconv.Atoi(r.FormValue("role"))

		if username != user.Username && !model.IsUsernameAvailable(username) {
			message = "Username (" + username + ") not available"
			log.Print("Username (", username, ") not available")

			has_error = true
		} else {
			model.UpdateAdminUser(id, username, password, name, lastname, age, gender, 1, status, task, role)

			message = "User Updated succesfully!"

			log_message := "User [User id: " + strconv.Itoa(id) + "] updated successfully"
			model.AddLog(model.GetUserId(r), log_message)
			log.Print(log_message)
		}

		user.Username = username
		user.Name = name
		user.Lastname = lastname
		user.Age = age
		user.Gender = gender
		user.Status = status
		user.Task = task
		user.Role = role
	}

	data := struct {
		Header    Header
		Message   string
		Has_error bool
		User      model.Admin_User
	}{getHeader(r, "Update Admin Details"), message, has_error, user}

	t, _ := template.ParseFiles("views/admin_user_edit.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "edit_admin", data)
}

func DeleteAdminUser(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	user_id, _ := strconv.Atoi(ids[0])

	model.DeleteAdminUser(user_id)

	log_message := "User [User id: " + strconv.Itoa(user_id) + "] deleted successfully"
	model.AddLog(model.GetUserId(r), log_message)
	log.Print(log_message)

	http.Redirect(w, r, "/list_admins", http.StatusSeeOther)
}
