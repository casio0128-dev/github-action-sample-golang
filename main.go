package main

import "fmt"

func main() {
	for _, r := range []rune("Hello,世界!?") {
		fmt.Print(r, " ")
		fmt.Println(string(r))
	}
}