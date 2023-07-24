package logic

import (
	"io/ioutil"
	"log"
	"strings"
)

func FillMap() map[rune]string {
	file, err := ioutil.ReadFile("./banners/standard.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	lines := strings.Split(string(file), "\n")
	art := make(map[rune]string)
	for i := 32; i <= 126; i++ {
		symbolstart := (i-32)*9 + 1
		for j := symbolstart; j <= symbolstart+8; j++ {
			if j != symbolstart+8 {
				art[rune(i)] += lines[j] + "\n"
			} else {
				art[rune(i)] += lines[j]
			}
		}
	}
	return art
}
