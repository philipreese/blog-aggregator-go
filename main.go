package main

import (
	"fmt"
	"log"

	cfg "github.com/philipreese/blog-aggregator-go/internal/config"
)

func main() {
	config, err := cfg.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = config.SetUser("Philip")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	config, err = cfg.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err);
	}
	fmt.Println(config)
}