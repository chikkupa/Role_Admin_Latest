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

func type2Dashboard(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	if model.GetUserRole(r) != library.Role_user_type_2 {
		http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
		return
	}

	users := model.GetUsersListForUserType2()

	data := struct {
		Header Header
		List   interface{}
	}{getHeader(r, "Dashboard | User Type 2"), users}

	t, _ := template.ParseFiles("views/type_2_dashboard.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "dashboard", data)
}

func Type2UserDetails(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	if model.GetUserRole(r) != library.Role_user_type_2 {
		http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
		return
	}

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	userID, _ := strconv.Atoi(ids[0])
	user := model.GetUserDetails(userID)

	if r.Method == "GET" {
		files := model.GetUserFiles(userID)

		message := library.GetCookie(r, "message")
		user.Message = message
		especialidadeFiles := model.GetEspecialidadeFiles(userID)
		library.PutCookie(w, "message", "")

		data := struct {
			Header             Header
			Files              interface{}
			User               model.User
			EspecialidadeFiles interface{}
		}{getHeader(r, "View User Files | User Type 2"), files, user, especialidadeFiles}
		t, _ := template.ParseFiles("views/type_2_user_details.tmpl", "views/templates/especialidade_files_template.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
		t.ExecuteTemplate(w, "user_details", data)
	} else {
		status, _ := strconv.Atoi(r.PostFormValue("status"))

		message := ""
		if user.Status != status {
			model.UpdateUserStatus(userID, status)
			message = "User status updated successfully!"

			SendStatusChangedEmail(user, user.Status, status)

		} else {
			message = "No change"
		}

		library.PutCookie(w, "message", message)
		http.Redirect(w, r, "/user_details?id="+r.PostFormValue("id"), http.StatusSeeOther)
	}
}

// Type2ApprovePdf Approve user uploaded PDF files
func Type2ApprovePdf(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	/*if model.GetUserRole(r) != library.Role_user_type_2 && model.GetUserRole(r) != library.Role_admin {
		http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
		return
	} */

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	respMessage := ""

	id, _ := strconv.Atoi(ids[0])
	file := model.GetUserFileDetails(id)
	files := model.GetUserFiles(file.UserID)
	previousIndex := getPreviousIndex(files, file.ID)
	nextIndex := getNextIndex(files, file.ID)

	if r.Method == "POST" {
		fileID, _ := strconv.Atoi(r.PostFormValue("id"))
		status, _ := strconv.Atoi(r.PostFormValue("status"))
		expiryDate := r.PostFormValue("expiry_date")
		message := r.PostFormValue("message")
		userID, _ := strconv.Atoi(r.PostFormValue("user_id"))

		model.ChangeUserFileStatus(fileID, message, status, expiryDate)

		isApproved := model.ChangeStatusIfUploaded(userID)

		if isApproved {
			user := model.GetUserDetails(userID)
			SendStatusChangedEmail(user, 4, 5)
			// log.Print("User [user_id: ", user_id, "] status changed to approved")
		}

		respMessage = "File status changed successfully!"

		statusList := []string{"Pending", "Approved", "Denied"}
		logMessage := "Status of user[" + strconv.Itoa(userID) + "] file type: " + file.FileCategory + " changed to " + statusList[status]
		model.AddLog(model.GetUserId(r), logMessage)

		if nextIndex != 0 {
			http.Redirect(w, r, "/approve_pdf?id="+strconv.Itoa(nextIndex), http.StatusSeeOther)
			return
		}
	}

	filename := config.Upload_location + file.Filename

	data := struct {
		Header        Header
		Filename      string
		File          model.UserFile
		Message       string
		PreviousIndex int
		NextIndex     int
	}{getHeader(r, "Verify Pdf | User Type 2"), filename, file, respMessage, previousIndex, nextIndex}
	t, _ := template.ParseFiles("views/type_2_approve_pdf.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "approve_pdf", data)
	return
}

func getNextIndex(files interface{}, current int) int {
	fileArray := files.([]interface{})

	for i := 0; i < len(fileArray); i++ {
		file := fileArray[i].(model.UserFile)

		if file.ID == current {
			if (i + 1) < len(fileArray) {
				file = fileArray[i+1].(model.UserFile)
				return file.ID
			}
		}
	}

	return 0
}

func getPreviousIndex(files interface{}, current int) int {
	fileArray := files.([]interface{})

	for i := 0; i < len(fileArray); i++ {
		file := fileArray[i].(model.UserFile)

		if file.ID == current {
			if (i - 1) >= 0 {
				file = fileArray[i-1].(model.UserFile)
				return file.ID
			}
		}
	}

	return 0
}
