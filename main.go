package main

import (
	"fmt"
	"os"

	"github.com/Mattwwilson34/blog_aggregator/internal/config"
)

const DBURL = "postgres://example"

type state struct {
	config   config.Config
	commands commands
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading config")
	}
	state := state{
		config: cfg, commands: commands{
			cmd: map[string]registeredCommand{},
		},
	}
	err = state.commands.register("login", handlerLogin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "cli requires at least 1 argument")
		os.Exit(1)
	}

	userCommand := command{
		name: args[1],
		args: args[2:],
	}

	config.Print()
	err = state.commands.run(&state, userCommand)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	config.Print()
}
