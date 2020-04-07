package sort

import (
	"github.com/go-playground/assert/v2"
	"sort"
	"testing"
)

func TestSortBubbe(t *testing.T) {
	arrs := []int{9,8,7,6,6,5,4,3,2,1}

	expected := make([]int, len(arrs))
	copy(expected, arrs)

	sort.Ints(expected)
	SortBubbe(arrs, 0, len(arrs))
	assert.Equal(t, expected, arrs)
}


func TestSortHeap(t *testing.T) {
	arrs := []int{9,8,7,6,6,5,4,3,2,1}

	expected := make([]int, len(arrs))
	copy(expected, arrs)

	sort.Ints(expected)
	SortHeap(arrs, 0, len(arrs))
	assert.Equal(t, expected, arrs)
}

func TestSortInsert(t *testing.T) {
	arrs := []int{9,8,7,6,6,5,4,3,2,1}

	expected := make([]int, len(arrs))
	copy(expected, arrs)

	sort.Ints(expected)
	SortInsert(arrs, 0, len(arrs))
	assert.Equal(t, expected, arrs)
}

func TestSortMerge(t *testing.T) {
	arrs := []int{9,8,7,6,6,5,4,3,2,1}

	expected := make([]int, len(arrs))
	copy(expected, arrs)

	sort.Ints(expected)
	SortMerge(arrs, 0, len(arrs) - 1)
	assert.Equal(t, expected, arrs)
}

func TestSortQuick(t *testing.T) {
	arrs := []int{9,8,7,6,6,5,4,3,2,1}

	expected := make([]int, len(arrs))
	copy(expected, arrs)

	sort.Ints(expected)
	SortQuick(arrs, 0, len(arrs) - 1)
	assert.Equal(t, expected, arrs)
}