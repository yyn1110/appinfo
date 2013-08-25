package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type ProjectInfo struct {
	Name string
}
type ResponseInfo struct {
	ResultCode int
	Msg        string
	Results    map[string][]*ProjectInfo
	Timeline   int64
}
type ListController struct {
	beego.Controller
}

func init() {

}
func (this *ListController) Get() {
	this.Ctx.WriteString(RequestRootDef)
}
func (this *ListController) initInfo() *ResponseInfo {
	data := []string{"企业管理", "学校管理", "社会管理", "社区管理", "国家管理"}
	response := make(map[string][]*ProjectInfo)
	rooms := make([]*ProjectInfo, len(data))
	for i := 0; i < len(data); i++ {
		r := new(ProjectInfo)
		r.Name = data[i]
		rooms[i] = r
	}
	response["经济"] = rooms
	resinfo := &ResponseInfo{
		ResultCode: 1,
		Msg:        "hello",
		Timeline:   time.Now().Unix(),
		Results:    response,
	}
	return resinfo

}
func (this *ListController) Post() {
	//id := this.Input().Get("uid")
	resinfo := this.initInfo()
	res, err := json.Marshal(resinfo)

	if err != nil {
		this.Ctx.WriteString(RequestRootDef)
		fmt.Println(string(res))
		return
	}
	this.Ctx.WriteString(string(res))
}
