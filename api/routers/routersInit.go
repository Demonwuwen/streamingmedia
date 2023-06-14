package routers

import (
	"Demonwuwen/streamingmedia/api/handlers"
	"github.com/gin-gonic/gin"
)

func RoutersInit(r *gin.Engine) {
	userLogRouters := r.Group("/user")
	{
		userLogRouters.POST("/login", handlers.Login)
		userLogRouters.POST("/")
	}
}
