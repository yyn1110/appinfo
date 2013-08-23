package controllers

const (
	AppleCNLookURL = "http://itunes.apple.com/cn/lookup?id="
	AppleENLookURL = "http://itunes.apple.com/lookup?id="
)
const (
	RequestDef    = `{"ResultCount":0,"Results": [],"Msg":"无数据"}`
	RequestLgErr  = `{"ResultCount":100,"Results": [],"Msg":"语言设置错误"}`
	RequestErr    = `{"ResultCount":900,"Results": [],"Msg":"请求错误"}`
	DecodeBodyErr = `{"ResultCount":901,"Results": [],"Msg":"解析数据错误"}`
	DecodeErr     = `{"ResultCount":902,"Results": [],"Msg":"解析json错误"}`
	AppIDErr      = `{"ResultCount":903,"Results": [],"Msg":"AppId必须是数字"}`
)
