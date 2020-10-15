package config

import (
	"encoding/json"
	"github.com/bhupeshpandey/cars/model"
	"io/ioutil"
)

func ReadConfig() (*model.AppConfig, error) {

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return nil, err
	}

	var appConfig *model.AppConfig
	err = json.Unmarshal(file, appConfig)
	if err != nil {
		return nil, err
	}
	return appConfig, nil
}