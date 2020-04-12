package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	res := NewAdapter(NewAdaptee()).Request()
	assert.Equal(t, expect, res)
}