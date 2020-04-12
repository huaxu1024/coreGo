package simpleFactory

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	assert.Equal(t, "Hi, Tom", s)
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	assert.Equal(t, "Hello, Tom", s)
}
