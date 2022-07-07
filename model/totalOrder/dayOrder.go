package totalOrder

type DayOrder struct {
	Order int64  `json:"order" bson:"order"`
	Day   string `json:"day" bson:"day"`
}
