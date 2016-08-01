package models

import (
	//	"errors"
	//	"strconv"
	"github.com/astaxie/beego/orm"
	"time"
)

type SfUser struct {
	Userid        int `orm:"pk;auto"`
	Username      string
	Password      string
	Registertime  time.Time
	Lastlogintime time.Time
	Email         string
	Telephone     string
	Idcard        string
}

func (this *SfUser) AddUser() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}

func GetUserByEmail(email string) (*SfUser, error) {
	u := &SfUser{}
	err := orm.NewOrm().QueryTable("sf_user").Filter("email", email).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetAllUsers() ([]*SfUser, error) {
	var Users []*SfUser
	o := orm.NewOrm()
	qs := o.QueryTable("sf_user")
	_, err := qs.OrderBy("-registertime").All(&Users)
	return Users, err
}
