package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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
	Desc  string
	Time  time.Time
	Del   bool
	Type  int8 //0 home 1 blog
}

func RegisterDB() {
	cnf, err := config.NewConfig("ini", "conf/hello.conf")
	if err != nil {
		beego.Error(err)
	}
	db := cnf.String("hello::db")
	ali_db := cnf.String("hello::ali_db")
	db_url := db + "?charset=utf8&loc=Local"
	// 本地 local 阿里 ali
	hello_runenv := cnf.String("hello::hello_runenv")
	beego.Debug("hello_runenv:", hello_runenv)
	if hello_runenv == "ali" {
		db_url = ali_db + "?charset=utf8&loc=Local"
	}
	beego.Debug("---db_url:", db_url)
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

func AddWeb(title string, weburl string, desc string, icon string, wtype int8) (*Web, error) {
	time := time.Now()
	o := orm.NewOrm()
	obj := &Web{Title: title, Url: weburl, Icon: icon, Desc: desc, Type: wtype, Time: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func GetWebFid(id string) (*Web, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var webs []Web
	_, err = o.Raw("SELECT * FROM web WHERE id = ? ", cid).QueryRows(&webs)
	web := &Web{}
	if len(webs) > 0 {
		web = &webs[0]
	}
	return web, err
}

func GetWebs() ([]Web, error) {
	o := orm.NewOrm()
	var objs []Web
	_, err := o.Raw("SELECT * FROM web  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}
func GetFTypeWebs(ftype int8) ([]Web, error) {
	o := orm.NewOrm()
	var objs []Web
	_, err := o.Raw("SELECT * FROM web  WHERE type = ? AND del = false ORDER BY id DESC", ftype).QueryRows(&objs)
	return objs, err
}
func UpWebInfo(id, title, weburl, desc string, wtype int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Web{Id: cid}
	obj.Title = title
	obj.Url = weburl
	obj.Desc = desc
	obj.Type = wtype
	_, err = o.Update(obj, "title", "url", "desc", "type")
	return err
}
func UpWebIcon(id, icon string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Web{Id: cid}
	obj.Icon = icon
	_, err = o.Update(obj, "icon")
	return err
}

func DelWebFid(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Web{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}
