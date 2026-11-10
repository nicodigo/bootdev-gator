package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nicodigo/blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: login <username>")
	}

	username := cmd.arguments[0]

	dbUser, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(dbUser.Name)
	if err != nil {
		return err
	}

	fmt.Println("Login successful! Username set to:", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: register <username>")
	}

	username := cmd.arguments[0]

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	currUser, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	s.cfg.SetUser(currUser.Name)
	fmt.Println("User registered successfuly!")
	fmt.Println()
	fmt.Println("Name:", currUser.Name)
	fmt.Println("ID:", currUser.ID)
	fmt.Println("Created at:", currUser.CreatedAt)
	fmt.Println("Updated at:", currUser.UpdatedAt)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}
	fmt.Println("All users have been deleted!")
	return nil
}
