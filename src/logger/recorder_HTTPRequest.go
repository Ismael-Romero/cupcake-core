package logger

import (
	"net/http"
	"time"
)

type httpRequestLogData struct {
	Id      string          `json:"id,omitempty"`
	Type    string          `json:"type,omitempty"`
	Level   string          `json:"level,omitempty"`
	Date    string          `json:"date,omitempty"`
	Message string          `json:"message,omitempty"`
	Body    httpRequestBody `json:"body,omitempty"`
}

type httpRequestBody struct {
	From     string `json:"from,omitempty"`
	Method   string `json:"method,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	URL      string `json:"URL,omitempty"`
	Status   int16  `json:"status,omitempty"`
}

func (l *Logger) Response200(r *http.Request, id, m string) {
	content := httpRequestLogData{
		Id:      id,
		Type:    "response",
		Level:   "info",
		Date:    time.Now().Format(l.dateFormat),
		Message: m,
		Body: httpRequestBody{
			From:     r.RemoteAddr,
			Method:   r.Method,
			Protocol: r.Proto,
			URL:      r.RequestURI,
			Status:   200,
		},
	}
	encoder(l, content)
}

func (l *Logger) Response404(r *http.Request, id, m string) {
	content := httpRequestLogData{
		Id:      id,
		Type:    "response",
		Level:   "error",
		Date:    time.Now().Format(l.dateFormat),
		Message: m,
		Body: httpRequestBody{
			From:     r.RemoteAddr,
			Method:   r.Method,
			Protocol: r.Proto,
			URL:      r.RequestURI,
			Status:   404,
		},
	}
	encoder(l, content)
}

func (l *Logger) Response500(r *http.Request, id, m string) {
	content := httpRequestLogData{
		Id:      id,
		Type:    "response",
		Level:   "error",
		Date:    time.Now().Format(l.dateFormat),
		Message: m,
		Body: httpRequestBody{
			From:     r.RemoteAddr,
			Method:   r.Method,
			Protocol: r.Proto,
			URL:      r.RequestURI,
			Status:   500,
		},
	}
	encoder(l, content)
}
