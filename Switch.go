package main

import (
	"fmt"
	"strings"
)

var res string ///for golabal variable in our application

func main() {
	var city string
	fmt.Scan(&city)
	hello()
	helloUser("Ahmad")
	switch city {
	case "New York ":
		fmt.Println("New york")
	case "sangapor":
		fmt.Println("sangapor")
	case "tehran":
		fmt.Println("tehran")
	default:
		fmt.Println("New york")
	}
}

///Functions
func hello() {
	fmt.Println("Welcome to my application ")
}

///Function white parameter
func helloUser(userName string) string {
	return "Welcome to my application " + userName + " "
}

///Many type for out put
func checkObject(inputOne string, inputTwo string, inputThree string) (bool, bool) {
	var lenghtValue bool
	var isEmailo bool
	if len(inputOne) < 2 || len(inputTwo) < 2 || len(inputThree) < 2 {
		lenghtValue = true
	}
	if strings.Contains(inputTwo, "@") {
		isEmailo = true
	}
	return lenghtValue, isEmailo

	///vaqti ke chand ta go file darim pas tasmim migrim ke le ona ro taqcim konim
	///bara hamin to chand ta file minevicimo
	///Hame ro ba ham run milonim
	//go run xxx.go  manin.go
	// in bara ine ke chang ta bashe
	//ama agar chanf ta majol bashe va har majol chand ta package bashe
	///yek package dorost mikonim
	//va hame ro on to be ham refrence midim
	//barye in mozo ma mirim to go.mod
	//esme koli ro bardashe
	//va be sorate zir dar poackege insert mikonim
	///{esmePackageToGo.Mod}/{esmeFolder}/{esme package}

	//Agar harfe avale Bozorg bashad on halate public darad
	//agar harfe aval kochock bash halate private darad

	//3 level of scope
	//local  ===> in {} for exampel loop es
	//Package{} ==> that we can accress inm the packeage Level
	//global ==> what we can use all of the majol is Global

}
