package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	operation := os.Args[2]
	switch operation {
	case "add":
		task_description := os.Args[3]
		new_task := task{
			Id:          1,
			Description: task_description,
		}
		str_task, _ := json.Marshal(new_task)
		fmt.Println(string(str_task))
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
