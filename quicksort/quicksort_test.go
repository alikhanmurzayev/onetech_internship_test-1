package quicksort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var size int = 1e6
	arr := makeRandomSlice(size)
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	QuickSort(arr)
	sort.Ints(arrCopy)
	if !reflect.DeepEqual(arr, arrCopy) {
		t.Fatal("slice is not sorted")
	}
}

func makeRandomSlice(size int) []int {
	result := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result = append(result, rand.Int())
	}
	return result
}
