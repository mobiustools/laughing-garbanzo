package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Running...")
	time.Sleep(30 * time.Second)

	// Exit with code 1 if the current time is even, flaky!
	if time.Now().Unix()%2 == 0 {
		fmt.Println("Exiting with code: ", 1)
		os.Exit(1)
	}

	os.Exit(0)
}
