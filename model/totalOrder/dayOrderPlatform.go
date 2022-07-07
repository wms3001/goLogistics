package totalOrder

type DayOrderPlatform struct {
	Order    int64  `json:"order" bson:"order"`
	Day      string `json:"day" bson:"day"`
	Platform string `json:"platform" bson:"platform"`
}
