package pkg

import "time"

type RouteStats struct {
	TotalRequestsCount int           `json:"totalRequestsCount"`
	RequestsFrequency  float64       `json:"requestsFrequency"`  // per second
	AverageProcessTime time.Duration `json:"averageProcessTime"` // in ms
	LastRequestTime    time.Time     `json:"lastRequestTime"`
}

type RouteStatsOutput struct {
	TotalRequestsCount string `json:"totalRequestsCount"`
	RequestsFrequency  string `json:"requestsFrequency"`  // per second
	AverageProcessTime string `json:"averageProcessTime"` // in ms
	LastRequestTime    string `json:"lastRequestTime"`
}

type ProcessData struct {
	RequestProcessTime time.Duration
}
