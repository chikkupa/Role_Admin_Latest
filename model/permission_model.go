package model

import(
	"encoding/json"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"../library"
	"../config"
)

type Permission struct{
	ID int      `json:id`
    Value int	`json:"value"`
}

type Permissions struct{
	Permission_staff_add, Permission_staff_view, Permission_staff_edit, Permission_staff_delete int
	Permission_user_add, Permission_user_view, Permission_user_edit, Permission_user_delete int
	Permission_files_add, Permission_files_view, Permission_files_edit, Permission_files_delete int
	Permission_status_add, Permission_status_view, Permission_status_edit, Permission_status_delete int
	Permission_roles_add, Permission_roles_view, Permission_roles_edit, Permission_roles_delete int
	Permission_permissions_add, Permission_permissions_view, Permission_permissions_edit, Permission_permissions_delete int
	Permission_status_1_add, Permission_status_1_view, Permission_status_1_edit, Permission_status_1_delete int
	Permission_status_2_add, Permission_status_2_view, Permission_status_2_edit, Permission_status_2_delete int
	Permission_status_3_add, Permission_status_3_view, Permission_status_3_edit, Permission_status_3_delete int
	Permission_status_4_add, Permission_status_4_view, Permission_status_4_edit, Permission_status_4_delete int
	Permission_status_5_add, Permission_status_5_view, Permission_status_5_edit, Permission_status_5_delete int
	Permission_status_6_add, Permission_status_6_view, Permission_status_6_edit, Permission_status_6_delete int
	Permission_status_7_add, Permission_status_7_view, Permission_status_7_edit, Permission_status_7_delete int
	Permission_status_8_add, Permission_status_8_view, Permission_status_8_edit, Permission_status_8_delete int
	Permission_status_9_add, Permission_status_9_view, Permission_status_9_edit, Permission_status_9_delete int
	Permission_status_10_add, Permission_status_10_view, Permission_status_10_edit, Permission_status_10_delete int
}

func ConverPermissionsToJSON(permissions interface{}) string {
	jsonValue, _ := json.Marshal(permissions)

	return string(jsonValue)
}

