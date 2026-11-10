package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdHandlers map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	handler, ok := cmds.cmdHandlers[cmd.name]
	if !ok {
		return fmt.Errorf("command '%s' not found\n", cmd.name)
	}
	return handler(s, cmd)
}

func (cmds *commands) register(name string, h func(*state, command) error) {
	cmds.cmdHandlers[name] = h
}
