package utils

// []int {9,8,7,6,5}
// []int {5,6,7,8,9}
func BubbleSort(element []int) []int {
	keepRUnning := true
	for keepRUnning {
		keepRUnning = false

		for i := 0; i < len(element)-1; i++ {
			if element[i] > element[i+1] {
				element[i], element[i+1] = element[i+1], element[i]
				keepRUnning = true
			}
		}
	}
	return element
}
