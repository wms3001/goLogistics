package order

type OrderAddress struct {
	Country `json:"country" bson:"country"`
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
}
