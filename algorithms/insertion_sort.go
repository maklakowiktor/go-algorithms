package algorithms

func InsertionSort(arr *[]int, n int) {
	pArr := *arr
	for k := 1; k < n; k++ {
		for i := k; i > 0 && pArr[i-1] < pArr[i]; i-- {
			pArr[i], pArr[i-1] = pArr[i-1], pArr[i]
		}
	}
}
