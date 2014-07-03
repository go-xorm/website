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

package actions

import (
	"github.com/go-xorm/website/models"
	"github.com/go-xweb/xweb"
)

type HomeAction struct {
	baseAction

	get        xweb.Mapper `xweb:"/"`
	about      xweb.Mapper
	page       xweb.Mapper
	quickStart xweb.Mapper
	donate     xweb.Mapper
}

// Get implemented Get method for HomeRouter.
func (this *HomeAction) Get() error {
	return this.Render("home.html", &xweb.T{
		"IsHome": true,
	})
}

func (this *HomeAction) About() error {
	// Get language.
	df := models.GetDoc("about", this.Lang)
	return this.Render("about.html", &xweb.T{
		"IsAbout": true,
		"Title":   df.Title,
		"Data":    string(df.Data),
		"Section": "about",
	})
}

func (this *HomeAction) Page() error {
	// Get language.
	df := models.GetDoc("team", this.Lang)
	return this.Render("team.html", &xweb.T{
		"IsTeam":  true,
		"Title":   df.Title,
		"Data":    string(df.Data),
		"Section": "team",
	})
}

func (this *HomeAction) QuickStart() error {
	df := models.GetDoc("quickstart", this.Lang)
	return this.Render("quickstart.html", &xweb.T{
		"IsQuickStart":  true,
		"Section":       "quickstart",
		"Title":         df.Title,
		"Data":          string(df.Data),
		"IsHasMarkdown": true,
	})
}

func (this *HomeAction) Donate() error {
	// Get language.
	df := models.GetDoc("donate", this.Lang)

	return this.Render("donate.html", &xweb.T{
		"IsDonate":      true,
		"Title":         df.Title,
		"Data":          string(df.Data),
		"IsHasMarkdown": true,
	})
}
