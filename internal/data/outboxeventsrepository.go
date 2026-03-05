package data

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/project-kessel/inventory-api/internal"
	"github.com/project-kessel/inventory-api/internal/biz/model_legacy"
)

// walOutboxMessage mirrors the outbox table column names for JSON serialization.
// This ensures Debezium's DecodeLogicalDecodingMessageContent SMT produces
// a record with the same field names as the original outbox table, making the
// transition transparent to downstream consumers.
type walOutboxMessage struct {
	ID            string              `json:"id"`
	AggregateType string              `json:"aggregatetype"`
	AggregateID   string              `json:"aggregateid"`
	Operation     string              `json:"operation"`
	TxID          string              `json:"txid"`
	Payload       internal.JsonObject `json:"payload"`
}

// PublishOutboxEvent is a function variable that emits outbox events. It defaults to
// publishing via PostgreSQL WAL logical decoding messages. Tests can replace this with
// a no-op or mock implementation (e.g., for SQLite which lacks pg_logical_emit_message).
var PublishOutboxEvent = publishOutboxEventWAL

// publishOutboxEventWAL emits a transactional logical decoding message to PostgreSQL WAL.
// Debezium captures these messages directly from the WAL stream, eliminating the need
// for an outbox table and the INSERT+DELETE pattern that caused SSI serialization conflicts.
// The message is transactional (first arg = true), meaning it only appears in the WAL
// if the surrounding transaction commits.
func publishOutboxEventWAL(tx *gorm.DB, event *model_legacy.OutboxEvent) error {
	// Generate UUID if not already set (previously handled by GORM BeforeCreate hook)
	if event.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			return fmt.Errorf("failed to generate uuid for WAL message: %w", err)
		}
		event.ID = id
	}

	msg := walOutboxMessage{
		ID:            event.ID.String(),
		AggregateType: string(event.AggregateType),
		AggregateID:   event.AggregateID,
		Operation:     string(event.Operation.OperationType()),
		TxID:          event.TxId,
		Payload:       event.Payload,
	}

	content, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal outbox event for WAL message: %w", err)
	}

	// Use aggregatetype as the message prefix for Debezium topic routing
	prefix := string(event.AggregateType)

	return tx.Exec(
		"SELECT pg_logical_emit_message(true, ?, ?)",
		prefix, string(content),
	).Error
}
