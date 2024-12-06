package utils

import "fmt"

// Stats stores bandwidth statistics
type Stats struct {
	Upload   int
	Download int
}

// Task represents a network task
type Task struct {
	ID   int
	Name string
}

// PrintTaskStatus logs the assignment of tasks to interfaces
func PrintTaskStatus(task Task, interfaceName string) {
	fmt.Printf("Task %d ('%s') assigned to %s\n", task.ID, task.Name, interfaceName)
}
