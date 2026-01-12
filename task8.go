package main

import "fmt"


func main() {
	users := map[string]int{
		"Denis": 30,
		"Bob":   25,
		"Carol": 27,		
	}
	for name, age := range users {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	delete(users, "Bob")

	age, ok := users["Denis"]
	if ok {
		fmt.Print(age)
	} else {
		fmt.Println("User not found")
	}
}