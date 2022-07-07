package totalOrder

type DayOrderPerson struct {
	Order int64  `json:"order" bson:"order"`
	Day   string `json:"day" bson:"day"`
	User  string `json:"user" bson:"user"`
}
