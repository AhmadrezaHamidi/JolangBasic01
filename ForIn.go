package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello word ")
	firstName := []string{"Ahmad", "Ali", "Nastran", "Nasrin", "Mahsa", "Mina", "Rima", "Tina"}
	firstNameOf := []string{}
	for index, name := range firstName {
		fmt.Println(index, name)
		firstNameOf = append(firstNameOf, strings.Fields(name)[0])
	}
	fmt.Printf("%v", firstNameOf)
	fmt.Println("After _")
	//If we dont want this index in our loop we using
	for _, Fname := range firstName {
		firstNameOf = append(firstNameOf, strings.Fields(Fname)[0])
	}
	fmt.Printf("%v", firstNameOf)
}
