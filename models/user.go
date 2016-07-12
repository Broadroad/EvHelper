package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id          int64 `orm:"pk;auto"`
	Regittime   time.Time
	Updatetime  time.Time
	Username    string
	Password    string
	Picture     string
	Points      int64
	Telephone   string
	IsDeleted   int64
	Description string
}

func (this *User) GetUser() (u *User, err error) {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err = qs.Filter("telephone", this.Telephone).One(&user)

	if err != nil {
		return &user, errors.New("User not exists")
	}
	return &user, nil
}

func (this *User) AddUser() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}
