package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../middleware"
	"../model"
)

// CreateRating Controller to create rating
func CreateRating(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	message := ""
	hasError := false

	var rating model.Rating

	valuationList := model.GetRatingValuationList()

	if r.Method == "POST" {
		r.ParseForm()

		valuationID, _ := strconv.Atoi(r.FormValue("valuation_id"))
		userID, _ := strconv.Atoi(r.FormValue("user_id"))
		description := r.FormValue("description")
		dateValuation := r.FormValue("date_of_valuation")
		givenScore, _ := strconv.Atoi(r.FormValue("given_score"))
		staffID, _ := strconv.Atoi(r.FormValue("staff_id"))
		comments := r.FormValue("comments")

		model.CreateRating(valuationID, userID, description, dateValuation, givenScore, staffID, comments)

		message = "Rating added succesfully!"

		logMessage := "New Rating [User Id: " + strconv.Itoa(userID) + "] created"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	data := struct {
		Header        Header
		Message       string
		HasError      bool
		Rating        model.Rating
		ValuationList interface{}
	}{getHeader(r, "Create New Rating"), message, hasError, rating, valuationList}

	t, _ := template.ParseFiles("views/rating_add.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "create_rating", data)
}

// CreateUserRating Controller to create rating
func CreateUserRating(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	message := ""
	hasError := false

	var rating model.Rating

	valuationList := model.GetRatingValuationList()

	if r.Method == "POST" {
		r.ParseForm()

		valuationID, _ := strconv.Atoi(r.FormValue("valuation_id"))
		userID := model.GetUserId(r)
		description := r.FormValue("description")
		dateValuation := r.FormValue("date_of_valuation")
		givenScore := 0
		staffID := 0
		comments := ""

		model.CreateRating(valuationID, userID, description, dateValuation, givenScore, staffID, comments)

		message = "Rating added succesfully!"

		logMessage := "New Rating [User Id: " + strconv.Itoa(userID) + "] created"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	data := struct {
		Header        Header
		Message       string
		HasError      bool
		Rating        model.Rating
		ValuationList interface{}
	}{getHeader(r, "Create New Rating"), message, hasError, rating, valuationList}

	t, _ := template.ParseFiles("views/rating_user_add.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "create_user_rating", data)
}

// ListRating Controller to list rating
func ListRating(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ratingList := model.GetRatingList()

	data := struct {
		Header Header
		List   interface{}
	}{getHeader(r, "Rating List"), ratingList}

	t, _ := template.ParseFiles("views/rating_list.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "rating_list", data)
}

// ListRatingUser Controller to list user rating
func ListRatingUser(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	userID := model.GetUserId(r)

	ratingList := model.GetRatingUserList(userID)

	data := struct {
		Header Header
		List   interface{}
	}{getHeader(r, "Rating List"), ratingList}

	t, _ := template.ParseFiles("views/rating_user_list.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "rating_user_list", data)
}

// UpdateRating Controller to update rating
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	message := ""
	hasError := false
	valuationList := model.GetRatingValuationList()

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		valuationID, _ := strconv.Atoi(r.FormValue("valuation_id"))
		description := r.FormValue("description")
		dateValuation := r.FormValue("date_of_valuation")
		givenScore, _ := strconv.Atoi(r.FormValue("given_score"))
		staffID := model.GetUserId(r)
		comments := r.FormValue("comments")

		model.UpdateRating(id, valuationID, description, dateValuation, givenScore, staffID, comments)

		message = "User Rating succesfully!"

		logMessage := "Rating [Rating id: " + strconv.Itoa(id) + "] updated successfully"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	rating := model.GetRatingDetails(id)

	data := struct {
		Header        Header
		Message       string
		HasError      bool
		Rating        model.Rating
		ValuationList interface{}
	}{getHeader(r, "Update Rating Details"), message, hasError, rating, valuationList}

	t, _ := template.ParseFiles("views/rating_edit.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "update_rating", data)
}

// UpdateUserRating Controller to update rating by user
func UpdateUserRating(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	message := ""
	hasError := false
	valuationList := model.GetRatingValuationList()

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.FormValue("id"))
		valuationID, _ := strconv.Atoi(r.FormValue("valuation_id"))
		description := r.FormValue("description")
		dateValuation := r.FormValue("date_of_valuation")

		model.UpdateRatingByUser(id, valuationID, description, dateValuation)

		message = "User Rating updated succesfully!"

		logMessage := "Rating [Rating id: " + strconv.Itoa(id) + "] updated successfully"
		model.AddLog(model.GetUserId(r), logMessage)
		log.Print(logMessage)
	}

	rating := model.GetRatingDetails(id)

	data := struct {
		Header        Header
		Message       string
		HasError      bool
		Rating        model.Rating
		ValuationList interface{}
	}{getHeader(r, "Update Rating Details"), message, hasError, rating, valuationList}

	t, _ := template.ParseFiles("views/rating_user_edit.tmpl", "views/templates/main_template.tmpl")
	t.ExecuteTemplate(w, "update_user_rating", data)
}

// DeleteRating Controller for deleting the rating
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	middleware.RedirectIfNotLoggedIn(w, r)

	ids, err := r.URL.Query()["id"]

	if !err || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, _ := strconv.Atoi(ids[0])

	model.DeleteRating(id)

	logMessage := "Rating [Rating id: " + strconv.Itoa(id) + "] deleted successfully"
	model.AddLog(model.GetUserId(r), logMessage)
	log.Print(logMessage)

	http.Redirect(w, r, "/list_especialidade", http.StatusSeeOther)
}
