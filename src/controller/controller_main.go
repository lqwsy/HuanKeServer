package controller

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

type HelloResponse struct {
	Code     int    `json:"code"`
	Data     string `json:"data"`
	ExtraMsg string `json:"extra_msg"`
}

func (this *IndexController) Post() {

	var resultRsp = HelloResponse{}
	resultRsp.Code = 101
	resultRsp.ExtraMsg = "Welcome to Huanke!"
	this.Data["json"] = &resultRsp
	this.ServeJSON()
}
