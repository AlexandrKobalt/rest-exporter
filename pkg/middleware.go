package pkg

import (
	"net/http"
	"time"

	"gitlab.smd.local/backend/memes.git/internal/stats"
)

func StatsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		endTime := time.Now()

		processTime := endTime.Sub(startTime)

		processData := stats.ProcessData{
			RequestProcessTime: processTime,
		}

		stats.UpdateStat(r.URL.Path, processData)
	})
}
