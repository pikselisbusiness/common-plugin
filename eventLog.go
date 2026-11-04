package shared

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Event log statuses the host shows.
const (
	EventLogSent     = "sent"
	EventLogRetrying = "retrying"
	EventLogFailed   = "failed"
	EventLogSkipped  = "skipped"
	EventLogPending  = "pending"
)

// Reference types the host links to an ERP record; anything else is shown
// as text.
const (
	EventLogRefOrder    = "order"
	EventLogRefInvoice  = "invoice"
	EventLogRefProduct  = "product"
	EventLogRefCompany  = "company"
	EventLogRefShipment = "shipment"
	EventLogRefContact  = "contact"
)

// EventLog describes, in the host's plugin event log, what the plugin did
// with an event it handled. Put it on the HandleResult:
//
//	return shared.HandleResult{Tags: shared.EventLog{
//		Title: "Placed order #123", ReferenceType: shared.EventLogRefOrder,
//		ReferenceID: "123", ResponseCode: 202,
//	}.Tags()}, nil
//
// The host logs the delivery with these details, its status from the result
// (sent, retrying, failed; Status "skipped" for an event deliberately not
// acted on) and keeps the event so a user can send it again. A successful
// result without an EventLog is not logged: the plugin ignored the event.
// Failed and retried deliveries are logged either way.
type EventLog struct {
	Title         string
	ReferenceType string
	ReferenceID   string
	ResponseCode  int
	// Response is the external service's answer, shown in the details (64 KiB).
	Response string
	// Status overrides the status read from the result; only "skipped" is used.
	Status string
}

// Tags renders the event log as HandleResult tags (eventLog.* keys).
func (l EventLog) Tags() map[string]string {
	tags := map[string]string{}
	set := func(key, value string) {
		if value != "" {
			tags["eventLog."+key] = value
		}
	}
	set("title", l.Title)
	set("referenceType", l.ReferenceType)
	set("referenceId", l.ReferenceID)
	if l.ResponseCode != 0 {
		set("responseCode", strconv.Itoa(l.ResponseCode))
	}
	set("response", l.Response)
	set("status", l.Status)
	if len(tags) == 0 {
		// An empty EventLog still asks for the delivery to be logged.
		tags["eventLog.title"] = ""
	}
	return tags
}

// EventLogEntry is an entry the plugin writes itself, for work that is not an
// event delivery (a catalogue sync, an import). EventID is the plugin's own
// key: writing again with the same key updates the entry.
type EventLogEntry struct {
	EventID       string `json:"eventId"`
	EventType     string `json:"eventType"`
	Title         string `json:"title"`
	ReferenceType string `json:"referenceType"`
	ReferenceID   string `json:"referenceId"`
	Status        string `json:"status"`
	ResponseCode  int    `json:"responseCode"`
	Error         string `json:"error"`
	Request       string `json:"request"`
	Response      string `json:"response"`
}

// APIEventLogger is the host call behind LogEvent. Optional, like APIStopper:
// not part of API, so plugin test mocks of API keep compiling.
type APIEventLogger interface {
	LogEvent(entryJSON []byte) error
}

// ErrEventLogUnsupported is returned by LogEvent when the API has no event log.
var ErrEventLogUnsupported = errors.New("host has no plugin event log")

// LogEvent writes an entry to the host's plugin event log. An error means it
// was not stored (an older host answers "method not found"); never fail the
// plugin's own work because of it.
func LogEvent(api API, entry EventLogEntry) error {
	logger, ok := api.(APIEventLogger)
	if !ok {
		return ErrEventLogUnsupported
	}
	raw, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return logger.LogEvent(raw)
}