func GetAllPermissions(permissionsJSON string) Permissions {
	permissionsArray := []Permission{}
    json.Unmarshal([]byte(string(permissionsJSON)), &permissionsArray)

    var permissions Permissions

    for i := 0; i < len(permissionsArray); i++ {
    	switch permissionsArray[i].ID{
    		case library.Permission_staff_add: permissions.Permission_staff_add = permissionsArray[i].Value
			case library.Permission_staff_view: permissions.Permission_staff_view = permissionsArray[i].Value
			case library.Permission_staff_edit: permissions.Permission_staff_edit = permissionsArray[i].Value
			case library.Permission_staff_delete: permissions.Permission_staff_delete = permissionsArray[i].Value

			case library.Permission_user_add: permissions.Permission_user_add = permissionsArray[i].Value
			case library.Permission_user_view: permissions.Permission_user_view = permissionsArray[i].Value
			case library.Permission_user_edit: permissions.Permission_user_edit = permissionsArray[i].Value
			case library.Permission_user_delete: permissions.Permission_user_delete = permissionsArray[i].Value

			case library.Permission_files_add: permissions.Permission_files_add = permissionsArray[i].Value
			case library.Permission_files_view: permissions.Permission_files_view = permissionsArray[i].Value
			case library.Permission_files_edit: permissions.Permission_files_edit = permissionsArray[i].Value
			case library.Permission_files_delete: permissions.Permission_files_delete = permissionsArray[i].Value

			case library.Permission_status_add: permissions.Permission_status_add = permissionsArray[i].Value
			case library.Permission_status_view: permissions.Permission_status_view = permissionsArray[i].Value
			case library.Permission_status_edit: permissions.Permission_status_edit = permissionsArray[i].Value
			case library.Permission_status_delete: permissions.Permission_status_delete = permissionsArray[i].Value

			case library.Permission_roles_add: permissions.Permission_roles_add = permissionsArray[i].Value
			case library.Permission_roles_view: permissions.Permission_roles_view = permissionsArray[i].Value
			case library.Permission_roles_edit: permissions.Permission_roles_edit = permissionsArray[i].Value
			case library.Permission_roles_delete: permissions.Permission_roles_delete = permissionsArray[i].Value

			case library.Permission_permissions_add: permissions.Permission_permissions_add = permissionsArray[i].Value
			case library.Permission_permissions_view: permissions.Permission_permissions_view = permissionsArray[i].Value
			case library.Permission_permissions_edit: permissions.Permission_permissions_edit = permissionsArray[i].Value
			case library.Permission_permissions_delete: permissions.Permission_permissions_delete = permissionsArray[i].Value

			case library.Permission_status_1_add: permissions.Permission_status_1_add = permissionsArray[i].Value
			case library.Permission_status_1_view: permissions.Permission_status_1_view = permissionsArray[i].Value
			case library.Permission_status_1_edit: permissions.Permission_status_1_edit = permissionsArray[i].Value
			case library.Permission_status_1_delete: permissions.Permission_status_1_delete = permissionsArray[i].Value

			case library.Permission_status_2_add: permissions.Permission_status_2_add = permissionsArray[i].Value
			case library.Permission_status_2_view: permissions.Permission_status_2_view = permissionsArray[i].Value
			case library.Permission_status_2_edit: permissions.Permission_status_2_edit = permissionsArray[i].Value
			case library.Permission_status_2_delete: permissions.Permission_status_2_delete = permissionsArray[i].Value

			case library.Permission_status_3_add: permissions.Permission_status_3_add = permissionsArray[i].Value
			case library.Permission_status_3_view: permissions.Permission_status_3_view = permissionsArray[i].Value
			case library.Permission_status_3_edit: permissions.Permission_status_3_edit = permissionsArray[i].Value
			case library.Permission_status_3_delete: permissions.Permission_status_3_delete = permissionsArray[i].Value

			case library.Permission_status_4_add: permissions.Permission_status_4_add = permissionsArray[i].Value
			case library.Permission_status_4_view: permissions.Permission_status_4_view = permissionsArray[i].Value
			case library.Permission_status_4_edit: permissions.Permission_status_4_edit = permissionsArray[i].Value
			case library.Permission_status_4_delete: permissions.Permission_status_4_delete = permissionsArray[i].Value

			case library.Permission_status_5_add: permissions.Permission_status_5_add = permissionsArray[i].Value
			case library.Permission_status_5_view: permissions.Permission_status_5_view = permissionsArray[i].Value
			case library.Permission_status_5_edit: permissions.Permission_status_5_edit = permissionsArray[i].Value
			case library.Permission_status_5_delete: permissions.Permission_status_5_delete = permissionsArray[i].Value

			case library.Permission_status_6_add: permissions.Permission_status_6_add = permissionsArray[i].Value
			case library.Permission_status_6_view: permissions.Permission_status_6_view = permissionsArray[i].Value
			case library.Permission_status_6_edit: permissions.Permission_status_6_edit = permissionsArray[i].Value
			case library.Permission_status_6_delete: permissions.Permission_status_6_delete = permissionsArray[i].Value

			case library.Permission_status_7_add: permissions.Permission_status_7_add = permissionsArray[i].Value
			case library.Permission_status_7_view: permissions.Permission_status_7_view = permissionsArray[i].Value
			case library.Permission_status_7_edit: permissions.Permission_status_7_edit = permissionsArray[i].Value
			case library.Permission_status_7_delete: permissions.Permission_status_7_delete = permissionsArray[i].Value

			case library.Permission_status_8_add: permissions.Permission_status_8_add = permissionsArray[i].Value
			case library.Permission_status_8_view: permissions.Permission_status_8_view = permissionsArray[i].Value
			case library.Permission_status_8_edit: permissions.Permission_status_8_edit = permissionsArray[i].Value
			case library.Permission_status_8_delete: permissions.Permission_status_8_delete = permissionsArray[i].Value

			case library.Permission_status_9_add: permissions.Permission_status_9_add = permissionsArray[i].Value
			case library.Permission_status_9_view: permissions.Permission_status_9_view = permissionsArray[i].Value
			case library.Permission_status_9_edit: permissions.Permission_status_9_edit = permissionsArray[i].Value
			case library.Permission_status_9_delete: permissions.Permission_status_9_delete = permissionsArray[i].Value

			case library.Permission_status_10_add: permissions.Permission_status_10_add = permissionsArray[i].Value
			case library.Permission_status_10_view: permissions.Permission_status_10_view = permissionsArray[i].Value
			case library.Permission_status_10_edit: permissions.Permission_status_10_edit = permissionsArray[i].Value
			case library.Permission_status_10_delete: permissions.Permission_status_10_delete = permissionsArray[i].Value
    	}
    }

    return permissions
}

func UpdateUserPermissionWithRole(user_id int, role_id int){
	role := GetRoleDetails(role_id)

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    stmt, err := db.Prepare("UPDATE user SET permissions=? WHERE id=?")
    // checkErr(err)

    _, err = stmt.Exec(role.Permission, user_id)

    if err != nil {
		log.Fatal(err)
	}
}


func GetUserPermissionString(user_id int) string{
	db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("SELECT permissions FROM user WHERE id=?", user_id)

    var permissionsJSON string

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    if results.Next() {

        err = results.Scan(&permissionsJSON)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    return permissionsJSON;
}

func UpdateUserPermission(user_id int, permissions interface{}) {

	jsonPermissions := ConverPermissionsToJSON(permissions)

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    stmt, err := db.Prepare("UPDATE user SET permissions=? WHERE id=?")
    // checkErr(err)

    _, err = stmt.Exec(jsonPermissions, user_id)

    if err != nil {
		log.Fatal(err)
	}
}