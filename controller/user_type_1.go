package controller

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"../library"
	"../middleware"
	"../model"
)

func userDashboard(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	userID, _ := strconv.Atoi(library.GetCookie(r, "user_id"))
	status, _ := strconv.Atoi(library.GetCookie(r, "status"))
	user := model.GetUserDetails(userID)
	fileCategories := model.GetUserFileCategoryList()
	files := model.GetUserFiles(userID)
	especialidadeFiles := model.GetEspecialidadeFiles(userID)

	if status >= 2 {
		if r.Method != http.MethodPost {
			data := struct {
				Header             Header
				User               model.User
				FileCategories     interface{}
				Files              interface{}
				EspecialidadeFiles interface{}
			}{getHeader(r, "Dashboard | User Type 1"), user, fileCategories, files, especialidadeFiles}
			t, _ := template.ParseFiles("views/type_1_dashboard.tmpl", "views/showcomponent.tmpl", "views/templates/especialidade_files_template.tmpl", "views/templates/main_template.tmpl")
			t.ExecuteTemplate(w, "dashboard", data)

			return
		}
		uploadFiles(w, r)
	}

}

// UploadFile uploads a file to the server
func uploadFiles(w http.ResponseWriter, r *http.Request) {

	fileCategories := model.GetUserFileCategoryList()
	categoryArray := fileCategories.([]interface{})
	sUserID := library.GetCookie(r, "user_id")
	userID, _ := strconv.Atoi(sUserID)

	for i := 1; i <= len(categoryArray); i++ {
		category := categoryArray[i-1].(model.UserCategory)
		index := strconv.Itoa(category.ID)
		fileCategory := "field" + index
		file, handle, err := r.FormFile(fileCategory)

		if err != nil {
			log.Print(err.Error())
			return
		}

		filename := sUserID + "-" + strconv.Itoa(category.ID) + "-" + handle.Filename
		destination := "./public/uploads/" + filename
		err = uploadPdfFile(file, handle, destination)

		if err != nil {
			log.Print(err.Error())

			t, err := template.ParseFiles("views/upload.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			t.Execute(w, err.Error())
			return
		}

		model.UserFileAdd(userID, category.ID, filename, category.ExpiryDate)

		model.AddUserLog(userID, "Uploaded file type: "+fileCategory)
	}

	model.UpdateUserStatus(userID, 4)
	library.PutCookie(w, "status", "4")
	log.Print("User (", userID, "): File uploaded successfully")

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func uploadImageFile(file multipart.File, handle *multipart.FileHeader, destination string) error {
	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg", "image/jpg", "image/png":
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(destination, data, 0666)

		if err != nil {
			return err
		}
		return nil

	default:
		return errors.New("File type not supported")
	}
}

func uploadPdfFile(file multipart.File, handle *multipart.FileHeader, destination string) error {
	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "application/pdf":
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(destination, data, 0666)

		if err != nil {
			return err
		}
		return nil

	default:
		return errors.New("File type not supported")
	}
}
