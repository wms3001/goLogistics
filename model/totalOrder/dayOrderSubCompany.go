package totalOrder

type DayOrderSubCompany struct {
	Order      int64  `json:"order" bson:"order"`
	Day        string `json:"day" bson:"day"`
	SubCompany string `json:"subCompany" bson:"subCompany"`
}
