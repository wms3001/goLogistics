package service

import (
	"github.com/wms3001/goMongo"
	"github.com/wms3001/goMongo/model"
)

//var conn *model.Conn

type Collect struct {
	Name string `json:"name" bson:"name"`
}

func (coll *Collect) add(data interface{}) *model.AddOne {
	conn := goMongo.OpenConn(coll.Name)
	//defer goMongo.CloseConn(conn.Client)
	addOne = goMongo.AddOne(conn.Collection, &data)
	return addOne
}

func (coll *Collect) find(mData *model.MData) *model.GetOne {
	conn := goMongo.OpenConn(coll.Name)
	//defer goMongo.CloseConn(conn.Client)
	getOne := goMongo.GetOne(conn.Collection, mData)
	return getOne
}

func (coll *Collect) update(mData *model.MData) *model.UpdateOne {
	conn := goMongo.OpenConn(coll.Name)
	//defer goMongo.CloseConn(conn.Client)
	update := goMongo.UpdateOne(conn.Collection, mData)
	return update
}
