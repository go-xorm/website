// Copyright 2013 Beego Web authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package routers implemented controller methods of beego.
package actions

import (
	"strings"
	"time"

	"github.com/Unknwon/i18n"
	"github.com/go-xweb/log"
	"github.com/go-xweb/xweb"
)

var (
	AppVer string
	IsPro  bool
)

var langTypes []*langType // Languages are supported.

// langType represents a language type.
type langType struct {
	Lang, Name string
}

// baseRouter implemented global settings for all other routers.
type baseAction struct {
	*xweb.Action
	i18n.Locale

	start time.Time
}

// Prepare implemented Prepare method for baseRouter.
func (this *baseAction) Init() {
	// Setting properties.
	this.AddTmplVars(&xweb.T{
		"AppVer":        AppVer,
		"IsPro":         IsPro,
		"PageStartTime": time.Now(),
		"loadtimes": func() int64 {
			return int64(time.Now().Sub(this.start) / time.Millisecond)
		},
	})

	this.start = time.Now()

	// Redirect to make URL clean.
	if this.setLangVer() {
		i := strings.Index(this.Request.RequestURI, "?")
		this.Redirect(this.Request.RequestURI[:i], 302)
		return
	}
}

// setLangVer sets site language version.
func (this *baseAction) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.GetString("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		cookie, _ := this.GetCookie("lang")
		if cookie != nil {
			lang = cookie.String()
			hasCookie = true
		}
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		cookie := xweb.NewCookie("lang", curLang.Lang, 1<<31-1)
		this.SetCookie(cookie)
	}

	log.Info(langTypes)

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	this.AddTmplVars(&xweb.T{
		"Lang":      curLang.Lang,
		"CurLang":   curLang.Name,
		"RestLangs": restLangs,
	})
	this.Lang = lang

	return isNeedRedir
}
