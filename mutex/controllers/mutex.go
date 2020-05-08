package controllers

import (
	"fmt"
	"sync"
	"time"
)

type Mutex struct {
	ApiController
}

var mutexx sync.Mutex
var rwmutexx sync.RWMutex

// @Description wenjianliu 互斥所，锁住所有携程*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /mutex [post]
func (this *Mutex) Mutex() {

	var form Zhi
	if !this.ParseForm(&form) {
		fmt.Println("进入参数退出")
		return
	}

	fmt.Println("----------进入---------", form.Id)

	//可以用两个程序调用此方法，第一次调用方法，循环的第一次被锁住，
	// 则如果在100秒内进行调用,第二次请求所有线程会进入等待中

	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {

		go func(i int) {
			// 使另外一个线程进入等待，必须要时此线程执行完，才能进入下个一线程，同一个方法中不开线程无效。
			mutexx.Lock()
			fmt.Println("go---", i)
			time.Sleep(2 * time.Second)
			c <- true
			mutexx.Unlock()
		}(i)

	}
	for i := 0; i < 10; i++ {
		<-c
	}
	fmt.Println("----------结束---------", form.Id)
}

// @Description wenjianliu 读写所，锁住所有携程*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /rwmutexx [post]
func (this *Mutex) Rwmutexx() {

	var form Zhi
	if !this.ParseForm(&form) {
		fmt.Println("进入参数退出")
		return
	}

	fmt.Println("----------进入---------", form.Id)

	//读锁 RLock 写锁Lock
	//读锁遇到读锁和普通流程没有区别
	//当读锁遇到写锁就需要等待,写优先读后
	//写锁就可以看做互斥锁

	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {

		go func(i int) {
			// 使另外一个线程进入等待，必须要时此线程执行完，才能进入下个一线程，同一个方法中不开线程无效。
			rwmutexx.RLock()
			fmt.Println("go---", i)
			time.Sleep(2 * time.Second)
			c <- true
			rwmutexx.RUnlock()
		}(i)

		rwmutexx.Lock()
		go func(i int) {
			// 使另外一个线程进入等待，必须要时此线程执行完，才能进入下个一线程，同一个方法中不开线程无效。
			fmt.Println("zz---", i)
			time.Sleep(2 * time.Second)
			c <- true
		}(i)
		rwmutexx.Unlock()

	}
	for i := 0; i < 20; i++ {
		<-c
	}
	fmt.Println("----------结束---------", form.Id)
}

// @Description 测试*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /ceshi [post]
func (this *Mutex) Ceshi() {
	var form Zhi
	if !this.ParseForm(&form) {
		fmt.Println("进入参数退出")
		return
	}
	for i := 0; i < 100000000000; i++ {
		fmt.Println(form.Id, "------", i)
	}

	/*	time.Sleep(5*time.Second)
		time.Sleep(5*time.Second)
		fmt.Println(form.Id,"------2")
		time.Sleep(5*time.Second)
		fmt.Println(form.Id,"------3")
		time.Sleep(5*time.Second)
		fmt.Println(form.Id,"------4")*/

}
