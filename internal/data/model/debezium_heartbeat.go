package model

import (
	"time"
)

// DebeziumHeartbeat is utilized to publish heartbeat events to prevent WAL growth in idle systems
// Heartbeat table events advance WAL but are not filtered by the outbox SMT.
type DebeziumHeartbeat struct {
	Id int       `gorm:"primaryKey"`
	Ts time.Time `gorm:"type:timestamptz;not null"`
}
