package main

import (
	"log"
	"os"
	"encoding/json"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Usage: framecc input.frame")
    }

	file, err := parseFrameDefinition(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
	jsonStr, _ := json.Marshal(file)
    log.Printf("%v", string(jsonStr))
}
