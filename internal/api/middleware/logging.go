package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		log.Printf(
			"[API] request method=%s path=%s",
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(rw, r)

		log.Printf(
			"[API] response method=%s path=%s status=%d duration=%s",
			r.Method,
			r.URL.Path,
			rw.status,
			time.Since(start),
		)
	})
}
