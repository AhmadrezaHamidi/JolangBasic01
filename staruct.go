package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func manin() {
	fmt.Println("this is struct")

	type userData struct {
		firstName string
		lastName  string
		email     string
		age       uint
	}
	var users = make([]userData, 0)
	var t = userData{
		firstName: "Ahmad",
		lastName:  "HAmidi",
		email:     "Aasdjsadjas",
		age:       22,
	}
	users = append(users, t)
	go senfTicket("Ahmad", 550, 20)
	wg.Add(1)
	wg.Wait()
}

//fmt.Sprintf for put charecter and create string of sentence to gather

func senfTicket(name string, allticket int, userTicket int) string {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("hi %v ypure tivket is %v of %v", name, allticket, userTicket)
	wg.Done()
	return ticket
}

///go key word  ==> bara ine ke age kari zaman mibare  ba go multi threading Mishe
///Ama na agar bekhahim kare maro bere anjam bede va badadn biyad iz wahiteGrpup estefade mikonim
