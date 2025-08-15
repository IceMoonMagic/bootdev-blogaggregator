package main

import (
	"context"
	"fmt"

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
