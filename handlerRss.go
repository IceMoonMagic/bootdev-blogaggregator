package main

import (
	"context"
	"fmt"

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

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := state.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			Name:     name,
			Url:      url,
			UserName: state.cfg.CurrentUserName,
		},
	)

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return handlerFollow(
		state,
		command{
			name: "following",
			args: []string{
				url,
			},
		},
	)
}

func handlerFeeds(state *state, cmd command) error {
	feeds, err := state.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Feed:\n\tName:\t%s\n\tURL:\t%s\n\tUser:\t%s\n",
			feed.Name, feed.Url, feed.User.String,
		)
	}
	return nil
}
