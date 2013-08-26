package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

const (
	chemTag = 0   //化学tag
	comTag  = 100 //计算机tag
	phyTag  = 200 //物理tag
	mathTag = 300
)

type ListController struct {
	beego.Controller
}

var response map[string][]*BookInfo

func init() {
	response = make(map[string][]*BookInfo)

}

type BookInfo struct {
	Name     string
	BookIntr string
	BookID   int
}
type ResponseInfo struct {
	ResultCode int
	ResultMsg  string
	Results    map[string][]*BookInfo
	Timeline   int64
}

func (this *ListController) Get() {
	this.Ctx.WriteString(RequestRootDef)
}
func (this *ListController) initMathInfo() {
	data := []string{"高等数学"}
	rooms := make([]*BookInfo, len(data))
	for i := 0; i < len(data); i++ {
		r := new(BookInfo)
		r.Name = data[i]
		rooms[i] = r
		r.BookID = i + mathTag
	}
	response["数学"] = rooms

}
func (this *ListController) initChemistryInfo() {

	data := []string{"无机化学", "有机化学", "物理化学", "化学分析", "结构化学"}
	rooms := make([]*BookInfo, len(data))
	for i := 0; i < len(data); i++ {
		r := new(BookInfo)
		r.Name = data[i]
		rooms[i] = r
		r.BookID = i + chemTag
	}
	response["化学"] = rooms

}
func (this *ListController) initComputerInfo() {

	//URLStr := "http://" + this.Ctx.Request.Host
	data := []string{"redis"}
	//dataurl := []string{"/static/books/redis.pdf"}
	datainfo := []string{"这是一个关于redis的学习手册，在这里你可以轻松的学习关于redis的操作和使用"}
	rooms := make([]*BookInfo, len(data))
	for i := 0; i < len(data); i++ {
		r := new(BookInfo)
		r.Name = data[i]
		r.BookIntr = datainfo[i]
		r.BookID = i + comTag
		//r.BookURL = URLStr + dataurl[i]
		rooms[i] = r
	}
	response["计算机"] = rooms
}

func (this *ListController) initPhysicsInfo() {
	data := []string{"力学",
		"热学",
		"原子物理",
		"电磁学",
		"光学",
		"热力学统计",
		"经典力学",
		"固体物理"}
	rooms := make([]*BookInfo, len(data))
	for i := 0; i < len(data); i++ {
		r := new(BookInfo)
		r.Name = data[i]
		r.BookID = phyTag + i
		rooms[i] = r
	}
	response["物理"] = rooms
}
func (this *ListController) initInfo() *ResponseInfo {
	this.initChemistryInfo()
	this.initPhysicsInfo()
	this.initComputerInfo()
	this.initMathInfo()
	resinfo := &ResponseInfo{
		ResultCode: 1,
		ResultMsg:  "hello",
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
