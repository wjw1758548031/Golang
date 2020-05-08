package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type ApiController struct {
	beego.Controller
	err       error
	data      interface{}
	startTime int64
}

func (baseApi *ApiController) Prepare() {
	/* url := baseApi.Ctx.Input.URL()
		if(strings.HasPrefix(url, "/1/uploads")){
	        acsToken := baseApi.GetString("acs_token")
	        loginUser, err := GetLoginUser(acsToken)
	        if err != nil {
	            errResponse := ErrResponseVo{ErrCode: config.LOGIN_ERR_CODE, ErrMsg: err.Error()}
	            baseApi.Ctx.Output.JSON(errResponse, false, true)
	            return
	        }
	        baseApi.loginUser = *loginUser
		}else if(strings.HasPrefix(url, "/1/report") || strings.HasPrefix(url, "/1/advice") || strings.HasPrefix(url, "/1/print")){
	        var baseForm BaseForm
	        baseForm.AcsToken = baseApi.GetString("acs_token")
	        loginUser, err := GetLoginUser(baseForm.AcsToken)
	        if err != nil {
	            errResponse := ErrResponseVo{ErrCode: config.LOGIN_ERR_CODE, ErrMsg: err.Error()}
	            baseApi.Ctx.Output.JSON(errResponse, false, true)
	            return
	        }
	        baseApi.loginUser = *loginUser
	    }else if !strings.HasPrefix(url, "/1/login/do") && !strings.HasPrefix(url, "/1/sysadmin") && !strings.HasPrefix(url, "/1/demos") {
	        var baseForm BaseForm
	        err := json.Unmarshal(baseApi.Ctx.Input.RequestBody, &baseForm); if err != nil {
	            utils.PrtErrError("页面请求转换json出错", err)
	            // utils.PrtError(fmt.Sprintf("原始请求：\n%s", string(ctx.Input.RequestBody))) 在logfiler中打印了,这里就不需要重复打印了
	            errResponse := ErrResponseVo{ErrCode: config.BIZ_ERR_CODE, ErrMsg: "请求json格式错误"}
	            baseApi.Ctx.Output.JSON(errResponse, false, true)
	            return
	        }
	        loginUser, err := GetLoginUser(baseForm.AcsToken)
	        if err != nil {
	            errResponse := ErrResponseVo{ErrCode: config.LOGIN_ERR_CODE, ErrMsg: err.Error()}
	            baseApi.Ctx.Output.JSON(errResponse, false, true)
	            return
	        }
	        baseApi.loginUser = *loginUser
	    }*/
	baseApi.startTime = time.Now().Unix()
}

func (baseApi *ApiController) Finish() {
	responseToClient(baseApi)
	processTimeTooLongWarn(baseApi)
}

func responseToClient(baseApi *ApiController) {
	if baseApi.err != nil {
		baseApi.Data["json"] = baseApi.data
	} else {
		baseApi.Data["json"] = baseApi.data
	}

}

func processTimeTooLongWarn(baseApi *ApiController) {
	endTime := time.Now().Unix()
	requestTime := endTime - baseApi.startTime
	if baseApi.startTime > 0 && requestTime > 3000 {
	}
}

func (baseApi *ApiController) ParseForm(form interface{}) bool {
	err := json.Unmarshal(baseApi.Ctx.Input.RequestBody, form)
	if err != nil {
		return false
	}
	err = validForm(form)
	if err != nil {
		baseApi.err = err
		return false
	}
	return true
}

func validForm(form interface{}) error {
	return nil
}

func (baseApi *ApiController) validGetForm(form interface{}) bool {
	err := validForm(form)
	if err != nil {
		baseApi.err = err
		return false
	}
	return true
}
