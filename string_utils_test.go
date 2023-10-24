package gqldefs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstN(t *testing.T) {
	assert.Equal(t, "TO", FirstN("TOXXE", 2))
	assert.Equal(t, "T", FirstN("T", 2))
	assert.Equal(t, "", FirstN("", 2))
}

func TestLastN(t *testing.T) {
	assert.Equal(t, "XYZ", LastN("TOXYZ", 3))
	assert.Equal(t, "XYZ", LastN("XYZ", 3))
	assert.Equal(t, "YZ", LastN("YZ", 3))
	assert.Equal(t, "", LastN("", 3))
}
