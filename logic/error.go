package logic

import "fmt"

func Error(n int) {
	if n == 1 {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
	}
}
