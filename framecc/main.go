package main

import (
	"log"
	"os"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Usage: framecc input.frame")
    }

	frames, err := parseFrameDefinition(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%+v", frames)
}
