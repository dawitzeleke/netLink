package config

import "fmt"

// Config stores system configuration values
var Interfaces = map[string]string{
	"Mobile": "cellular",
	"WiFi":   "wifi",
}

// LogConfig prints configuration for debugging
func LogConfig() {
	fmt.Println("Using Interfaces:")
	for key, value := range Interfaces {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
