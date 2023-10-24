package main

import (
	algo "algorithms/algorithms"
	"fmt"
)

func main() {
	//items := []int{95,78,46,58,45,86,99,251,320}
	//
	//fmt.Println(algorithms.LinearSearch(items, 320))
	//fmt.Println(algorithms.BinarySearch(items, 320))

	// Linked list: 10 -> 20
	//list := &algo.ListNode{
	//	Next: &algo.ListNode{
	//		Next: nil,
	//		Val: 20,
	//	},
	//	Val: 10,
	//}
	//
	//reversed := list.Reverse() // 2 -> 1
	//
	//fmt.Printf("%#v\n", list) // 1 -> 2, because original list is unchanged
	//fmt.Printf("%#v\n", reversed)

	result := algo.EffectivePow(2, 1)
	fmt.Println("result:", result)
}
