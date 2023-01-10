package logger

import (
	"log"
	"time"
)

/*
Backup backs up the log file every time the time interval specified
in the backupTime field of the Logger structure elapses.
The backup is created with the current name in the format "YYYYY-MM-DD_HH-MM-SS-file_name".
The function also uses a mutex to protect concurrent access
to the log file and the backup file.
*/
func (l *Logger) Backup() {
	for {
		l.mutex.Lock()
		now := time.Now().Format("2006-01-02_15-04-05")
		nameLog := now + "-" + l.name
		_, err := openLogFile(parserPath(l.pathOrigin, nameLog), flagCreate)

		if err != nil {
			log.Println("Failed backup>", err.Error())
			return
		}

		l.path = parserPath(l.pathOrigin, nameLog)
		l.mutex.Unlock()
		time.Sleep(l.backupTime)
	}
}
