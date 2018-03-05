package handlers

import (
	"net/http"
	"encoding/json"
	job "github.com/managef/api/models/job"
	pb "github.com/managef/models/rpc"
	"github.com/gorilla/mux"
	"github.com/managef/models/log"
)

func GetJob(w http.ResponseWriter, r *http.Request) {
	log.Info("Requested Get Job")
	params := mux.Vars(r)
	response, err := job.GetJob(pb.JobRequest{Id: params["id"], Number: 4, Name: "Hello Worker"})
	if err!= nil{
		log.Errorf("could not get job: %v", err)
		return
	}
	log.Info("Job Returned")
	json.NewEncoder(w).Encode(response)
}