package main

import (
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewStorage(config.S3)
	if err != nil {
		log.Fatal(err)
	}

	w, err := Watcher(config.Source, config.Delay)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Watching %s for new files...", config.Source)
	for {
		name, more := <-w
		if !more {
			break
		}
		path, err := Resize(config.Resize, config.Source, config.Target, name)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		log.Printf("Uploading %s...", path)
		url, err := s.Upload(path)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		err = PostContent(config.Webhook, url)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
	}
}
