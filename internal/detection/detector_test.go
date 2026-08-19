package detection

import (
	"testing"

	"github.com/hami7-glitch/NetGhost/internal/model"
)

func TestAnalyzeTelnet(t *testing.T) {
	event := model.NetworkEvent{
		SourceIP: "192.168.1.10",
		DestIP:   "192.168.1.20",
		DestPort: 23,
		Protocol: "TCP",
	}

	alert := Analyze(event)

	if alert == nil {
		t.Fatal("expected an alert for Telnet port 23")
	}
}

func TestAnalyzeNormalHTTPS(t *testing.T) {
	event := model.NetworkEvent{
		SourceIP: "192.168.1.10",
		DestIP:   "8.8.8.8",
		DestPort: 443,
		Protocol: "TCP",
	}

	alert := Analyze(event)

	if alert != nil {
		t.Fatal("did not expect an alert for HTTPS port 443")
	}
}
