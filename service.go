package restexporter

import (
	"time"
)

var routeStats map[string]*RouteStats
var initTime time.Time

func init() {
	routeStats = make(map[string]*RouteStats)
	initTime = time.Now()
}

func GetCurrentRouteStats() map[string]RouteStats {
	destMap := make(map[string]RouteStats)

	for key, value := range routeStats {
		destMap[key] = *value
	}

	return destMap
}

func getStats(url string) *RouteStats {
	stats, exists := routeStats[url]
	if !exists {
		stats = &RouteStats{}
		routeStats[url] = stats
	}

	return stats
}

func UpdateStat(url string, processData ProcessData) {
	stats := getStats(url)
	stats.Update(processData)
}

func (stats *RouteStats) Update(processData ProcessData) {
	stats.updateTotalRequestsCount()
	stats.updateLastRequestTime()
	stats.updateRequestsFrequency()
	stats.updateAverageProcessTime(processData.RequestProcessTime)
}

func (stats *RouteStats) updateTotalRequestsCount() {
	stats.TotalRequestsCount++
}

func (stats *RouteStats) updateRequestsFrequency() {
	upTime := stats.LastRequestTime.Sub(initTime).Seconds()
	stats.RequestsFrequency = float64(stats.TotalRequestsCount) / upTime
}

func (stats *RouteStats) updateAverageProcessTime(value time.Duration) {
	totalRequests := float64(stats.TotalRequestsCount)
	totalProcessTime := float64(stats.AverageProcessTime * time.Duration(stats.TotalRequestsCount))

	totalProcessTime += float64(value)

	stats.AverageProcessTime = time.Duration(totalProcessTime / totalRequests)
	stats.TotalRequestsCount++
}

func (stats *RouteStats) updateLastRequestTime() {
	stats.LastRequestTime = time.Now()
}
