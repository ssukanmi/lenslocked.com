package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Dog  string
	Age  map[int]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Dog:  "Fido",
		Age: map[int]string{
			1: "one",
			2: "two",
			3: "three",
			4: "four",
			5: "five",
		},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	data = User{
		Name: "Mary Jane",
		Dog:  "Spot",
		Age: map[int]string{
			30: "thirty",
			35: "thirtyfive",
			40: "forty",
			45: "fortyfive",
		},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
