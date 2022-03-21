package controllers

import (
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type AboutusController struct {
	beego.Controller  // 1
	i18n.Locale  // 2
	
  }
  func (aboutus *AboutusController) Get(){
	val:=  aboutus.Ctx.Input.Param(":value")
	fmt.Println(val)

	if(val==""){
	c:=strings.Split(aboutus.Ctx.Request.Header.Get("Accept-Language"),",")
	aboutus.Data["langTemplateKey"] =c[0]
	// fmt.Println(aboutus.Data["langTemplateKey"] )
	aboutus.TplName = "aboutus.tpl" 
  }else{
	// c:=strings.Split(aboutus.Ctx.Request.Header.Get("Accept-Language"),",")
	aboutus.Data["langTemplateKey"] =val
	// fmt.Println(aboutus.Data["langTemplateKey"] )
	aboutus.TplName = "aboutus.tpl" 
  }
}