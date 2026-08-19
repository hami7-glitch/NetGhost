package main

import (
	"fmt"

	"github.com/hami7-glitch/NetGhost/internal/collector"
	"github.com/hami7-glitch/NetGhost/internal/detection"
)

func main() {
	fmt.Println("🛡️ NetGhost v0.1")
	fmt.Println("Network monitoring system starting...")
	fmt.Println()

	event := collector.CollectSample()

	fmt.Println("Network Event:")
	fmt.Println("Source IP:", event.SourceIP)
	fmt.Println("Destination IP:", event.DestIP)
	fmt.Println("Destination Port:", event.DestPort)
	fmt.Println("Protocol:", event.Protocol)
	fmt.Println()

	alert := detection.Analyze(event)

	if alert != nil {
		fmt.Println("🚨 ALERT:", alert.Message)
	} else {
		fmt.Println("✅ No suspicious activity detected.")
	}
}
