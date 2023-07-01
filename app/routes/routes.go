package routes

import (
	"university-dean-and-student/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	public := router.Group("/api")

	public.POST("/login", controllers.Login)
	public.POST("/register", controllers.Register)

	private := router.Group("/api", TokenAuthMiddleware())
	private.POST("/session", controllers.Session)
	private.GET("/sessions", controllers.Sessions)
	private.GET("/dean_list", controllers.DeanList)

	router.Run("localhost:8080")
}
