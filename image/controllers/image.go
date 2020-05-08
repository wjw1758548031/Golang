package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

type ImageController struct {
	ApiController
}

// @Description wenjianliu 做的时候记得把登录注释了*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /testImageA [post]
func (this *ImageController) TestImageA() {

	//打开图片
	ff, _ := os.Open("../image/controllers/dd.jpg")
	defer ff.Close()
	sourcebuffer := make([]byte, 50000000)
	//算出图片的大小
	n, _ := ff.Read(sourcebuffer)
	fmt.Println("n:", n)
	//dist := make([]byte, 50000000)
	//base64压缩
	//转换成base64
	ddd1 := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	//转换成[]byte
	ddd := []byte(ddd1)
	ioutil.WriteFile("../image/dd.jpg", ddd, 0666)

	fmt.Println("ddd", ddd)
	//fmt.Println("ff",ff)
	//_ = ioutil.WriteFile("../image/controllers/1548660223(1).jpg", dist, 0667)
	// 文件写成 base64   imgFile -> base64
	//ddd, _ := base64.StdEncoding.DecodeString("../image/controllers/1548660223(1).jpg")
	//把转码后的图片移到../image/1548660223(1).jpg位置
	//ioutil.WriteFile("../image/1548660223(1).jpg", ddd, 0666)
	//fmt.Println("ddd:",ddd)

	//将base64码转成buffer
	bbb := bytes.NewBuffer(ddd)
	fmt.Println("bbb:", bbb)

}
