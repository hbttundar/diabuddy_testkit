package factory_test

import (
	"github.com/hbttundar/diabuddy_testkit/factory"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type dummy struct {
	Name string
}

func TestGenerateMany(t *testing.T) {
	items := factory.GenerateMany(3, func(i int) *dummy {
		return &dummy{Name: "user-" + strconv.Itoa(i)}
	})

	assert.Len(t, items, 3)
	assert.Equal(t, "user-0", items[0].Name)
	assert.Equal(t, "user-1", items[1].Name)
	assert.Equal(t, "user-2", items[2].Name)
}
