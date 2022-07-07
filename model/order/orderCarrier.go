package order

type OrderCarrier struct {
	Carrier            string `json:"carrier" bson:"carrier" binding:"required"`
	TrackNumber        string `json:"trackNumber" bson:"trackNumber"`
	CarrierCompany     string `json:"carrierCompany" bson:"carrierCompany"`
	CarrierCompanyCode string `json:"carrierCompanyCode" bson:"carrierCompanyCode"`
}
