package actions

import (
	"errors"
	"path/filepath"
	"strings"
	"time"
	"github.com/howeyc/fsnotify"

	"github.com/Unknwon/i18n"
	"github.com/go-xweb/log"
	"github.com/go-xweb/xweb"

	"github.com/go-xorm/website/models"
)

var (
	CompressConfPath = "conf/compress.json"
)

func initLocales() {
	// Initialized language type list.
	langs := strings.Split(models.Cfg.MustValue("lang", "types"), "|")
	names := strings.Split(models.Cfg.MustValue("lang", "names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		log.Debug("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			log.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func initTemplates() {
	xweb.AddTmplVars(&xweb.T{
		"dict":      dict,
		"loadtimes": loadtimes,
	})
}

func InitApp() {
	initTemplates()
	initLocales()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					log.Info(event)

					if err := i18n.ReloadLangs(); err != nil {
						log.Error("Conf Reload: ", err)
					}

					log.Info("Config Reloaded")

				case ".json":
					if event.Name == CompressConfPath {
						log.Info("Beego Compress Reloaded")
					}
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		log.Error(err)
	}
}
