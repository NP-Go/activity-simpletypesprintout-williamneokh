package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {

	text := "The following is the account information."
	firstName := "Luke"
	lastName := "Skywalker"
	age := 20
	weight := 73.0
	height := 1.72
	remainingCredits := 123.55
	accountName := "admin"
	accountPassword := "password"
	isSubscribe := true

	fmt.Printf("%T %T %T %T %T %T %T %T %T %T", text, firstName, lastName, age, weight,
		height, remainingCredits, accountName, accountPassword, isSubscribe)

}

func main() {
	tellMeTypes()
}
