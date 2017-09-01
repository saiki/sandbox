package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"os"
)

func main() {
	var dirs []string
	if len(os.Args) < 1 {
		dirs = []string{"."}
	} else {
		dirs = os.Args
	}
	ExampleNewWatcher(dirs)
}

func ExampleNewWatcher(dirs []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		log.Println("start watch.")
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				log.Println("event.Op:", event.Op)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, d := range dirs {
		err = watcher.Add(d)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}
