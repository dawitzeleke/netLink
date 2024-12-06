package ui

import (
	"bandwidth-aggregator/monitor"
	"fmt"
)

func DisplayStats(interfaceName string, stats *monitor.Stats) {
	fmt.Printf("Interface: %s\n", interfaceName)
	fmt.Printf("Upload: %d KB, Download: %d KB\n", stats.Upload, stats.Download)
}
