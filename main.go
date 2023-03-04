package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Email string
	Phone string
}

func main() {
	p1 := Person{
		First: "First",
		Email: "rrrj@gmail.com",
		Phone: "8874444",
	}
	psjon, err := json.Marshal(p1)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(psjon))
}
