package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watcher(path string, delay time.Duration) (chan string, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}

	nameChan := make(chan string)

	go func() {
	WatchLoop:
		for {
			select {
			case event, more := <-watcher.Events:
				if !more {
					break WatchLoop
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					name, err := filepath.Rel(path, event.Name)
					if err != nil {
						log.Printf("Watcher: error: %v", err)
						break WatchLoop
					}
					// Wait for writes to finish
					// TODO: Non-blocking timers
					time.Sleep(delay)
					nameChan <- name
					log.Printf("Watcher: file %s created", name)
				}
			case err, more := <-watcher.Errors:
				if !more {
					break WatchLoop
				}
				log.Printf("Watcher: error: %v", err)
			}
		}
		watcher.Close()
		close(nameChan)
	}()

	return nameChan, nil
}
