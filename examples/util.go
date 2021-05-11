package examples

import (
	"fmt"
	"github.com/murlokito/ccex/models/ws"
	"time"
)

func PrettyFormat(value float64) string {
	if value > 1000000000 {
		return fmt.Sprintf("%.2fB", value / 1000000000)
	}
	if value > 1000000 {
		return fmt.Sprintf("%.2fM", value / 1000000)
	}
	if value > 1000 {
		return fmt.Sprintf("%.2fK", value / 1000)
	}
	return "%.2f"
}

func GetDeltaTrades(trades []ws.Trade, duration time.Duration, moment time.Time) int {
	var count int
	deltaTime := moment.Add(-duration)
	for _, item := range trades {
		if deltaTime.Before(item.Time) {
			count++
		}
	}
	return count
}

func GetDeltaVol(trades []ws.Trade, duration time.Duration, moment time.Time) float64 {
	var count float64
	deltaTime := moment.Add(-duration)
	for _, item := range trades {
		if deltaTime.Before(item.Time) {
			count += item.Size * item.Price
		}
	}
	return count
}
