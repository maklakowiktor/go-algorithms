package misc

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DoHTTP(url string) {
	t := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get <%s>: %s\n", url, err.Error())
	}

	defer resp.Body.Close()

	fmt.Printf("<%s> - Status Code [%d] - Latency %d ms\n", url, resp.StatusCode, time.Since(t).Milliseconds())
}

func GetSumFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}, errors.Unwrap(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return []string{}, errors.Unwrap(err)
		}
		sum += num
	}
	return []string{strconv.Itoa(sum)}, nil
}

func SubtractSlice[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range b {
		m[item] = true
	}

	diff := []T{}
	for _, item := range a {
		if !m[item] {
			diff = append(diff, item)
		}
	}
	return diff
}

func IntersectSlices[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	intersection := []T{}

	// Заполняем карту элементами из первого слайса
	for _, item := range a {
		m[item] = true
	}

	// Проверяем, существуют ли элементы второго слайса в карте
	for _, item := range b {
		if m[item] {
			intersection = append(intersection, item)
		}
	}
	return intersection
}
