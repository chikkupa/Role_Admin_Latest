package model

import (
	"database/sql"
	"log"

	"../config"
)

// Rating Structure for Rating
type Rating struct {
	ID              int
	ValuationID     int
	ValuationName   string
	UserID          int
	Description     string
	DateValuation   string
	GivenScore      int
	StaffID         int
	Comments        string
	ScorePercentage float32
	Status          int
	CreateDate      string
}

//GetRatingList Get Rating list
func GetRatingList() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, valuation_id, valuation_name, user_id, description, date_valuation, given_score, staff_id, comments, score_percentage, status, create_date FROM view_rating"
	results, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	list := []interface{}{}
	for results.Next() {
		var item Rating
		err = results.Scan(&item.ID, &item.ValuationID, &item.ValuationName, &item.UserID, &item.Description, &item.DateValuation, &item.GivenScore, &item.StaffID, &item.Comments, &item.ScorePercentage, &item.Status, &item.CreateDate)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, item)
	}
	return list
}

//GetRatingUserList Get Rating list of a user
func GetRatingUserList(userID int) interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, valuation_id, valuation_name, user_id, description, date_valuation, given_score, staff_id, comments, score_percentage, status, create_date FROM view_rating WHERE user_id=?"
	results, err := db.Query(sql, userID)
	if err != nil {
		panic(err.Error())
	}

	list := []interface{}{}
	for results.Next() {
		var item Rating
		err = results.Scan(&item.ID, &item.ValuationID, &item.ValuationName, &item.UserID, &item.Description, &item.DateValuation, &item.GivenScore, &item.StaffID, &item.Comments, &item.ScorePercentage, &item.Status, &item.CreateDate)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, item)
	}
	return list
}

//CreateRating Add a new Rating
func CreateRating(valuationID int, userID int, description string, dateValuation string, givenScore int, staffID int, comments string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "INSERT INTO rating (valuation_id, user_id, description, date_valuation, given_score, staff_id, comments) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(valuationID, userID, description, dateValuation, givenScore, staffID, comments)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteRating Delete a Rating
func DeleteRating(id int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "DELETE FROM rating WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateRating Update a new Rating
func UpdateRating(id int, valuationID int, description string, dateValuation string, givenScore int, staffID int, comments string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "UPDATE rating SET valuation_id=?, description=?, date_valuation=?, given_score=?, staff_id=?, comments=?, status=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(valuationID, description, dateValuation, givenScore, staffID, comments, 1, id)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateRatingByUser Update a new Rating by user
func UpdateRatingByUser(id int, valuationID int, description string, dateValuation string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "UPDATE rating SET valuation_id=?, description=?, date_valuation=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(valuationID, description, dateValuation, id)
	if err != nil {
		log.Fatal(err)
	}
}

//GetRatingDetails Get Rating details
func GetRatingDetails(id int) Rating {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, valuation_id, valuation_name, user_id, description, date_valuation, given_score, staff_id, comments FROM view_rating WHERE id=?"
	results, err := db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}
	var item Rating
	for results.Next() {
		err = results.Scan(&item.ID, &item.ValuationID, &item.ValuationName, &item.UserID, &item.Description, &item.DateValuation, &item.GivenScore, &item.StaffID, &item.Comments)
		if err != nil {
			panic(err.Error())
		}
	}
	return item
}
