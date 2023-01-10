package logger

import "time"

type restLogData struct {
	Id      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Level   string `json:"level,omitempty"`
	Date    string `json:"date,omitempty"`
	Message string `json:"message,omitempty"`
}

func (l *Logger) Info(id, m string) {
	content := restLogData{
		Id:      id,
		Type:    "",
		Level:   "info",
		Date:    time.Now().Format(l.dateFormat),
		Message: m,
	}
	encoder(l, content)
}

func (l *Logger) Error(id, m string) {
	content := restLogData{
		Id:      id,
		Type:    "",
		Level:   "error",
		Date:    time.Now().Format(l.dateFormat),
		Message: m,
	}
	encoder(l, content)
}

func (l *Logger) Warning(id, m string) {
	content := restLogData{
		Id:      id,
		Type:    "",
		Level:   "warning",
		Date:    time.Now().Format(l.dateFormat),
		Message: m,
	}
	encoder(l, content)
}
