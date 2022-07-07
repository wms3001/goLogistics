package order

type OrderWarehouse struct {
	Warehouse     string `json:"warehouse" bson:"warehouse"`
	WarehouseName string `json:"warehouseName" bson:"warehouseName"`
}
