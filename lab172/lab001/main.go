package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	example()
}

func example() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("fsnotidy.NewWatcher error:%v", err)
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
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	//添加监控
	path := "../data"
	if err := watcher.Add(path); err != nil {
		log.Fatalf("watcher.Add error:%v", err)
	} else {
		log.Println("开始监控:", path)
	}

	<-done
	log.Println("ending...")
}
