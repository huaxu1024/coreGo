package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	res := sub.Do()
	assert.Equal(t, "pre:real:after", res)
}