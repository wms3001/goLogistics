package totalOrder

type DayOrderAccount struct {
	Order   int64  `json:"order" bson:"order"`
	Day     string `json:"day" bson:"day"`
	Account string `json:"account" bson:"account"`
}
