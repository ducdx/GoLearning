package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}

type item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type order struct {
	TotalPrice  int    `json:"total"`
	IsPaid      bool   `json:"paid"`
	Fragile     bool   `json:",omitempty"`
	OrderDetail []item `json:"orderdetail"`
}

type customer struct {
	UserName      string  `json:"username"`
	Password      string  `json:"_"`
	Token         string  `json:"_"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}

func main() {
	jsonData := []byte(`
	{
		"username" :"blackhat",
		"shipto":  
			{
		    		"street": "Sulphur Springs Rd",
	    			"city": "Park City",
	    			"state": "VA",
	    			"zipcode": 12345
			},
		"order":
			{
				"paid":false,
				"orderdetail" : 
					[{
						"itemname":"A Guide to the World of zeros and ones",
						"desc": "book",
						"qty": 3,
						"price": 50
					}]
			
			}
		
	}
	`)

	if !json.Valid(jsonData) {
		fmt.Println("Data is invalid")
		os.Exit(1)
	}

	var cus customer
	err := json.Unmarshal(jsonData, &cus)
	if err != nil {
		log.Fatal("Decode Error: ", err)
	}

	item1 := item{
		Name:        "Final Fantasy The Zodiac Age",
		Description: "Nintendo Switch Game",
		Quantity:    3,
		Price:       50,
	}

	item2 := item{
		Name:     "Crystal Drinking Glass",
		Quantity: 11,
		Price:    25,
	}
	cus.PurchaseOrder.OrderDetail = append(cus.PurchaseOrder.OrderDetail, item1)
	cus.PurchaseOrder.OrderDetail = append(cus.PurchaseOrder.OrderDetail, item2)
	cus.PurchaseOrder.IsPaid = true
	cus.PurchaseOrder.Fragile = true
	updateTotalPrices(&cus)

	jsonString, err := json.MarshalIndent(cus, "", "    ")
	if err != nil {
		fmt.Println("Encode Error: ", err)
	} else {
		fmt.Println(string(jsonString))
	}
}

func updateTotalPrices(cus *customer) {
	var sum int
	listItem := cus.PurchaseOrder.OrderDetail
	for _, it := range listItem {
		sum += it.Price * it.Quantity
	}
	cus.PurchaseOrder.TotalPrice = sum
}
