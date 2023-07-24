package logic

import (
	"fmt"
	"log"
)

func Function(a []string, art map[rune]string) {
	for i := 0; i < len(a); i++ {
		if a[i] == "\n" && len(a[i]) == 1 {
			fmt.Println()
			continue
		}

		cols := len(a[i])/8 - 1
		rows := 8

		str := a[i]
		arr := make([][]byte, rows)
		for j := range arr {
			arr[j] = make([]byte, cols)
		}

		x := 0
		y := 0
		for j := 0; j < len(str); j++ {
			if str[j] != '\n' {
				arr[x][y] = str[j]
				y++
			} else {
				x++
				y = 0
			}
		}

		space := 0
		index := []int{}
		for c := 0; c < cols; c++ {
			space = 0
			for r := 0; r < rows; r++ {
				if arr[r][c] == ' ' {
					space++
				}

				if space == 8 {
					index = append(index, c)
				}
			}
		}

		for j := 0; j < len(index); j++ {
			if j != 0 && j+4 < len(index) && index[j-1] == index[j]-1 && index[j]+1 == index[j+1] && index[j]+2 == index[j+2] && index[j]+3 == index[j+3] && index[j]+4 == index[j+4] {
				index = append(index[:j], index[j+5:]...)
				j--
			} else if j == 0 && j+4 < len(index) && index[j] == index[j+1]-1 && index[j] == index[j+2]-2 && index[j] == index[j+3]-3 && index[j] == index[j+4]-4 {
				index = append(index[:j], index[j+5:]...)
				j--
			}
		}

		for j := 0; j < len(index); j++ {
			symbol := ""
			for r := 0; r < 8; r++ {
				if j == 0 {
					for c := 0; c < index[j]; c++ {
						symbol += string(arr[r][c])
					}
					symbol += " \n"
				} else {
					for c := index[j-1] + 1; c < index[j]; c++ {
						symbol += string(arr[r][c])
					}
					symbol += " \n"
				}
			}

			var flag bool
			for i, v := range art {
				if v == symbol {
					fmt.Print(string(i))
					flag = true
					break
				}
			}
			if !flag {
				fmt.Println()
				log.Fatalln("Not correct input")
			}
		}
		fmt.Println()
	}
}
