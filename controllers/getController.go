package controllers

import (
	"math"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetPoints(c *gin.Context) {
	
	receiptID := c.Param("id")
	receipt, exists := receipts[receiptID]

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt not found"})
		return
	}

	points := 0

	// Points for alphanumeric characters in retailer name
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	// 50 points if total ends in ".00"
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 25 points if total is a multiple of 0.25
	totalValue, _ := strconv.ParseFloat(receipt.Total, 64)
	if int(totalValue*100)%25 == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	itemPairPoints := (len(receipt.Items) / 2) * 5
	points += itemPairPoints

	// Points for items with shortDescription length multiple of 3
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description)%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := int(math.Ceil(itemPrice * 0.2))
			points += itemPoints
		}
	}

	// 6 points if purchase day is odd
	day, _ := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2])
	if day%2 != 0 {
		points += 6
	}

	// 10 points if time is between 2 PM and 4 PM
	hour, _ := strconv.Atoi(strings.Split(receipt.PurchaseTime, ":")[0])
	if hour >= 14 && hour < 16 {
		points += 10
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
