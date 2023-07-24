package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git/ssengerb/ascii-art-reverse/logic"
)

func main() {
	if len(os.Args) == 2 {
		if len(os.Args[1]) < 14 || os.Args[1][:10] != "--reverse=" || os.Args[1][len(os.Args[1])-4:] != ".txt" {
			logic.Error(1)
		} else {
			file := os.Args[1][10:]
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			art := logic.FillMap()
			logic.Function(logic.Length(string(data)), art)
		}
	} else {
		logic.Error(1)
		return
	}
}
