package shared

import (
	"database/sql"
	"time"
)

// Clause types for query building
type Clause struct {
	Type   string // "where", "order", "limit", "offset", "group", "having", "join"
	Query  string
	Args   []interface{}
	Column string
	Desc   bool
}

// Model represents a database model with common fields
type Model struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `gorm:"index" json:"deletedAt,omitempty"`
}

// QueryResult represents the result of a database operation
type QueryResult struct {
	Error        error
	RowsAffected int64
	LastInsertID int64
}

// ScanResult for scanning query results
type ScanResult struct {
	Rows  []map[string]interface{}
	Error error
}

// Association types for preloading
type Association struct {
	Name       string
	ForeignKey string
	References string
}

// JoinType for different join operations
type JoinType string

const (
	InnerJoin JoinType = "INNER"
	LeftJoin  JoinType = "LEFT"
	RightJoin JoinType = "RIGHT"
)

// OrderDirection for ordering
type OrderDirection string

const (
	ASC  OrderDirection = "ASC"
	DESC OrderDirection = "DESC"
)
