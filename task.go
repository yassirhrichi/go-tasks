package main

import (
	"encoding/json"
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

	var tasks []task

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	defer file.Close()

	dec := json.NewDecoder(file)
	enc := json.NewEncoder(file)

	err = dec.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}

	operation := os.Args[2]
	switch operation {
	case "add":
		task_description := os.Args[3]
		new_task := task{
			Id:          len(tasks) + 1,
			Description: task_description,
		}
		tasks = append(tasks, new_task)

	case "update":
		//	updated_task_id :=int(os.Args[3])
	}

	err = file.Truncate(0)
	check(err)
	file.Seek(0, 0)
	err = enc.Encode(tasks)
	check(err)
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
