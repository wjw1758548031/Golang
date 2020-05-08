package src

import (
	"bytes"
	"os"

	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ParkingPlate struct {
	ParkingId       string `json:"parking_id"`
	Channel         string `json:"Channel"`
	Flate           string `json:"Flate"`
	Image           string `json:"Image"`
	BackgroundImage string `json:"BackgroundImage"`
	InOutTime       string `json:"InOutTime"`
	ChannelType     string `json:"ChannelType"`
}

type Parking struct {
	IDPlate        string `json:"iDPlate"`
	PlateColor     string `json:"plate_color"`
	CarColor       string `json:"car_color"`
	CarType        string `json:"car_type"`
	ImagePlate     string `json:"image_plate"`
	ImageFullPlate string `json:"image_full_plate"`
	RecordTime     string `json:"record_time"`
}

var plate = func(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("参数传入失败：" + err.Error()))
		return
	}
	var form Parking
	err = json.Unmarshal(body, &form)
	var parkingId = req.FormValue("parkingId")
	var channel = req.FormValue("channel")
	var inOutType = req.FormValue("inOutType")
	fmt.Println(form)

	parkingPlate := &ParkingPlate{
		ParkingId:       parkingId,
		Channel:         channel,
		Flate:           form.IDPlate,
		Image:           form.ImagePlate,
		BackgroundImage: form.ImageFullPlate,
		InOutTime:       inOutType,
		ChannelType:     form.CarType,
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			res.Write([]byte("请求接口访问不到"))
		}
	}()

	//发送请求
	bodya, _ := json.Marshal(parkingPlate)
	addr := os.Getenv("SIPADDURL")
	resa, errs := http.Post(addr, "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodya)))
	if errs != nil {
		fmt.Println("Fatal error ", errs.Error())
	}
	defer resa.Body.Close()
	content, errs := ioutil.ReadAll(resa.Body)
	if errs != nil {
		fmt.Println("Fatal error ", errs.Error())
	}
	fmt.Println(string(content))

	fmt.Println("----------结束---------")
	res.Write([]byte("成功"))
}
