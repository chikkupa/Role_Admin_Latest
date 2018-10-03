package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"../config"
)

type Especialidade struct {
	ID          int
	Name        string
	Expiry_date string
	Description string
}

type EspecialidadeFile struct {
	ID               int
	User_id          int
	Especialidade_id string
	Filename         string
	Expiry_date      string
	Status           int
	Message          string
}

func CreateEspecialidade(name string, expiry_date string, description string) {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO especialidade (name, expiry_date, description) values (?, ?, ?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmt.Exec(name, expiry_date, description)

	if err != nil {
		log.Fatal(err)
	}

}

func GetEspecialidadeList() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, name, expiry_date, description FROM especialidade")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	list := []interface{}{}

	for results.Next() {
		var item Especialidade
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item.ID, &item.Name, &item.Expiry_date, &item.Description)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		list = append(list, item)
	}

	return list
}

func GetEspecialidadeDetails(id int) Especialidade {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, name, expiry_date, description FROM especialidade where id=?", id)

	item := Especialidade{}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		err = results.Scan(&item.ID, &item.Name, &item.Expiry_date, &item.Description)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return item
}

func UpdateEspecialidade(id int, name string, expiry_date string, description string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE especialidade SET name=?, expiry_date=?, description=? WHERE id=?")

	_, err = stmt.Exec(name, expiry_date, description, id)

	if err != nil {
		log.Print("[UpdateEspecialidade] Error executing query")
	}
}

func DeleteEspecialidade(id int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM especialidade WHERE id=?")
	// checkErr(err)

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}
}

func EspecialidadeFileAdd(user_id int, especialidade_id int, filename string, expiry_date string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM especialidade_user_files where user_id=? and especialidade_id=?")
	_, err = stmt.Exec(user_id, especialidade_id)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT especialidade_user_files SET user_id=?, especialidade_id=?, filename=?, message='', expiry_date=?")
	_, err = stmt.Exec(user_id, especialidade_id, filename, expiry_date)

	if err != nil {
		log.Fatal(err)
	}
}

func EspecialidadeFileUpdate(user_id int, especialidade_id string, filename string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE especialidade_user_files SET filename=?, message='', status=0 where user_id=? and especialidade_id=?")
	_, err = stmt.Exec(filename, user_id, especialidade_id)

	if err != nil {
		log.Fatal(err)
	}
}

func GetEspecialidadeFiles(user_id int) interface{} {

	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, user_id, especialidade_id, filename, expiry_date, status, message from especialidade_user_files where user_id=?", user_id)
	if err != nil {
		log.Print("Query Execution Error: user_model : GetUserFiles") // proper error handling instead of panic in your app
	}

	files := []interface{}{}

	for results.Next() {
		var file EspecialidadeFile
		// for each row, scan the result into our tag composite object
		err = results.Scan(&file.ID, &file.User_id, &file.Especialidade_id, &file.Filename, &file.Expiry_date, &file.Status, &file.Message)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		files = append(files, file)
	}

	return files
}

func GetEspecialidadeFileDetails(id int) EspecialidadeFile {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, user_id, especialidade_id, filename, expiry_date, status, message from especialidade_user_files where id=?", id)

	var file EspecialidadeFile

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		err = results.Scan(&file.ID, &file.User_id, &file.Especialidade_id, &file.Filename, &file.Expiry_date, &file.Status, &file.Message)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return file
}

func EspecialidadeFileExpiryDateUpdate(id int, expiry_date string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE especialidade_user_files SET expiry_date=? where id=?")
	_, err = stmt.Exec(expiry_date, id)

	if err != nil {
		log.Fatal(err)
	}
}
