package main

import (
	jobReq "github.com/managef/api/requests/job"
	serviceReq "github.com/managef/api/requests/service"
	"github.com/gorilla/mux"
	"github.com/managef/models/log"
	"github.com/golang/glog"
	"net/http"
)


func main() {
	defer glog.Flush()
	router := mux.NewRouter()
	router.HandleFunc("/", serviceReq.GetApi).Methods("GET")
	router.HandleFunc("/job/{id}", jobReq.GetJob).Methods("GET")

	log.Error(http.ListenAndServe(":8080", router))
}