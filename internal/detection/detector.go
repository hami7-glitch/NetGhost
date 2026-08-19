package detection

import "github.com/hami7-glitch/NetGhost/internal/model"

type Alert struct {
	Message string
	Event   model.NetworkEvent
}

func Analyze(event model.NetworkEvent) *Alert {
	if event.DestPort == 23 {
		return &Alert{
			Message: "Suspicious connection to Telnet port",
			Event:   event,
		}
	}

	return nil
}
