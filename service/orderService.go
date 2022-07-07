package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wms3001/goMongo"
	"github.com/wms3001/goMongo/model"
	"go.mongodb.org/mongo-driver/bson"
	//"logistics/db"
	//"logistics/model/modb"
	orderInfo "logistics/model/order"
)

var order *orderInfo.Order
var conn *model.Conn
var addOne *model.AddOne
var addMany *model.AddMany
var mData *model.MData
var collect string = "order1"

func AddOrder(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	order = &orderInfo.Order{}
	c.ShouldBindBodyWith(order, binding.JSON)
	addOne = goMongo.AddOne(conn.Collection, &order)
	c.JSON(200, &addOne)
}

func AddOrders(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	orders := []interface{}{}
	c.ShouldBindBodyWith(&orders, binding.JSON)
	addMany = goMongo.AddMany(conn.Collection, orders)
	c.JSON(200, &addMany)
}

func GetOrderByOrderId(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	orderId := c.Query("orderId")
	ser := bson.D{{"orderId", orderId}}
	one := goMongo.GetOneSingl(conn.Collection, ser)
	c.JSON(200, &one)
}

func GetOrderById(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	id := c.Query("id")
	ser := bson.D{{"_id", id}}
	one := goMongo.GetOneSingl(conn.Collection, ser)
	c.JSON(200, &one)
}

func GetOrder(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	one := goMongo.GetOne(conn.Collection, mData)
	c.JSON(200, &one)
}

func GetOrders(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	many := goMongo.GetMany(conn.Collection, mData)
	c.JSON(200, &many)
}

func GetManyOrders(mData *model.MData) *model.GetMany {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	many := goMongo.GetMany(conn.Collection, mData)
	return many
}

func UpdateOrder(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	up := goMongo.UpdateOne(conn.Collection, mData)
	c.JSON(200, &up)
}

func UpdateOrderInfo(mData *model.MData) *model.UpdateOne {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	up := goMongo.UpdateOne(conn.Collection, mData)
	return up
}

func UpdateOrders(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	up := goMongo.UpdateMany(conn.Collection, mData)
	c.JSON(200, &up)
}

func DeleteOrder(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	del := goMongo.DeleteOne(conn.Collection, mData)
	c.JSON(200, &del)
}

func DeleteOrders(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	del := goMongo.DeleteMany(conn.Collection, mData)
	c.JSON(200, &del)
}

func CountOrder(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	cou := goMongo.Count(conn.Collection, mData)
	c.JSON(200, &cou)
}

func DistinctOrder(c *gin.Context) {
	conn := goMongo.OpenConn(collect)
	defer goMongo.CloseConn(conn.Client)
	mData = &model.MData{}
	c.ShouldBindBodyWith(mData, binding.JSON)
	dis := goMongo.Distinct(conn.Collection, mData, "carrier")
	c.JSON(200, &dis)
}
