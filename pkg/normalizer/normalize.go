package normalizer

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type LogEntry struct {
	Timestamp string
	Alert     string
	Details   string
}

func ProcessNewLogEntries(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		// handle error
		return
	}
	defer file.Close()

	// Use a file offset or other mechanism to remember the last processed position
	// For simplicity, here we read the entire file, but in reality, you'll want to
	// only read new entries since the last processed position.

	entries, _ := readLogEntries(file)

	// Further processing can be done on the entries if needed.
	for _, entry := range entries {
		// For example: print them out for now
		log.Printf("%s: %s - %s", entry.Timestamp, entry.Alert, entry.Details)
	}
}

func readLogEntries(file *os.File) ([]LogEntry, error) {
	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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
