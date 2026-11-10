package main

import (
	"fmt"
	"os"

	"github.com/nicodigo/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	myConfig, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading config: %v\n", err)
		os.Exit(1)
	}

	programState := state{cfg: myConfig}
	cmds := loadCommands()

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
	return &cmds
}
