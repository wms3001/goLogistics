package service

import (
	"encoding/json"
	"github.com/wms3001/goMongo/model"
	"log"
	orderInfo "logistics/model/order"
	"logistics/model/totalOrder"
)

var OrderChan chan totalOrder.DayOrder = make(chan totalOrder.DayOrder, 5000)
var OrderDepartmentChan chan totalOrder.DayOrderDepartment = make(chan totalOrder.DayOrderDepartment, 5000)
var OrderCountryChan chan totalOrder.DayOrderCountry = make(chan totalOrder.DayOrderCountry, 5000)
var OrderPlatformChan chan totalOrder.DayOrderPlatform = make(chan totalOrder.DayOrderPlatform, 5000)
var OrderAccountChan chan totalOrder.DayOrderAccount = make(chan totalOrder.DayOrderAccount, 5000)
var OrderSubCompanyChan chan totalOrder.DayOrderSubCompany = make(chan totalOrder.DayOrderSubCompany, 5000)
var OrderSubCompanyDepartmentChan chan totalOrder.DayOrderSubCompanyDepartment = make(chan totalOrder.DayOrderSubCompanyDepartment, 5000)
var OrderPersonChan chan totalOrder.DayOrderPerson = make(chan totalOrder.DayOrderPerson, 5000)
var cData *model.MData
var con *model.MData

func StatisticsInChannel() {
	cData := &model.MData{}
	cData.FilterMap = map[string]interface{}{"status": 0}
	cData.Type = "and"
	cData.MOption.Limit = 5000
	many := GetManyOrders(cData)
	if many.Code == 1 {
		for _, v := range many.Data {
			var order *orderInfo.Order
			mJson, _ := json.Marshal(v)
			json.Unmarshal(mJson, &order)
			go DayOrder(order)
			go DayOrderDepartment(order)
			go DayOrderCountry(order)
			go DayOrderPlatform(order)
			go DayOrderAccount(order)
			go DayOrderSubCompany(order)
			go DayOrderSubCompanyDepartment(order)
			go DayOrderPerson(order)
			go DaySku(order)
			go DaySkuDepartment(order)
			go DaySkuPlatform(order)
			go DaySkuCountry(order)
			go DaySkuAccount(order)
			go DaySkuSubCompany(order)
			go DaySkuSubCompanyDepartment(order)
			go DaySkuPerson(order)
			cData.FilterMap["orderId"] = v["orderId"]
			cData.UpMap = map[string]interface{}{"status": 1}
			UpdateOrderInfo(cData)
		}
		//close(OrderChan)
		//close(OrderDepartmentChan)
	}
}

func StatisticsOutChannel() {
	//timeStr:=time.Now().Format("2006-01-02")
	var coll = &Collect{}
	var con = &model.MData{}
	for {
		select {
		case o, ok := <-OrderChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.FilterMap["day"] = o.Day
				con.Type = "and"
				coll.Name = "dayOrder"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderChan is closed!")
			}
		case o, ok := <-OrderDepartmentChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["department"] = o.Department
				coll.Name = "dayOrderDepartment"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderDepartmentChan is closed!")
			}
		case o, ok := <-OrderCountryChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["country"] = o.Country
				coll.Name = "dayOrderCountry"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderCountryChan is closed!")
			}
		case o, ok := <-OrderPlatformChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["platform"] = o.Platform
				coll.Name = "dayOrderPlatform"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderPlatformChan is closed!")
			}
		case o, ok := <-OrderAccountChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["account"] = o.Account
				coll.Name = "dayOrderAccount"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderAccountChan is closed!")
			}
		case o, ok := <-OrderSubCompanyChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["subCompany"] = o.SubCompany
				coll.Name = "dayOrderSubCompany"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderSubCompanyChan is closed!")
			}
		case o, ok := <-OrderSubCompanyDepartmentChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["subCompany"] = o.SubCompany
				con.FilterMap["department"] = o.Department
				coll.Name = "dayOrderSubCompanyDepartment"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderSubCompanyDepartmentChan is closed!")
			}
		case o, ok := <-OrderPersonChan:
			if ok {
				con.FilterMap = map[string]interface{}{}
				con.UpMap = map[string]interface{}{}
				con.Type = "and"
				con.FilterMap["day"] = o.Day
				con.FilterMap["user"] = o.User
				coll.Name = "dayOrderPerson"
				one := coll.find(con)
				if one.Code != 1 {
					coll.add(o)
				} else {
					con.UpMap["order"] = o.Order + one.Data["order"].(int64)
					coll.update(con)
				}
			} else {
				log.Printf("Channel OrderPersonChan is closed!")
			}
		default:
			//fmt.Println("waiting...")
		}
	}
}

