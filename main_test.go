package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_display(t *testing.T) {
	n := randomNumbers()

	// Test 1: Random number should be greater than 0
	assert.GreaterOrEqual(t, n, 0)

	// Test 2: Random number should be lesser than 1000
	assert.LessOrEqual(t, n, 1000)
}
