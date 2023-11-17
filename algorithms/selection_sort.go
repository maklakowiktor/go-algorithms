package algorithms

func SelectionSort(arr *[]int, n int) {
	pArr := (*arr)
	for k := 0; k < n-1; k++ {
		for j := k + 1; j < n; j++ {
			if pArr[k] > pArr[j] {
				pArr[j], pArr[k] = pArr[k], pArr[j]
			}
		}
	}
}
