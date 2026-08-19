package detection

import (
	"testing"
	"time"

	"github.com/hami7-glitch/NetGhost/internal/model"
)

func TestDetectPortScan(t *testing.T) {
	baseTime := time.Now()

	events := []model.NetworkEvent{
		{SourceIP: "192.168.1.50", DestPort: 21, Timestamp: baseTime},
		{SourceIP: "192.168.1.50", DestPort: 22, Timestamp: baseTime.Add(1 * time.Second)},
		{SourceIP: "192.168.1.50", DestPort: 23, Timestamp: baseTime.Add(2 * time.Second)},
		{SourceIP: "192.168.1.50", DestPort: 80, Timestamp: baseTime.Add(3 * time.Second)},
		{SourceIP: "192.168.1.50", DestPort: 443, Timestamp: baseTime.Add(4 * time.Second)},
	}

	alert := Analyze(events)

	if alert == nil {
		t.Fatal("expected a port scan alert")
	}
}

func TestNormalTraffic(t *testing.T) {
	baseTime := time.Now()

	events := []model.NetworkEvent{
		{SourceIP: "192.168.1.50", DestPort: 443, Timestamp: baseTime},
		{SourceIP: "192.168.1.50", DestPort: 443, Timestamp: baseTime.Add(1 * time.Second)},
		{SourceIP: "192.168.1.50", DestPort: 443, Timestamp: baseTime.Add(2 * time.Second)},
	}

	alert := Analyze(events)

	if alert != nil {
		t.Fatal("did not expect a port scan alert")
	}
}
