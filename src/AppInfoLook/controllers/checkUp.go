package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UpdateController struct {
	beego.Controller
}
type checkinfo struct {
	IsUpdate bool
	Msg      string
}
type checkResponse struct {
	ResultCode int
	ResultMsg  string
	Results    map[string]*checkinfo
	Timeline   int64
}

var responseMap map[string]*checkinfo

func init() {
	responseMap = make(map[string]*checkinfo)
}
func (this *UpdateController) initInfo() *checkResponse {
	info := &checkinfo{
		IsUpdate: false,
		Msg:      "增加了高等数学",
	}
	responseMap["data"] = info
	resinfo := &checkResponse{
		ResultCode: 1,
		ResultMsg:  "ok",
		Timeline:   time.Now().Unix(),
		Results:    responseMap,
	}
	return resinfo
}
func (this *UpdateController) Get() {
	this.Ctx.WriteString(RequestRootDef)
}
func (this *UpdateController) Post() {
	//id := this.Input().Get("uid")
	str := this.Ctx.Request.Form.Get("DeviceSystemName")
	fmt.Println(str)
	resinfo := this.initInfo()
	res, err := json.Marshal(resinfo)

	if err != nil {
		this.Ctx.WriteString(RequestRootDef)
		fmt.Println(string(res))
		return
	}
	this.Ctx.WriteString(string(res))
}
