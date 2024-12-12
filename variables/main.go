package main

import "fmt"

const LoginToken string = "8r8rwruwdfcweru48ct" // public token as first letter is capital

func main() {
	var username string = "Sajal"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n" , username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n" , isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n" , smallVal)

	var smallFloat float64 = 256.3435454857498572495
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n" , smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T \n" , anotherVariable)

	var website = "learn.in"
	fmt.Println(website)

	name := "sajal"
	fmt.Println(name)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n" , LoginToken)




}
