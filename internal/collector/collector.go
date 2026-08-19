package collector

import (
	"time"

	"github.com/hami7-glitch/NetGhost/internal/model"
)

func CollectSample() model.NetworkEvent {
	return model.NetworkEvent{
		Timestamp: time.Now(),
		SourceIP:  "192.168.1.10",
		DestIP:    "8.8.8.8",
		DestPort:  23,
		Protocol:  "TCP",
	}
}
