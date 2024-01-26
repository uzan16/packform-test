package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/uzan16/packform-test/helpers"
)

func TestHelpers(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("created_at", ToSnakeCase("createdAt"))
	assert.Equal("begin_with", ToSnakeCase("beginWith"))
	assert.Equal("it_is_snake_case", ToSnakeCase("itIsSnakeCase"))
}
