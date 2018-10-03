package controller

import(
    "log"
    "net/http"
    "html/template"
    "strconv"

    "../library"
    "../model"
    "../config"
    "../middleware"
)

func FilesShowEspecialidadeFile(w http.ResponseWriter, r *http.Request){
    middleware.RedirectIfNotLoggedIn(w, r)

    ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    sId := ids[0]
    id, _ := strconv.Atoi(sId)

    resp_message := ""

    if r.Method == "POST" {
        r.ParseForm()

        id, _ := strconv.Atoi(r.FormValue("id"))
        expiry_date := r.FormValue("expiry_date")

        model.EspecialidadeFileExpiryDateUpdate(id, expiry_date)

        resp_message = "Expiry date updated successfully!"
        log_message := "Especialidade file expiry date changed"
        model.AddLog(model.GetUserId(r), log_message)
        log.Print(log_message)
    }

    file_details := model.GetEspecialidadeFileDetails(id)

    filename := config.Especialidade_upload_location + file_details.Filename

    role, _ := strconv.Atoi(library.GetCookie(r, "role"))

    data := struct{
                Header Header
                Filename string
                File model.EspecialidadeFile
                Role int
                Message string
            }{getHeader(r, "View Pdf"), filename, file_details, role, resp_message }

    t, _ := template.ParseFiles("views/especialidade_pdf_view.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "pdf_view", data)
}