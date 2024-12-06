package monitor

import "testing"

func TestMonitorInterface(t *testing.T) {
	stats := MonitorInterface("eth0")

	if stats.Upload < 0 || stats.Download < 0 {
		t.Errorf("Invalid stats: %+v", stats)
	}
}
