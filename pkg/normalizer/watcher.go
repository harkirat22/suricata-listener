package normalizer

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func WatchFile(logFilePath string, processFunc func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file:", event.Name)
					processFunc() // Call the provided function to process new log entries.
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
