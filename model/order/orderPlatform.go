package order

type OrderPlatform struct {
	Platform        string `json:"platform" bson:"platform" binging:"required"`
	PlatformOrderId string `json:"platformOrderId" bson:"platformOrderId" binging:"required"`
	Account         string `json:"account" bson:"account"`
}
