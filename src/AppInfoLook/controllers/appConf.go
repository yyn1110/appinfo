package controllers

const (
	AppleCNLookURL = "http://itunes.apple.com/cn/lookup?id="
	AppleENLookURL = "http://itunes.apple.com/lookup?id="
)
const (
	RequestDef    = `{"resultCount":0,"results": [],"msg":"无数据"}`
	RequestErr    = `{"resultCount":900,"results": [],"msg":"请求错误"}`
	DecodeBodyErr = `{"resultCount":901,"results": [],"msg":"解析数据错误"}`
	DecodeErr     = `{"resultCount":902,"results": [],"msg":"解析json错误"}`
	AppIDErr      = `{"resultCount":903,"results": [],"msg":"AppId必须是数字"}`
)
