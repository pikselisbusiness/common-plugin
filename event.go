package shared

import "time"

type EventEnvelope struct {
	// routing / meta
	EventID       string `json:"eventId"`      // unique id (UUID)
	EventName     string `json:"eventName"`    // "cart.created", "order.paid", "product.updated"
	EventVersion  uint32 `json:"eventVersion"` // semantic version *of the payload schema*

	AggregateID   string `json:"aggregateId"`  // CartID/OrderID/ProductID
	TenantID      string `json:"tenantId,omitempty"`
	PartitionKey string            `json:"partitionKey,omitempty"` // default AggregateID

	OccurredAt    time.Time `json:"occurredAt"`
	CorrelationID string `json:"correlationId,omitempty"`
	TraceID       string `json:"traceId,omitempty"`
	// delivery hints
	Headers      map[string]string `json:"headers,omitempty"`

	// data
	ContentType string `json:"contentType"` // "application/json;type=CartCreatedV2"
	Payload     []byte `json:"payload"`     // JSON/Proto bytes of the versioned payload
}

// To be determined when will be used
type EventEnvelopeV2 struct {
	// Core identity (CloudEvents-friendly)
	EventID      string `json:"eventId"`           // unique id (UUID)
	EventName    string `json:"eventName"`         // "cart.created", "order.paid"
	EventVersion string `json:"eventVersion"`      // semver of payload/schema, e.g. "2.1.0"
	Source       string `json:"source,omitempty"`  // service/topic URI or name
	Subject      string `json:"subject,omitempty"` // sub-entity (e.g., "lineItem:abc")

	// Domain routing
	AggregateID  string            `json:"aggregateId"`
	TenantID     string            `json:"tenantId,omitempty"`
	PartitionKey string            `json:"partitionKey,omitempty"` // defaults to AggregateID
	Headers      map[string]string `json:"headers,omitempty"`

	// Time
	OccurredAt time.Time `json:"occurredAt"`           // business time
	ProducedAt time.Time `json:"producedAt,omitempty"` // producer emit time
	ReceivedAt time.Time `json:"receivedAt,omitempty"` // ingress time

	// Correlation / causality / tracing
	CorrelationID string `json:"correlationId,omitempty"`
	CausationID   string `json:"causationId,omitempty"`
	ParentID      string `json:"parentId,omitempty"`
	TraceParent   string `json:"traceParent,omitempty"` // W3C traceparent
	Baggage       string `json:"baggage,omitempty"`

	// Actor / audit
	ActorID     string `json:"actorId,omitempty"` // system or user
	UserID      string `json:"userId,omitempty"`
	SessionID   string `json:"sessionId,omitempty"`
	Environment string `json:"environment,omitempty"` // prod/stage/dev
	Region      string `json:"region,omitempty"`

	// Delivery / replay
	DeliveryAttempt uint32 `json:"deliveryAttempt,omitempty"`
	AckPolicy       string `json:"ackPolicy,omitempty"` // "at-least-once" (default)
	TTLSeconds      uint32 `json:"ttlSeconds,omitempty"`
	Replay          bool   `json:"replay,omitempty"`
	RetryAfterMs    uint32 `json:"retryAfterMs,omitempty"`

	// Data
	ContentType     string `json:"contentType"`               // "application/json;type=CartCreatedV2"
	ContentEncoding string `json:"contentEncoding,omitempty"` // "gzip" or ""
	ContentSchema   string `json:"contentSchema,omitempty"`   // URL/URN/registry ref
	PayloadChecksum string `json:"payloadChecksum,omitempty"` // "sha256:..."
	Encrypted       bool   `json:"encrypted,omitempty"`
	EncryptionKeyID string `json:"encryptionKeyId,omitempty"`
	PIIMinimized    bool   `json:"piiMinimized,omitempty"`
	Payload         []byte `json:"payload"`
}
type HandleResult struct {
	// If set, the host should requeue with backoff.
	Retry bool `json:"retry"`
	// Optional next attempt delay hint.
	RetryAfterMs uint32 `json:"retryAfterMs,omitempty"`
	// For DLQ / logs.
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Optional plugin-produced side effects or metrics.
	Metrics map[string]float64 `json:"metrics,omitempty"`
	Tags    map[string]string  `json:"tags,omitempty"`
}
