package controllers

import (
	"fmt"
	"golang_mysql_lock/model/drivers"
	"log"
	"time"
)

type GolangMysqlLock struct {
	ApiController
}

// @Description 乐观锁*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /lock_optimi [post]
func (this *GolangMysqlLock) LockOptimi() {
	//有道云 mysql文件夹里有资料
}

//排它锁和共享锁如果用到增删改上，之前的sql有一个用到锁，这个增删改就会报错，如果是没有则能正常运行。
//一般可以用在查询语句上，多少个查询语句使用锁都没事，并且只会执行第一个食物里的增删改，后面的必须要等待第一个
//提交才能执行增删改

// @Description 悲观锁 的 共享*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /lock_pessimi [post]
func (this *GolangMysqlLock) LockPessimi() {

	//此方法需要被调用两次
	//lock in share mode 共享锁，共享锁是需要全部被释放完才能进行其他操作
	//第一次update user_one set name = '1' where name = ? 会执行
	//第二次则不会，需要共享锁提交释放才能执行

	session := drivers.Engine.NewSession()
	defer session.Clone()
	session.Begin()
	user := []UserOne{}
	err := session.SQL("select * from user_one where Name in (?,?,?) lock in share mode ", "1", "2", "3").Find(&user)
	if err != nil {
		log.Println("err:", err)
		return
	}
	_, err = session.SQL("update user_one set name = '1' where name = ? ", "1").Execute()
	if err != nil {
		log.Println("err:", err)
		return
	}
	fmt.Println("user:", user)
	chanT := time.Tick(100 * time.Second)
	<-chanT
	fmt.Println("zz")
	session.Commit()

}

// @Description 悲观锁 的 排他*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /lock_pessimi_for [post]
func (this *GolangMysqlLock) LockPessimiFor() {

	//目前知道的排它锁和共享锁所实现的效果差不多

	session := drivers.Engine.NewSession()
	defer session.Clone()
	session.Begin()
	user := []UserOne{}
	err := session.SQL("select * from user_one where Name in (?,?,?) for update ", "1", "2", "3").Find(&user)
	if err != nil {
		log.Println("err:", err)
		return
	}
	_, err = session.SQL("update user_one set name = '1' where name = ?   ", "1").Execute()
	if err != nil {
		log.Println("err:", err)
		return
	}
	fmt.Println("user:", user)
	chanT := time.Tick(100 * time.Second)
	<-chanT
	fmt.Println("zz")
	session.Commit()

}
