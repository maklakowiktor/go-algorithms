package algorithms

func BubbleSort(arr *[]int, n int) {
	pArr := (*arr)
	for k := 0; k < n; k++ {
		for i := 0; i < n-1; i++ {
			if pArr[i] > pArr[i+1] {
				pArr[i], pArr[i+1] = pArr[i+1], pArr[i]
			}
		}
	}
}
