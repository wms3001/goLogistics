package totalOrder

type DayOrderDepartment struct {
	Order      int64  `json:"order" bson:"order"`
	Day        string `json:"day" bson:"day"`
	Department string `json:"department" bson:"department"`
}
