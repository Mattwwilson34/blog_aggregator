package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Mattwwilson34/blog_aggregator/internal/config"
	"github.com/Mattwwilson34/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, <login> <username>")
	}

	username := cmd.args[0]
	fmt.Printf("logging in as %s...\n", username)

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
	fmt.Printf("successfully logged in as %s\n", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
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
	fmt.Printf("attempting to register %s...\n", username)
	registeredUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to register %s\n", username)
		return err
	}
	fmt.Printf("user %s has been registered\n", registeredUser.Name)
	fmt.Printf("logging in as %s...\n", registeredUser.Name)
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("successfully logged in as %s\n", registeredUser.Name)

	return nil
}

func handlerUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to get users\n")
		return err
	}
	for _, user := range users {
		if s.config.CurrentUsername == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
			continue
		}
		fmt.Printf("%s\n", user.Name)
	}
	return nil
}

// reset user table and config
func handlerReset(s *state, _ command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to reset user table\n")
		return err
	}
	resetConfig := config.Config{
		DBURL:           s.config.DBURL,
		CurrentUsername: "",
	}
	err = config.Write(&resetConfig)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to reset config\n")
		return err
	}
	fmt.Print("user table has been reset\n")
	return nil
}
