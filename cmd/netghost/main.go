package main

import (
	"fmt"
	"time"

	"github.com/hami7-glitch/NetGhost/internal/collector"
	"github.com/hami7-glitch/NetGhost/internal/detection"
	"github.com/hami7-glitch/NetGhost/internal/model"
)

func main() {
	fmt.Println("🛡️ NetGhost v0.1")
	fmt.Println("Network monitoring system starting...")
	fmt.Println()

	var events []model.NetworkEvent

	for i := 0; i < 5; i++ {
		event := collector.CollectSample()
		event.DestPort = uint16(20 + i)
		event.Timestamp = time.Now()
		events = append(events, event)
	}

	alert := detection.Analyze(events)

	if alert != nil {
		fmt.Println("🚨 ALERT:", alert.Message)
	} else {
		fmt.Println("✅ No suspicious activity detected.")
	}
}
