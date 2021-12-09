package main

import (
	"log"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "sakura endo")
	mySlice = append(mySlice, "sakura miyawaki")
	mySlice = append(mySlice, "aki suda")

	log.Println(mySlice[0:1])

	log.Println(mySlice)
}
