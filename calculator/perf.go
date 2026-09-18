package calculator

import "time"

// CalculateStats accumulates the cost of Calculate calls. It measures Go work
// only, so comparing TotalNanos against the wall time the caller observes
// isolates the JS/WASM boundary from the calculation itself.
type CalculateStats struct {
	Calls      int
	TotalNanos int64
	MaxNanos   int64
}

var (
	calculateStatsEnabled bool
	calculateStats        CalculateStats
)

// SetCalculateTracking turns Calculate timing on or off. It is off by default so
// the timing calls stay out of the reverse-search sweep.
func SetCalculateTracking(enabled bool) {
	calculateStatsEnabled = enabled
}

func GetCalculateStats() CalculateStats {
	return calculateStats
}

func ResetCalculateStats() {
	calculateStats = CalculateStats{}
}

func recordCalculate(start time.Time) {
	elapsed := time.Since(start).Nanoseconds()

	calculateStats.Calls++
	calculateStats.TotalNanos += elapsed

	if elapsed > calculateStats.MaxNanos {
		calculateStats.MaxNanos = elapsed
	}
}
