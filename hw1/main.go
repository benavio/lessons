package main

import "fmt"

func main() {

	Link("httsp://weqwehttp://qweqwe.com qwdqw http://qefqef*")

}

func Link(str string) {
	newList := []byte(str)
	marker := []byte("http://")

	findLink(newList, marker)

	fmt.Println(newList)
	fmt.Println(string(newList))
}

func findLink(newList []byte, marker []byte) {
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
			hideLink(i+len(marker), newList)
			i += len(marker)
		}
	}

}

func hideLink(point int, newList []byte) {
	for k := point; k < len(newList); k++ {
		if newList[k] != 32 {
			newList[k] = 42
		} else {
			break
		}
	}
}
