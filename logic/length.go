package logic

import (
	"log"
)

func Length(file string) []string {
	arr := []int{}
	var height int
	var count int

	for i := 0; i < len(file); i++ {
		if height == 8 {
			height = 0
			count = 0
		}
		if file[i] == '\n' {
			arr = append(arr, count)
			count = 0
			height++
		} else if file[i] != '\r' {
			count++
		}
	}

	var okay bool
	var length int

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			if i+7 < len(arr) && arr[i] == arr[i+1] && arr[i] == arr[i+2] && arr[i] == arr[i+3] && arr[i] == arr[i+4] && arr[i] == arr[i+5] && arr[i] == arr[i+6] && arr[i] == arr[i+7] {
				i += 7
				length++
				okay = true
			} else {
				okay = false
				break
			}
		} else {
			length++
		}
	}

	if !okay {
		log.Fatalln("Invalid input")
	}

	res := make([]string, length)
	height = 0
	count = 0
	var c int

	for i := 0; i < len(file); i++ {
		if file[i] == '\n' && i != 0 && file[i-1] == '\n' {
			c++
			res[c] += "\n"
			continue
		}

		if height == 8 {
			height = 0
			count = 0
			c++
		}

		if file[i] == '\n' {
			res[c] += "\n"
			count = 0
			height++
		} else if file[i] != '\r' {
			res[c] += string(file[i])
			count++
		}
	}

	return res
}
