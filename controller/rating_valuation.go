package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../middleware"
	"../model"
)

// CreateRatingValuation Controller to create valuation
func CreateRatingValuation(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	message := ""
	hasError := false
	valuationTypes := model.GetRatingValuationTypeList()

	var valuation model.RatingValuation

	if r.Method == "POST" {
		r.ParseForm()

		name := r.FormValue("name")
		typeID, _ := strconv.Atoi(r.FormValue("type_id"))
		minimumScore, _ := strconv.Atoi(r.FormValue("minimum_score"))
		maximumScore, _ := strconv.Atoi(r.FormValue("maximum_score"))
		importance := r.FormValue("importance")
		userID := model.GetUserId(r)

		model.CreateRatingValuation(name, typeID, minimumScore, maximumScore, importance, userID)

		message = "Valuation added succesfully!"

		logMessage := "New Valuation [Name: " + name + "] created"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	r.Method = "GET"

	data := struct {
		Header         Header
		Message        string
		HasError       bool
		ValuationTypes interface{}
		Valuation      model.RatingValuation
	}{getHeader(r, "Create New Rating Valuation"), message, hasError, valuationTypes, valuation}

	t, _ := template.ParseFiles("views/rating_valuation_add.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "create_rating_valuation", data)
}

// ListRatingValuation Controller to list rating valuation
func ListRatingValuation(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	valuationList := model.GetRatingValuationList()

	data := struct {
		Header Header
		List   interface{}
	}{getHeader(r, "Rating Valuation List"), valuationList}

	t, _ := template.ParseFiles("views/rating_valuation_list.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "rating_valuation_list", data)
}

// UpdateRatingValuation Controller to update rating valuation
func UpdateRatingValuation(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	message := ""
	hasError := false
	valuationTypes := model.GetRatingValuationTypeList()

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		typeID, _ := strconv.Atoi(r.FormValue("type_id"))
		minimumScore, _ := strconv.Atoi(r.FormValue("minimum_score"))
		maximumScore, _ := strconv.Atoi(r.FormValue("maximum_score"))
		importance := r.FormValue("importance")

		model.UpdateRatingValuation(id, name, typeID, minimumScore, maximumScore, importance)

		message = "Updated Rating Valuatioon succesfully!"

		logMessage := "Rating [Rating id: " + strconv.Itoa(id) + "] updated successfully"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	valuation := model.GetRatingValuationDetails(id)

	data := struct {
		Header         Header
		Message        string
		HasError       bool
		ValuationTypes interface{}
		Valuation      model.RatingValuation
	}{getHeader(r, "Update Rating Valuation Details"), message, hasError, valuationTypes, valuation}

	t, _ := template.ParseFiles("views/rating_valuation_edit.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "edit_rating_valuation", data)
}

// DeleteRatingValuation Controller for deleting the rating valuation
func DeleteRatingValuation(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	model.DeleteRatingValuation(id)

	logMessage := "Rating Valuation [Id: " + strconv.Itoa(id) + "] deleted successfully"
	model.AddLog(model.GetUserId(r), logMessage)
	log.Print(logMessage)

	http.Redirect(w, r, "/list_rating_valuation", http.StatusSeeOther)
}
