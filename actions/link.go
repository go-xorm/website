package actions

import (
	"github.com/go-xweb/xweb"
)

type LinkAction struct {
	baseAction

	get xweb.Mapper `xweb:"/"`
}

func (l *LinkAction) Get() error {
	return l.Render("link.html", &xweb.T{
		"IsLink": true,
	})
}
