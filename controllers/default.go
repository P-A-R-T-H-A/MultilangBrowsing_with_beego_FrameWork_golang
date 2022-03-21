package controllers

import (
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type MainController struct {
	beego.Controller
	i18n.Locale 
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	val:=  c.Ctx.Input.Param(":value")
	fmt.Println(val)

	
	if(val==""){
	cx:=strings.Split(c.Ctx.Request.Header.Get("Accept-Language"),",")
	c.Data["langTemplateKey"] =cx[0]
	// fmt.Println(c.Data["langTemplateKey"] )
	c.TplName = "index.tpl"
  }else{
	// c:=strings.Split(c.Ctx.Request.Header.Get("Accept-Language"),",")
	c.Data["langTemplateKey"] =val
	// fmt.Println(c.Data["langTemplateKey"] )
	c.TplName = "index.tpl"
  }

}
