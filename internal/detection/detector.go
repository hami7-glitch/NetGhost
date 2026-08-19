package detection

import (
	"fmt"
	"time"

	"github.com/hami7-glitch/NetGhost/internal/model"
)

type Alert struct {
	Message string
	Event   model.NetworkEvent
}

const (
	portScanThreshold = 5
	portScanWindow    = 10 * time.Second
)

func Analyze(events []model.NetworkEvent) *Alert {
	if len(events) == 0 {
		return nil
	}

	portsByIP := make(map[string]map[uint16]bool)

	for _, event := range events {
		if _, exists := portsByIP[event.SourceIP]; !exists {
			portsByIP[event.SourceIP] = make(map[uint16]bool)
		}

		portsByIP[event.SourceIP][event.DestPort] = true
	}

	for sourceIP, ports := range portsByIP {
		if len(ports) >= portScanThreshold {
			return &Alert{
				Message: fmt.Sprintf(
					"Possible port scan detected from %s: %d different ports",
					sourceIP,
					len(ports),
				),
				Event: events[0],
			}
		}
	}

	return nil
}
