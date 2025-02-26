package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func readTasks() ([]Task, error) {
	fileData, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		return []Task{}, nil
	} else if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func getNextTaskID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	return maxID + 1
}

func findTaskIndex(tasks []Task, id int) (int, error) {
	for i, t := range tasks {
		if t.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("task not found")
}

func addTask(title string) {
	tasks, err := readTasks()
	check(err)

	newTask := Task{
		Id:        getNextTaskID(tasks),
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	check(writeTasks(tasks))
	fmt.Println("Task added successfully!")
}

func updateTask(id int, title string) {
	tasks, err := readTasks()
	check(err)

	index, err := findTaskIndex(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	tasks[index].Title = title
	tasks[index].UpdatedAt = time.Now()
	check(writeTasks(tasks))
	fmt.Println("Task updated successfully!")
}

func updateTaskStatus(id int, status string) {
	tasks, err := readTasks()
	check(err)

	index, err := findTaskIndex(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now()
	check(writeTasks(tasks))
	fmt.Println("Task status updated successfully!")
}

func listTasks() {
	tasks, err := readTasks()
	check(err)

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Id < tasks[j].Id })

	for _, t := range tasks {
		fmt.Printf("[%d] %s (%s)\n", t.Id, t.Title, t.Status)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task [command] [arguments]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task add \"Task description\"")
			return
		}
		addTask(os.Args[2])

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task update [id] \"New description\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		check(err)
		updateTask(id, os.Args[3])

	case "list":
		listTasks()

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task mark-in-progress [id]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		check(err)
		updateTaskStatus(id, StatusInProgress)

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task mark-done [id]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		check(err)
		updateTaskStatus(id, StatusDone)

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
