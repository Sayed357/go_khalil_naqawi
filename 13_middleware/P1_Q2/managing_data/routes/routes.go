package routes

import (
	"managing_data/controllers"
	"managing_data/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.RegisterUser)
		v1.POST("/login", controllers.LoginUser)
	}

	protected := r.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware())
	protected.Use(middlewares.RateLimiterMiddleware())
	{
		protected.GET("/posts", controllers.GetPost)
		protected.GET("/posts/:id", controllers.GetPost)
		protected.POST("/posts", controllers.CreatePost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)
	}
}
