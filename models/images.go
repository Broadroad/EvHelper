package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Images struct {
	Id          int64 `orm:"pk;auto"`
	Uploadtime  time.Time
	Userid      int64
	Picname     string
	Key         string
	Point       int64
	IsDeleted   int64
	Description string
}

func (this *Images) GetImage() (*Images, error) {
	var images Images
	o := orm.NewOrm()
	qs := o.QueryTable("images")
	err := qs.Filter("id", this.Id).One(&images)

	if err != nil {
		return nil, errors.New("Image not exists")
	}
	return &images, nil
}

func (this *Images) GetImages() ([]*Images, error) {
	var images []*Images
	o := orm.NewOrm()
	qs := o.QueryTable("images")
	_, err := qs.All(&images)
	return images, err
}

func (this *Images) UploadImage() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}
