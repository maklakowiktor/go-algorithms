package problems

import "fmt"

func FindQuadruples(a1, a2, a3, a4, m int) [][4]int {
	var result [][4]int
	for x1 := 1; x1 < m+1; x1++ {
		for x2 := 1; x2 < m+1; x2++ {
			for x3 := 1; x3 < m+1; x3++ {
				for x4 := 1; x4 < m+1; x4++ {
					if a1*x1+a2*x2+a3*x3+a4*x4 == m {
						fmt.Printf("x1 = %d, x2 = %d, x3 = %d, x4 = %d\n", x1, x2, x3, x4)
						result = append(result, [4]int{x1, x2, x3, x4})
					}
				}
			}
		}
	}
	return result
}

func FindQuadruplesRecursive(a [4]int, m int, currentSolution *[4]int, index int) {
	if currentSolution == nil {
		currentSolution = &[4]int{0, 0, 0, 0}
	}

	if index == 4 {
		var sum int
		for x := 1; x < len(a); x++ {
			sum += a[x] * currentSolution[x]
		}
		if sum == m {
			fmt.Printf("X1 = %d, X2 = %d, X3 = %d, X4 = %d\n", currentSolution[0], currentSolution[1], currentSolution[2], currentSolution[3]) // fmt.Printf
		}
		return
	}
	for x := 1; x < m+1; x++ {
		currentSolution[index] = x
		FindQuadruplesRecursive(a, m, currentSolution, index+1)
	}
}
