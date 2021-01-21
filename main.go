package main

import (
	"fmt"
	"go-rest-crud/views"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func main(){ 

	r :=mux.NewRouter()
	r.HandleFunc("/", views.Index)

    srv := &http.Server{
		Handler: r,
        Addr:         ":8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
	} 
	fmt.Println(srv.ListenAndServe())
}