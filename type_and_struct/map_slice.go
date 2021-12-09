package main

import "log"

type User struct {
	Firtname string
	Lastname string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		Firtname: "sakura",
		Lastname: "endo",
	}

	myMap["me"] = me
	log.Println(myMap["me"].Firtname)

}
