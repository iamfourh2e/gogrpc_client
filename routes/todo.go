package routes

import (
	"gogrpc_client/controller"
	"gogrpc_client/pb"

	"github.com/gin-gonic/gin"
)

func NewTodoRoute(r *gin.Engine, client *pb.TodoServiceClient) {
	c := controller.NewTodoController(client)
	g := r.Group("/todo")
	g.GET("/", c.GetTodoFN)
}
