package middleware

import (
	"log"
	"log/slog"
	"net/http"
	"time"
)


func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Println(r.Method, time.Since(start))
        log.Printf("hello form %s request", r.URL.Path)
		slog.Info("Thank you dear")
	})
}