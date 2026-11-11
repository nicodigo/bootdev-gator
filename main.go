package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/nicodigo/blog-aggregator/internal/config"
	"github.com/nicodigo/blog-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db *database.Queries
}

func main() {
	myConfig, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading config: %v\n", err)
		os.Exit(1)
	}

	programState := state{cfg: myConfig}
	cmds := loadCommands()

	db, err := sql.Open("postgres", myConfig.DbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening database: %v\n", err)
		os.Exit(1)
	}

	programState.db = database.New(db)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: <command name> [args]")
		os.Exit(1)
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func loadCommands() *commands {
	cmds := commands{
		cmdHandlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", handleAddFeed)
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", handleFollow)
	cmds.register("following", handleFollowing)
	return &cmds
}
