package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
	Inventory int     `json:"inventory"`
}

type Order struct {
	Products     []OrderedProduct `json:"products"`
	OrderValue   float64          `json:"orderValue"`
	DispatchDate time.Time        `json:"dispatchDate"`
	OrderStatus  string           `json:"orderStatus"`
}

type OrderedProduct struct {
	ProductID string `json:"productID"`
	Quantity  int    `json:"quantity"`
}

var products = []Product{
	{
		ID:        "1",
		Name:      "Premium Savings Account",
		Price:     100.00,
		Category:  "premium",
		Inventory: 10,
	},
	{
		ID:        "2",
		Name:      "Regular Checking Account",
		Price:     5.00,
		Category:  "regular",
		Inventory: 20,
	},
	{
		ID:        "3",
		Name:      "Budget Credit Card",
		Price:     0.00,
		Category:  "budget",
		Inventory: 30,
	},
}

func main(){

	http.HandleFunc("/products",productsHandler)
	http.HandleFunc("/orders",ordersHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))

	func productsHandler(w http.ResponseWriter, r *http.Request)
	{
		if r.Method==http.MethodGet{
			productsJson,err:=json.Marshal(products)
			if err!=nil{
				http.Error(w,err.Error(),
			http..StatusInternalServerError)
			return
			}
			w.Header().Set("Content-Type","application/json")
			w.Write(products.Json)
			return
		}
		http.Error(w,"Method not allowed", http.StatusMethodNotAllowed)
	}

	function ordersHandler(w http.ResponseWriter, r *http.Request)const{
		if r.Method==http.MethodPost{
			var order Order
			err:= json.NewDecoder(r.Body).Decpde(&order)
			if err!=nil{
				http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
			}
		
			orervalye:=0.0
			for _,orderedProduct:=range order.Products{
				product,err:=getProductByID(orderdProduct.ProductID)
				if err !=nil{
					http.Error(w, err.Error(),http.StatusBadRequest)
					return
				}
				if product.Inventory<orderedProduct,Quantity{
					http.Error(w, "Insufficient inventory  for "+product.Name,http.StatusBadRequest)
					return 
				}
				orderValue +=product.Price*float64(orderedProduct.Quantity)

				}

				order.Ordervalue=orderValue
				order.DispatchDate=time.Now().Add(24*time.Hour)
				order.OrderStatus="Dispatched"
			}
			}
		}
	}
}

	


