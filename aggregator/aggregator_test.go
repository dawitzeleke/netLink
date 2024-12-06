package aggregator

import (
	"bandwidth-aggregator/monitor"
	"testing"
)

func TestAggregateBandwidth(t *testing.T) {
	stats1 := monitor.InterfaceStats{Upload: 500, Download: 1000}
	stats2 := monitor.InterfaceStats{Upload: 700, Download: 1500}

	expected := BandwidthStats{TotalUpload: 1200, TotalDownload: 2500}
	result := AggregateBandwidth(stats1, stats2)

	if result != expected {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
