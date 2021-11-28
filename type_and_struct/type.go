package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	Firtsname   string
	Lastname    string
	PhoneNumber string
	Age         int
	Birtdate    time.Time
}

func main() {
	//s2 := "six"

	//	s := "eight"
	//	log.Println("s is ", s)
	//	log.Println("s2 is ", s2)
	//
	//	saySomeThings("xxx")
	user := User{
		Firtsname:   "aki",
		Lastname:    "suda",
		PhoneNumber: "980000",
		Age:         16,
	}

	log.Println(user.Firtsname, user.Lastname)

}

func saySomeThings(s3 string) (string, string) {
	log.Println("s from the say samoething func is", s3)
	return s3, "hello"
}
