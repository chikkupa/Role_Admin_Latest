package controller

import(
    "log"
    "net/http"
    "html/template"
    "strconv"
    "time"
    "math/rand"

    "../library"
    "../model"
    "../config"
    "../middleware"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:];

    if path != "" && path != "login" {
        t, _ := template.ParseFiles("views/404.tmpl")
        t.Execute(w, nil)
    } else if r.Method == "GET" {
        error_message := ""
        t, _ := template.ParseFiles("views/login.tmpl", "views/templates/login_template.tmpl")
        t.ExecuteTemplate(w, "login", error_message)
    } else {
        r.ParseForm()
        name := r.PostFormValue("username")
        password := library.ConvertToSha1(r.PostFormValue("password"))

        user, err := model.IsValidUser(name, password)

        if err != nil {
            error_message := "Invalid username or password"
            t, _ := template.ParseFiles("views/login.tmpl", "views/templates/login_template.tmpl")
            t.ExecuteTemplate(w, "login", error_message)
            log.Print(name, ": Login Failed => Invalid credentials")
            return
        }
        t := time.Now()
        library.PutCookie(w, "user_id", strconv.Itoa(user.ID))
        library.PutCookie(w, "username", user.Username)
        library.PutCookie(w, "name", user.Name)
        library.PutCookie(w, "email", user.Email)
        library.PutCookie(w, "role", strconv.Itoa(user.Role))
        library.PutCookie(w, "password", user.Password)
        library.PutCookie(w, "status", strconv.Itoa(user.Status))
        library.PutCookie(w, "last_login", t.Format("03:04 PM / 02-01-2006"))
        model.SetUserPermission(w, user.Permissions)

        middleware.SetLoggedIn(w)

        // writeSuccess("Lgin Success!", w)

        log.Print(name, ": User Login Success")
        model.AddLog(user.ID, "Successfully logged in")

        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
    }
}

