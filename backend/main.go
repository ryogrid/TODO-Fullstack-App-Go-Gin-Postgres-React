package main

import (
	api "github.com/ryogrid/TODO-Fullstack-App-Go-Gin-Postgres-React/backend/main/backend/api"
	"net/http"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

// Function called for index
func indexView(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "TODO APP"})
}

// Setup Gin Routes
func SetupRoutes() *gin.Engine {
	// Use Gin as router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Set route for index
	router.GET("/", indexView)

	// Set routes for API
	// Update to POST, UPDATE, DELETE etc
	router.Static("/todo", "./moved/frontend/build")
	router.Static("/static", "./moved/frontend/build/static")
	router.Static("/manifest.json", "./moved/frontend/build/manifest.json")
	router.GET("/items", api.TodoItems)
	router.GET("/item/create/:item", api.CreateTodoItem)
	router.GET("/item/update/:id/:done", api.UpdateTodoItem)
	router.GET("/item/delete/:id", api.DeleteTodoItem)

	// Set up Gin Server
	return router
}

// Main function
func main() {
	api.SetupPostgres()
	router := SetupRoutes()
	//router.Run(":8080")
	//port := os.Getenv("PORT")
	port := "8088"
	http.ListenAndServe(":"+port, router)
}
