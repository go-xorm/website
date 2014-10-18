package actions

import "github.com/go-xweb/xweb"

type HomeAction struct {
	baseAction

	get        xweb.Mapper `xweb:"/"`
	about      xweb.Mapper
	team       xweb.Mapper
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
	return this.Render("about.html", &xweb.T{
		"IsAbout": true,
		"Section": "about",
	})
}

func (this *HomeAction) Team() error {
	return this.Render("team.html", &xweb.T{
		"IsTeam":  true,
		"Section": "team",
	})

}

func (this *HomeAction) Donate() error {
	return this.Render("donate.html", &xweb.T{
		"IsDonate":      true,
		"IsHasMarkdown": true,
	})
}