func Logout(w http.ResponseWriter, r *http.Request){
    middleware.Logout(w, r)
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        user := model.User{Name: "", Username: "", Email: "", Gender: "", Message : ""}
        especialidadeList := model.GetEspecialidadeList()
        terms_and_conditions := model.GetDetails("TERMS_AND_CONDITIONS")

        data := struct{
                Header Header
                User model.User
                EspecialidadeList interface{}
                Terms string
            }{getHeader(r, "Update Admin Details"), user, especialidadeList, terms_and_conditions}

        t, _ := template.ParseFiles("views/templates/login_template.tmpl", "views/register.tmpl")
        t.ExecuteTemplate(w, "registration", data)
    } else {
        r.ParseMultipartForm(0)

        username := r.FormValue("username")
        password := library.ConvertToSha1(r.FormValue("password"))
        name := r.FormValue("name")
        midname := r.FormValue("midname")
        lastname := r.FormValue("lastname")
        dob := r.FormValue("dob")
        email := r.FormValue("email")
        //status := r.FormValue("status")
        stcivil := r.FormValue("stcivil")
        gender := r.FormValue("gender")
        cpf := r.FormValue("cpf")
        rg := r.FormValue("rg")
        ufemissor := r.FormValue("ufemissor")
        orgemissor := r.FormValue("orgemissor")
        crm := r.FormValue("crm")
        ufcrm := r.FormValue("ufcrm")
        especialidademed := r.FormValue("especialidademed")
        enderecocasa := r.FormValue("enderecocasa")
        numero := r.FormValue("numero")
        complemento := r.FormValue("complemento")
        bairro := r.FormValue("bairro")
        cidade := r.FormValue("cidade")
        ufresidencia := r.FormValue("ufresidencia")
        cep := r.FormValue("cep")
        celular := r.FormValue("celular")
        telresidencial := r.FormValue("telresidencial")
        emailfeed := r.FormValue("emailfeed")
        enderecocomercial := r.FormValue("enderecocomercial")
        complementocomercial := r.FormValue("complementocomercial")
        ufcomercial := r.FormValue("ufcomercial")
        telconsultorio := r.FormValue("telconsultorio")
        telconsultorioramal := r.FormValue("telconsultorioramal")
        emailconsultorio := r.FormValue("emailconsultorio")
        medicoresponsavelnome := r.FormValue("medicoresponsavelnome")
        medicoresponsavelcrmm := r.FormValue("medicoresponsavelcrmm")
        tipohonoravelmedico := r.FormValue("tipohonoravelmedico")
        tipopagamento := r.FormValue("tipopagamento")
        codigobanco := r.FormValue("codigobanco")
        agencia := r.FormValue("agencia")
        conta := r.FormValue("conta")
        recolheinssoutrainst := r.FormValue("recolheinssoutrainst")
        numeropis := r.FormValue("numeropis")
        numeropisautonomo := r.FormValue("numeropisautonomo")
        carromarca := r.FormValue("carromarca")
        carromodelo := r.FormValue("carromodelo")
        carrocor := r.FormValue("carrocor")
        carroano := r.FormValue("carroano")
        renavam := r.FormValue("renavam")
        placa := r.FormValue("placa")

        log.Print("Username: ", username)

        user := model.User{Name: name, Username: username, Email: email, Gender: "", Message : ""}

        errors := false

        if(!model.IsUsernameAvailable(username)){
            user.Message = "Username not available"
            log.Print("Username not available!")
            errors = true
        }
        if(!errors && !model.IsEmailAvailable(email)){
            user.Message = "Email not available"
            log.Print("Email not available!")
            errors = true
        }

        if(errors){
            t, _ := template.ParseFiles("views/templates/login_template.tmpl", "views/register.tmpl")
            t.ExecuteTemplate(w, "registration", user)
            return
        }

        user_id := model.Insert(username, password, name, midname, lastname, dob, email, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa)

        especialidadeList := model.GetEspecialidadeList()

        especialidadeArray := especialidadeList.([]interface{})

        for i := 0; i < len(especialidadeArray); i++ {
            espItem := especialidadeArray[i].(model.Especialidade)
            esp_id := strconv.Itoa(espItem.ID)

            especialidadeCheck := getEspecialidadeCheckboxValue(r.FormValue("especialidade_" + esp_id))

            if especialidadeCheck > 0{
                file, handle, err := r.FormFile("especialidade_file_" + esp_id)

                if err != nil {
                    log.Print(err.Error())
                    return
                }
                filename := strconv.Itoa(user_id) + "-" + esp_id + "-" +  handle.Filename
                destination := "./public/uploads/especialidade/" + filename
                err = uploadPdfFile(file, handle, destination)

                model.EspecialidadeFileAdd(user_id, espItem.ID, filename, espItem.Expiry_date)
            }
        }

        random_string := strconv.Itoa(rand.Intn(999)) + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(9999))
        random_string = library.ConvertToSha1(random_string)

        verify_email_link := config.BaseUrl + "/verify_email?key=" + random_string

        model.UpdateVarifyString(username, random_string)

        library.SendMail(email, "Confirm Email", "Please click the link below to verify email<br><br><a href='" + verify_email_link + "''>Verify<a>")

        data := struct{
                    Message_1 string; 
                    Message_2 string;
                }{"Your registration has been completed successfully!", "Please check your email for verification"}

        t, _ := template.ParseFiles("views/templates/success_template.tmpl", "views/success.tmpl")
        t.ExecuteTemplate(w, "success", data)
    }
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
    middleware.RedirectIfNotLoggedIn(w, r)

    ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    user_id, _ := strconv.Atoi(ids[0])

    message := ""
    has_error := false

    user_details := model.GetUserDetails(user_id)

    if r.Method == "POST" {
        r.ParseMultipartForm(0)

        username := r.FormValue("username")
        password := library.ConvertToSha1(r.FormValue("password"))
        name := r.FormValue("name")
        midname := r.FormValue("midname")
        lastname := r.FormValue("lastname")
        dob := r.FormValue("dob")
        email := r.FormValue("email")
        status := r.FormValue("status")
        stcivil := r.FormValue("stcivil")
        gender := r.FormValue("gender")
        cpf := r.FormValue("cpf")
        rg := r.FormValue("rg")
        ufemissor := r.FormValue("ufemissor")
        orgemissor := r.FormValue("orgemissor")
        crm := r.FormValue("crm")
        ufcrm := r.FormValue("ufcrm")
        especialidademed := r.FormValue("especialidademed")
        enderecocasa := r.FormValue("enderecocasa")
        numero := r.FormValue("numero")
        complemento := r.FormValue("complemento")
        bairro := r.FormValue("bairro")
        cidade := r.FormValue("cidade")
        ufresidencia := r.FormValue("ufresidencia")
        cep := r.FormValue("cep")
        celular := r.FormValue("celular")
        telresidencial := r.FormValue("telresidencial")
        emailfeed := r.FormValue("emailfeed")
        enderecocomercial := r.FormValue("enderecocomercial")
        complementocomercial := r.FormValue("complementocomercial")
        ufcomercial := r.FormValue("ufcomercial")
        telconsultorio := r.FormValue("telconsultorio")
        telconsultorioramal := r.FormValue("telconsultorioramal")
        emailconsultorio := r.FormValue("emailconsultorio")
        medicoresponsavelnome := r.FormValue("medicoresponsavelnome")
        medicoresponsavelcrmm := r.FormValue("medicoresponsavelcrmm")
        tipohonoravelmedico := r.FormValue("tipohonoravelmedico")
        tipopagamento := r.FormValue("tipopagamento")
        codigobanco := r.FormValue("codigobanco")
        agencia := r.FormValue("agencia")
        conta := r.FormValue("conta")
        recolheinssoutrainst := r.FormValue("recolheinssoutrainst")
        numeropis := r.FormValue("numeropis")
        numeropisautonomo := r.FormValue("numeropisautonomo")
        carromarca := r.FormValue("carromarca")
        carromodelo := r.FormValue("carromodelo")
        carrocor := r.FormValue("carrocor")
        carroano := r.FormValue("carroano")
        renavam := r.FormValue("renavam")
        placa := r.FormValue("placa")

        if(user_details.Username != username && !model.IsUsernameAvailable(username)){
            message = "Username not available"
            log.Print("User Edit: Username not available!")
            has_error = true
        }
        if(user_details.Email != email && !has_error && !model.IsEmailAvailable(email)){
            message = "Email not available"
            log.Print("User Edit: Email not available!")
            has_error = true
        }

        if(user_details.Cpf != cpf && !has_error && !model.IsCPFAvailable(cpf)){
            message = "CPF not available"
            log.Print("User Edit: CPF not available!")
            has_error = true
        }

        if(!has_error){
            message = "User details updated successfully"

            model.EditUser(user_id, username, password, name, midname, lastname, dob, email, status, stcivil, gender, cpf, rg, ufemissor, orgemissor, crm, ufcrm, especialidademed, enderecocasa, numero, complemento, bairro, cidade, ufresidencia, cep, celular, telresidencial, emailfeed, enderecocomercial, complementocomercial, ufcomercial, telconsultorio, telconsultorioramal, emailconsultorio, medicoresponsavelnome, medicoresponsavelcrmm, tipohonoravelmedico, tipopagamento, codigobanco, agencia, conta, recolheinssoutrainst, numeropis, numeropisautonomo, carromarca, carromodelo, carrocor, carroano, renavam, placa)

            user_details = model.GetUserDetails(user_id)
        }
    }

    data := struct{
                Header Header
                Message string
                Has_error bool
                User model.User
            }{getHeader(r, "Update Especialidade Details"), message, has_error, user_details}

    t, _ := template.ParseFiles("views/user_edit.tmpl", "views/showcomponent.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "user_edit", data)
        
}

