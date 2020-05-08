package controllers

import "errors"

type GlobalErrorController struct {
	ApiController
}

func (c *GlobalErrorController) Error404() {
	c.err = errors.New("请求的服务不存在")
}
