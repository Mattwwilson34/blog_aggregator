package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Mattwwilson34/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, <login> <username>")
	}

	username := cmd.args[0]
	fmt.Fprintf(os.Stderr, "logging in as %s...\n", username)

	// check db for user
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unknown user: %q\n", username)
		return err
	}

	// update config
	err = s.config.SetUser(user.Name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to set user in config to %s\n", user.Name)
		return err
	}
	fmt.Fprintf(os.Stderr, "successfully logged in as %s\n", user.Name)
	return nil
}

func hanlderRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, <register> <username>")
	}
	username := cmd.args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	fmt.Fprintf(os.Stderr, "attempting to register %s...\n", username)
	registeredUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to regsiter %s\n", username)
		return err
	}
	fmt.Fprintf(os.Stderr, "user %s has been registered\n", registeredUser.Name)
	fmt.Fprintf(os.Stderr, "logging in as %s...\n", registeredUser.Name)
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "successfully logged in as %s\n", registeredUser.Name)

	return nil
}
