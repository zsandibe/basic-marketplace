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
			users.POST("/users")
			users.GET("/users/{id}")
			users.PUT("/users/{id}")
			users.DELETE("/users/{id}")
		}
		products := api.Group("/products")
		{
			products.POST("/products")
			products.GET("/products/{id}")
			products.PUT("/products/{id}")
			products.DELETE("/products/{id}")
		}
		orders := api.Group("/orders")
		{
			orders.POST("/orders")
			orders.GET("/orders/{id}")
			orders.PUT("/orders/{id}")
			orders.DELETE("/orders/{id}")
		}
		auth := api.Group("/auth")
		{
			auth.POST("/login")
			auth.POST("/register")
			auth.POST("/logout")
		}
		payments := api.Group("/payments")
		{
			payments.POST("/payments")
			payments.POST("/payments/{id}")
		}
	}
	return router
}
