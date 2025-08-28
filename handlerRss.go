package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/icemoonmagic/bootdev-blogaggregator/internal/database"
	"github.com/icemoonmagic/bootdev-blogaggregator/internal/rss"
)

func handlerAgg(state *state, cmd command) error {
	if err := checkCommandArgsCount(cmd, 1, 1); err != nil {
		return err
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		if err := rss.ScrapeFeeds(state.db); err != nil {
			return err
		}
	}
}

func handlerAddFeed(state *state, cmd command, user database.User) error {
	if err := checkCommandArgsCount(cmd, 2, 2); err != nil {
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := state.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			Name:   name,
			Url:    url,
			UserID: user.ID,
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
		user,
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

func handlerBrowse(state *state, cmd command, user database.User) error {
	if err := checkCommandArgsCount(cmd, 0, 1); err != nil {
		return err
	}

	var limit32 int32
	limit32 = 2
	if len(cmd.args) > 0 {
		limit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit32 = int32(limit)
	}

	posts, err := state.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  limit32,
		},
	)
	if err != nil {
		return nil
	}

	for _, post := range posts {
		fmt.Printf(
			"%v\t%v\t%v\n\t%v\n\t%v\n",
			post.PublishedAt,
			post.FeedName,
			post.Title,
			post.Description,
			post.Url,
		)
	}

	return nil
}
