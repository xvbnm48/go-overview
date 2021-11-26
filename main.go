package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var wantToSay string

	umur := 18

	wantToSay = "sakura endo sangat cantik"

	fmt.Println(wantToSay)

	fmt.Println("umur sakura endo adalah ", umur)
	whatWeSaid := saySomething()
	fmt.Println(whatWeSaid)
}

func saySomething() string {
	return "someting"
}
