package main

import (
	"context"
	"fmt"
	"os"
	"slices"
)

func handlerReset(s *state, cmd command) error {
	all := len(cmd.args) == 0

	if all || slices.Contains(cmd.args, "users") {
		if err := s.db.DeleteUsers(context.Background()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Cleared `users` Table")
	}
	return nil
}
