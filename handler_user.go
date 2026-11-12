package main

import (
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, <login> <username>")
	}

	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "current user set to: %q\n", username)

	return nil
}
