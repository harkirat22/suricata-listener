package normalizer

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type LogEntry struct {
	Type      string `json:"event_type"`
	Timestamp string `json:"timestamp"`
	SrcIP     string `json:"src_ip"`
	Alert     Alert  `json:"alert"`
}

type Alert struct {
	Signature string `json:"signature"`
}

// ReadLogEntries reads new log entries from a file starting from the given position and returns them as a slice.
func ReadLogEntries(file *os.File, startPos int64) ([]LogEntry, int64, error) {
	var entries []LogEntry

	file.Seek(startPos, io.SeekStart)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var logEntry LogEntry
		err := json.Unmarshal([]byte(line), &logEntry)
		if err != nil {
			// you may decide to skip the line and continue or handle it differently
			continue
		}
		if logEntry.Type == "alert" {
			entries = append(entries, logEntry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, startPos, err
	}

	newPosition, _ := file.Seek(0, io.SeekCurrent)
	return entries, newPosition, nil
}
