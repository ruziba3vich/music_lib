package main

import (
	"io"
	"log"
	"os"

	"github.com/ruziba3vich/music_lib/cmd/app"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "[MusicLib] ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Fatal(app.Run(logger))
}
