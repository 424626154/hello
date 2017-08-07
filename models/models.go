package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Admin struct {
	Id       int64
	Username string
	Password string
	Time     time.Time
}

type Web struct {
	Id    int64
	Url   string
	Icon  string
	Title string
	Time  time.Time
	Del   bool
}

func RegisterDB() {
	db_url := "root:890503@/hello?charset=utf8&loc=Local"
	orm.RegisterDataBase("default", "mysql", db_url)
	// register model
	orm.RegisterModel(new(Admin)) //官网管理员表
	orm.RegisterModel(new(Web))   //产品
}

func GetOneAdmin(account string) (*Admin, error) {
	o := orm.NewOrm()
	var admins []Admin
	_, err := o.Raw("SELECT * FROM admin WHERE username = ? ", account).QueryRows(&admins)
	admin := &Admin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}
