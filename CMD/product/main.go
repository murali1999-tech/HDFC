package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Available bool    `json:"available"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
}

var products = []Product{
	{
		ID:        "1",
		Name:      "Premium Savings Account",
		Available: true,
		Price:     100.00,
		Category:  "premium",
	},
	{
		ID:        "2",
		Name:      "Regular Checking Account",
		Available: true,
		Price:     5.00,
		Category:  "regular",
	},
	{
		ID:        "3",
		Name:      "Budget Credit Card",
		Available: false,
		Price:     0.00,
		Category:  "budget",
	},
}

func main() {
	http.HandleFunc("/products", getProducts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
