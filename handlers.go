package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user %s has been set\n", cmd.args[0])
	return nil
}
