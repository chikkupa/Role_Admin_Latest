package controller

import (
	"log"
	"net/http"

	"../library"
	"../model"
)

type Header struct {
	Title       string
	Name        string
	Role        int
	Last_login  string
	Permissions model.Permissions
}

func getHeader(r *http.Request, title string) Header {
	name := library.GetCookie(r, "name")
	role := model.GetUserRole(r)
	last_login := library.GetCookie(r, "last_login")
	permissionString := model.GetUserPermissionString(model.GetUserId(r))
	permissions := model.GetAllPermissions(permissionString)

	log.Print("Permission Header => ", permissions)

	return Header{Title: title, Name: name, Role: role, Last_login: last_login, Permissions: permissions}
}
