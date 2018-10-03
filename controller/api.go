package controller

import(
	"net/http"
	"encoding/json"

	"../model"
)

type Response struct{
	Result bool	`json:"result"`
}

func API_IsCPFAvailable(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	cpf := r.FormValue("cpf")

	var resp Response

	resp.Result = model.IsCPFAvailable(cpf)

	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}