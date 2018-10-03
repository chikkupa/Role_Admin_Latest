package model

import(
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"../config"
)

func SetDetails(name string, value string){
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

   stmt, err := db.Prepare("UPDATE details SET value=? WHERE name=?")

    _, err = stmt.Exec(value, name)

    if err != nil {
        log.Fatal(err)
    }
}

func GetDetails(name string) string {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("SELECT value FROM details WHERE name=?", name)

    var value string

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    if results.Next() {
        err = results.Scan(&value)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    return value;
}