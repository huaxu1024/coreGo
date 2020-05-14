package test

import (
	"coreGo/gobook/ch3"
	"testing"
)


func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	res := ch3.Sum(numbers)
	if res != expected {
		t.Errorf("Expected the sum of %v to be %d bu instedd got %d !", numbers, expected, res)
	}
}