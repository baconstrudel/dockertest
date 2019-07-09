package main

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	lr := io.LimitReader(r.Body, 1024*1024)
	w.WriteHeader(201)
	io.Copy(w, lr)
}

func rootHandler(l logrus.FieldLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l.WithFields(logrus.Fields{
			"method":     r.Method,
			"URI":        r.RequestURI,
			"Protocol":   r.Proto,
			"RemoteAddr": r.RemoteAddr,
		}).Info("Incoming Request")
		switch r.Method {
		case "POST":
			echoHandler(w, r)
		default:
			switch {
			case r.RequestURI == "/version":
				w.Write([]byte("latest"))
			case r.RequestURI == "/health-check":
				w.Write([]byte("ok"))
			case r.RequestURI == "/ready":
				w.Write([]byte("I was born ready"))
			default:
				w.Write([]byte("Hello World!"))
			}
		}
	}
}

func main() {
	l := logrus.StandardLogger()
	l.SetLevel(logrus.DebugLevel)
	l.Debug("Logger started.")
	http.HandleFunc("/", rootHandler(l))
	l.Info(http.ListenAndServe("0.0.0.0:8080", nil))
}
