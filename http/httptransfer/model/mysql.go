package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"httptransfer/server"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "root"
	dbName   = "mysql"
	dataPool = "mysql"
)

func init() {
	orm.RegisterDriver(dataPool, orm.DRMySQL)
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", user, password, host, port, dbName)
	orm.RegisterDataBase("default", dataPool, psqlInfo)
	orm.SetMaxIdleConns("default", 1000)
	orm.SetMaxOpenConns("default", 2000)

	o := orm.NewOrm()
	goods := []server.Goods{}
	_, err := o.Raw("select * from t_goods").QueryRows(&goods)
	if err != nil {
		fmt.Println("查询出错", err)
		return
	}
	//fmt.Println("init ： goods集合:",goods)
}
