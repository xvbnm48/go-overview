package main

import "log"

func main() {
	myMap2 := make(map[string]int)
	myMap2["one"] = 1

	myMap := make(map[string]string) 

	myMap["dog"] = "samson"
	log.Println(myMap["dog"])
	log.Println(myMap2["one"])
}
