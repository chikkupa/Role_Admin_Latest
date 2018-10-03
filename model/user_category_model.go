package model

import (
	"database/sql"
	"log"

	"../config"
)

// UserCategory User Categry structure
type UserCategory struct {
	ID          int
	Name        string
	ExpiryDate  string
	Description string
}

// GetUserFileCategoryList To get all user categories
func GetUserFileCategoryList() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, name, expiry_date, description FROM user_file_category WHERE deleted = 0")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	list := []interface{}{}

	for results.Next() {
		var item UserCategory
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item.ID, &item.Name, &item.ExpiryDate, &item.Description)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		list = append(list, item)
	}

	return list
}
