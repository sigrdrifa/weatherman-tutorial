package logger

import "fmt"

type Logger struct {
	msgs []string
  logToConsole bool
}

func NewLogger(logToConsole bool) *Logger {
	return &Logger{
		msgs: []string{},
    logToConsole: logToConsole,
	}
}

func (l *Logger) Flush() []string {
  msgsCopy := make([]string, len(l.msgs))
  copy(msgsCopy, l.msgs)
  l.msgs = []string{}
  return msgsCopy
}

func (l *Logger) Debug(msg string) {
  l.msgs = append(l.msgs, msg)
  if l.logToConsole {
    fmt.Printf("[%s] %s \n", "DEBUG", msg)
  }
}

func (l *Logger) Error(err error) {
  l.msgs = append(l.msgs, err.Error())
  if l.logToConsole {
    fmt.Printf("[%s] %s \n", "ERROR", err.Error())
  }
}
