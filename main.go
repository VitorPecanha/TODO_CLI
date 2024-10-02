package main

import "fmt"

func main() {
	todos := Todos{}

	todos.add("Buy meat")
	todos.add("Buy bread")
	todos.add("Learn Go")

	fmt.Printf("%+v\n\n", todos)

	todos.delete(0)
	fmt.Printf("%+v\n\n", todos)
}
