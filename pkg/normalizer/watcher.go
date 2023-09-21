package normalizer

import (
	"fmt"
	"io"
	"os"

	"github.com/fsnotify/fsnotify"
)

func WatchLog(filePath string, processNewEntries func([]LogEntry)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Go to the end of the file immediately upon startup to only get new entries.
	file.Seek(0, io.SeekEnd)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					entries, err := ReadLogEntries(file)
					if err != nil {
						// Consider logging or handling this error
						fmt.Println("Error reading log entries:", err)
					} else if len(entries) > 0 {
						processNewEntries(entries)
					}

					// Reset the file pointer to the end for the next entries.
					file.Seek(0, io.SeekEnd)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		fmt.Println("Error adding file to watcher:", err)
		return
	}
	<-done
}
