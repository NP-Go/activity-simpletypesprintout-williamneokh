package main

import "fmt"

/*
a. text of “The following is
the account information.”
b. first name “Luke”
c. last name “Skywalkter”
d. age of 20 yrs old
e. weight of 73.0 kg
f. height of 1.72 m
g. remaining credits of $123.55
h. account name of “admin”
i. account password of “password”
j. subscribed user as “true”
*/
//Insert variables declarations here based on activity

func tellMeTypes() {
	/*
		a. text of “The following is
		the account information.”
		b. first name “Luke”
		c. last name “Skywalkter”
		d. age of 20 yrs old
		e. weight of 73.0 kg
		f. height of 1.72 m
		g. remaining credits of $123.55
		h. account name of “admin”
		i. account password of “password”
		j. subscribed user as “true”
	*/

	text := "The following is the account information."
	firstName := "Luke"
	lastName := "Skywalker"
	age := 20
	weight := 73.0
	height := 1.72
	remainingCredits := 123.55
	accountName := "admin"
	accountPassword := "password"
	subscribe := true

	fmt.Printf("%T %T %T %T %T %T %T %T %T %T", text, firstName, lastName, age, weight,
		height, remainingCredits, accountName, accountPassword, subscribe)

}

func main() {
	tellMeTypes()
}
