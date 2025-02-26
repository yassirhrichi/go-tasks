package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

// TODO : Test `json : "id"` behaviore
type task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//To Format time layout must use the reference time Mon Jan 2 15:04:05 MST 2006. eg. Time.Now().Local().Format("02-01-2006 15:04:05")

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Wrong usage. For help use the command \"task help\"")
		return
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
			return
		}
		taskDescription := os.Args[3]
		var taskId int
		if len(tasks) == 0 {
			taskId = 1
		} else {
			taskId = tasks[len(tasks)-1].Id + 1
		}
		newTask := task{
			Id:          taskId,
			Description: taskDescription,
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, newTask)
		fmt.Println("Task added successfully!")
	case "update":
		if len(args) != 5 {
			fmt.Println("Wrong usage. for the update command, use \"task update [id] [new_descrition]\"")
			return
		}
		taskId, err := strconv.Atoi(args[3])
		check(err)
		upTaskIndx := slices.IndexFunc(tasks, func(t task) bool { return t.Id == taskId })
		upDesc := args[4]
		tasks[upTaskIndx].Description = upDesc
		tasks[upTaskIndx].UpdatedAt = time.Now()

		fmt.Println("Task updated successfully!")
	case "delete":
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the add command, use \"task delete [id]\"")
			return
		}
		taskId, err := strconv.Atoi(args[3])
		check(err)
		tasks = slices.DeleteFunc(tasks, func(t task) bool { return t.Id == taskId })
		fmt.Println("Task deleted successfully!")

	case "list":
		if len(args) == 3 {
			fmt.Println("Tasks list :")
			for _, t := range tasks {
				fmt.Printf("Id : %v | Description : %v | Status : %v \n", t.Id, t.Description, t.Status)
			}
		} else if len(args) == 4 {
			status := args[3]
			switch status {
			case "todo":
				fmt.Println("The todo tasks list :")
				for _, t := range tasks {
					if t.Status == "todo" {
						fmt.Printf("Id : %v | Description : %v\n", t.Id, t.Description)
					}
				}
			case "in-progress":
				fmt.Println("The in-progress tasks list :")
				for _, t := range tasks {
					if t.Status == "in-progress" {
						fmt.Printf("Id : %v | Description : %v\n", t.Id, t.Description)
					}
				}
			case "done":
				fmt.Println("The done tasks list :")
				for _, t := range tasks {
					if t.Status == "done" {
						fmt.Printf("Id : %v | Description : %v\n", t.Id, t.Description)
					}
				}
			default:
				fmt.Println("Wrong usage. for the list command, use \"task list [status] [id]\"")
				return
			}

		} else {
			fmt.Println("Wrong usage. for the list command, use \"task list [status]\"")
			return
		}

	case "mark-in-progress":
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the list command, use \"task mark-in-progress [id]\"")
			return
		}
		taskId, err := strconv.Atoi(args[3])
		check(err)
		upTaskIndx := slices.IndexFunc(tasks, func(t task) bool { return t.Id == taskId })
		tasks[upTaskIndx].Status = "in-progress"
		tasks[upTaskIndx].UpdatedAt = time.Now()

		fmt.Println("Task' status updated successfully!")

	case "mark-done":
		if len(args) != 4 {
			fmt.Println("Wrong usage. for the list command, use \"task mark-in-progress [id]\"")
			return
		}
		taskId, err := strconv.Atoi(args[3])
		check(err)
		upTaskIndx := slices.IndexFunc(tasks, func(t task) bool { return t.Id == taskId })
		tasks[upTaskIndx].Status = "done"
		tasks[upTaskIndx].UpdatedAt = time.Now()

		fmt.Println("Task' status updated successfully!")
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
