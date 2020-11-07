package util

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

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
					time.Sleep(3000 * time.Millisecond)
					pathf := strings.ReplaceAll(event.Name, `\`, "/")
					file, err := os.Open("broadcast_dir/" + path.Base(pathf))
					if err != nil {
						fmt.Println(err)
					}
					s := bufio.NewScanner(file)
					file.Close()
					sum := sha256.Sum256(s.Bytes())
					client.Files[path.Base(pathf)] = sum
					client.BroadCast(path.Base(pathf))
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					pathf := strings.ReplaceAll(event.Name, `\`, "/")
					fmt.Println(pathf)
					delete(client.Files, path.Base(pathf))
					fmt.Println(client.Files)
				} else {
					fmt.Println(event.Op)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(gwd() + "/broadcast_dir")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
