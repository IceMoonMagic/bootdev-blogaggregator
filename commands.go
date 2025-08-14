package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	allcmds map[string]func(*state, command) error
}

func (cmds commands) run(s *state, cmd command) error {
	f, ok := cmds.allcmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return f(s, cmd)
}

func (cmds commands) register(name string, f func(*state, command) error) {
	cmds.allcmds[name] = f
}

func checkCommandArgsCount(cmd command, min int, max int) error {
	count := len(cmd.args)
	return checkArgsCount(count, min, max)
}

func checkArgsCount(count int, min int, max int) error {
	if count < min {
		return fmt.Errorf("too few argments: expected %d - %d arguments, got %d", min, max, count)
	}
	if count > max {
		return fmt.Errorf("too many argments: expected %d - %d arguments, got %d", min, max, count)
	}
	return nil
}
