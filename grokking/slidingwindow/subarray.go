package slidingwindow

func AverageOfSubarrayOfSizeK(arr []int, k int) []float64 {
	var result []float64
	for i := 0; i < len(arr)-k+1; i++ {
		var sum float64
		for j := i; j < (i + k); j++ {
			sum += float64(arr[j])
		}
		result = append(result, sum/float64(k))
	}
	return result
}
