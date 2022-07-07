package order

type Order struct {
	OrderId                         string  `json:"orderId" bson:"orderId" binding:"required"`
	Status                          int     `json:"status" bson:"status" binging:"required"`
	OrderStatus                     int     `json:"orderStatus" bson:"orderStatus"`
	OrderDepartmentStatus           int     `json:"orderDepartmentStatus" bson:"orderDepartmentStatus"`
	OrderCountryStatus              int     `json:"orderCountryStatus" bson:"orderCountryStatus"`
	OrderPlatformStatus             int     `json:"orderPlatformStatus" bson:"orderPlatformStatus"`
	OrderAccountStatus              int     `json:"orderAccountStatus" bson:"orderAccountStatus"`
	OrderSubCompanyStatus           int     `json:"orderSubCompanyStatus" bson:"orderSubCompanyStatus"`
	OrderSubCompanyDepartmentStatus int     `json:"orderSubCompanyDepartmentStatus" bson:"orderSubCompanyDepartmentStatus"`
	OrderPersonStatus               int     `json:"orderPersonStatus" bson:"orderPersonStatus"`
	SkuStatus                       int     `json:"skuStatus" bson:"skuStatus"`
	SkuDepartmentStatus             int     `json:"skuDepartmentStatus" bson:"skuDepartmentStatus"`
	SkuPlatformStatus               int     `json:"skuPlatformStatus" bson:"skuPlatformStatus"`
	SkuCountryStatus                int     `json:"skuCountryStatus" bson:"skuCountryStatus"`
	SkuAccountStatus                int     `json:"skuAccountStatus" bson:"skuAccountStatus"`
	SkuSubCompanyStatus             int     `json:"skuSubCompanyStatus" bson:"skuSubCompanyStatus"`
	SkuSubCompanyDepartmentStatus   int     `json:"skuSubCompanyDepartmentStatus" bson:"skuSubCompanyDepartmentStatus"`
	SkuPersonStatus                 int     `json:"skuPersonStatus" bson:"skuPersonStatus"`
	OrderTotal                      float64 `json:"orderTotal" bson:"orderTotal"`
	OrderDay                        string  `json:"orderDay" bson:"orderDay"`
	OrderWarehouse                  `json:"orderWarehouse" bson:"orderWarehouse"`
	OrderPlatform                   `json:"orderPlatform" bson:"orderPlatform"`
	OrderCarrier                    `json:"orderCarrier" bson:"orderCarrier"`
	OrderAddress                    `json:"orderAddress" bson:"orderAddress"`
	OrderDetails                    []OrderDetail `json:"orderDetails" bson:"orderDetails"`
	OrderOwner                      `json:"orderOwner" bson:"orderOwner"`
}
