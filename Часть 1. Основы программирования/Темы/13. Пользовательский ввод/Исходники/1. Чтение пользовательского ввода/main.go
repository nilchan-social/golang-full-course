package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Введите команду: добавить лук
// вы хотите добавить лук

// Введите команду: удалить морковь
// вы кажется хотите удалить морковь

// неивестная команда
// |
// help
// вот список команд!

// выйти

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			return
		}

		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Println("До скорого!")
			return
		}

		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i != len(fields)-1 {
					str += " "
				}
			}

			fmt.Println("вы хотите добавить:", str)
			fmt.Println("")
		} else if cmd == "удалить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i != len(fields)-1 {
					str += " "
				}
			}

			fmt.Println("вы кажется хотите удалить:", str)
			fmt.Println("")
		} else if cmd == "help" {
			fmt.Println("Команда: help")
			fmt.Println("-- эта команда выводит список доступных команд")
			fmt.Println("")
			fmt.Println("Команда: добавить {что нужно добавить}")
			fmt.Println("-- эта команда позволяет добавлять что-то")
			fmt.Println("")
			fmt.Println("Команда: удалить {что нужно удалить}")
			fmt.Println("-- эта команда позволяет удалять что-то")
			fmt.Println("")
			fmt.Println("Команда: выйти")
			fmt.Println("-- эта команда позволяет завершить программу")
			fmt.Println("")
		} else {
			fmt.Println("Вы ввели неизвестную команду")
			fmt.Println("")
		}
	}

}
