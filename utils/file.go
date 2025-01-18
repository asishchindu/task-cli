package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var filename = "task-cli-data.json"

func init() {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		createFile()
	}
}

type TaskData struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// AddTask adds a new task to the task list and saves it to the file
func AddTask(description string) {
	// Read existing tasks
	tasks, err := readFile()
	if err != nil {
		fmt.Println("Error reading existing tasks:", err)
		return
	}

	// fmt.Println(existingTasks)
	taskId := getNextID(tasks)

	tasks[taskId] = TaskData{
		Id:          taskId, // Generate a simple ID
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// // Convert the updated task list to JSON
	taskJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling task data:", err)
		return
	}

	// // Save the updated task list back to the file
	err = os.WriteFile(filename, taskJSON, 0644)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Println("Task added successfully")
}

func ListTasks() {
	existingTasks, err := readFile()
	if err != nil {
		fmt.Println("Error reading existing tasks:", err)
		return
	}

	fmt.Println("Id", "Description")
	for key, task := range existingTasks {
		fmt.Println(key, task.Description)
	}
}

func UpdateTask(id string, description string) {
	tasks, err := readFile()
	if err != nil {
		//
	}

	task := tasks[id]
	task.Description = description
	tasks[id] = task

	// // Convert the updated task list to JSON
	taskJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling task data:", err)
		return
	}

	// // Save the updated task list back to the file
	err = os.WriteFile(filename, taskJSON, 0644)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Println("Task Updated successfully")
}

func DeleteTask(id string) {
	tasks, err := readFile()
	if err != nil {
		//
	}

	delete(tasks, id)

	// // Convert the updated task list to JSON
	taskJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling task data:", err)
		return
	}

	// // Save the updated task list back to the file
	err = os.WriteFile(filename, taskJSON, 0644)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Println("Task Deleted successfully")
}

// createFile creates an empty JSON file with the initial structure
func createFile() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write an empty task list to the file
	initialContent := `{}`
	_, err = file.WriteString(initialContent)
	if err != nil {
		fmt.Printf("Error initializing file content: %v\n", err)
	}
}

// readFile reads the task list from the file and returns it
func readFile() (map[string]TaskData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read the file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Parse the JSON into a map
	var taskList map[string]TaskData
	err = json.Unmarshal(bytes, &taskList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return taskList, nil
}

func getNextID(tasks map[string]TaskData) string {
	maxID := 0
	for id := range tasks {
		intID, err := strconv.Atoi(id)
		if err == nil && intID > maxID {
			maxID = intID
		}
	}

	return fmt.Sprintf("%d", maxID+1)
}
