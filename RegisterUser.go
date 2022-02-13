package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var userName string
	var mail string
	var city string
	fmt.Printf("enter youre name : ")
	fmt.Scan(&name)
	fmt.Printf("enter youre userName : ")
	fmt.Scan(&userName)
	fmt.Printf("enter youre email : ")
	fmt.Scan(&mail)
	fmt.Printf("enter youre City : ")
	fmt.Scan(&city)
	isValid := len(name) > 2 || len(userName) > 2 || len(mail) > 2
	validEmail := strings.Contains(mail, "@")
	validCity := city != "Londen" || city != "Tehran"

	if isValid && validEmail && validCity {
		fmt.Printf("valid Input")
	} else {
		if !isValid {
			fmt.Printf("Invalid lenght of data ")
		} else if !validEmail {
			fmt.Printf("Invalid validEmail")
		} else {
			fmt.Printf("validCity")
		}
	}
}
