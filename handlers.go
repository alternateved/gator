package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alternateved/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatal(err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user %s has been set\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	uuid := uuid.New()
	now := time.Now()
	params := database.CreateUserParams{
		ID:        uuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user was created: %v\n", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("users reset\n")
	return nil
}
