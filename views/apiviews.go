package views

import (
	"encoding/json"
	"go-rest-crud/models"
	"go-rest-crud/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Product struct { 
	Name_Product string `json:"name_product"`
	Price string `json:"price"`
}


func Index(w http.ResponseWriter, r *http.Request) { 
		resp := models.ReadAlldata()
		response.JsonResponse(w,resp, 200)
	}

func AddData(w http.ResponseWriter, r *http.Request ){ 
	name_product :=	r.FormValue("name_product")
	price        :=	r.FormValue("price")
	product := Product{name_product, price}
	json_resp,_ :=json.Marshal(product)
	log.Println("Add Data")	

	models.Insert(product.Name_Product, product.Price)
	w.Header().Add("content-type", "application/json")
	response.JsonResponse(w, string(json_resp), 201)
}

func DeleteData(w http.ResponseWriter, r *http.Request) { 
	vars :=mux.Vars(r)
	del_resp , status := models.DeleteData(vars["id"])
	response.JsonResponse(w, del_resp, status)
} 

func DetailsData(w http.ResponseWriter, r *http.Request) { 
		vars :=mux.Vars(r)
		detalis_resp,status := models.DetailsData(vars["id"])
		response.JsonResponse(w, detalis_resp, status)
}

func UpdateData(w http.ResponseWriter, r *http.Request) { 
	vars :=mux.Vars(r)
	name_product :=	r.FormValue("name_product")
	price        :=	r.FormValue("price")
	product := Product{name_product, price}
	json_resp,_ :=json.Marshal(product)
	_, status := models.UpdateData(name_product, price, vars["id"])
	response.JsonResponse(w, string(json_resp), status)



}