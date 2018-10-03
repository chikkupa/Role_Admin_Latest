package config

var Host = "http://localhost"
var Port = "9090"
var Host_location = Host + ":" + Port + "/"
var Upload_location = Host_location + "uploads/"
var Especialidade_upload_location = Upload_location + "especialidade/"
var Reports_upload_location = Upload_location + "reports/"

/* Database config */
var Db_name = "go_login_app"
var Db_user = "root"
var Db_password = "123"
var Mysql = "mysql"

var Dbconnection = Db_user + ":" + Db_password + "@tcp(127.0.0.1:3306)/" + Db_name;
var BaseUrl = Host + ":" + Port

/* Mail server setup */
var Mail_smtp_addr = "smtp.gmail.com:587"
var Mail_smtp_host = "smtp.gmail.com"
var Mail_username = "chikku.pa@gmail.com"
var Mail_Password = "9387324"
