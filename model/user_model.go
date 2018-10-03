package model

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"../config"
	"../library"
)

// User User struct
type User struct {
	ID                       int
	Username                 string
	Password                 string
	Name                     string
	Midname                  string
	Lastname                 string
	Email                    string
	Dob                      string
	Ethnicity                string
	Status                   int
	Stcivil                  string
	Gender                   string
	Cpf                      string
	Rg                       string
	Ufemissor                string
	Orgemissor               string
	Crm                      string
	Ufcrm                    string
	Especialidademed         string
	crmedica                 string
	Descricrmedica           string
	Tituloespecialista       string
	descritituloespecialista string
	Enderecocasa             string
	Numero                   string
	Complemento              string
	Bairro                   string
	Cidade                   string
	Ufresidencia             string
	Cep                      string
	Celular                  string
	Telresidencial           string
	Emailfeed                string
	Enderecocomercial        string
	Complementocomercial     string
	Ufcomercial              string
	Telconsultorio           string
	Telconsultorioramal      string
	Emailconsultorio         string
	Medicoresponsavelnome    string
	Medicoresponsavelcrmm    string
	Tipohonoravelmedico      string
	Tipopagamento            string
	Codigobanco              string
	Agencia                  string
	Conta                    string
	Recolheinssoutrainst     string
	Numeropis                string
	Numeropisautonomo        string
	Carromarca               string
	Carromodelo              string
	Carrocor                 string
	Carroano                 string
	Renavam                  string
	Placa                    string
	Role                     int
	Task                     string
	Message                  string
	Logs                     string
	Permissions              string
}

// UserFile Structure for user file data
type UserFile struct {
	ID           int
	UserID       int
	CategoryID   int
	FileCategory string
	Filename     string
	Status       int
	ExpiryDate   string
	Message      string
}

// GetUserId To get the login user id
func GetUserId(r *http.Request) int {
	sUser_id := library.GetCookie(r, "user_id")
	user_id, _ := strconv.Atoi(sUser_id)
	return user_id
}

func GetUserStatus(r *http.Request) int {
	sStatus := library.GetCookie(r, "status")
	status, _ := strconv.Atoi(sStatus)
	return status
}

func GetUserRole(r *http.Request) int {
	sRole := library.GetCookie(r, "role")
	role, _ := strconv.Atoi(sRole)
	return role
}

func SetUserPermission(w http.ResponseWriter, permissions string) {
	library.PutCookie(w, "permissions", permissions)
}

func GetUserPermissions(r *http.Request) string {
	permissions := library.GetCookie(r, "permissions")
	return permissions
}

func Initialize() User {
	return User{ID: 0, Name: "", Lastname: "", Username: "", Gender: "", Dob: "", Email: "", Password: "", Role: 0, Status: 0, Message: ""}
}

func FilesInitialize() UserFile {
	return UserFile{ID: 0, UserID: 0, FileCategory: "", Filename: "", Status: 0, Message: ""}
}

func Register(username string, name string, email string, password string, gender string) {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT user SET username=?, name=?, email=?, password=?, gender=?")
	// checkErr(err)

	res, err := stmt.Exec(username, name, email, password, gender)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	log.Print("New user added [Id: ", id, " Username: ", username, " Name: ", name, ", Email: ", email, " Gender: ", gender, "]")
}

