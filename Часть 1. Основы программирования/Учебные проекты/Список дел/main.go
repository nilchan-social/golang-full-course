package main

import (
	"todoapp/scanner"
	"todoapp/todo"
)

func main() {
	list := todo.NewList()

	scanner := scanner.NewScanner(list)

	scanner.Start()
}
