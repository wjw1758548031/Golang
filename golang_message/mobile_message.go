package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// bingone
func main() {

	// 修改为您的apikey(https://www.yunpian.com)登录官网后获取
	apikey := "5a76868775c4fe53158517866cac4a69"
	// 修改为您要发送的手机号码，多个号码用逗号隔开
	mobile := "15821936548"
	// 发送内容
	text := "无敌是多么的寂寞1"
	// 发送模板编号
	tpl_id := 1
	// 语音验证码
	code := "11112" //string(1000+i)
	company := "云片网"
	// 发送模板内容
	tpl_value := url.Values{"#code#": {code}, "#company#": {company}}.Encode()

	// 获取user信息url
	url_get_user := "https://sms.yunpian.com/v2/user/get.json"
	// 智能模板发送短信url
	url_send_sms := "https://sms.yunpian.com/v2/sms/single_send.json"
	// 指定模板发送短信url
	url_tpl_sms := "https://sms.yunpian.com/v2/sms/tpl_single_send.json"
	// 发送语音短信url
	url_send_voice := "https://voice.yunpian.com/v2/voice/send.json"

	data_get_user := url.Values{"apikey": {apikey}}
	data_send_sms := url.Values{"apikey": {apikey}, "mobile": {mobile}, "text": {text}}
	data_tpl_sms := url.Values{"apikey": {apikey}, "mobile": {mobile},
		"tpl_id": {fmt.Sprintf("%d", tpl_id)}, "tpl_value": {tpl_value}}
	data_send_voice := url.Values{"apikey": {apikey}, "mobile": {mobile}, "code": {code}}

	httpsPostForm(url_get_user, data_get_user)
	httpsPostForm(url_send_sms, data_send_sms)
	httpsPostForm(url_tpl_sms, data_tpl_sms)
	httpsPostForm(url_send_voice, data_send_voice)

}

func httpsPostForm(url string, data url.Values) {
	resp, err := http.PostForm(url, data)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
