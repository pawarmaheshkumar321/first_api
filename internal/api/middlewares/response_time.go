package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received Request in ResponseTime")
		start := time.Now()

		//create a  custom ResponseWriter to capture the status code
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		//calculate Duration
		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())

		next.ServeHTTP(wrappedWriter, r)

		duration = time.Since(start)
		//Log the request Details
		fmt.Printf("Method : %v, URL : %v, Status : %v, Duration : %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response From Response Time Middleware")

	})
}

// responseWriter
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
