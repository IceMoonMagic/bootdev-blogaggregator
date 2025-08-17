package main

import (
	"context"
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if err := checkCommandArgsCount(cmd, 1, 1); err != nil {
		return err
	}

	name := cmd.args[0]

	user, err := s.db.CreateUser(
		context.Background(),
		name,
	)
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("Registered user %v\n", user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if err := checkCommandArgsCount(cmd, 1, 1); err != nil {
		return err
	}

	user, err := s.db.GetUser(
		context.Background(),
		cmd.args[0],
	)
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("Set user to `%s`\n", cmd.args[0])

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if err := checkCommandArgsCount(cmd, 0, 0); err != nil {
		return err
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("- %s (current)\n", user.Name)
		} else {
			fmt.Printf("- %s\n", user.Name)
		}
	}

	return nil
}
