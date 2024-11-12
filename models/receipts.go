package models

// Item represents an item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"` // Price as string
}

// Receipt represents the structure of the receipt
type Receipt struct {
	Retailer    string  `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Total       string `json:"total" binding:"required"` // Total as string
	Items       []Item `json:"items" binding:"required"`
}
