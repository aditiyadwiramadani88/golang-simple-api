package views

import (
	"go-rest-crud/response"
	"net/http")

func Index(w http.ResponseWriter, r *http.Request) { 
	response.JsonResponse(w, `{"hello":"world"}`)
	
}