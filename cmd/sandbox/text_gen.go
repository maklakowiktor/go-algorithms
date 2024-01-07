package main

import (
	"fmt"
	"math/rand"
)

func GenerateText(word string) string {
	var charset string = "abcdefghijklmnopqrstuvwxyz"
	var output string = ""
	for i := 0; i < 1000; i++ {
		printWord := rand.Intn(2) != 0
		length := rand.Intn(50) + 10
		txt := ""
		randomWordPos := rand.Intn(length)
		for i := 0; i < length+len(word); i++ {
			if printWord && randomWordPos == i {
				txt += word
			}
			c := charset[rand.Intn(len(charset))]
			txt += string(c)
		}
		output = fmt.Sprintf("%s%s\n", output, txt)
	}
	return output
}
