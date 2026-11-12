package main

import (
	"fmt"
	"os"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]registeredCommand
}

type registeredCommand struct {
	name     string
	callback func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.cmd[cmd.name]
	if !ok {
		return fmt.Errorf("command %q not found", cmd.name)
	}
	return command.callback(s, cmd)
}

func (c *commands) register(name string, f func(s *state, cmd command) error) error {
	_, ok := c.cmd[name]
	if ok {
		fmt.Fprintf(os.Stderr, "the command %s has already been registered", name)
	}
	c.cmd[name] = registeredCommand{
		name:     name,
		callback: f,
	}
	return nil
}
