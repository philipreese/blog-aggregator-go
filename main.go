package main

import (
	"log"
	"os"

	"github.com/philipreese/blog-aggregator-go/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state { cfg: &cfg }
	cmds := commands { registeredCommands: make(map[string]func(*state, command) error) }

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("usage: cli <command> [args...]")
	}

	if err:= cmds.run(s, command { Name: os.Args[1], Args: os.Args[2:] }); err != nil {
		log.Fatalf("error running command %s", os.Args[1])
	}
}