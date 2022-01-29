package logger

import (
	"net/http"
	"time"

	LOG "github.com/sirupsen/logrus"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	LOG.Infof("%s %s %v", r.Method, r.RequestURI, time.Since(start))
}

func LoggerMiddleware(handleToWrap http.Handler) *Logger {
	return &Logger{handleToWrap}
}
