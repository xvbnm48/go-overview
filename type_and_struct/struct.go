package main

import "log"

type myStruct struct {
	Firtname string
}

func (m *myStruct) printFirstsName() string {
	return m.Firtname
}

func main() {
	var myVar myStruct

	myVar.Firtname = "sakura"

	myvar2 := myStruct{
		Firtname: "aki suda",
	}

	log.Println("myvar is set to", myVar.printFirstsName())
	log.Println("myvar2 is set to", myvar2.printFirstsName())
}
