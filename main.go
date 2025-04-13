package main

func main() {
	todoList := TodoList{}
	s := NewStorage[TodoList]("todos.json")
	s.Load(&todoList)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todoList)
	// todoList.Add("first todo")
	// todoList.Add("second todo")

	// fmt.Printf("%v+\n\n", todoList)

	// todoList.Print()
	// todoList.Delete(0)

	// fmt.Printf("%v+\n\n", todoList)

	// todoList.toggleCompleted(0)

	// fmt.Printf("%v+\n\n", todoList)

	// todoList.Update(0, "updated todo")
	// fmt.Printf("%v+\n\n", todoList)

	// todoList.Print()

	s.Save(todoList)
	// todo, err := todoList.Get(0)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("todo: ", todo)
}
