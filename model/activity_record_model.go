package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"../config"
)

type ActivityRecord struct {
	ID             int
	User_id        int
	UserName       string
	Name           string
	Lastname       string
	Crm            string
	Activity       string
	Date           string
	Duration       int
	Assistant      int
	Assistant_name string
	Assistant_crm  string
	Description    string
}

func CreateActivityRecord(user_id int, activity string, date string, duration int, assistant int, description string) {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO activity_record (user_id, activity, date, duration, assistant, description) VALUES (?, ?, ?, ?, ?, ?)")
	// checkErr(err)

	_, err = stmt.Exec(user_id, activity, date, duration, assistant, description)

	if err != nil {
		log.Fatal(err)
	}

}

func GetActivityRecordList(date_start string, date_end string, crm string, name string, lastname string, assistant string) interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT id, user_id, name, lastname, crm, activity, date, duration, assistant, assistant_name, assistant_crm, description FROM view_activity_records WHERE id > 0"

	if date_start != "" {
		query += " AND date >= '" + date_start + "'"
	}

	if date_end != "" {
		query += " AND date <= '" + date_end + "'"
	}

	if crm != "" {
		query += " AND crm like '%" + crm + "%'"
	}

	if name != "" {
		query += " AND name like '%" + name + "%'"
	}

	if lastname != "" {
		query += " AND lastname like '%" + lastname + "%'"
	}

	if assistant != "" {
		query += " AND assistant_name like '%" + assistant + "%'"
	}

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	list := []interface{}{}

	for results.Next() {
		var item ActivityRecord
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item.ID, &item.User_id, &item.Name, &item.Lastname, &item.Crm, &item.Activity, &item.Date, &item.Duration, &item.Assistant, &item.Assistant_name, &item.Assistant_crm, &item.Description)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		list = append(list, item)
	}

	return list
}

func GetActivityRecordDetails(id int) ActivityRecord {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, user_id, name, lastname, crm, activity, date, duration, assistant, assistant_name, assistant_crm, description FROM view_activity_records WHERE id=?", id)

	item := ActivityRecord{}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		err = results.Scan(&item.ID, &item.User_id, &item.Name, &item.Lastname, &item.Crm, &item.Activity, &item.Date, &item.Duration, &item.Assistant, &item.Assistant_name, &item.Assistant_crm, &item.Description)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return item
}

func UpdateActivityRecord(id int, user_id int, activity string, date string, duration int, assistant int, description string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE activity_record SET user_id=?, activity=?, date=?, duration=?, assistant=?, description=? WHERE id=?")

	_, err = stmt.Exec(user_id, activity, date, duration, assistant, description, id)

	if err != nil {
		panic(err.Error())
	}
}

func DeleteActivityRecord(id int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM activity_record WHERE id=?")
	// checkErr(err)

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}
}
