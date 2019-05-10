package main

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World!"))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	lr := io.LimitReader(r.Body, 1024*1024)
	w.WriteHeader(201)
	io.Copy(w, lr)
}

func rootHandler(l logrus.FieldLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l.WithFields(logrus.Fields{
			"method":     r.Method,
			"URL":        r.RequestURI,
			"Protocol":   r.Proto,
			"RemoteAddr": r.RemoteAddr,
		}).Info("Incoming Request")
		switch r.Method {
		case "POST":
			echoHandler(w, r)
		default:
			helloWorldHandler(w, r)
		}
	}
}

func main() {
	l := logrus.StandardLogger()
	l.Debug("Logger started.")
	http.HandleFunc("/", rootHandler(l))
	l.Info(http.ListenAndServe(":8080", nil))
}
