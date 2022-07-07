package service

import (
	"github.com/wms3001/goMongo"
	"github.com/wms3001/goMongo/model"
	"logistics/model/totalOrder"
)

func AddTotalOrder(totalOrder totalOrder.TotalOrder) *model.AddOne {
	conn := goMongo.OpenConn("totalOrder")
	defer goMongo.CloseConn(conn.Client)
	addOne = goMongo.AddOne(conn.Collection, totalOrder)
	return addOne
}
