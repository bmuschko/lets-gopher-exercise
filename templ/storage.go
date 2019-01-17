package templ

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bmuschko/lets-gopher-exercise/utils"
	"github.com/mitchellh/go-homedir"
)

var TemplateDir = filepath.Join(ConfigHomeDir, "templates")

func homeDir() string {
	homeDir, err := homedir.Dir()
	utils.CheckIfError(err)
	return homeDir
}

func ListTemplates() map[string]string {
	var templates = make(map[string]string)

	if _, err := os.Stat(TemplateDir); os.IsNotExist(err) {
		err := utils.CreateDir(TemplateDir)
		utils.CheckIfError(err)
	} else {
		files, err := ioutil.ReadDir(TemplateDir)
		utils.CheckIfError(err)
		for _, f := range files {
			if f.IsDir() {
				templates[f.Name()] = filepath.Join(TemplateDir, f.Name())
			}
		}
	}

	return templates
}
