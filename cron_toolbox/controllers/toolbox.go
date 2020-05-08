package controllers

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"time"
)

type Toolbox struct {
	ApiController
}

// @Description wenjianliu 做的时候记得把登录注释了*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /toolbox [post]
func (this *Toolbox) Toolbox() {
	fmt.Println("----------进入1---------")
	//秒钟：0-59、分钟：0-59、小时：1-23、日期：1-31、月份：1-12、星期：0-6（0 表示周日）
	tk := toolbox.NewTask("cronTimer", "* * * * * *", func() error {
		CronTimer()
		return nil
	})
	toolbox.AddTask("cronTimer", tk)
	toolbox.StartTask()
	fmt.Println("----------结束---------")
}

func CronTimer() {
	ticker := time.NewTicker(1 * time.Second)
	//无线循环，每隔1秒运行一次,但可以被break中断
	for _ = range ticker.C {
		fmt.Println("----------进入定时器方法-------", ticker.C)
		break
	}
	//关闭名称为cronTimer的定时器
	delay := time.NewTicker(4 * time.Second)
	//延迟四秒
	<-delay.C
	toolbox.DeleteTask("cronTimer")
}
