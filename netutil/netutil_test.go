package netutil

import (
	"github.com/frozzare/go-assert"
	"testing"
)

func TestRandomPort(t *testing.T) {
	p, err := RandomPort()
	assert.Nil(t, err)
	assert.Equal(t, true, p > 0)
}
