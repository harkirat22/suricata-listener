package normalizer

import (
	"bufio"
	"os"
	"strings"
)

// LogEntry represents a structured format of the Suricata fast.log entry.
type LogEntry struct {
	Timestamp string
	Alert     string
	Details   string
}

// Normalize will take the path to the fast.log, read it, and return structured log entries.
func Normalize(logFilePath string) ([]LogEntry, error) {
	file, err := os.Open(logFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split and structure the log entry. This is a basic split and will need adjustments based on the exact format of fast.log.
		parts := strings.SplitN(line, " ", 3)
		if len(parts) >= 3 {
			entry := LogEntry{
				Timestamp: parts[0],
				Alert:     parts[1],
				Details:   parts[2],
			}
			entries = append(entries, entry)
		}
	}
	return entries, scanner.Err()
}
