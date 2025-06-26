package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter string:  ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	newList := []byte(str)
	marker := []byte("http://")

	for i := 0; i < len(newList); i++ {
		if i+len(marker) > len(newList) {
			break
		}

		match := true
		for j := 0; j < len(marker); j++ {
			if newList[i+j] != marker[j] {
				match = false
				break
			}
		}

		if match {
			for k := i + len(marker); k < len(newList); k++ {
				if newList[k] != 32 {
					newList[k] = 42
				} else {
					break
				}
			}
			i += len(marker)
		}
	}

	fmt.Println(newList)
	fmt.Println(string(newList))
}
