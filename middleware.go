package main

import (
	"net/http"
	"time"
)

func StatsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		endTime := time.Now()

		processTime := endTime.Sub(startTime)

		processData := ProcessData{
			RequestProcessTime: processTime,
		}

		UpdateStat(r.URL.Path, processData)
	})
}
