package util

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/p1ck0/broadcast_dir/client"
)

func Notify(client *client.Client) {
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
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					pathf := strings.ReplaceAll(event.Name, `\`, "/")
					fmt.Println(pathf)
					client.BroadCast(path.Base(pathf))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("C:/Users/хихи/broadcast_dir/broadcast_dir")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
