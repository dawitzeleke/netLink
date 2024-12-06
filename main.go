package main

import (
	"fmt"
	"time"
	"bandwidth-aggregator/monitor"
	"bandwidth-aggregator/aggregator"
)

func main() {
	fmt.Println("[System] Initializing Bandwidth Aggregator...")

	// Initialize monitoring and aggregator components
	netMonitor := monitor.NewMonitorInterface()
	taskAggregator := aggregator.NewAggregator()

	// Start monitoring network interfaces
	netMonitor.StartMonitoring()

	// Continuously distribute tasks across interfaces
	taskAggregator.ContinuouslyDistributeTasks(netMonitor)

	// Continuously print tasks every few seconds
	for {
		taskAggregator.PrintStats()
		time.Sleep(5 * time.Second)
	}
}
