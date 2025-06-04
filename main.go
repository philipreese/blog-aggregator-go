package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/philipreese/blog-aggregator-go/internal/config"
	"github.com/philipreese/blog-aggregator-go/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error connecting to database")
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := commands { registeredCommands: make(map[string]func(*state, command) error) }

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

	if len(os.Args) < 2 {
		log.Fatalf("usage: cli <command> [args...]")
	}

	err = cmds.run(programState, command { Name: os.Args[1], Args: os.Args[2:] })
	if err != nil {
		log.Fatal(err)
	}
}