package actions

import "github.com/go-xweb/xweb"

// DocsRouter serves about page.
type DocsAction struct {
	baseAction

	get xweb.Mapper `xweb:"/"`
}

// Get implemented Get method for DocsRouter.
func (this *DocsAction) Get() error {
	return this.Render("docs.html", &xweb.T{
		"IsDocs": true,
	})
}