func Verifyemail(w http.ResponseWriter, r *http.Request){
    keys, err := r.URL.Query()["key"]
    
    if !err || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]

    result := model.VerifyEmail(key)

    if(result){
         data := struct{
                    Message_1 string; 
                    Message_2 template.HTML;
                }{"Your email verified successfully!", template.HTML("Please login <a href='/'>here <a>")}

        t, _ := template.ParseFiles("views/templates/success_template.tmpl", "views/success.tmpl")
        t.ExecuteTemplate(w, "success", data)
    }else{
        data := struct{
                    Message_1 string; 
                    Message_2 string;
                }{"Error verifying your email!", "Please contact the administrator for more details"}

        t, _ := template.ParseFiles("views/templates/success_template.tmpl", "views/error.tmpl")
        t.ExecuteTemplate(w, "error", data)
    }
}

func ViewUserDetails(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)
    
    role, _ := strconv.Atoi(library.GetCookie(r, "role"))

    if role != library.Role_admin {
        // http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
    }

    if r.Method == "GET" {
        ids, err := r.URL.Query()["id"];

        if !err || len(ids[0]) < 1 {
            log.Println("Url Param 'id' is missing")
            return
        }

        id,_ := strconv.Atoi(ids[0])

        user := model.GetUserDetails(id)
        message:= library.GetCookie(r, "message")
        user.Message = message;
        library.PutCookie(w, "message", "")
        files := model.GetUserFiles(id)
        especialidade_files := model.GetEspecialidadeFiles(id)

        data := struct{
                    Header Header
                    User model.User
                    Files interface{}
                    Especialidade_files interface{}
                }{getHeader(r, "View User | Administrator"), user, files, especialidade_files }

        t, _ := template.ParseFiles("views/view_user.tmpl", "views/showcomponent.tmpl", "views/templates/especialidade_files_template.tmpl", "views/templates/main_template.tmpl")
        t.ExecuteTemplate(w, "view_user", data)
    }else{
        user_id, _ := strconv.Atoi(r.PostFormValue("id"))
        status, _ := strconv.Atoi(r.PostFormValue("status"))

        user := model.GetUserDetails(user_id)

        message := "";
        if user.Status != status {
            model.UpdateUserStatus(user_id, status)
            message = "User status updated successfully!"

            SendStatusChangedEmail(user, user.Status, status)

        }else{
            message = "No change"
        }

        library.PutCookie(w, "message", message)
        http.Redirect(w, r, "/view_user?id=" + r.PostFormValue("id"), http.StatusSeeOther)
    }
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
    w.WriteHeader(status)
    w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(message))
}

func getEspecialidadeCheckboxValue(value string) int {
    if(value == ""){
        return 0
    }
    int_value, _ := strconv.Atoi(value)
    return int_value
}
