package controller

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/wonderivan/logger"
	"HuanKeServer/src/model"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

type RegisterRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Code     string `json:"code"`
}

func (this *RegisterController) Post() {
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

	if err == nil && user.Userid != 0 {
		logger.Debug("register email:" + registerRequest.Username + " exist!")
		resultRsp.Code = 104
		resultRsp.ExtraMsg = "register error : email had registered!"
		this.Data["json"] = &resultRsp
		this.ServeJSON()
		return
	}

	user.Email = registerRequest.Username
	user.Pwd = registerRequest.Password
	user.Salt = registerRequest.Salt
	user.Code = registerRequest.Code

	num, err := o.Insert(&user)
	if num == 0 || err != nil {
		if err != nil {
			logger.Error("register email:" + registerRequest.Username + " error:" + err.Error())
		}
		resultRsp.Code = 105
		resultRsp.ExtraMsg = "register error : insert database error!"
		this.Data["json"] = &resultRsp
		this.ServeJSON()
		return
	}

	resultRsp.Code = 101
	resultRsp.ExtraMsg = "register success!"
	this.Data["json"] = &resultRsp
	this.ServeJSON()
}
