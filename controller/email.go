package controller

import(
	"log"
	"strconv"

	"../model"
	"../library"
)

func SendStatusChangedEmail(user model.User, previous_status int, next_status int){
	if previous_status == next_status {
		return;
	}
	
	status := []string {"Pending Email Verification", "Active", "Pendent", "OK", "Uploaded", "Verified", "Approved", "Meet", "Accepted", "Denied", "Deactivated"}

	email_id := user.Email

	subject := "Your Status Changed"
	message := "Your status has been changed to '" + status[next_status] + "' from '" + status[previous_status] + "'"

	if next_status == 3 {
		message = message + "<br><br> Please upload PDF files"
	}

	library.SendMail(email_id, subject, message)

	user_log := "User [user_id: " + strconv.Itoa(user.ID) + "] status has been changed to '" + status[next_status] + "' from '" + status[previous_status] + "'"

	model.AddLog(user.ID, user_log)

	log.Print("User [user_id: ", user.ID, "] status has been changed to '" + status[next_status] + "' from '" + status[previous_status] + "'")
}