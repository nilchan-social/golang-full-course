package scanner

import (
	"fmt"
	"todoapp/todo"

	"github.com/k0kubun/pp"
)

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
	fmt.Println("")
}

func printResult(result string) {
	pp.Println("result:", result)
	fmt.Println("")
}

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы... До скорого!")
}

func printAdd(taskDescription string) {
	fmt.Println("Задача " + "'" + taskDescription + "'" + " успешно добавлена")
	fmt.Println("")
}

func printDone(taskDescription string) {
	fmt.Println("Задача " + "'" + taskDescription + "'" + " помечена как выполненная")
	fmt.Println("")
}

func printDel(taskDescription string) {
	fmt.Println("Задача " + "'" + taskDescription + "'" + " успешно удалена")
	fmt.Println("")
}

func printEvents(events []Event) {
	pp.Println("События:", events)
	fmt.Println("")
}

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("1. help")
	fmt.Println("-- эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("")
	fmt.Println("2. add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов}")
	fmt.Println("-- эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("")
	fmt.Println("3. list")
	fmt.Println("-- эта команда позволяет получить полный список всех задач")
	fmt.Println("")
	fmt.Println("4. del {заголовок существующей задачи}")
	fmt.Println("-- эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("5. done {заголовок существующей задачи}")
	fmt.Println("-- эта команда позволяет отменить задачу как выполненную")
	fmt.Println("")
	fmt.Println("6. events")
	fmt.Println("-- эта команда позволяет получить список всех событий")
	fmt.Println("")
	fmt.Println("7. exit")
	fmt.Println("-- эта команда позволяет завершить выполнение программы")
	fmt.Println("")
}
