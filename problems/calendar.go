package problems

import (
	"fmt"
	"log"
)

func PrintCalendar(startDay uint, numDays uint) {
	if startDay < 1 || startDay > 7 {
		log.Fatal("Start day должен быть от 1 до 7")
		return
	}
	if numDays < 1 || numDays > 99 {
		log.Fatal("NumDays должен быть от 1 до 99")
		return
	}
	//fmt.Println("Mon Tue Wed Thu Fri Sat Sun")
	var day uint = 1
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 7; j++ {
			if i == 1 && uint(j) < startDay {
				fmt.Print("   ")
			} else if day <= numDays {
				// TODO: Implement print with std::set in C++
				//fmt.Printf("%*s\n", 3, " ")
				fmt.Printf(" %v ", day)
				day++
			}
		}
		fmt.Println("")
		if day > numDays {
			break
		}
	}
}
