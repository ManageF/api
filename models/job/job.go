package job

import(
	worker "github.com/managef/api/resources/worker"
	pb "github.com/managef/models/rpc"
	"github.com/managef/models/log"
	"context"
	"time"
)

func GetJob(request pb.JobRequest)(pb.JobResponse, error){
	conn, err := worker.Conn()
	if err!=nil{
		return pb.JobResponse{}, err
	}
	defer conn.Close()
	c := pb.NewJobClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.GetJob(ctx, &request)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return pb.JobResponse{}, err
	}
	return *response, nil
}

