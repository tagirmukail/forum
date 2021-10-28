package forum

import (
	"net/http"
	"strings"
	"time"

	"github.com/tagirmukail/forum/internal/utils/ctxvars"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	reqIDHeaderName = "Request-Id"
	// 4 MB
	contentLengthSizeLimit = 4 << 20
)

type customResponseWriter struct {
	http.ResponseWriter
	ERR error
}

func (s *Service) requestContentLimitMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, contentLengthSizeLimit)

			h.ServeHTTP(w, r)
		})
	}
}

func (s *Service) requestIDMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(reqIDHeaderName)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			r = r.WithContext(ctxvars.WithRequestID(r.Context(), requestID))
			w.Header().Set(reqIDHeaderName, requestID)

			h.ServeHTTP(w, r)
		})
	}
}

func (s *Service) loggingMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := ctxvars.GetRequestID(r.Context())

			level := logrus.InfoLevel
			if strings.Contains(r.URL.Path, "swagger") {
				level = logrus.DebugLevel
			}

			now := time.Now()
			logStartEntry := logrus.WithField("start_time", now).
				WithField("method", r.Method).
				WithField("request_id", requestID).
				WithField("path", r.URL.String())

			logStartEntry.Log(level, "started")

			wr := &customResponseWriter{
				ResponseWriter: w,
			}

			wr.Header().Set(reqIDHeaderName, requestID)

			h.ServeHTTP(wr, r)

			if wr.ERR != nil {
				logrus.
					WithError(wr.ERR).
					WithField("work_time", time.Since(now)).
					WithField("method", r.Method).
					WithField("request_id", requestID).
					WithField("path", r.URL.String()).Error("failed")
				return
			}

			logFinishEntry := logrus.WithField("work_time", time.Since(now)).
				WithField("method", r.Method).
				WithField("request_id", requestID).
				WithField("path", r.URL.String())

			logFinishEntry.Log(level, "finished")
		})
	}
}

func corsHandler() func(handler http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
			reqIDHeaderName, "Origin", "Referer"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))
}
