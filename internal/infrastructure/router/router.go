package router

import (
	"gin-sample/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(todoHander *handler.TodoHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.POST("", todoHander.Create)
			todos.GET("", todoHander.GetAll)
			todos.GET("/:id", todoHander.GetByID)
			todos.PUT("/:id", todoHander.UpdateStatus)
			todos.DELETE("/:id", todoHander.Delete)

		}
	}

	return router
}
