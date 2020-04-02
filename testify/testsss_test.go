package testify

import (
	"coreGo/run"
	"testing"
)


func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	res := run.Sum(numbers)
	if res != expected {
		t.Errorf("Expected the sum of %v to be %d bu instedd got %d !", numbers, expected, res)
	}
}