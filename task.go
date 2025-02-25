package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
)

// TODO : Test `json : "id"` behaviore
type task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Wrong usage. For help use the command \"task help\"")
	}

	command := args[1]

	if command != "task" {
		fmt.Println("Wrong usage. For help use the command \"task help\"")
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
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the add command, use \"task add [descrition]\"")
		}
		task_description := os.Args[3]
		new_task := task{
			Id:          tasks[len(tasks)-1].Id + 1,
			Description: task_description,
			Status:      "todo",
		}
		tasks = append(tasks, new_task)
	case "update":
		if len(args) != 5 {
			fmt.Println("Wrong usage. for the update command, use \"task update [id] [new_descrition]\"")
		}
		task_id, err := strconv.Atoi(args[3])
		check(err)
		upd_task_indx := slices.IndexFunc(tasks, func(t task) bool { return t.Id == task_id })
		upd_desc := args[4]
		tasks[upd_task_indx].Description = upd_desc
	case "delete":
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the add command, use \"task delete [id]\"")
		}
		task_id, err := strconv.Atoi(args[3])
		check(err)
		tasks = slices.DeleteFunc(tasks, func(t task) bool { return t.Id == task_id })

	case "list":
		if len(args) == 3 {
			fmt.Println("Tasks list")
			for _, t := range tasks {
				fmt.Println("Id : %V | Description : %V | Status : ", t.Id, t.Description, t.Status)
			}
		} else if len(args) == 4 {
			// switch statement for the status
		} else {
			fmt.Println("Wrong usage. for the list command, use \"task list [status]\"")
		}

	case "mark-in-progress":
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the list command, use \"task mark-in-progress [id]\"")
		}
		task_id, err := strconv.Atoi(args[3])
		check(err)
		upd_task_indx := slices.IndexFunc(tasks, func(t task) bool { return t.Id == task_id })
		tasks[upd_task_indx].Status = "in-progress"
	case "mark-done":
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the list command, use \"task mark-in-progress [id]\"")
		}
		task_id, err := strconv.Atoi(args[3])
		check(err)
		upd_task_indx := slices.IndexFunc(tasks, func(t task) bool { return t.Id == task_id })
		tasks[upd_task_indx].Status = "done"
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
