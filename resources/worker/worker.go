package worker

import (
	"google.golang.org/grpc"
	"github.com/managef/models/log"
)

const (
	worker_address     = "worker:50051"
)


func Conn()(*grpc.ClientConn, error){
	conn, err := grpc.Dial(worker_address, grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect: %v", err)
		return nil, err
	}
	log.Info("Api connected to Worker")
	return conn,nil
}