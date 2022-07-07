package totalOrder

type DayOrderCountry struct {
	Order       int64  `json:"order" bson:"order"`
	Day         string `json:"day" bson:"day"`
	Country     string `json:"country" bson:"country"`
	CountryName string `json:"countryName" bson:"countryName"`
}
