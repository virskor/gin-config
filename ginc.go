package ginc

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mingzhehao/goutils/filetool"
	"github.com/tidwall/gjson"
)

type Config struct {
}

type ConfigOptions struct {
	Key  string
	File string
}

// Get configrations file, this mothod return map interface
func (this *Config) Get(conf *ConfigOptions) (string, error) {

	env := gin.Mode()

	if conf.File == "" {
		conf.File = env
	}

	path := "configs/app/" + conf.File + ".json"

	json, err := filetool.ReadFileToString(path)
	if err != nil {
		return "", err
	}

	find := gjson.Get(json, conf.Key)
	if find.String() == "" {
		return "", errors.New("bad key: " + conf.Key)
	}

	return find.String(), nil
}
