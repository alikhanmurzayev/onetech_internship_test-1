package quicksort

func partion(a []int, left, right int) int {
	// my pivot element always rightWall
	pivot_element := a[right]
	wall := left
	for i := left; i < right; i++ {
		if a[i] < pivot_element {
			// we swap greater than pivot with less than pivot, if i != wall
			a[wall], a[i] = a[i], a[wall]
			wall++
		}
	}
	// as a result wall's left side have elements less or equal pivot, wall' right side have elements greater than pivot

	// swap pivot element with wall
	a[right], a[wall] = a[wall], a[right]
	return wall
}
func QuickSortWithLeftAndRightBoundary(a []int, left, right int) {
	// we do recursion while left < right
	if left < right {
		// middle is a pivot location
		middle := partion(a, left, right)
		QuickSortWithLeftAndRightBoundary(a, left, middle-1)
		QuickSortWithLeftAndRightBoundary(a, middle+1, right)
	}
}
func QuickSort(a []int) {
	QuickSortWithLeftAndRightBoundary(a, 0, len(a)-1)
}
