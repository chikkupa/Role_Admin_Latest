package model

import(
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"../config"
)

type ReportEspecialidateUserFile struct{
	ID int 
	Especialidade_id int
	Especialidade_name string
	User_id int
	Username string
	Name string
	Filename string
	Expiry_date string
	Status int
	Message string
}

type ReportUser struct{
	Username string
	Name string
	Lastname string
	Create_date string
	Status int 
	Especialidades string
}

type ReportEspecialidade struct{
	ID int
	Name string
	Expiry_date string
	Description string
}

func GetDocumentsExpiredReport(date_start string, date_end string) interface{}{
	db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    query := "select id, especialidade_id, especialidade_name, user_id, username, name, filename, expiry_date, status, message from view_especialidade_user_files where expiry_date < CURRENT_DATE"

    if date_start != "" {
    	query += " and expiry_date >= '" + date_start + "'"
    }

    if date_end != "" {
    	query += " and expiry_date <= '" + date_end + "'"
    }

    results, err := db.Query(query)

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    list := []interface{}{}

    for results.Next() {
        var item ReportEspecialidateUserFile
        // for each row, scan the result into our tag composite object
        err = results.Scan(&item.ID, &item.Especialidade_id, &item.Especialidade_name, &item.User_id, &item.Username, &item.Name, &item.Filename, &item.Expiry_date, &item.Status, &item.Message)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        list = append(list, item)
    }
    
    return list
}

func GetUserReport(date_start string, date_end string) interface{}{
	db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    query := "select username, name, lastname, create_date, status, especialidades from view_user_report where user_id > 0"

    if date_start != "" {
    	query += " and create_date >= '" + date_start + "'"
    }

    if date_end != "" {
    	query += " and create_date <= '" + date_end + "'"
    }

    results, err := db.Query(query)

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    list := []interface{}{}

    for results.Next() {
        var item ReportUser
        // for each row, scan the result into our tag composite object
        err = results.Scan(&item.Username, &item.Name, &item.Lastname, &item.Create_date, &item.Status, &item.Especialidades)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        list = append(list, item)
    }
    
    return list
}

func GetEspecialidadeReport(date_start string, date_end string) interface{}{
	db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    query := "select id, name, expiry_date, description from especialidade where id > 0"

    if date_start != "" {
    	query += " and expiry_date >= '" + date_start + "'"
    }

    if date_end != "" {
    	query += " and expiry_date <= '" + date_end + "'"
    }

    results, err := db.Query(query)

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    list := []interface{}{}

    for results.Next() {
        var item ReportEspecialidade
        // for each row, scan the result into our tag composite object
        err = results.Scan(&item.ID, &item.Name, &item.Expiry_date, &item.Description)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        list = append(list, item)
    }
    
    return list
}