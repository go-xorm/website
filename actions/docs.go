package actions

import (
	"strings"

	"github.com/go-xweb/xweb"
)

// DocsRouter serves about page.
type DocsAction struct {
	baseAction

	get xweb.Mapper `xweb:"/"`
}

func toLower(l string) string {
	return strings.ToLower(l)
}

// Get implemented Get method for DocsRouter.
func (this *DocsAction) Get() error {
	return this.Render("docs.html", &xweb.T{
		"IsDocs":  true,
		"toLower": toLower,
	})
}
