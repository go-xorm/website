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
	"github.com/go-xweb/xweb"

	"github.com/go-xorm/website/models"
)

// DocsRouter serves about page.
type DocsAction struct {
	baseAction

	get xweb.Mapper `xweb:"/(.*)"`
}

// Get implemented Get method for DocsRouter.
func (this *DocsAction) Get(link string) error {
	dRoot := models.GetDocByLocale(this.Lang)
	if dRoot == nil {
		return this.NotFound("Not Found Your Language")
	}

	var doc *models.DocNode
	if len(link) == 0 {
		if dRoot.Doc.HasContent() {
			doc = dRoot.Doc
		} else {
			return this.Redirect("/docs/intro/", 302)
		}
	} else {
		doc, _ = dRoot.GetNodeByLink(link)
		if doc == nil {
			doc, _ = dRoot.GetNodeByLink(link + "/")
			/*if doc != nil {
				fmt.Println(link, doc)
				return this.Redirect("/docs/"+link+"/", 301)
			}*/
		}
	}

	if doc == nil {
		return this.NotFound("Not Found Doc")
	}

	return this.Render("docs.html", &xweb.T{
		"IsDocs":  true,
		"DocRoot": dRoot,
		"Doc":     doc,
		"Title":   doc.Name,
		"Data":    doc.GetContent(),
	})
}

/*
func DocsStatic(ctx *context.Context) {
	if uri := ctx.Input.Params[":all"]; len(uri) > 0 {
		lang := ctx.GetCookie("lang")
		if !i18n.IsExist(lang) {
			lang = "en-US"
		}

		f, err := os.Open("docs/images/" + uri)
		if err != nil {
			ctx.WriteString(err.Error())
			return
		}
		defer f.Close()

		_, err = io.Copy(ctx.ResponseWriter, f)
		if err != nil {
			ctx.WriteString(err.Error())
			return
		}
	}
}
*/
