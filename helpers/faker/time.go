package faker

import (
	"math/rand"
	"time"
)

var timeRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomPastTime returns a time before now within given duration.
func RandomPastTime(maxAgo time.Duration) time.Time {
	delta := timeRand.Int63n(int64(maxAgo))
	return time.Now().Add(-time.Duration(delta))
}

// RandomFutureTime returns a time after now within given duration.
func RandomFutureTime(maxAhead time.Duration) time.Time {
	delta := timeRand.Int63n(int64(maxAhead))
	return time.Now().Add(time.Duration(delta))
}

// RandomTimeRange returns a random [start, end] pair with min gap.
func RandomTimeRange(minGap, maxGap time.Duration) (time.Time, time.Time) {
	start := RandomPastTime(30 * 24 * time.Hour) // within past month
	gap := time.Duration(timeRand.Int63n(int64(maxGap-minGap))) + minGap
	end := start.Add(gap)
	return start, end
}
