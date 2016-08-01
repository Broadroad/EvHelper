package models

import (
	"github.com/astaxie/beego/orm"
)

type EvcardUser struct {
	Userid   int `orm:"pk"`
	Username string
	Password string
}

func (this *EvcardUser) AddUser() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}

func (this *EvcardUser) GetUserByUserId(userid int) (*EvcardUser, error) {
	u := &EvcardUser{}
	err := orm.NewOrm().QueryTable("evcard_user").Filter("userid", userid).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
