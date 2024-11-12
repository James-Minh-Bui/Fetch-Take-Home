package controllers

import (
	"FETCH-TAKE-HOME/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var receipts = make(map[string]models.Receipt) 

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()

	receipts[id] = receipt

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
	})
}

