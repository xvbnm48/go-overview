package main

import "log"

func main() {
	//for i := 0; i <= 10; i++ {
	//	log.Println(i)
	//}

	//	idol := []string{"sakura endo", "sakura miyawaki", "aki suda"}
	idols := make(map[string]string)

	idols["HKT48"] = "sakura miyawaki"
	idols["tokimeki sendenby"] = "aki suda"
	idols["nogizaka46"] = "sakura endo"

	for idolsType, idols := range idols {
		log.Println(idolsType, idols)
	}
}
