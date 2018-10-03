package controller

import (
	"net/http"
    "html/template"
    "strconv"
    "log"
    "strings"

    "github.com/360EntSecGroup-Skylar/excelize"
    "github.com/jung-kurt/gofpdf"

    "../middleware"
    "../model"
    "../library"
)

func ReportDocumentsExpired(w http.ResponseWriter, r *http.Request){
	middleware.RedirectIfNotLoggedIn(w, r)

	var documentsExpired interface{}

	output_type := "0"
	date_from := ""
	date_to := ""

	if r.Method == "GET" {
    	documentsExpired = model.GetDocumentsExpiredReport("", "")
    } else {
    	r.ParseForm()

    	output_type = r.FormValue("output_type")
    	date_from = r.FormValue("date_from")
        date_to = r.FormValue("date_to")

        documentsExpired = model.GetDocumentsExpiredReport(date_from, date_to)

        if output_type == "1"{
        	ReportDocumentsExpiredTOExcel(w, r, documentsExpired)
        } else if output_type == "2" {
            ReportDocumentsExpiredTOPDF(w, r, documentsExpired)
        }
    }

    data := struct{
                Header Header
                List interface{}
                Date_from string
                Date_to string
            }{getHeader(r, "Document Expired Report"), documentsExpired, date_from, date_to }

    t, _ := template.ParseFiles("views/report_document_expired.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "report_document_expired", data)
}

func ReportUser(w http.ResponseWriter, r *http.Request){
	middleware.RedirectIfNotLoggedIn(w, r)

	var users interface{}

	output_type := ""
	date_from := ""
	date_to := ""

	if r.Method == "GET" {
    	users = model.GetUserReport("", "")
    } else {
    	r.ParseForm()

    	output_type = r.FormValue("output_type")
    	date_from = r.FormValue("date_from")
        date_to = r.FormValue("date_to")

        users = model.GetUserReport(date_from, date_to)

        if output_type == "1"{
        	ReportUserTOExcel(w, r, users)
        } else if output_type == "2"{
            ReportUserTOPDF(w, r, users)
        }
    }

    data := struct{
                Header Header
                List interface{}
                Date_from string
                Date_to string
                Replace func(string, string, string) template.HTML
            }{getHeader(r, "User Report"), users, date_from, date_to, func(value string, key string, replace string) template.HTML{ return template.HTML(strings.Replace(value, key, replace, -1)) } }

    t, _ := template.ParseFiles("views/report_user.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "report_user", data)
}

func ReportEspecialidade(w http.ResponseWriter, r *http.Request){
	middleware.RedirectIfNotLoggedIn(w, r)

	var especialidades interface{}

	output_type := ""
	date_from := ""
	date_to := ""

	if r.Method == "GET" {
    	especialidades = model.GetEspecialidadeReport("", "")
    } else {
    	r.ParseForm()

    	output_type = r.FormValue("output_type")
    	date_from = r.FormValue("date_from")
        date_to = r.FormValue("date_to")

        especialidades = model.GetEspecialidadeReport(date_from, date_to)

        if output_type == "1"{
        	ReportEspecialidadeTOExcel(w, r, especialidades)
        } else if output_type == "2" {
            log.Print("output_type", 2)
            ReportEspecialidadeTOPDF(w, r, especialidades)
        }
    }

    data := struct{
                Header Header
                List interface{}
                Date_from string
                Date_to string
            }{getHeader(r, "User Report"), especialidades, date_from, date_to }

    t, _ := template.ParseFiles("views/report_especialidade.tmpl", "views/templates/main_template.tmpl")
    t.ExecuteTemplate(w, "report_especialidade", data)
}

func ReportDocumentsExpiredTOExcel(w http.ResponseWriter, r *http.Request, report interface{}){
	xlsx := excelize.NewFile()
    
    table_header_style, err := xlsx.NewStyle(`{"font":{"bold":true,"size":8,"color":"#000000"}}`)
    xlsx.SetCellStyle("Sheet1", "A1", "A1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "B1", "B1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "C1", "C1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "D1", "D1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "E1", "E1", table_header_style)

    xlsx.SetCellValue("Sheet1", "A1", "Especialidade Name")
    xlsx.SetCellValue("Sheet1", "B1", "Username")
    xlsx.SetCellValue("Sheet1", "C1", "Name")
    xlsx.SetCellValue("Sheet1", "D1", "Filename")
    xlsx.SetCellValue("Sheet1", "E1", "Expiry Date")

	reportArray := report.([]interface{})

	for i := 0; i < len(reportArray); i++ {
		row := reportArray[i].(model.ReportEspecialidateUserFile)
		rowIndex := strconv.Itoa(i + 2)

		xlsx.SetCellValue("Sheet1", "A" + rowIndex, row.Especialidade_name)
		xlsx.SetCellValue("Sheet1", "B" + rowIndex, row.Username)
		xlsx.SetCellValue("Sheet1", "C" + rowIndex, row.Name)
		xlsx.SetCellValue("Sheet1", "D" + rowIndex, row.Filename)
		xlsx.SetCellValue("Sheet1", "E" + rowIndex, row.Expiry_date)
	}

	filename := "./public/uploads/reports/documents_expired.xlsx"
	err = xlsx.SaveAs(filename)
    if err != nil {
        log.Print(err)
    }

    library.ForceDownload(w, r, filename, "documents_expired.xlsx")
}

func ReportUserTOExcel(w http.ResponseWriter, r *http.Request, report interface{}){
	statusList := []string{"Pending Email Verification", "Active", "Pendent", "OK", "Uploaded", "Verified", "Approved", "Meet", "Accepted", "Denied", "Deactivated"}
	xlsx := excelize.NewFile()
    
    table_header_style, err := xlsx.NewStyle(`{"font":{"bold":true,"size":8,"color":"#000000"}}`)
    xlsx.SetCellStyle("Sheet1", "A1", "A1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "B1", "B1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "C1", "C1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "D1", "D1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "E1", "E1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "F1", "F1", table_header_style)

    xlsx.SetCellValue("Sheet1", "A1", "Username")
    xlsx.SetCellValue("Sheet1", "B1", "Name")
    xlsx.SetCellValue("Sheet1", "C1", "Lastname")
    xlsx.SetCellValue("Sheet1", "D1", "Create Date")
    xlsx.SetCellValue("Sheet1", "E1", "Status")
    xlsx.SetCellValue("Sheet1", "F1", "Especialidades")

	reportArray := report.([]interface{})

	for i := 0; i < len(reportArray); i++ {
		row := reportArray[i].(model.ReportUser)
		rowIndex := strconv.Itoa(i + 2)

		especialidades := strings.Replace(row.Especialidades, ",", ", ", -1)
		xlsx.SetCellValue("Sheet1", "A" + rowIndex, row.Username)
		xlsx.SetCellValue("Sheet1", "B" + rowIndex, row.Name)
		xlsx.SetCellValue("Sheet1", "C" + rowIndex, row.Lastname)
		xlsx.SetCellValue("Sheet1", "D" + rowIndex, row.Create_date)
		xlsx.SetCellValue("Sheet1", "E" + rowIndex, statusList[row.Status])
		xlsx.SetCellValue("Sheet1", "F" + rowIndex, especialidades)
	}

	filename := "./public/uploads/reports/users.xlsx"
	err = xlsx.SaveAs(filename)
    if err != nil {
        log.Print(err)
    }

    library.ForceDownload(w, r, filename, "users.xlsx")
}

func ReportEspecialidadeTOExcel(w http.ResponseWriter, r *http.Request, report interface{}){
	xlsx := excelize.NewFile()
    
    table_header_style, err := xlsx.NewStyle(`{"font":{"bold":true,"size":8,"color":"#000000"}}`)
    xlsx.SetCellStyle("Sheet1", "A1", "A1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "B1", "B1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "C1", "C1", table_header_style)
    xlsx.SetCellStyle("Sheet1", "D1", "D1", table_header_style)

    xlsx.SetCellValue("Sheet1", "A1", "Id")
    xlsx.SetCellValue("Sheet1", "B1", "Name")
    xlsx.SetCellValue("Sheet1", "C1", "Expiry Date")
    xlsx.SetCellValue("Sheet1", "D1", "Description")

	reportArray := report.([]interface{})

	for i := 0; i < len(reportArray); i++ {
		row := reportArray[i].(model.ReportEspecialidade)
		rowIndex := strconv.Itoa(i + 2)

		xlsx.SetCellValue("Sheet1", "A" + rowIndex, row.ID)
		xlsx.SetCellValue("Sheet1", "B" + rowIndex, row.Name)
		xlsx.SetCellValue("Sheet1", "C" + rowIndex, row.Expiry_date)
		xlsx.SetCellValue("Sheet1", "D" + rowIndex, row.Description)
	}

	filename := "./public/uploads/reports/especialidades.xlsx"
	err = xlsx.SaveAs(filename)
    if err != nil {
        log.Print(err)
    }

    library.ForceDownload(w, r, filename, "especialidades.xlsx")
}

func ReportDocumentsExpiredTOPDF(w http.ResponseWriter, r *http.Request, report interface{}){
    pdf := gofpdf.New("L", "mm", "Letter", "")
    pdf.AddPage()
    pdf.SetFont("Times", "B", 16)
    pdf.Cell(40, 10, "Documents Expired Report")
    pdf.Ln(12)

    pdf.SetFont("Times", "B", 16)
    pdf.SetFillColor(240, 240, 240)
    pdf.CellFormat(60, 7, "Especialidade Name", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Username", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Name", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Expiry Date", "1", 0, "", true, 0, "")
    pdf.CellFormat(80, 7, "Filename", "1", 0, "", true, 0, "")
    pdf.Ln(-1)

    reportArray := report.([]interface{})

    pdf.SetFont("Times", "", 16)
    pdf.SetFillColor(255, 255, 255)
    
    for i := 0; i < len(reportArray); i++ {
        row := reportArray[i].(model.ReportEspecialidateUserFile)

        pdf.CellFormat(60, 7, row.Especialidade_name, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, row.Username, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, row.Name, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, row.Expiry_date, "1", 0, "L", false, 0, "")
        pdf.CellFormat(80, 7, row.Filename, "1", 0, "L", false, 0, "")
        pdf.Ln(-1)
    }

    filename := "./public/uploads/reports/documents_expired.pdf"
    err := pdf.OutputFileAndClose(filename)
    if err != nil {
        log.Print(err)
    }

    library.ForceDownload(w, r, filename, "documents_expired.pdf")
}

func ReportUserTOPDF(w http.ResponseWriter, r *http.Request, report interface{}){
    statusList := []string{"Pending Email Verification", "Active", "Pendent", "OK", "Uploaded", "Verified", "Approved", "Meet", "Accepted", "Denied", "Deactivated"}

    pdf := gofpdf.New("L", "mm", "Letter", "")
    pdf.AddPage()
    pdf.SetFont("Times", "B", 16)
    pdf.Cell(40, 10, "Documents Expired Report")
    pdf.Ln(12)

    pdf.SetFont("Times", "B", 12)
    pdf.SetFillColor(240, 240, 240)
    pdf.CellFormat(40, 7, "Username", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Name", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Lastname", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Create Date", "1", 0, "", true, 0, "")
    pdf.CellFormat(40, 7, "Status", "1", 0, "", true, 0, "")
    pdf.CellFormat(60, 7, "Especialidades", "1", 0, "", true, 0, "")
    pdf.Ln(-1)

    reportArray := report.([]interface{})

    pdf.SetFont("Times", "", 10)
    pdf.SetFillColor(255, 255, 255)
    for i := 0; i < len(reportArray); i++ {
        row := reportArray[i].(model.ReportUser)

        especialidades := strings.Replace(row.Especialidades, ",", ", ", -1)
        pdf.CellFormat(40, 7, row.Username, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, row.Name, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, row.Lastname, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, row.Create_date, "1", 0, "L", false, 0, "")
        pdf.CellFormat(40, 7, statusList[row.Status], "1", 0, "L", false, 0, "")
        pdf.CellFormat(60, 7, especialidades, "1", 0, "L", false, 0, "")
        pdf.Ln(-1)
    }

    filename := "./public/uploads/reports/users.pdf"
    err := pdf.OutputFileAndClose(filename)
    if err != nil {
        log.Print(err)
    }

    library.ForceDownload(w, r, filename, "users.pdf")
}

func ReportEspecialidadeTOPDF(w http.ResponseWriter, r *http.Request, report interface{}){
    pdf := gofpdf.New("L", "mm", "Letter", "")
    pdf.AddPage()
    pdf.SetFont("Times", "B", 16)
    pdf.Cell(40, 10, "Especialidade Report")
    pdf.Ln(12)

    pdf.SetFont("Times", "B", 12)
    pdf.SetFillColor(240, 240, 240)
    pdf.CellFormat(40, 7, "Id", "1", 0, "", true, 0, "")
    pdf.CellFormat(80, 7, "Name", "1", 0, "", true, 0, "")
    pdf.CellFormat(80, 7, "Expiry Date", "1", 0, "", true, 0, "")
    pdf.Ln(-1)    

    reportArray := report.([]interface{})

    pdf.SetFont("Times", "", 16)
    pdf.SetFillColor(255, 255, 255)

    for i := 0; i < len(reportArray); i++ {
        row := reportArray[i].(model.ReportEspecialidade)

        pdf.CellFormat(40, 7, strconv.Itoa(row.ID), "1", 0, "L", false, 0, "")
        pdf.CellFormat(80, 7, row.Name, "1", 0, "L", false, 0, "")
        pdf.CellFormat(80, 7, row.Expiry_date, "1", 0, "L", false, 0, "")
        pdf.Ln(-1)
    }

    filename := "./public/uploads/reports/especialidades.pdf"
    err := pdf.OutputFileAndClose(filename)
    if err != nil {
        log.Print(err)
    }

    library.ForceDownload(w, r, filename, "especialidades.pdf")
}