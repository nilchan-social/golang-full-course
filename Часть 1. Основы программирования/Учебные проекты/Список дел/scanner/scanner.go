package scanner

import (
	"bufio"
	"os"
	"strings"
	"todoapp/todo"
)

type Scanner struct {
	todoList *todo.List

	bufioScanner *bufio.Scanner

	events []Event
}

func NewScanner(todoList *todo.List) Scanner {
	return Scanner{
		todoList: todoList,

		bufioScanner: bufio.NewScanner(os.Stdin),
	}
}

func (s *Scanner) Start() {
	for {
		printPrompt()

		ok := s.bufioScanner.Scan()
		if !ok {
			break
		}

		inputString := s.bufioScanner.Text()

		result := s.process(inputString)

		if result != "" {
			if result == needExis {
				printExit()
				return
			}

			printResult(result)

			printHelp()
		}

		event := NewEvent(result, inputString)
		s.events = append(s.events, event)
	}
}

func (s *Scanner) process(inputString string) string {
	fields := strings.Fields(inputString)

	if len(fields) == 0 {
		return emptyInput
	}

	cmd := fields[0]

	if cmd == "exit" {
		return needExis
	}

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "del" {
		return s.cmdDel(fields)
	}

	if cmd == "done" {
		return s.cmdDone(fields)
	}

	if cmd == "events" {
		return s.cmdEvents(fields)
	}

	if cmd == "help" {
		return s.cmdHelp(fields)
	}

	return unknownCommand
}

func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}

	taskDescription := fields[1]

	taskText := ""
	for i := 2; i < len(fields); i++ {
		taskText += fields[i]

		if i != len(fields)-1 {
			taskText += " "
		}
	}

	task := todo.NewTask(taskDescription, taskText)

	s.todoList.AddTask(task)

	printAdd(taskDescription)

	return ""
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	tasks := s.todoList.ListTasks()

	printTasks(tasks)

	return ""
}

func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	taskDescription := fields[1]

	doneTaskResult := s.todoList.DoneTask(taskDescription)
	if doneTaskResult != "" {
		return doneTaskResult
	}

	printDone(taskDescription)

	return ""
}

func (s *Scanner) cmdDel(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	taskDescription := fields[1]

	delTaskResult := s.todoList.DeleteTask(taskDescription)
	if delTaskResult != "" {
		return delTaskResult
	}

	printDel(taskDescription)

	return ""
}

func (s *Scanner) cmdEvents(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printEvents(s.events)

	return ""
}

func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printHelp()

	return ""
}
