package totalOrder

type DayOrderSubCompanyDepartment struct {
	Order      int64  `json:"order" bson:"order"`
	Day        string `json:"day" bson:"day"`
	SubCompany string `json:"subCompany" bson:"subCompany"`
	Department string `json:"department" bson:"department"`
}
