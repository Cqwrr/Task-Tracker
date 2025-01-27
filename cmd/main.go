package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("buy milk")
	todos.add("Купить ГАВНА")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)
}
