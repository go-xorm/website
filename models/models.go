package models

import (
	"errors"
	"os"

	"github.com/Unknwon/com"
	"github.com/Unknwon/goconfig"
	"github.com/go-xweb/log"
	"github.com/slene/blackfriday"
)

const (
	_CFG_PATH        = "conf/app.ini"
	_CFG_CUSTOM_PATH = "conf/custom.ini"
)

var (
	Cfg *goconfig.ConfigFile
)

func InitModels() {
	var err error
	Cfg, err = goconfig.LoadConfigFile(_CFG_PATH)
	if err != nil {
		log.Fatalf("Fail to load config file(%s): %v", _CFG_PATH, err)
	}
	if com.IsFile(_CFG_CUSTOM_PATH) {
		if err = Cfg.AppendFiles(_CFG_CUSTOM_PATH); err != nil {
			log.Fatalf("Fail to load config file(%s): %v", _CFG_CUSTOM_PATH, err)
		}
	}
}

// loadFile returns []byte of file data by given path.
func loadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return []byte(""), errors.New("Fail to open file: " + err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		return []byte(""), errors.New("Fail to get file information: " + err.Error())
	}

	d := make([]byte, fi.Size())
	f.Read(d)
	return d, nil
}

func markdown(raw []byte) []byte {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	htmlFlags |= blackfriday.HTML_GITHUB_BLOCKCODE
	htmlFlags |= blackfriday.HTML_OMIT_CONTENTS
	htmlFlags |= blackfriday.HTML_COMPLETE_PAGE
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	// set up the parser
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	body := blackfriday.Markdown(raw, renderer, extensions)
	return body
}
