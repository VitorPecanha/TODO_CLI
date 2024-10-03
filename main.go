package main

func main() {
	todos := Todos{}

	todos.add("Buy meat")
	todos.add("Buy bread")
	todos.add("Learn Go")
	todos.toggle(0)
	todos.print()
}
