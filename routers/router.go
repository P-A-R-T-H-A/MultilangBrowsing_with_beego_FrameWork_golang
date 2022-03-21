package routers

import (
	"fmt"
	
	"multilang_browsing/controllers"
	
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

func init() {
	

	
    beego.Router("/home/?:value", &controllers.MainController{})

	beego.Router("/home/?:value/aboutus", &controllers.AboutusController{})
	//
	langs, err := beego.AppConfig.String("langs")  // 1
if err != nil {  // 2
  fmt.Println("Failed to load languages from the app.conf")
  return
}
langsArr := strings.Split(langs, "|")  // 3
for _, lang := range langsArr {  // 4
  if err := i18n.SetMessage(lang, "conf/"+lang+".ini"); err != nil {  // 5
    fmt.Println("Failed to set message file for l10n")
    return
  }
}
}
