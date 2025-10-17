package task1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CmdClient() {
	ReadFromFile("output.txt")
	scanner := bufio.NewScanner(os.Stdin)

	exit := false
	for {
		fmt.Println("Список доступных функций:")
		fmt.Println("1. Вывести список всех задач")
		fmt.Println("2. Вывести список задач с определённым статусом (to do, in progress, done)")
		fmt.Println("3. Добавить новую задачу")
		fmt.Println("4. Обновить информацию о задаче ")
		fmt.Println("5. Удалить задачу ")
		fmt.Println("6. Выход")
		fmt.Print("Ввод: ")
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "1":
			ShowAllTasks()
		case "2":
			back := false
			for {
				fmt.Print("Введите задачи какого статуса вы хотите вывести(to do, in progress, done): ")
				scanner.Scan()
				sortType := scanner.Text()
				switch sortType {
				case "to do":
					ShowTodoTasks()
					back = true
				case "in progress":
					ShowTodoTasks()
					back = true
				case "done":
					ShowDoneTasks()
					back = true
				default:
					fmt.Println("Неверный ввод.")
				}
				if back {
					break
				}
			}
		case "3":
			fmt.Print("Введите имя задачи: ")
			scanner.Scan()
			name := scanner.Text()
			t, err := CreateTask(name)
			if err != nil {
				fmt.Println(err)
			}
			AddTask(t)

		case "4":
			back := false
			for {
				fmt.Print("Введите id задачи, которую вы хотите поменять: ")
				scanner.Scan()
				id, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println(err)
				}
				_, ok := taskStorage[id]
				if !ok {
					fmt.Println("нет такого id")
				}
				fmt.Print("Введите, что изменить(name, status): ")
				scanner.Scan()
				upd := scanner.Text()
				if upd == "name" {
					fmt.Print("Введите новое имя: ")
					scanner.Scan()
					newName := scanner.Text()
					UpdateTaskName(id, newName)
					back = true
				} else if upd == "status" {
					for {
						fmt.Print("Введите новый статус (to do, in progress, done): ")
						scanner.Scan()
						newStatus := status(scanner.Text())
						if newStatus == Todo {
							UpdateTaskStatus(id, Todo)
							break
						} else if newStatus == Done {
							UpdateTaskStatus(id, newStatus)
							break
						} else if newStatus == InProgres {
							UpdateTaskStatus(id, InProgres)
							break
						} else {
							fmt.Println("Неверный статус задачи.")
						}
					}
					back = true
				} else {
					fmt.Print("Неверный ввод")
				}
				if back {
					break
				}
			}
		case "5":
			for {
				fmt.Print("Введите id задачи для удаления: ")
				scanner.Scan()
				id, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Print(err)
					continue
				}
				_, ok := taskStorage[id]
				if !ok {
					fmt.Println("Такой задачи не сущетсвует")
					continue
				}

				DeleteTask(id)
				break
			}
		case "6":
			fmt.Println("пока пока")
			WriteInFile("output.txt")
			exit = true
		}
		if exit {
			break
		}
	}

}
