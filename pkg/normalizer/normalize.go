package normalizer

import (
	"encoding/json"
	"os"
)

// Update the LogEntry to match the structure of an alert in eve.json
type LogEntry struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	SrcIP     string `json:"src_ip"`
	Alert     Alert  `json:"alert"`
}

type Alert struct {
	Signature string `json:"signature"`
}

// ReadLogEntries reads the log entries from a file and returns them as a slice.
func ReadLogEntries(file *os.File) ([]LogEntry, error) {
	// Step 1: Read the entire file into memory.
	fileContent, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	// Step 2: Unmarshal the file content into a slice of LogEntry.
	var entries []LogEntry
	err = json.Unmarshal(fileContent, &entries)
	if err != nil {
		return nil, err
	}

	// Step 3: Filter for entries of type "alert".
	filteredEntries := make([]LogEntry, 0)
	for _, entry := range entries {
		if entry.Type == "alert" {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries, nil
}
