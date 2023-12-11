package main

import (
	"gogrpc_client/pb"
	"gogrpc_client/routes"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()
	c, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer c.Close()
	//reigister service client
	todoServiceClient := pb.NewTodoServiceClient(c)
	routes.NewTodoRoute(r, &todoServiceClient)
	r.Run(":8081")

}
