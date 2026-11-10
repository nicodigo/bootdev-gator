package main

import "fmt"

func handlerLogin(s *state, cmd command) error{
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: login <username>")
	}

	username := cmd.arguments[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("Login successful! Username set to:", username)
	return  nil
}
