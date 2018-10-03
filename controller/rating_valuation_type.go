package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../middleware"
	"../model"
)

// CreateRatingValuationType Controller to create valuation
func CreateRatingValuationType(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	message := ""
	hasError := false

	var valuationType model.RatingValuationType

	if r.Method == "POST" {
		r.ParseForm()

		name := r.FormValue("name")
		description := r.FormValue("description")

		model.CreateRatingValuationType(name, description)

		message = "Valuation Type added succesfully!"

		logMessage := "New Valuation Type [Name: " + name + "] created"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	r.Method = "GET"

	data := struct {
		Header        Header
		Message       string
		HasError      bool
		ValuationType model.RatingValuationType
	}{getHeader(r, "Create New Rating Valuation"), message, hasError, valuationType}

	t, _ := template.ParseFiles("views/rating_valuation_type_add.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "create_rating_valuation_type", data)
}

// ListRatingValuationType Controller to list rating valuation
func ListRatingValuationType(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	valuationTypeList := model.GetRatingValuationTypeList()

	data := struct {
		Header Header
		List   interface{}
	}{getHeader(r, "Valuation Type List"), valuationTypeList}

	t, _ := template.ParseFiles("views/rating_valuation_type_list.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "rating_valuation_type_list", data)
}

// UpdateRatingValuationType Controller to update rating valuation
func UpdateRatingValuationType(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	message := ""
	hasError := false

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")

		model.UpdateRatingValuationType(id, name, description)

		message = "Updated Valuatioon Type succesfully!"

		logMessage := "Valuation Type [Id: " + strconv.Itoa(id) + "] updated successfully"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	valuationType := model.GetRatingValuationTypeDetails(id)

	data := struct {
		Header        Header
		Message       string
		HasError      bool
		ValuationType model.RatingValuationType
	}{getHeader(r, "Update Rating Valuation Details"), message, hasError, valuationType}

	t, _ := template.ParseFiles("views/rating_valuation_type_edit.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "edit_rating_valuation_type", data)
}

// DeleteRatingValuationType Controller for deleting the rating valuation
func DeleteRatingValuationType(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	model.DeleteRatingValuationType(id)

	logMessage := "Valuation Type [Id: " + strconv.Itoa(id) + "] deleted successfully"
	model.AddLog(model.GetUserId(r), logMessage)
	log.Print(logMessage)

	http.Redirect(w, r, "/rating_valuation_type_list", http.StatusSeeOther)
}
