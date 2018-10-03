package controller

import(
	"net/http"
    "html/template"
    "strconv"
    "log"

    "../middleware"
    "../library"
    "../model"
)

func UserPermissions(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    user_id, _ := strconv.Atoi(ids[0])

    message := ""
    has_error := false

    if r.Method == "POST" {
		r.ParseForm()

		permissions := []interface{}{}

		permissions = append(permissions, model.Permission{ID: library.Permission_staff_add, Value: getCheckboxValue(r.FormValue("staff_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_staff_view, Value: getCheckboxValue(r.FormValue("staff_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_staff_edit, Value: getCheckboxValue(r.FormValue("staff_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_staff_delete, Value: getCheckboxValue(r.FormValue("staff_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_user_add, Value: getCheckboxValue(r.FormValue("user_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_user_view, Value: getCheckboxValue(r.FormValue("user_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_user_edit, Value: getCheckboxValue(r.FormValue("user_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_user_delete, Value: getCheckboxValue(r.FormValue("user_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_files_add, Value: getCheckboxValue(r.FormValue("files_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_files_view, Value: getCheckboxValue(r.FormValue("files_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_files_edit, Value: getCheckboxValue(r.FormValue("files_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_files_delete, Value: getCheckboxValue(r.FormValue("files_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_add, Value: getCheckboxValue(r.FormValue("status_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_view, Value: getCheckboxValue(r.FormValue("status_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_edit, Value: getCheckboxValue(r.FormValue("status_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_delete, Value: getCheckboxValue(r.FormValue("status_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_roles_add, Value: getCheckboxValue(r.FormValue("roles_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_roles_view, Value: getCheckboxValue(r.FormValue("roles_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_roles_edit, Value: getCheckboxValue(r.FormValue("roles_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_roles_delete, Value: getCheckboxValue(r.FormValue("roles_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_permissions_add, Value: getCheckboxValue(r.FormValue("permissions_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_permissions_view, Value: getCheckboxValue(r.FormValue("permissions_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_permissions_edit, Value: getCheckboxValue(r.FormValue("permissions_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_permissions_delete, Value: getCheckboxValue(r.FormValue("permissions_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_1_add, Value: getCheckboxValue(r.FormValue("status_1_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_1_view, Value: getCheckboxValue(r.FormValue("status_1_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_1_edit, Value: getCheckboxValue(r.FormValue("status_1_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_1_delete, Value: getCheckboxValue(r.FormValue("status_1_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_2_add, Value: getCheckboxValue(r.FormValue("status_2_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_2_view, Value: getCheckboxValue(r.FormValue("status_2_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_2_edit, Value: getCheckboxValue(r.FormValue("status_2_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_2_delete, Value: getCheckboxValue(r.FormValue("status_2_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_3_add, Value: getCheckboxValue(r.FormValue("status_3_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_3_view, Value: getCheckboxValue(r.FormValue("status_3_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_3_edit, Value: getCheckboxValue(r.FormValue("status_3_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_3_delete, Value: getCheckboxValue(r.FormValue("status_3_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_4_add, Value: getCheckboxValue(r.FormValue("status_4_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_4_view, Value: getCheckboxValue(r.FormValue("status_4_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_4_edit, Value: getCheckboxValue(r.FormValue("status_4_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_4_delete, Value: getCheckboxValue(r.FormValue("status_4_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_5_add, Value: getCheckboxValue(r.FormValue("status_5_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_5_view, Value: getCheckboxValue(r.FormValue("status_5_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_5_edit, Value: getCheckboxValue(r.FormValue("status_5_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_5_delete, Value: getCheckboxValue(r.FormValue("status_5_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_6_add, Value: getCheckboxValue(r.FormValue("status_6_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_6_view, Value: getCheckboxValue(r.FormValue("status_6_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_6_edit, Value: getCheckboxValue(r.FormValue("status_6_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_6_delete, Value: getCheckboxValue(r.FormValue("status_6_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_7_add, Value: getCheckboxValue(r.FormValue("status_7_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_7_view, Value: getCheckboxValue(r.FormValue("status_7_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_7_edit, Value: getCheckboxValue(r.FormValue("status_7_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_7_delete, Value: getCheckboxValue(r.FormValue("status_7_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_8_add, Value: getCheckboxValue(r.FormValue("status_8_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_8_view, Value: getCheckboxValue(r.FormValue("status_8_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_8_edit, Value: getCheckboxValue(r.FormValue("status_8_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_8_delete, Value: getCheckboxValue(r.FormValue("status_8_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_9_add, Value: getCheckboxValue(r.FormValue("status_9_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_9_view, Value: getCheckboxValue(r.FormValue("status_9_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_9_edit, Value: getCheckboxValue(r.FormValue("status_9_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_9_delete, Value: getCheckboxValue(r.FormValue("status_9_delete"))})

		permissions = append(permissions, model.Permission{ID: library.Permission_status_10_add, Value: getCheckboxValue(r.FormValue("status_10_add"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_10_view, Value: getCheckboxValue(r.FormValue("status_10_view"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_10_edit, Value: getCheckboxValue(r.FormValue("status_10_edit"))})
		permissions = append(permissions, model.Permission{ID: library.Permission_status_10_delete, Value: getCheckboxValue(r.FormValue("status_10_delete"))})

		model.UpdateUserPermission(user_id, permissions)

		log_message := "Updated user[Id: " + strconv.Itoa(user_id) +"] permissions"
		model.AddLog(model.GetUserId(r), log_message)
		log.Print(log_message)
		message = "Permissions updated successfully"
	}

	permissionsJSON := model.GetUserPermissionString(user_id)

	data := struct{
                Header Header
                Message string
                Has_error bool
                User_id int
                PermissionsJSON string
            }{getHeader(r, "User Permission"), message, has_error, user_id, permissionsJSON}

    t, _ := template.ParseFiles("views/user_permissions.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "permissions", data)
}