package main

import (
	"log"

	"go-rest-crud/models"
	"go-rest-crud/views"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func main(){ 
	r :=mux.NewRouter()
	r.HandleFunc("/", views.Index).Methods("GET")
	r.HandleFunc("/", views.AddData).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", views.DetailsData).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", views.DeleteData).Methods("DELETE")
	r.HandleFunc("/{id:[0-9]+}", views.UpdateData).Methods("PUT")
    srv := &http.Server{
		Handler: r,
        Addr:         ":8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
		} 
	models.CreateTable()
	log.Println(srv.ListenAndServe())

}