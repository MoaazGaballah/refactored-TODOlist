package main

func main() {
	a := App{}
	a.Initialize("todouser", "1", "todos")

	a.Run(":8010")
}
