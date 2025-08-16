package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/icemoonmagic/bootdev-blogaggregator/internal/database"
	"github.com/icemoonmagic/bootdev-blogaggregator/internal/rss"
)

func handlerAgg(state *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(state *state, cmd command) error {
	if err := checkCommandArgsCount(cmd, 2, 2); err != nil {
		return err
	}

	user, err := state.db.GetUser(context.Background(), state.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := state.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
