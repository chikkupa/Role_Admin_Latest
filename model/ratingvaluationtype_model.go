package model

import (
	"database/sql"
	"log"

	"../config"
)

// RatingValuationType Structure for RatingValuationType
type RatingValuationType struct {
	ID          int
	Name        string
	Description string
}

//GetRatingValuationTypeList Get RatingValuationType list
func GetRatingValuationTypeList() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, name, description FROM rating_valuation_type"
	results, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	list := []interface{}{}
	for results.Next() {
		var item RatingValuationType
		err = results.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, item)
	}
	return list
}

//CreateRatingValuationType Add a new RatingValuationType
func CreateRatingValuationType(name string, description string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "INSERT INTO rating_valuation_type (name, description) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(name, description)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteRatingValuationType Delete a RatingValuationType
func DeleteRatingValuationType(id int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "DELETE FROM rating_valuation_type WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateRatingValuationType Update a new RatingValuationType
func UpdateRatingValuationType(id int, name string, description string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "UPDATE rating_valuation_type SET name=?, description=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(name, description, id)
	if err != nil {
		log.Fatal(err)
	}
}

//GetRatingValuationTypeDetails Get RatingValuationType details
func GetRatingValuationTypeDetails(id int) RatingValuationType {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, name, description FROM rating_valuation_type WHERE id=?"
	results, err := db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}
	var item RatingValuationType
	for results.Next() {
		err = results.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			panic(err.Error())
		}
	}
	return item
}
