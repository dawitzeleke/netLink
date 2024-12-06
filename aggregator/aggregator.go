package aggregator

import (
	"fmt"
	"sync"
	"time"
	"bandwidth-aggregator/monitor"
	"os/exec"
)

// Aggregator stores and distributes tasks across interfaces
type Aggregator struct {
	TaskQueue   []*Task
	mu          sync.Mutex
	TaskCounter int
}

// NewAggregator initializes the Aggregator
func NewAggregator() *Aggregator {
	return &Aggregator{
		TaskQueue: make([]*Task, 0),
	}
}

// DistributeTask assigns a task to the interface dynamically
func (a *Aggregator) DistributeTask(interfaceName string, monitor *monitor.MonitorInterface) {
	a.mu.Lock()

	a.TaskCounter++
	stats := monitor.GetStats()[interfaceName]

	task := &Task{
		ID:        a.TaskCounter,
		Interface: interfaceName,
		Details:   fmt.Sprintf("Interface %s - Upload: %d KB, Download: %d KB", interfaceName, stats.Upload, stats.Download),
		Status:    "Active",
		Action: func() error {
			// Detect actual network activity
			cmd := exec.Command("netstat", "-i")
			output, err := cmd.Output()

			if err != nil {
				return fmt.Errorf("netstat error: %s", err)
			}
			
			fmt.Printf("[Detected Network Activity on Interface %s]\n%s\n", interfaceName, string(output))

			return nil
		},
	}

	a.TaskQueue = append(a.TaskQueue, task)

	a.mu.Unlock()

	fmt.Printf("[Task Distributor] Interface: %s assigned dynamic activity scanning\n", interfaceName)
}

// ContinuouslyDistributeTasks across interfaces dynamically
func (a *Aggregator) ContinuouslyDistributeTasks(monitor *monitor.MonitorInterface) {
	go func() {
		for {
			interfaces := monitor.GetStats()

			for interfaceName := range interfaces {
				a.DistributeTask(interfaceName, monitor)
			}

			time.Sleep(2 * time.Second)  // Regular interface monitoring
		}
	}()
}

// PrintStats shows currently active interface assignments
func (a *Aggregator) PrintStats() {
	a.mu.Lock()

	fmt.Printf("\nCurrently Running Network Interface Activities:\n")
	for _, task := range a.TaskQueue {
		fmt.Printf("\nTaskID: %d Status %s across Interface %s -> Activity: %s\n", task.ID, task.Status, task.Interface, task.Details)
	}

	a.mu.Unlock()
}
