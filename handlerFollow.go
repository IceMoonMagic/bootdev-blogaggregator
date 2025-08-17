package main

import (
	"context"
	"fmt"

	"github.com/icemoonmagic/bootdev-blogaggregator/internal/database"
)

func handlerFollow(state *state, cmd command) error {
	if err := checkCommandArgsCount(cmd, 1, 1); err != nil {
		return nil
	}

	url := cmd.args[0]

	follow, err := state.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			Name: state.cfg.CurrentUserName,
			Url:  url,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(follow)

	return nil
}

func handlerFollowing(state *state, cmd command) error {
	follows, err := state.db.GetFeedFollowsForUser(
		context.Background(),
		state.cfg.CurrentUserName,
	)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.Name)
	}

	return nil
}
