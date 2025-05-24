package faker_test

import (
	"github.com/hbttundar/diabuddy_testkit/helpers/faker"
	"github.com/stretchr/testify/assert"
	_ "regexp"
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	value := faker.RandomString(10)
	assert.Len(t, value, 10)
	assert.Regexp(t, `^[a-zA-Z]+$`, value)
}

func TestRandomEmail(t *testing.T) {
	email := faker.RandomEmail()
	assert.True(t, strings.Contains(email, "@"))
	assert.True(t, strings.HasSuffix(email, ".test"))
}

func TestRandomSlug(t *testing.T) {
	slug := faker.RandomSlug("foo")
	assert.True(t, strings.HasPrefix(slug, "foo-"))
	assert.Regexp(t, `^foo-[0-9]+$`, slug)
}
