package faker_test

import (
	"github.com/hbttundar/diabuddy_testkit/helpers/faker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomInt(t *testing.T) {
	val := faker.RandomInt(5, 10)
	assert.GreaterOrEqual(t, val, 5)
	assert.LessOrEqual(t, val, 10)
}

func TestRandomFloat(t *testing.T) {
	val := faker.RandomFloat(1.5, 3.0)
	assert.GreaterOrEqual(t, val, 1.5)
	assert.LessOrEqual(t, val, 3.0)
}

func TestRandomIntSlice(t *testing.T) {
	slice := faker.RandomIntSlice(5, 1, 100)
	assert.Len(t, slice, 5)
	for _, val := range slice {
		assert.GreaterOrEqual(t, val, 1)
		assert.LessOrEqual(t, val, 100)
	}
}
