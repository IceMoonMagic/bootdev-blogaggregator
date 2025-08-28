package main

import (
	"context"
	"fmt"

	"github.com/icemoonmagic/bootdev-blogaggregator/internal/database"
)

func handlerFollow(state *state, cmd command, user database.User) error {
	if err := checkCommandArgsCount(cmd, 1, 1); err != nil {
		return nil
	}

	url := cmd.args[0]

	follow, err := state.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			UserID: user.ID,
			Url:    url,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(follow)

	return nil
}

func handlerFollowing(state *state, cmd command, user database.User) error {
	follows, err := state.db.GetFeedFollowsForUser(
		context.Background(),
		user.ID,
	)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.Name)
	}

	return nil
}

func handlerUnfollow(state *state, cmd command, user database.User) error {
	if err := checkCommandArgsCount(cmd, 1, 1); err != nil {
		return nil
	}

	url := cmd.args[0]

	err := state.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			Url:    url,
		},
	)
	return err
}
