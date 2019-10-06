package model

type TsUser struct {
	Userid   int    `json:"userid" orm:"pk"`
	Pwd      string `json:"pwd"`
	Salt     string `json:"salt"`
	Email    string `json:"email"`
	Resetpwd string `json:"resetpwd"`
	Code     string `json:"code"`
}

