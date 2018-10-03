package model

import(
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"../config"
)

type Log struct{
	ID int
	User_id int
	Username string
	Name string
	Time string
	Actions string
}

func InitializeLog() Log {
	return Log{ID: 0, User_id: 0, Username: "", Time: "", Actions: ""}
}

func AddLog(user_id int, actions string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Print("Error connecting database")
	}
	defer db.Close()

    stmt, err := db.Prepare("INSERT user_log set user_id=?, actions=?")
    // checkErr(err)

    _, err = stmt.Exec(user_id, actions)

    if err != nil {
    	log.Print("Error adding log")
	}
}

func AddUserLog(user_id int, actions string) {
    db, err := sql.Open(config.Mysql, config.Dbconnection)
    if err != nil {
        log.Print("Error connecting database")
    }
    defer db.Close()

    stmt, err := db.Prepare("INSERT user_log set user_id=?, user_type=1, actions=?")
    // checkErr(err)

    _, err = stmt.Exec(user_id, actions)

    if err != nil {
        log.Print("Error adding log")
    }
}

func GetAllLogs() interface{} {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, user_id, username, name, date_time, actions from view_user_log order by id desc")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    logs := []interface{}{}

    for results.Next() {
        var userlog Log
        // for each row, scan the result into our tag composite object
        err = results.Scan(&userlog.ID, &userlog.User_id, &userlog.Username, &userlog.Name, &userlog.Time, &userlog.Actions)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        logs = append(logs, userlog)
    }
    
    return logs
}

func GetLogsForUserType3(user_id int) interface{} {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query(" select id, user_id, username, name, date_time, actions from view_user_log where role = 2 or id=? order by id desc", user_id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    logs := []interface{}{}

    for results.Next() {
        var userlog Log
        // for each row, scan the result into our tag composite object
        err = results.Scan(&userlog.ID, &userlog.User_id, &userlog.Username, &userlog.Name, &userlog.Time, &userlog.Actions)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        logs = append(logs, userlog)
    }
    
    return logs
}

func GetAllLogUsers() interface{}{
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, user_id, username, name, date_time from view_log_user_list")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    logs := []interface{}{}

    for results.Next() {
        var userlog Log
        // for each row, scan the result into our tag composite object
        err = results.Scan(&userlog.ID, &userlog.User_id, &userlog.Username, &userlog.Name, &userlog.Time)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        logs = append(logs, userlog)
    }
    
    return logs
}

func GetUserLogs(user_id int) interface{}{
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, user_id, username, name, date_time, actions from view_user_log where user_id=? order by id desc", user_id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    logs := []interface{}{}

    for results.Next() {
        var userlog Log
        // for each row, scan the result into our tag composite object
        err = results.Scan(&userlog.ID, &userlog.User_id, &userlog.Username, &userlog.Name, &userlog.Time, &userlog.Actions)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        logs = append(logs, userlog)
    }
    
    return logs
}

func GetAllLogUsersForType3(user_id int) interface{}{
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, user_id, username, name, date_time from view_log_user_list where role=2 or user_id=?", user_id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    logs := []interface{}{}

    for results.Next() {
        var userlog Log
        // for each row, scan the result into our tag composite object
        err = results.Scan(&userlog.ID, &userlog.User_id, &userlog.Username, &userlog.Name, &userlog.Time)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        logs = append(logs, userlog)
    }
    
    return logs
}