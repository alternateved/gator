package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command %s doesn't exist", cmd.name)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error {
	_, ok := c.handlers[name]
	if ok {
		return fmt.Errorf("command %s already registered", name)
	}

	c.handlers[name] = f
	return nil
}
