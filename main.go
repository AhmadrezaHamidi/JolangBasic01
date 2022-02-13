//package awesomeProject

package main

import "fmt"

// func main() {
// 	// var confrenaceName = "Ahmadreza Hamid fard "
// 	// fmt.Println(confrenaceName)
// 	// const conferenceTicket = 50
// 	// var remaningTecket = 50
// 	// fmt.Println("of the ", conferenceTicket, "remanin\n", remaningTecket)
// 	// var username string
// 	// var ticketUSer int
// 	// username = "Mahsa"
// 	// ticketUSer = 2
// 	// fmt.Printf("hello world %v and youre ticket %v", confrenaceName, ticketUSer)
// 	// fmt.Println(username)
// 	// fmt.Printf("hello world %t and youre ticket %t,", username, ticketUSer)
// 	// var nameOne string = "ahmadreza"
// 	// nameTwo := "Ahmadreza"
// 	// ///nameOne === name two
// 	// ///we van user := jaust for Variable
// 	// fmt.Println(nameOne)
// 	// fmt.Println(nameTwo)
// 	// //fmt.Scan() ///For input
// 	// var userNameInput string
// 	// fmt.Scan(userNameInput)
// 	// fmt.Println(&userNameInput)
// 	//& is pointer to show memori location
// }

func main() {
	var userName string
	var lastName string
	var age int
	const allTicket = 50
	var remaningTicket = 50
	var userTicket int
	fmt.Println("pleasr entr youreName")
	fmt.Scan(&userName)
	fmt.Println("Please enter LastName")
	fmt.Scan(&lastName)
	fmt.Println("Please enter Age")
	fmt.Scan(&age)
	fmt.Println("Please enter ticket ")
	fmt.Scan(&userTicket)
	remaningTicket = remaningTicket - userTicket
	fmt.Printf("Hello %v %v ,youre welcome", userName, lastName)
	fmt.Printf("we have %v \n", remaningTicket)

	//Declare Arrayes in Goland
	//	var customerName = [50]string{"Ahmad", "Mahsa", "Farnoush"}
	//var customerName = [50]string{}
	var customerName [50]string
	customerName[0] = "Ahmad"
	customerName[1] = userName + " " + lastName
	//fmt.Printf("the whole Arraye is %v \n", customerName)
	fmt.Printf("the First Element  Arraye is %v \n", customerName[0])
	///fmt.Printf("the type of  Arraye is %t \n", customerName)
	fmt.Printf("the lenght of Arraye is %v \n", len(customerName))

	//Declare Slice In GoLang
	///Look like arrayes but no size and that is daynamic
	var names []string
	names = append(names, "Ahmad")
	names = append(names, "Ahmad1")
	names = append(names, "Ahmad2")
	names = append(names, "Ahmad3")
	fmt.Printf("the whole slice is %v \n", names)
	fmt.Printf("the First Element  slice is %v \n", names[0])
	fmt.Printf("the type of  slice is %t \n", names)
	fmt.Printf("the lenght of slice is %v \n", len(names))

	//Loopes in goLang
	var allName []string
	for {
		var namett string
		fmt.Printf("entr youre name :")
		fmt.Scan(&namett)
		allName = append(allName, namett)
		fmt.Println(allName)
	}

}
