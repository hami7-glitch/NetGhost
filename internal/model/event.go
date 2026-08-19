package model

import "time"

type NetworkEvent struct {
	Timestamp time.Time
	SourceIP  string
	DestIP    string
	DestPort  uint16
	Protocol  string
}