func Insert(username string, password string, name string, midname string, lastname string, dob string, email string, stcivil string, gender string, cpf string, rg string, ufemissor string, orgemissor string, crm string, ufcrm string, especialidademed string, enderecocasa string, numero string, complemento string, bairro string, cidade string, ufresidencia string, cep string, celular string, telresidencial string, emailfeed string, enderecocomercial string, complementocomercial string, ufcomercial string, telconsultorio string, telconsultorioramal string, emailconsultorio string, medicoresponsavelnome string, medicoresponsavelcrmm string, tipohonoravelmedico string, tipopagamento string, codigobanco string, agencia string, conta string, recolheinssoutrainst string, numeropis string, numeropisautonomo string, carromarca string, carromodelo string, carrocor string, carroano string, renavam string, placa string) int {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(username, password, user_type) VALUES (?,?,?)")

	res, err := stmt.Exec(username, password, 2)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	stmt, err = db.Prepare("INSERT INTO user_details(user_id, name, midname, lastname, dob, email, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	_, err = stmt.Exec(id, name, midname, lastname, dob, email, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa)

	if err != nil {
		log.Fatal(err)
	}

	UpdateUserPermissionWithRole(int(id), library.Role_user_type_1)

	return int(id)
}

func EditUser(user_id int, username string, password string, name string, midname string, lastname string, dob string, email string, status string, stcivil string, gender string, cpf string, rg string, ufemissor string, orgemissor string, crm string, ufcrm string, especialidademed string, enderecocasa string, numero string, complemento string, bairro string, cidade string, ufresidencia string, cep string, celular string, telresidencial string, emailfeed string, enderecocomercial string, complementocomercial string, ufcomercial string, telconsultorio string, telconsultorioramal string, emailconsultorio string, medicoresponsavelnome string, medicoresponsavelcrmm string, tipohonoravelmedico string, tipopagamento string, codigobanco string, agencia string, conta string, recolheinssoutrainst string, numeropis string, numeropisautonomo string, carromarca string, carromodelo string, carrocor string, carroano string, renavam string, placa string) {

	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user set username=?, password=?, user_type=?, status=? WHERE id=?")

	_, err = stmt.Exec(username, password, 2, status, user_id)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("UPDATE user_details SET name=?, midname=?, lastname=?, dob=?, email=?, stcivil=?, gender=?, cpf=?, rg=?, ufemissor=?, orgemissor=?, crm=?, ufcrm=?, especialidademed=?, enderecocasa=?, numero=?, complemento=?, bairro=?, cidade=?, ufresidencia=?, cep=?, celular=?, telresidencial=?, emailfeed=?, enderecocomercial=?, complementocomercial=?, ufcomercial=?, telconsultorio=?, telconsultorioramal=?, emailconsultorio=?, medicoresponsavelnome=?, medicoresponsavelcrmm=?, tipohonoravelmedico=?, tipopagamento=?, codigobanco=?, agencia=?, conta=?, recolheinssoutrainst=?, numeropis=?, numeropisautonomo=?, carromarca=?, carromodelo=?, carrocor=?, carroano=?, renavam=?, placa=? WHERE user_id=?")

	_, err = stmt.Exec(name, midname, lastname, dob, email, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa, user_id)

	if err != nil {
		log.Fatal(err)
	}
}

func IsValidUser(username string, password string) (User, error) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, name, username, email, password, role, status, permissions from view_users where username=? and password=? and status!=0", username, password)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	user := User{ID: 0, Name: "", Username: "", Email: "", Password: "", Role: 0, Status: 0}

	if results.Next() {
		err = results.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.Status, &user.Permissions)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		return user, nil
	}

	results, err = db.Query("select id, name, username, password, role, status, permissions from view_staffs where username=? and password=? and status!=0", username, password)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		err = results.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Role, &user.Status, &user.Permissions)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		return user, nil
	}

	return user, errors.New("INVALID_USERNAME_OR_PASSWORD")
}

func IsUsernameAvailable(username string) bool {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id from user where username=? ", username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		return false
	}

	return true
}

func IsEmailAvailable(email string) bool {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id from view_users where email=? ", email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		return false
	}

	return true
}

func IsCPFAvailable(cpf string) bool {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id from view_users where cpf=? ", cpf)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		return false
	}

	return true
}

func GetAllUsers() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, name, lastname, username, gender, email, password, role, status from view_users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	users := []interface{}{}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.Role, &user.Status)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, user)
	}

	return users
}

// GetAllStaffs List all staffs
func GetAllStaffs() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, name, lastname, username, gender, email, password, role, status from view_staffs")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	users := []interface{}{}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.Role, &user.Status)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, user)
	}

	return users
}

func GetUsersWithStatus(status string) interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, name, lastname, username, gender, email, password, role, status from user where status in (?)", status)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	users := []interface{}{}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.Role, &user.Status)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, user)
	}

	return users
}

func GetUsersListForUserType2() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, name, lastname, username, gender, email, password, role, status from view_users where role=2") // status in (2, 4)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	users := []interface{}{}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.Role, &user.Status)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, user)
	}

	return users
}

func GetUsersListForUserType3() interface{} {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id, name, lastname, username, gender, email, password, role, status from view_users where role=2 and status in (5, 8)") // status in (2, 4)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	users := []interface{}{}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.Role, &user.Status)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, user)
	}

	return users
}

