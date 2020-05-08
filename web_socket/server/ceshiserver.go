package server

import "time"

type goods struct {
}

type Goods struct {
	Name       string    `orm:"column(name)"`
	GoodsId    int       `orm:"column(goodsId)"`
	Price      string    `orm:"column(price)"`
	UpdateTime time.Time `orm:"column(updateTime)"`
	InsertTime time.Time `orm:"column(insertTime)"`
	Unit       string    `orm:"column(unit)"`
}

func Ceshi() goods {

	return goods{}
}
