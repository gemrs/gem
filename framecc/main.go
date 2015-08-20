package main

import (
	"fmt"
	"os"
	"encoding/json"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Print("Usage: framecc input.frame")
		os.Exit(1)
    }

	file, err := parseFrameDefinition(os.Args[1])
    if err != nil {
        fmt.Printf("%v\n", err)
		os.Exit(1)
    }
	//	jsonStr, _ := json.Marshal(file)
	jsonStr, _ := json.MarshalIndent(file, "", "  ")
    fmt.Printf("%v", string(jsonStr))
}
