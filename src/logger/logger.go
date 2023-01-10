package logger

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	ext        = ".json"
	flagCreate = os.O_CREATE | os.O_RDWR
	flagOpen   = os.O_RDWR
)

var payload []interface{}

// RestLogger defines the methods to perform server logging
type RestLogger interface {
	Backup()
	Info(id, message string)
	Error(id, message string)
	Warning(id, message string)
}

// HTTPRequestLogger defines methods for logging HTTP requests
type HTTPRequestLogger interface {
	Backup()
	Response200(r *http.Request, id, message string)
	Response404(r *http.Request, id, message string)
	Response500(r *http.Request, id, message string)
}

type Logger struct {
	name       string
	path       string
	pathOrigin string
	backupTime time.Duration
	dateFormat string
	data       []interface{}
	mutex      *sync.RWMutex
}

/*
New creates a new Logger structure and opens a log file in the specified path.
The function also initializes the fields of the Logger
structure and sets a timer for backing up the log file.
*/
func New(path, name string, time time.Duration, mtx *sync.RWMutex) *Logger {

	createLogFolders(path)
	_, err := openLogFile(parserPath(path, name), flagCreate)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &Logger{
		name:       name,
		path:       parserPath(path, name),
		pathOrigin: path,
		backupTime: time,
		dateFormat: "2006-01-02 15:04:05",
		data:       payload,
		mutex:      mtx,
	}
}

/*
encoder encodes the input data in JSON format and
writes it to a log file at the specified path.
The function also uses a mutex to protect concurrent access to the log file.
*/
func encoder(l *Logger, content interface{}) {
	l.mutex.Lock()
	l.data = append(l.data, content)
	file, err := openLogFile(l.path, flagOpen)

	if err != nil {
		log.Println(err.Error())
		return
	}

	_encoder := json.NewEncoder(file)
	err = _encoder.Encode(l.data)

	if err != nil {
		log.Println(err.Error())
		return
	}
	err = file.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	l.mutex.Unlock()
}

/*
createLogFolders creates a folder and all intermediate folders
necessary to complete the specified path.
If the folder already exists, no action is taken.
If an error occurs when creating the folder, an error message is printed on the screen.
*/
func createLogFolders(path string) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, 0766)
		if err != nil {
			log.Println("Error creating folder: ", path, err.Error())
			return
		}
	}
}

/*
openLogFile opens a log file or creates a log file depending on the flag passed as a parameter
if an error occurs when opening the file, a null file pointer and the error are returned.
*/
func openLogFile(path string, flag int) (file *os.File, err error) {
	file, err = os.OpenFile(path, flag, 0766)
	if err != nil {
		return nil, err
	}
	return file, nil
}

/*
parserPath returns the full path of a file given its name and a folder path.
If the folder path is "./", the file is assumed to be in the current folder.
If the folder path is not "./", the file is assumed to be in a subfolder of the current folder.
*/
func parserPath(path, name string) string {
	if path == "./" {
		return path + name + ext
	}
	return path + "/" + name + ext
}
