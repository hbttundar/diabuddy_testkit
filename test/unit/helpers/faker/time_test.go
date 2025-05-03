package faker_test

import (
	"github.com/hbttundar/diabuddy_testkit/helpers/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRandomPastTime(t *testing.T) {
	past := faker.RandomPastTime(24 * time.Hour)
	assert.WithinDuration(t, time.Now(), past, 24*time.Hour)
	assert.True(t, past.Before(time.Now()))
}

func TestRandomFutureTime(t *testing.T) {
	future := faker.RandomFutureTime(48 * time.Hour)
	assert.WithinDuration(t, time.Now().Add(48*time.Hour), future, 48*time.Hour)
	assert.True(t, future.After(time.Now()))
}

func TestRandomTimeRange(t *testing.T) {
	minGap := 2 * time.Hour
	maxGap := 6 * time.Hour
	start, end := faker.RandomTimeRange(minGap, maxGap)
	diff := end.Sub(start)
	assert.True(t, end.After(start))
	assert.GreaterOrEqual(t, diff, minGap)
	assert.LessOrEqual(t, diff, maxGap)
}
