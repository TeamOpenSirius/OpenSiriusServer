package main

import (
	"log"
	"ymst_server/core"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
