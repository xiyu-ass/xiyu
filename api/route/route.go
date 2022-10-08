package route

import (
	"myproject3/api/controller"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	group := r.Group("/api/")
	{
		users := group.Group("/user")
		{
			users.POST("/login", controller.UserLogin)
			users.GET("/index", controller.UserIndex)
			users.POST("/publish", controller.UserFeedback)
			users.GET("/test", controller.ListUser)
		}
	}

}
