package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserControllers(t *testing.T) {
	t.Run("dummy test", func(t *testing.T) {
		x := 2 + 2
		assert.Equal(t, x, 4)
	})
}
