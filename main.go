package main

import (
	"algorithms/problems"
	"fmt"
	"net/http"
	"time"
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

	//result := algo.EffectivePow(2, 1)
	//fmt.Println("result:", result)

	//message := make(chan string)

	//go func() {
	//	time.Sleep(2 * time.Second)
	//	message <- "Message 1"
	//}()
	//
	//fmt.Println(<-message)
	//fmt.Println(<-message)

	//go func() {
	//	for i := 1; i <= 10; i++ {
	//		message <- fmt.Sprintf("%d", i)
	//		time.Sleep(time.Millisecond * 500)
	//	}
	//
	//	close(message)
	//}()

	// 1 case
	//for {
	//	msg, open := <-message
	//	if !open {
	//		break
	//	}
	//	fmt.Println(msg)
	//}

	// 2 case (syntax sugar)
	//for msg := range message {
	//	fmt.Println(msg)
	//}

	//msgBuf := make(chan string, 2)
	//msgBuf <- "hello"
	//msgBuf <- "world"
	//
	//fmt.Println(<-msgBuf)
	//msgBuf <- "deadlock!"
	//
	//fmt.Println(<-msgBuf)
	//fmt.Println(<-msgBuf)

	//urls := []string{
	//	"http://google.com/",
	//	"http://youtube.com/",
	//	"http://github.com/",
	//	"http://medium.com/",
	//	"http://instagram.com/",
	//}
	//
	//var wg sync.WaitGroup
	//
	//for _, url := range urls {
	//	wg.Add(1)
	//
	//	go func(url string) {
	//		doHTTP(url)
	//		wg.Done()
	//	}(url)
	//}
	//
	//wg.Wait()

	//message1 := make(chan string)
	//message2 := make(chan string)
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Millisecond * 500)
	//		message1 <- "Gone half of second"
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Second * 2)
	//		message2 <- "Gone half of 2 seconds"
	//	}
	//}()
	//
	//for {
	//	// non-blocking getting from channels
	//	select {
	//	case msg := <-message1:
	//		fmt.Println(msg)
	//	case msg := <-message2:
	//		fmt.Println(msg)
	//	}
	//	//fmt.Println(<-message1)
	//	//fmt.Println(<-message2)
	//}

	//fmt.Println(algorithms.IsAnagram("anagram", "nagaram"))
	//fmt.Println(algorithms.IsAnagram("anmagram", "nagaram"))

	problems.PrintCalendar(5, 31)
}

func doHTTP(url string) {
	t := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get <%s>: %s\n", url, err.Error())
	}

	defer resp.Body.Close()

	fmt.Printf("<%s> - Status Code [%d] - Latency %d ms\n", url, resp.StatusCode, time.Since(t).Milliseconds())
}
