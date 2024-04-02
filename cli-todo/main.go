package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type TodoItem struct {
	id   int
	name string
	done bool
}

var NEXT_ID = 0

func init() {
	file, err := os.OpenFile("data/todo.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	NEXT_ID = len(records) + 1

	file.Close()

}

func main() {
	fmt.Println("Welcome to the Todo CLI")

	for {
		file, err := os.OpenFile("data/todo.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

		if err != nil {
			fmt.Println("Error reading file", err)
			break
		}

		defer file.Close()

		var operation string
		var id int
		var action string

		fmt.Println("What would you like to do? \n a (add) \n r (remove) \n l (list) \n u (update) \n q (quit)")
		fmt.Scanln(&operation)

		if operation == "q" {
			break
		}

		switch operation {

		case "l", "list":
			printTodos(getTodoFromFile(file))
			fmt.Println("\n >>>> Operation Complete <<<<\n ")

		case "a", "add":
			addTodoToFile(file)
			fmt.Println("\n >>>> Operation Complete <<<<\n ")

		case "r", "remove":
			fmt.Println("Which todo to remove?")
			fmt.Scanf("%d", &id)
			updateTodo(file, "d", id)
			file.Close()
			renameFile()
			fmt.Println("\n >>>> Operation Complete <<<<\n ")

		case "u", "update":
			fmt.Println("Which todo to update?")
			fmt.Scanf("%d", &id)

			fmt.Println("Choose an action c for marking complete, i for incomplete.")
			fmt.Scanf("%s", &action)

			updateTodo(file, action, id)
			file.Close()
			renameFile()
			fmt.Println("\n >>>> Operation Complete <<<<\n ")

		default:
			file.Close()
			fmt.Println("\n >>>> Unknown Operation <<<<\n ")
		}
	}

}

func renameFile() {
	if err := os.Rename("data/temp.csv", "data/todo.csv"); err != nil {
		panic(err)
	}
}

func getTodoFromFile(file io.Reader) []TodoItem {
	reader := csv.NewReader(file)

	var todoRecords []TodoItem
	records, _ := reader.ReadAll()

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		done := record[1] == "true"

		var todo = TodoItem{
			id:   id,
			done: done,
			name: record[2],
		}
		todoRecords = append(todoRecords, todo)
	}
	return todoRecords
}

func printTodos(todos []TodoItem) {
	for _, todo := range todos {
		mark := "[ ]"
		if todo.done {
			mark = "[x]"
		}
		fmt.Println(mark + " " + todo.name)
	}
}

func addTodoToFile(file io.Writer) {
	fmt.Println("Add your todo and press enter")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	todoToWrite := []string{
		strconv.Itoa(NEXT_ID), "false", name,
	}
	NEXT_ID++

	w := csv.NewWriter(file)

	err := w.Write(todoToWrite)
	if err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func updateTodo(file io.Reader, action string, id int) {
	tempFile, err := os.OpenFile("data/temp.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening temporary file")
	}

	tempFileWriter := csv.NewWriter(tempFile)
	fileReader := csv.NewReader(file)

	oldRecords, err := fileReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var DELETED = false
	for _, oldRecord := range oldRecords {
		oldRecordId, _ := strconv.Atoi(oldRecord[0])
		if id == oldRecordId {
			switch action {
			case "i":
				tempFileWriter.Write([]string{
					oldRecord[0], "false", oldRecord[2],
				})
			case "c":
				tempFileWriter.Write([]string{
					oldRecord[0], "true", oldRecord[2],
				})
			case "d":
				DELETED = true
				NEXT_ID = NEXT_ID - 1
				continue
			default:
				err := tempFileWriter.Write(oldRecord)
				if err != nil {
					log.Fatalln("error writing record to csv:", err)
				}
			}
		} else {
			record := oldRecord

			if DELETED {
				id, _ := strconv.Atoi(oldRecord[0])
				id = id - 1
				record[0] = strconv.Itoa(id)
			}

			err := tempFileWriter.Write(record)
			if err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		}
	}

	tempFileWriter.Flush()

	if err := tempFileWriter.Error(); err != nil {
		log.Fatal(err)
	}

	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
	}
}
