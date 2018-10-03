package model

import (
	"database/sql"
	"log"

	"../config"
)

// RatingValuation Structure for RatingValuation
type RatingValuation struct {
	ID           int
	Name         string
	TypeID       int
	TypeName     string
	MinimumScore int
	MaximumScore int
	Importance   string
	AddedBy      int
	CreateDate   string
}

//GetRatingValuationList Get RatingValuation list
func GetRatingValuationList() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, name, type_id, type_name, minimum_score, maximum_score, importance, added_by_id, create_date FROM view_rating_valuation"
	results, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	list := []interface{}{}
	for results.Next() {
		var item RatingValuation
		err = results.Scan(&item.ID, &item.Name, &item.TypeID, &item.TypeName, &item.MinimumScore, &item.MaximumScore, &item.Importance, &item.AddedBy, &item.CreateDate)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, item)
	}
	return list
}

//CreateRatingValuation Add a new RatingValuation
func CreateRatingValuation(name string, typeID int, minimumScore int, maximumScore int, importance string, addedBy int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "INSERT INTO rating_valuation (name, type_id, minimum_score, maximum_score, importance, added_by) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(name, typeID, minimumScore, maximumScore, importance, addedBy)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteRatingValuation Delete a RatingValuation
func DeleteRatingValuation(id int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "DELETE FROM rating_valuation WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateRatingValuation Update a new RatingValuation
func UpdateRatingValuation(id int, name string, typeID int, minimumScore int, maximumScore int, importance string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "UPDATE rating_valuation SET name=?, type_id=?, minimum_score=?, maximum_score=?, importance=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	_, err = stmt.Exec(name, typeID, minimumScore, maximumScore, importance, id)
	if err != nil {
		log.Fatal(err)
	}
}

//GetRatingValuationDetails Get RatingValuation details
func GetRatingValuationDetails(id int) RatingValuation {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, name, type_id, minimum_score, maximum_score, importance FROM rating_valuation WHERE id=?"
	results, err := db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}
	var item RatingValuation
	for results.Next() {
		err = results.Scan(&item.ID, &item.Name, &item.TypeID, &item.MinimumScore, &item.MaximumScore, &item.Importance)
		if err != nil {
			panic(err.Error())
		}
	}
	return item
}
