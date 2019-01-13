package templ

import (
	"github.com/bmuschko/lets-gopher/utils"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"path/filepath"
)

var TemplateDir = filepath.Join(ConfigHomeDir, "templates")

func homeDir() string {
	homeDir, err := homedir.Dir()
	utils.CheckIfError(err)
	return homeDir
}

func ListTemplates() map[string]string {
	var templates = make(map[string]string)

	files, err := ioutil.ReadDir(TemplateDir)
	utils.CheckIfError(err)
	for _, f := range files {
		if f.IsDir() {
			templates[f.Name()] = filepath.Join(TemplateDir, f.Name())
		}
	}

	return templates
}