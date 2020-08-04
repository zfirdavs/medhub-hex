package rest

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
}

func NewMiddleware(zaplogger *zap.Logger) *Middleware {
	return &Middleware{logger: zaplogger}
}

func (m *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		m.logger.Info("request", zap.String("method", r.Method), zap.String("route", r.URL.Path), zap.Duration("time", t2.Sub(t1)))
	})
}

func (m *Middleware) ContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers",
			"accept, accept-encoding, authorization, content-type, dnt, origin, user-agent, x-csrftoken, x-requested-with, Access-Control-Allow-Origin")
		w.Header().Set("Access-Control-Allow-Methods",
			"DELETE, GET, OPTIONS, PATCH, POST, PUT, PROPFIND, DELETE, HEAD, COPY, MKCOL")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				m.logger.Error("panic recovered", zap.Any("cause", rec), zap.String("method", r.Method), zap.String("route", r.URL.Path))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
