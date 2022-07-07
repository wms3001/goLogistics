package order

type OrderOwner struct {
	OwnerName  string `json:"ownerName" bson:"ownerName"`
	OwnerId    string `json:"ownerId" bson:"ownerId"`
	SubCompany string `json:"subCompany" bson:"subCompany"`
	Department string `json:"department" bson:"department"`
}
