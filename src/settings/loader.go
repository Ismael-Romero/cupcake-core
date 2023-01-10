package settings

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func Loader() *Settings {
	file, err := os.Open(configurationFile)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	var pointSetting Settings
	str := string(bytes)
	err = json.Unmarshal([]byte(str), &pointSetting)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &pointSetting
}
