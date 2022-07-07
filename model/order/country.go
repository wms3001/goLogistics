package order

type Country struct {
	CountryCnName string `json:"countryCnName" bson:"countryCnName"`
	CountryEnName string `json:"countryEnName" bson:"countryEnName"`
	CountryCode   string `json:"countryCode" bson:"countryCode"`
}
