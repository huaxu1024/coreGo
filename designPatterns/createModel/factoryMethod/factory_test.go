package factoryMethod

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory

	factory = PlusOperatorFactory{}
	assert.Equal(t, 3, compute(factory, 1, 2))

	factory = MinusOperatorFactory{}
	assert.Equal(t, 2, compute(factory, 4,2))
}