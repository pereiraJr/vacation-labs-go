package test

import (
	"testing"
	"gotest.tools/assert"
)


func TestBasic(t *testing.T) {
	assert.DeepEqual(t, 1, 1)
}