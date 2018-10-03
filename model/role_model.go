package model

import(
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"../config"
)

type Role struct{
	ID int 					`json:"id"`
	Role string				`json:"role"`
	Value int 				`json:"value"`
	Description string		`json:"description"`
	Permission string		`json:"permission"`
}

func UpdateRole(id int, description string, permissions interface{}) {

	jsonPermissions := ConverPermissionsToJSON(permissions)

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    stmt, err := db.Prepare("UPDATE user_roles SET description=?, permission=? WHERE id=?")
    // checkErr(err)

    _, err = stmt.Exec(description, jsonPermissions, id)

    if err != nil {
		log.Fatal(err)
	}
}

func GetAllRoles() interface{} {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("SELECT id, role, value, description, permission FROM user_roles")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    roles := []interface{}{}

    for results.Next() {
        var role Role
        // for each row, scan the result into our tag composite object
        err = results.Scan(&role.ID, &role.Role, &role.Value, &role.Description, &role.Permission)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        roles = append(roles, role)
    }
    
    return roles
}

func GetRoleDetails(role_id int) Role {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("SELECT id, role, value, description, permission FROM user_roles WHERE value=?", role_id)

    role := Role{}

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    if results.Next() {
        err = results.Scan(&role.ID, &role.Role, &role.Value, &role.Description, &role.Permission)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    return role;
}