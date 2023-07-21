package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var commandFlag = flag.String("command", "", "command to run when a file changes")

func main() {
	flag.Parse()
	if *commandFlag == "" {
		log.Fatal("command flag must be set")
	}
	args := strings.Split(*commandFlag, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					cmd.Process.Kill()
					cmd = exec.Command(args[0], args[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Start()
				}
				if event.Has(fsnotify.Create) {
					log.Println("created file:", event.Name)
					watcher.Add(event.Name)
					cmd.Process.Kill()
					cmd = exec.Command(args[0], args[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Start()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	recursivleyAddWatchers(watcher, ".")
	cmd.Start()
	// Block main goroutine forever.
	<-make(chan struct{})
}

func recursivleyAddWatchers(watcher *fsnotify.Watcher, path string) {
	err := watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			recursivleyAddWatchers(watcher, path+"/"+file.Name())
		}
	}
}
