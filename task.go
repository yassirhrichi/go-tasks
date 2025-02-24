package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TODO : Test `json : "id"` behaviore
type task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	// Status      string    `json:"status"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	command := os.Args[1]

	if command != "task" {
		return
	}
	data, err := os.ReadFile("tasks.json")
	check(err)
	dec := json.NewDecoder(strings.NewReader(string(data)))
	var tasks []task
	dec.Decode(&tasks)
	fmt.Print(tasks)

	// operation := os.Args[2]
	// switch operation {
	// case "add":
	// 	task_description := os.Args[3]
	// 	new_task := task{
	// 		Id:          1,
	// 		Description: task_description,
	// 	}
	// 	fmt.Println(new_task)

	// 	str_task, err := json.Marshal(new_task)
	// 	check(err)
	// 	// err2 := os.WriteFile("tasks.json", str_task, 0644)
	// 	// check(err2)
	// 	fmt.Println(string(str_task))
	// }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// case "update":
// 	fmt.Printf("%v %v", command, operation)
// case "delete":
// 	fmt.Printf("%v %v", command, operation)
// case "list":
// 	fmt.Printf("%v %v", command, operation)
// case "mark-in-progress":
// 	fmt.Printf("%v %v", command, operation)
// case "mark-done":
// 	fmt.Printf("%v %v", command, operation)
// default:
// 	fmt.Println("Wrong usage. For help use the command \"task help\"")
// }
