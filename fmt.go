package main

import "fmt"

func main() {
	firstName := "Muhtarom"
	lastName  := "Zain"

	fmt.Println("Hello", firstName, lastName, "!")
	fmt.Println("Hello", "'", firstName, lastName, "'")

	fmt.Printf("Hello %s %s!\n", firstName, lastName)
	fmt.Printf("Hello '%s %s'", firstName, lastName)
}