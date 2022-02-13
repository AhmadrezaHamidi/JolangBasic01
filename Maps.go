package main

import "strconv"

func main() {
	var listuserData = make([]map[string]string, 0)
	var userData = make(map[string]string)
	userData["firstName"] = "Ahmad"
	userData["lastName"] = "Hamidi"
	userData["email"] = "Hamidi@"
	userData["Age"] = strconv.FormatInt(int64(40), 10) ///Converting
	listuserData = append(listuserData, userData)
}
