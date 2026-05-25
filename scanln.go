package main

import "fmt"

func main() {
	var ism string
	fmt.Print("Ismingiz nima")
	fmt.Scanln(&ism)
	fmt.Println("Ismingiz", ism)
}
