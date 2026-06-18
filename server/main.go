package main

import (
	"os"

	"github.com/SunilKotte/go-react-calorie-tracker-yt/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)

	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryByID)
	router.PUT("/entry/update/:id/", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id/", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	router.Run(":" + port)
}
