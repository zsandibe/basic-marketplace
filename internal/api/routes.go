package api

import "github.com/gin-gonic/gin"

func Routes() *gin.Engine {
	router := gin.Default()

	// router.Use(func(c *gin.Context){
	// 	if c.Request.Method == "GET"
	// }

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/")
			users.GET("/:id")
			users.PUT("/:id")
			users.DELETE("/:id")
		}
		products := api.Group("/products")
		{
			products.POST("/")
			products.GET("/:id")
			products.PUT("/:id")
			products.DELETE("/:id")
		}
		orders := api.Group("/orders")
		{
			orders.POST("/")
			orders.GET("/:id")
			orders.PUT("/:id")
			orders.DELETE("/:id")
		}
		auth := api.Group("/auth")
		{
			auth.POST("/login")
			auth.POST("/register")
			auth.POST("/logout")
		}
		payments := api.Group("/payments")
		{
			payments.POST("/")
			payments.POST("/:id")
		}
	}
	return router
}
