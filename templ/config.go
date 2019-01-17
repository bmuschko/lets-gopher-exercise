package templ

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ConfigHomeDir = filepath.Join(homeDir(), ".lets-gopher")

type GenConfig struct {
	Domain string `yaml:",omitempty"`
}

func Load() (config GenConfig, err error) {
	configFile := filepath.Join(ConfigHomeDir, ".lets-gopher.yml")
	f, err := os.Open(configFile)
	if err != nil {
		return
	}
	return LoadReader(f)
}

func LoadReader(fd io.Reader) (config GenConfig, err error) {
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		return config, err
	}
	err = yaml.UnmarshalStrict(data, &config)
	return config, err
}
