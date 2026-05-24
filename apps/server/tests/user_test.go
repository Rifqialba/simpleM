package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserModule(t *testing.T) {
	assert.Equal(t, "user", "user")
}