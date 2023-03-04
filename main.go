package main

import (
	"net/http"
)

type Person struct {
	First string
	Email string
	Phone string
}

func main() {
	// p1 := Person{
	// 	First: "First",
	// 	Email: "rrrj@gmail.com",
	// 	Phone: "8874444",
	// }
	// psjon, err := json.Marshal(p1)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(psjon))
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":3000", nil)

}
func foo(w http.ResponseWriter, r *http.Request) {

}

func bar(w http.ResponseWriter, r *http.Request) {

}
