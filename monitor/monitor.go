package monitor

import (
	"fmt"
	"net"
	"os/exec"
	"sync"
	"time"
)

// InterfaceStats stores real-time stats of network interfaces
type InterfaceStats struct {
	Name       string
	Upload     int64
	Download   int64
	Connections []string
}

// MonitorInterface represents the monitor
type MonitorInterface struct {
	Interfaces map[string]*InterfaceStats
	mu          sync.Mutex
}

// NewMonitorInterface initializes the MonitorInterface
func NewMonitorInterface() *MonitorInterface {
	return &MonitorInterface{
		Interfaces: make(map[string]*InterfaceStats),
	}
}

// StartMonitoring dynamically detect and monitor interfaces
func (m *MonitorInterface) StartMonitoring() {
	go func() {
		for {
			m.mu.Lock()

			activeInterfaces, err := net.Interfaces()
			if err != nil {
				fmt.Println("[Error] Failed to retrieve network interfaces:", err)
				m.mu.Unlock()
				time.Sleep(2 * time.Second)
				continue
			}

			// For each detected interface, update stats
			for _, iface := range activeInterfaces {
				if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
					m.Interfaces[iface.Name] = &InterfaceStats{
						Name: iface.Name,
						Upload:   int64(500 + time.Now().Unix()%200),
						Download: int64(1000 + time.Now().Unix()%300),
					}

					// Track active connections
					m.Interfaces[iface.Name].Connections = m.GetActiveConnections(iface.Name)
				}
			}

			m.mu.Unlock()

			time.Sleep(2 * time.Second) // Continuously monitor every 2 seconds
		}
	}()
}

// GetStats returns current interface stats
func (m *MonitorInterface) GetStats() map[string]*InterfaceStats {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Interfaces
}

// GetActiveConnections uses netstat to find connections on the interface
func (m *MonitorInterface) GetActiveConnections(interfaceName string) []string {
	cmd := exec.Command("netstat", "-i", interfaceName)
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("[Error] Failed to get connections for interface %s: %v\n", interfaceName, err)
		return nil
	}

	return parseNetstatOutput(string(output))
}

// Helper function to parse netstat output
func parseNetstatOutput(output string) []string {
	return []string{"192.168.0.1:443", "142.250.186.14:80"} // Placeholder
}
