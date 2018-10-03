package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../middleware"
	"../model"
)

func CreateActivityRecord(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	message := ""
	has_error := false
	var activityRecord model.ActivityRecord

	if r.Method == "POST" {
		r.ParseForm()

		userId, _ := strconv.Atoi(r.FormValue("user_id"))
		activity := r.FormValue("activity")
		date := r.FormValue("date")
		duration, _ := strconv.Atoi(r.FormValue("duration"))
		assistant, _ := strconv.Atoi(r.FormValue("assistant"))
		description := r.FormValue("description")

		model.CreateActivityRecord(userId, activity, date, duration, assistant, description)

		message = "Activity Record added succesfully!"

		log_message := "New Activity Record [user_id: " + r.FormValue("user_id") + "] created"
		model.AddLog(model.GetUserId(r), log_message)
		log.Print(log_message)
	}

	users := model.GetAllUsers()

	data := struct {
		Header         Header
		Message        string
		Has_error      bool
		ActivityRecord model.ActivityRecord
		Users          interface{}
	}{getHeader(r, "Create New Activity Record"), message, has_error, activityRecord, users}

	t, _ := template.ParseFiles("views/activity_record_add.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "create_activity_record", data)
}

func ListActivityRecords(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	date_start := ""
	date_end := ""
	crm := ""
	name := ""
	lastname := ""
	assistant := ""

	if r.Method == "POST" {
		r.ParseForm()

		date_start = r.FormValue("date_start")
		date_end = r.FormValue("date_end")
		crm = r.FormValue("crm")
		name = r.FormValue("name")
		lastname = r.FormValue("lastname")
		assistant = r.FormValue("assistant")
	}

	activityRecords := model.GetActivityRecordList(date_start, date_end, crm, name, lastname, assistant)

	data := struct {
		Header     Header
		List       interface{}
		Date_start string
		Date_end   string
		Crm        string
		Name       string
		Lastname   string
		Assistant  string
	}{getHeader(r, "Activity Record List"), activityRecords, date_start, date_end, crm, name, lastname, assistant}

	t, _ := template.ParseFiles("views/activity_record_list.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "activity_record_list", data)
}

func UpdateActivityRecord(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	user_id, _ := strconv.Atoi(ids[0])

	message := ""
	hasError := false

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		user_id, _ := strconv.Atoi(r.FormValue("user_id"))
		activity := r.FormValue("activity")
		date := r.FormValue("date")
		duration, _ := strconv.Atoi(r.FormValue("duration"))
		assistant, _ := strconv.Atoi(r.FormValue("assistant"))
		description := r.FormValue("description")

		model.UpdateActivityRecord(id, user_id, activity, date, duration, assistant, description)

		message = "User Activity Record succesfully!"

		log_message := "Activity Record [Id: " + strconv.Itoa(id) + "] updated successfully"
		model.AddLog(model.GetUserId(r), log_message)
		log.Print(log_message)
	}

	activityRecord := model.GetActivityRecordDetails(user_id)
	users := model.GetAllUsers()

	data := struct {
		Header         Header
		Message        string
		Has_error      bool
		ActivityRecord model.ActivityRecord
		Users          interface{}
	}{getHeader(r, "Update Activity Record"), message, hasError, activityRecord, users}

	t, _ := template.ParseFiles("views/activity_record_edit.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "activity_record_edit", data)
}

// DeleteActivityRecord Delete Activity Record
func DeleteActivityRecord(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	model.DeleteActivityRecord(id)

	logMessage := "Delete Activity Record [Id: " + strconv.Itoa(id) + "] deleted successfully"
	model.AddLog(model.GetUserId(r), logMessage)
	log.Print(logMessage)

	http.Redirect(w, r, "/list_activity_record", http.StatusSeeOther)
}
