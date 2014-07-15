package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Unknwon/i18n"
	"github.com/go-xweb/log"
	"github.com/go-xweb/xweb"

	"github.com/go-xorm/website/actions"
	"github.com/go-xorm/website/models"
)

const (
	APP_VER = "0.2.0715"
)

func main() {
	models.InitModels()

	mode, _ := models.Cfg.GetValue("app", "run_mode")
	var isPro bool = true
	if mode == "dev" {
		log.SetOutputLevel(log.Ldebug)
		isPro = false
	}
	log.Info("run in " + mode + " mode")

	actions.InitApp()

	// Register routers.
	xweb.AddAction(&actions.HomeAction{})
	xweb.AutoAction(&actions.DocsAction{}, &actions.LinkAction{})
	xweb.AddTmplVars(&xweb.T{
		"i18n":    i18n.Tr,
		"IsPro":   isPro,
		"AppVer":  APP_VER,
		"XwebVer": xweb.Version,
		"GoVer":   strings.Trim(runtime.Version(), "go"),
	})
	port, _ := models.Cfg.GetValue("app", "http_port")
	usessl, _ := models.Cfg.GetValue("app", "ssl")
	if usessl == "true" {
		tlsCfg, _ := xweb.SimpleTLSConfig("cert.pem", "key.pem")
		xweb.RunTLS(fmt.Sprintf(":%v", port), tlsCfg)
	} else {
		xweb.Run(fmt.Sprintf(":%v", port))
	}
}
