package aggregator

import (
	"fmt"
	// "os/exec"
	// "time"
)

// Task represents a real-world task
type Task struct {
	ID        int
	Interface string
	Details   string // Task description (e.g., "Download a file")
	Action    func() error // Function to execute the task
	Status    string       // Task status: Pending, Completed, or Failed
}

// Execute runs the task and updates its status
func (t *Task) Execute() {
	fmt.Printf("[Task Executor] Executing Task ID %d on %s: %s\n", t.ID, t.Interface, t.Details)
	err := t.Action()
	if err != nil {
		t.Status = "Failed"
		fmt.Printf("[Task Executor] Task ID %d failed: %v\n", t.ID, err)
	} else {
		t.Status = "Completed"
		fmt.Printf("[Task Executor] Task ID %d completed successfully\n", t.ID)
	}
}
