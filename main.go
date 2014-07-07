package main

import (
	"github.com/Unknwon/i18n"
	"github.com/go-xweb/log"
	"github.com/go-xweb/xweb"

	"github.com/go-xorm/website/actions"
	"github.com/go-xorm/website/models"
)

const (
	APP_VER = "0.1.0627"
)

func main() {
	models.InitModels()

	/*mode, _ := models.Cfg.GetValue("mode", "debug")
	if mode == "debug" {*/
	log.SetOutputLevel(log.Ldebug)
	//}

	actions.InitApp()

	// Register routers.
	xweb.AddAction(&actions.HomeAction{})
	xweb.AutoAction(&actions.DocsAction{})
	xweb.AddTmplVar("i18n", i18n.Tr)
	xweb.Run(":9999")
}
