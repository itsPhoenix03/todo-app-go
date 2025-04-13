package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Toggle int
	Delete int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task to your todo list")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task in your todo list based on the ID provided")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task in your todo list based on the ID provided")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a task in your todo list based on the ID provided")
	flag.BoolVar(&cf.List, "list", false, "List all tasks in your todo list")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todoList *TodoList) {
	switch {
	case cf.Add != "":
		todoList.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid format for edit command. Use ID:Task_Title")
			os.Exit(1)
		}

		id, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Invalid ID provided for edit command")
			os.Exit(1)
		}

		todoList.Update(id, parts[1])

	case cf.Toggle != -1:
		todoList.toggleCompleted(cf.Toggle)
	case cf.Delete != -1:
		todoList.Delete(cf.Delete)
	case cf.List:
		todoList.Print()
	default:
		fmt.Println("Invalid command provided")

	}
}
