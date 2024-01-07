package problems

import "fmt"

// https://new.contest.yandex.ru/41233/problem?id=149944/2022_10_11/cpwScRzQ0S

func CalculateDeliveryTime(n, m, t int) string {
	totalMinutesStart := n*60 + m
	totalMinutesEnd := totalMinutesStart + t

	hours := (totalMinutesEnd / 60) % 24
	minutes := totalMinutesEnd % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
