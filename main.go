package main

import (
	_ "multilang_browsing/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

func main() {
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()

}

