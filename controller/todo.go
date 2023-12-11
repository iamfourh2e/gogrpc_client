package controller

import (
	"gogrpc_client/pb"

	"github.com/gin-gonic/gin"
)

type TodoControllerImpl struct {
	Client pb.TodoServiceClient
}

func NewTodoController(client *pb.TodoServiceClient) *TodoControllerImpl {
	return &TodoControllerImpl{
		Client: *client,
	}
}

func (t *TodoControllerImpl) GetTodoFN(c *gin.Context) {
	res, err := t.Client.Create(c, &pb.CreateRequest{
		Title:       "test",
		Description: "thkeam",
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, res)
	}
}
