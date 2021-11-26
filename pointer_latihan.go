package main

import "log"

func main() {
	var mystring string
	mystring = "green"

	log.Println("mystring is set to ", mystring)

	changeusingpointer(&mystring)
	log.Println("after call function change using pointer ", mystring)

}

func changeusingpointer(s *string) {
	newValue := "red"
	*s = newValue
}
