package main

import (
	"fmt"

	"github.com/hami7-glitch/NetGhost/internal/collector"
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
	fmt.Println("Timestamp:", event.Timestamp)
}