func GetUserDetails(id int) User {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, username, password, name, midname, lastname, dob, email, status, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa, role FROM view_users WHERE id=?", id)

	n := User{}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		var id, status, role int
		var username, password, name, midname, lastname, dob, email, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa string

		err = results.Scan(&id, &username, &password, &name, &midname, &lastname, &dob, &email, &status, &stcivil, &gender, &cpf, &rg, &ufemissor, &orgemissor, &crm, &ufcrm, &especialidademed, &enderecocasa, &numero, &complemento, &bairro, &cidade, &ufresidencia, &cep, &celular, &telresidencial, &emailfeed, &enderecocomercial, &complementocomercial, &ufcomercial, &telconsultorio, &telconsultorioramal, &emailconsultorio, &medicoresponsavelnome, &medicoresponsavelcrmm, &tipohonoravelmedico, &tipopagamento, &codigobanco, &agencia, &conta, &recolheinssoutrainst, &numeropis, &numeropisautonomo, &carromarca, &carromodelo, &carrocor, &carroano, &renavam, &placa, &role)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		n.ID = id
		n.Username = username
		n.Password = password
		n.Name = name
		n.Midname = midname
		n.Lastname = lastname
		n.Dob = dob
		n.Email = email
		n.Status = status
		n.Stcivil = stcivil
		n.Gender = gender
		n.Cpf = cpf
		n.Rg = rg
		n.Ufemissor = ufemissor
		n.Orgemissor = orgemissor
		n.Crm = crm
		n.Ufcrm = ufcrm
		n.Especialidademed = especialidademed
		n.Enderecocasa = enderecocasa
		n.Numero = numero
		n.Complemento = complemento
		n.Bairro = bairro
		n.Cidade = cidade
		n.Ufresidencia = ufresidencia
		n.Cep = cep
		n.Celular = celular
		n.Telresidencial = telresidencial
		n.Emailfeed = emailfeed
		n.Enderecocomercial = enderecocomercial
		n.Complementocomercial = complementocomercial
		n.Ufcomercial = ufcomercial
		n.Telconsultorio = telconsultorio
		n.Telconsultorioramal = telconsultorioramal
		n.Emailconsultorio = emailconsultorio
		n.Medicoresponsavelnome = medicoresponsavelnome
		n.Medicoresponsavelcrmm = medicoresponsavelcrmm
		n.Tipohonoravelmedico = tipohonoravelmedico
		n.Tipopagamento = tipopagamento
		n.Codigobanco = codigobanco
		n.Agencia = agencia
		n.Conta = conta
		n.Recolheinssoutrainst = recolheinssoutrainst
		n.Numeropis = numeropis
		n.Numeropisautonomo = numeropisautonomo
		n.Carromarca = carromarca
		n.Carromodelo = carromodelo
		n.Carrocor = carrocor
		n.Carroano = carroano
		n.Renavam = renavam
		n.Placa = placa
		n.Role = role
	}

	return n
}

func UpdateVarifyString(username string, verify_string string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update user set verify_string=? where username=?")
	// checkErr(err)

	_, err = stmt.Exec(verify_string, username)

	if err != nil {
		log.Fatal(err)
	}
}

func VerifyEmail(key string) bool {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id from user where verify_string=? ", key)
	if err != nil {
		return false // proper error handling instead of panic in your app
	}

	if results.Next() {
		stmt, err := db.Prepare("update user set status=2, verify_string='' where verify_string=?")
		// checkErr(err)

		_, err = stmt.Exec(key)

		if err != nil {
			log.Fatal(err)
		}
		return true
	}

	return false
}

func UpdateUserStatus(user_id int, status int) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update user set status=? where id=?")
	// checkErr(err)

	_, err = stmt.Exec(status, user_id)

	if err != nil {
		log.Fatal(err)
	}
}

// UserFileAdd Add a user uploaded file details to database
func UserFileAdd(userID int, fileCategory int, filename string, expiryDate string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM user_files where user_id=? and file_category=?")
	_, err = stmt.Exec(userID, fileCategory)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT user_files SET user_id=?, file_category=?, filename=?, expiry_date=?, message=''")
	_, err = stmt.Exec(userID, fileCategory, filename, expiryDate)

	if err != nil {
		log.Fatal(err)
	}
}

func UserFileUpdate(user_id int, file_category string, filename string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE user_files SET filename=?, message='', status=0 where user_id=? and file_category=?")
	_, err = stmt.Exec(filename, user_id, file_category)

	if err != nil {
		log.Fatal(err)
	}
}

// GetUserFiles Get the list of all user files
func GetUserFiles(userID int) interface{} {

	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, user_id, file_category, filename, status, expiry_date, message FROM view_user_files WHERE user_id=?", userID)
	if err != nil {
		log.Print("Query Execution Error: user_model : GetUserFiles") // proper error handling instead of panic in your app
	}

	files := []interface{}{}

	for results.Next() {
		var file UserFile
		// for each row, scan the result into our tag composite object
		err = results.Scan(&file.ID, &file.UserID, &file.FileCategory, &file.Filename, &file.Status, &file.ExpiryDate, &file.Message)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		files = append(files, file)
	}

	return files
}

// GetUserFileDetails To get user file details
func GetUserFileDetails(id int) UserFile {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, user_id, category_id, file_category, filename, status, expiry_date, message FROM view_user_files WHERE id=?", id)

	file := FilesInitialize()

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		err = results.Scan(&file.ID, &file.UserID, &file.CategoryID, &file.FileCategory, &file.Filename, &file.Status, &file.ExpiryDate, &file.Message)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return file
}

// ChangeStatusIfUploaded Change status if the files uploaded
func ChangeStatusIfUploaded(userID int) bool {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("select id from user_files where status=2 or status = 0 and user_id=?", userID)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if results.Next() {
		return false
	}

	UpdateUserStatus(userID, 5)
	return true
}

// GetAllUserCount Get user count
func GetAllUserCount() int {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT count(id) as count FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	count := 0

	if results.Next() {
		err = results.Scan(&count)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		return count
	}

	return count
}

// ChangeUserFileStatus Change user file status
func ChangeUserFileStatus(fileID int, message string, status int, expiryDate string) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update user_files set status=?, message=?, expiry_date=? where id=?")
	// checkErr(err)

	_, err = stmt.Exec(status, message, expiryDate, fileID)

	if err != nil {
		log.Fatal(err)
	}
}
