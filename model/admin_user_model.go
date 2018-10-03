package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"../config"
)

type Admin_User struct {
	ID       int
	Username string
	Password string
	Name     string
	Lastname string
	Age      int
	Gender   string
	Usertype int
	Status   int
	Task     string
	Role     int
}

func InitializeAdminUser() Admin_User {
	return Admin_User{ID: 0, Username: "", Password: "", Name: "", Lastname: "", Age: 0, Gender: "", Usertype: 0, Status: 0, Task: "", Role: 0}
}

func AdminUserRegister(username string, password string, name string, lastname string, age int, gender string, user_type int, status int, task string, role int) {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(username, password, user_type, role, status) VALUES (?,?,?,?,?)")

	res, err := stmt.Exec(username, password, 1, role, status)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	log.Print("Last Insert Id: ", id)

	stmt, err = db.Prepare("INSERT INTO admin_user_details (user_id, name, lastname, age, gender, task) VALUES (?, ?, ?, ?, ?, ?)")
	// checkErr(err)

	_, err = stmt.Exec(id, name, lastname, age, gender, task)

	if err != nil {
		log.Fatal(err)
	}

	UpdateUserPermissionWithRole(int(id), role)
}

func GetAllAdminUsers() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, username, name, lastname, age, gender, status, task, role from view_staffs")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	users := []interface{}{}

	for results.Next() {
		var user Admin_User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Username, &user.Name, &user.Lastname, &user.Age, &user.Gender, &user.Status, &user.Task, &user.Role)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, user)
	}

	return users
}

func GetAdminUserDetails(id int) Admin_User {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, username, password, name, lastname, age, gender, status, task, role from view_staffs where id=?", id)

	user := Admin_User{}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		err = results.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Lastname, &user.Age, &user.Gender, &user.Status, &user.Task, &user.Role)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return user
}

func UpdateAdminUser(id int, username string, password string, name string, lastname string, age int, gender string, user_type int, status int, task string, role int) {

	user := GetAdminUserDetails(id)

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user SET username=?, password=?, role=?, status=? where id=?")

	_, err = stmt.Exec(username, password, role, status, id)

	if err != nil {
		log.Print("[UpdateAdminUser] Error executing query")
	}

	stmt, err = db.Prepare("UPDATE admin_user_details SET name =?, lastname=?, age=?, gender=?, task=? where user_id=?")
	// checkErr(err)

	_, err = stmt.Exec(name, lastname, age, gender, task, id)

	if err != nil {
		log.Print("[UpdateAdminUser] Error executing query")
	}

	if user.Role != role {
		UpdateUserPermissionWithRole(id, role)
	}
	// id, err := res.LastInsertId()
}

func DeleteAdminUser(id int) {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE from user where id=?")
	// checkErr(err)

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("DELETE from admin_user_details where user_id=?")
	// checkErr(err)

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	// id, err := res.LastInsertId()
}
