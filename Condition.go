package main

import "fmt"

func main() {
	var numberTry int
	fmt.Printf("give me youre number that want to try that ")

	fmt.Scan(&numberTry)

	//for {
	//	numberTry--
	//	fmt.Printf("%v", numberTry)
	//	continue
	//	if numberTry == 0 {
	//		break
	//	}
	//}
	for {
		fmt.Println("whats you choseies ?")
		var input int
		fmt.Scan(&input)
		noRemaningTicket := input == 0
		RemaningTicket := input == 1

		if noRemaningTicket {
			fmt.Printf("noRemaningTicket !")
			break
		} else if RemaningTicket {
			fmt.Printf("RemaningTicket !")
			continue
		} else {
			fmt.Printf("youre Result is %v !", 2)
			continue
		}
	}
}
