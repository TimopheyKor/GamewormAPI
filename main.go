package main

import (
	"flag"
	"fmt"
)

// TODO: Create functionalityh to send requests to Google Docs
func main() {
	lvlFlag := flag.String("l", "info", "description of log level")

	flag.Parse()

	fmt.Printf("level is %s\n", *lvlFlag)
	fmt.Println("Gameworm API Blank Commit")
	fmt.Println("Hello, World!")
}
