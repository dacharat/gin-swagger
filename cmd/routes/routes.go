package routes

import (
	"github.com/dacharat/gin-swagger/cmd/handlers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Init init router
func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userV1 := r.Group("/user")
	{
		userHandler := handlers.NewUserHandler()
		userV1.POST("", AuthMiddleware(), userHandler.CreateUserHandler)

		userV1.GET("/:user_id", userHandler.GetUserHandler)
		userV1.PUT("/:user_id", userHandler.UpdateUserHandler)
	}

	return r
}
