package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../config"
	"../library"
	"../middleware"
	"../model"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)
	// name := library.GetCookie(r, "name")
	role, _ := strconv.Atoi(library.GetCookie(r, "role"))

	if role == library.Role_anonymous {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if role == library.Role_admin {
		users := model.GetAllUsers()

		data := struct {
			Header Header
			List   interface{}
		}{getHeader(r, "Dashboard | Administrator"), users}

		t, _ := template.ParseFiles("views/admin_dashboard.tmpl", "views/templates/main_template.tmpl")
		t.ExecuteTemplate(w, "dashboard", data)
	}
	if role == library.Role_user_type_1 {
		userDashboard(w, r)
	} else if role == library.Role_user_type_2 {
		type2Dashboard(w, r)
	} else if role == library.Role_user_type_3 {
		type3Dashboard(w, r)
	}

}

func DashboarShowPdf(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	sId := ids[0]
	id, _ := strconv.Atoi(sId)

	file_details := model.GetUserFileDetails(id)
	filename := config.Upload_location + file_details.Filename
	role, _ := strconv.Atoi(library.GetCookie(r, "role"))
	files := model.GetUserFiles(file_details.UserID)
	previousIndex := getPreviousIndex(files, file_details.ID)
	nextIndex := getNextIndex(files, file_details.ID)

	resp_message := ""

	if r.Method == "POST" {
		sUser_id := library.GetCookie(r, "user_id")
		user_id, _ := strconv.Atoi(sUser_id)

		file_category := r.PostFormValue("file_category")
		file, handle, err := r.FormFile("file")

		if err != nil {
			log.Print(err.Error())
			return
		}

		filename := sUser_id + "-" + file_category + "-" + handle.Filename
		destination := "./public/uploads/" + filename
		err = uploadPdfFile(file, handle, destination)

		if err != nil {
			log.Print(err.Error())
		}

		model.UserFileUpdate(user_id, file_category, filename)
		resp_message = "Status updated successfully!"
	}

	file_details = model.GetUserFileDetails(id)

	data := struct {
		Header        Header
		Filename      string
		File          model.UserFile
		Role          int
		Message       string
		PreviousIndex int
		NextIndex     int
	}{getHeader(r, "View Pdf"), filename, file_details, role, resp_message, previousIndex, nextIndex}

	t, _ := template.ParseFiles("views/pdf_view.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "pdf_view", data)
}
