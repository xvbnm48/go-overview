package main

import "log"

var s = "seven"

func main() {
	s2 := "six"

	s := "eight"
	log.Println("s is ",s)
	log.Println("s2 is ",s2)

	saySomeThings("xxx")
}

func saySomeThings(s3 string) (string, string) {
	log.Println("s from the say samoething func is" , s3)
	return s3, "hello"
}