package main

import (
	"FETCH-TAKE-HOME/controllers" // Correct import path
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Serve HTML file for user interaction
	r.LoadHTMLFiles("index.html")

	// Serve HTML page at the root URL
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// POST request to process receipt (using controller)
	r.POST("/receipts/process", controllers.ProcessReceipt)

	// GET request to fetch points for a specific receipt (using controller)
	r.GET("/receipts/:id/points", controllers.GetPoints)

	// Start the server
	r.Run(":8080")
}