func DayOrder(order *orderInfo.Order) {
	var orderDay = totalOrder.DayOrder{}
	orderDay.Day = order.OrderDay
	orderDay.Order = 1
	OrderChan <- orderDay
}

func DayOrderDepartment(order *orderInfo.Order) {
	var orderDepartment = totalOrder.DayOrderDepartment{}
	orderDepartment.Day = order.OrderDay
	orderDepartment.Order = 1
	orderDepartment.Department = order.Department
	OrderDepartmentChan <- orderDepartment
}
func DayOrderCountry(order *orderInfo.Order) {
	var orderCountry = totalOrder.DayOrderCountry{}
	orderCountry.Day = order.OrderDay
	orderCountry.Order = 1
	orderCountry.Country = order.CountryCode
	orderCountry.CountryName = order.CountryEnName
	OrderCountryChan <- orderCountry
}
func DayOrderPlatform(order *orderInfo.Order) {
	var orderPlatform = totalOrder.DayOrderPlatform{}
	orderPlatform.Day = order.OrderDay
	orderPlatform.Order = 1
	orderPlatform.Platform = order.Platform
	OrderPlatformChan <- orderPlatform
}
func DayOrderAccount(order *orderInfo.Order) {
	var orderAccount = totalOrder.DayOrderAccount{}
	orderAccount.Account = order.Account
	orderAccount.Day = order.OrderDay
	orderAccount.Order = 1
	OrderAccountChan <- orderAccount
}
func DayOrderSubCompany(order *orderInfo.Order) {
	var orderSubCompany = totalOrder.DayOrderSubCompany{}
	orderSubCompany.SubCompany = order.SubCompany
	orderSubCompany.Day = order.OrderDay
	orderSubCompany.Order = 1
	OrderSubCompanyChan <- orderSubCompany
}
func DayOrderSubCompanyDepartment(order *orderInfo.Order) {
	var orderSubCompanyDepartment = totalOrder.DayOrderSubCompanyDepartment{}
	orderSubCompanyDepartment.Day = order.OrderDay
	orderSubCompanyDepartment.Order = 1
	orderSubCompanyDepartment.Department = order.Department
	orderSubCompanyDepartment.SubCompany = order.SubCompany
	OrderSubCompanyDepartmentChan <- orderSubCompanyDepartment
}
func DayOrderPerson(order *orderInfo.Order) {
	var orderPerson = totalOrder.DayOrderPerson{}
	orderPerson.Order = 1
	orderPerson.Day = order.OrderDay
	orderPerson.User = order.OwnerName
	OrderPersonChan <- orderPerson
}
func DaySku(order *orderInfo.Order) {
	log.Printf("Sku %s", order.OrderId)
}
func DaySkuDepartment(order *orderInfo.Order) {
	log.Printf("SkuDepartment %s", order.OrderId)
}
func DaySkuPlatform(order *orderInfo.Order) {
	log.Printf("SkuPlatform %s", order.OrderId)
}
func DaySkuCountry(order *orderInfo.Order) {
	log.Printf("SkuCountry %s", order.OrderId)
}
func DaySkuAccount(order *orderInfo.Order) {
	log.Printf("SkuAccount %s", order.OrderId)
}
func DaySkuSubCompany(order *orderInfo.Order) {
	log.Printf("SkuSubCompany %s", order.OrderId)
}
func DaySkuSubCompanyDepartment(order *orderInfo.Order) {
	log.Printf("SkuSubCompanyDepartment %s", order.OrderId)
}
func DaySkuPerson(order *orderInfo.Order) {
	log.Printf("SkuPerson %s", order.OrderId)
}
