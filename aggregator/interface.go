// aggregator/interface.go
package aggregator

import (
	"fmt"
	"net"
)

// ListInterfaces returns a list of all available network interfaces
func ListInterfaces() ([]net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return interfaces, nil
}

// GetInterfaceByName returns a network interface by its name
func GetInterfaceByName(name string) (*net.Interface, error) {
	interfaces, err := ListInterfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if iface.Name == name {
			return &iface, nil
		}
	}
	return nil, fmt.Errorf("interface %s not found", name)
}
