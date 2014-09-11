package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/Unknwon/i18n"
	"github.com/go-xweb/log"
	"github.com/go-xweb/xweb"

	"github.com/go-xorm/website/actions"
	"github.com/go-xorm/website/models"
)

const (
	APP_VER = "0.2.0715"
)

func updateDoc(tmpPath, dstpath string) {
	os.RemoveAll(tmpPath)
	os.MkdirAll(tmpPath, os.ModePerm)

	resp, err := http.Get("https://github.com/go-xorm/xorm/archive/master.zip")
	if err != nil {
		log.Warn("download docs failed")
		return
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn("read zip docs failed")
		return
	}

	zrd, err := zip.NewReader(bytes.NewReader(bs), int64(len(bs)))
	if err != nil {
		log.Warn("open zip docs failed")
		return
	}

	for _, f := range zrd.File {
		if !strings.HasPrefix(f.Name, "xorm-master/docs") {
			continue
		}

		fPath := filepath.Join(tmpPath, f.Name)
		os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
		fInfo := f.FileInfo()
		if fInfo.IsDir() {
			os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		ff, err := os.Create(fPath)
		if err != nil {
			log.Warnf("create file %s failed", f.Name)
			return
		}
		defer ff.Close()

		rd, err := f.Open()
		if err != nil {
			log.Warnf("open zip file %s failed", f.Name)
			return
		}

		_, err = io.Copy(ff, rd)
		if err != nil {
			log.Warnf("unzip file %s failed", f.Name)
			return
		}
	}

	os.RemoveAll(dstpath)

	err = os.Rename(filepath.Join(tmpPath, "xorm-master/docs"), dstpath)
	if err != nil {
		log.Warnf("move docs from %s to %s failed", filepath.Join(tmpPath, "xorm-master/docs"), dstpath)
	}
}

func timeUpdate(tmpPath, dstpath string) {
	time.AfterFunc(time.Minute, func() {
		updateDoc(tmpPath, dstpath)
		timeUpdate(tmpPath, dstpath)
	})
}

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

	go timeUpdate("./tmp", "./docs")

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
