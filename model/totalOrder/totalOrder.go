package totalOrder

type TotalOrder struct {
	TotalOrder int64  `json:"totalOrder" bson:"totalOrder"`
	Day        string `json:"day" bson:"day"`
}
