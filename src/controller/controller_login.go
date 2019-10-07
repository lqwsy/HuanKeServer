package controller

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/wonderivan/logger"
	"HuanKeServer/src/model"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Code     string `json:"code"`
}

func (this *LoginController) Post() {
	var registerRequest RegisterRequestBody
	var resultRsp = HelloResponse{}
	resultRsp.Code = 100

	request := this.Ctx.Input.RequestBody
	err := json.Unmarshal(request, &registerRequest)
	if err != nil {
		logger.Error(err)
		resultRsp.Code = 102
		resultRsp.ExtraMsg = "request param error!"
		this.Data["json"] = &resultRsp
		this.ServeJSON()
		return
	}

	var user model.TsUser
	o := orm.NewOrm()

	user.Email = registerRequest.Username
	err = o.Read(&user, "email")
	if err != nil && err != orm.ErrNoRows {
		logger.Error(err)
		resultRsp.Code = 103
		resultRsp.ExtraMsg = "database error : query user error!"
		this.Data["json"] = &resultRsp
		this.ServeJSON()
		return
	}

	if err == orm.ErrNoRows || user.Userid == 0 {
		logger.Debug("login email:" + registerRequest.Username + " doesn't exist!")
		resultRsp.Code = 104
		resultRsp.ExtraMsg = "login error : email doesn't exist!"
		this.Data["json"] = &resultRsp
		this.ServeJSON()
		return
	}



	resultRsp.Code = 101
	resultRsp.ExtraMsg = "register success!"
	this.Data["json"] = &resultRsp
	this.ServeJSON()
}
