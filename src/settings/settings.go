package settings

import "os"

type Settings struct {
	Services Services `json:"services"`
	Database Database `json:"database"`
	Logs     Logs     `json:"logs"`
}

type Services struct {
	Rest RestService `json:"rest"`
}

type RestService struct {
	Enable bool     `json:"enable"`
	Host   string   `json:"host"`
	Https  HttpConf `json:"https"`
}

type HttpConf struct {
	Host   string `json:"host"`
	Enable bool   `json:"enable"`
	Cert   string `json:"cert"`
	Key    string `json:"key"`
}

type Database struct {
	Engine DatabaseEngine `json:"engine"`
}

type DatabaseEngine struct {
	MySQL MySQLEngine `json:"mySQL"`
}

type MySQLEngine struct {
	Addr     string `json:"addr"`
	Net      string `json:"net"`
	Scheme   string `json:"scheme"`
	User     string `json:"user"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

type Logs struct {
	HTTPRequest LogSetting `json:"http_request"`
	RestService LogSetting `json:"rest"`
}

type LogSetting struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

const configurationFile = "./.json"
const flagCreate = os.O_CREATE | os.O_APPEND | os.O_RDWR

func (s *Settings) GetDatabaseEngines() DatabaseEngine {
	return s.Database.Engine
}

func (s *Settings) GetServices() *Services {
	return &s.Services
}

func (s *Settings) GetConfigurationLogs() *Logs {
	return &s.Logs
}
