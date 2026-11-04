package shared

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestEventLogTags(t *testing.T) {
	tags := EventLog{Title: "Placed order #7", ReferenceType: EventLogRefOrder, ReferenceID: "7", ResponseCode: 202, Response: "{}"}.Tags()
	want := map[string]string{
		"eventLog.title": "Placed order #7", "eventLog.referenceType": "order", "eventLog.referenceId": "7",
		"eventLog.responseCode": "202", "eventLog.response": "{}",
	}
	if len(tags) != len(want) {
		t.Fatalf("tags: %v", tags)
	}
	for key, value := range want {
		if tags[key] != value {
			t.Fatalf("%s: %q, want %q", key, tags[key], value)
		}
	}
	if empty := (EventLog{}).Tags(); len(empty) != 1 {
		t.Fatalf("an empty EventLog must still ask to be logged: %v", empty)
	}
}

type eventLogProbe struct {
	API
	entries []EventLogEntry
	fail    bool
}

func (p *eventLogProbe) LogEvent(entryJSON []byte) error {
	if p.fail {
		return errors.New("store failed")
	}
	var entry EventLogEntry
	if err := json.Unmarshal(entryJSON, &entry); err != nil {
		return err
	}
	p.entries = append(p.entries, entry)
	return nil
}

func TestLogEventOverRPC(t *testing.T) {
	probe := &eventLogProbe{}
	client := featureReadClient(t, probe)
	entry := EventLogEntry{EventID: "sync-1", EventType: "catalog.sync", Title: "Catalogue sync", Status: EventLogSent}
	if err := LogEvent(client, entry); err != nil {
		t.Fatal(err)
	}
	if len(probe.entries) != 1 || probe.entries[0] != entry {
		t.Fatalf("entry: %+v", probe.entries)
	}
	probe.fail = true
	if err := LogEvent(client, entry); err == nil || err.Error() != "store failed" {
		t.Fatalf("error lost: %v", err)
	}
}

type noEventLogHost struct{ API }

func TestLogEventOnHostWithoutLog(t *testing.T) {
	client := featureReadClient(t, &noEventLogHost{})
	if err := LogEvent(client, EventLogEntry{EventType: "x", Status: EventLogSent}); err == nil {
		t.Fatal("a host without the log must answer an error")
	}
	if err := LogEvent(&noEventLogHost{}, EventLogEntry{}); !errors.Is(err, ErrEventLogUnsupported) {
		t.Fatalf("an API without LogEvent: %v", err)
	}
}
