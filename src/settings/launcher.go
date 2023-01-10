package settings

import (
	"encoding/json"
	"log"
	"os"
)

func Launcher() {
	_, err := os.Stat(configurationFile)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(configurationFile, flagCreate, 0766)
		if err != nil {
			log.Println(err.Error())
			return
		}

		data, err := json.Marshal(setSetting())
		if err != nil {
			log.Println(err.Error())
			return
		}

		if _, err = file.WriteString(string(data)); err != nil {
			log.Println(err.Error())
			return
		}

		if err = file.Close(); err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func setSetting() Settings {
	return Settings{
		Services: Services{
			Rest: RestService{
				Enable: false,
				Host:   "127.0.0.1:8080",
				Https: HttpConf{
					Enable: false,
					Host:   "127.0.0.1:443",
					Cert:   "./tls/cert.pem",
					Key:    "./tls/key.pem",
				},
			},
		},
		Database: Database{
			Engine: DatabaseEngine{
				MySQL: MySQLEngine{
					Addr:     "127.0.0.1:3306",
					Net:      "tcp",
					Scheme:   "scheme",
					User:     "user",
					Password: "password",
					Key:      "key",
				},
			},
		},
		Logs: Logs{
			HTTPRequest: LogSetting{
				Name: "http_request",
				Path: "./logs/http_request",
			},
			RestService: LogSetting{
				Name: "rest",
				Path: "./logs/services/rest",
			},
		},
	}
}
