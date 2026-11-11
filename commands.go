package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.cmd[cmd.name]
	if !ok {
		fmt.Println("command not found")
	}
	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(s *state, cmd command) error) error {
	_, ok := c.cmd[name]
	if ok {
		fmt.Printf("the command %s has already been registered", name)
	}
	c.cmd["name"] = f
	return nil
}
