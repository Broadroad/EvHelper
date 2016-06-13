package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/*
func init() {
	OrderList = make(map[string]*Order)
	u := Order{1, "useId", "jobId", "description", "timestamp", "location", "lat", "lot", "telephone"}
	OrderList["orderId"] = &u
}
*/

type SfOrder struct {
	Orderid              int `orm:"pk;auto"`
	Userid               int `orm:"not null"`
	Senderid             int `orm:"not null"`
	Sendertime           time.Time
	Ordertime            time.Time
	Usergetpackagetime   time.Time
	Sendergetpackagetime time.Time
	Expresslocation      string
	Senderlocation       string
	Expresslatitude      string `orm:"not null"`
	Expresslogitude      string `orm:"not null"`
	Senderlatitude       string
	Senderlogitude       string
	Price                float64
	Telephone            string
	IsDeleted            int
	Description          string
	IsOrdered            int
	IsCompleted          int
}

func (this *SfOrder) GetAllOrders() ([]*SfOrder, error) {
	var Orders []*SfOrder
	o := orm.NewOrm()
	qs := o.QueryTable("sf_order")
	qs = qs.Filter("is_deleted", 0)
	_, err := qs.OrderBy("-sendertime").All(&Orders)
	return Orders, err
}

func (this *SfOrder) SaveOrder() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}
