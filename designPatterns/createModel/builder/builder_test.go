package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	assert.Equal(t, "123", res)
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	assert.Equal(t, 6, res)
}
