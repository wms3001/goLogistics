package order

type OrderDetail struct {
	Sku    string  `json:"sku" bson:"sku"`
	Amount int     `json:"amount" bson:"amount"`
	Price  float64 `json:"price" bson:"price"`
}
