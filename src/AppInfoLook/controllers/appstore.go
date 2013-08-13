package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

const (
	AppStoreURL = "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewTop?genreId=6016&id=29091&popId=44"
)
const (
	ErrGetDataFromStore    = `{"encode":900,"msg":"get data err form appstore"}`
	ErrDecodeDataFromStore = `{"encode":901,"msg":"decode data err form appstore"}`
)

type AppStoreController struct {
	beego.Controller
}

func (this *AppStoreController) Get() {
	resp, err := http.Get(AppStoreURL)
	if err != nil {
		this.Ctx.WriteString(ErrGetDataFromStore)
		return
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			this.Ctx.WriteString(ErrDecodeDataFromStore)
			return
		}
		fmt.Println(string(body))
		this.Ctx.WriteString(string(body))

	}
}
