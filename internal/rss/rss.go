package rss

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"

	"github.com/icemoonmagic/bootdev-blogaggregator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "gator")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var feed RSSFeed
	decoder := xml.NewDecoder(response.Body)
	if err := decoder.Decode(&feed); err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	return &feed, nil
}

func ScrapeFeeds(db *database.Queries) error {
	next, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err := FetchFeed(context.Background(), next.Url)
	if err != nil {
		return err
	}

	db.MarkFeedFetched(context.Background(), next.ID)

	for _, post := range feed.Channel.Item {
		if err := db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				Title:       post.Title,
				Url:         post.Link,
				Description: post.Description,
				PublishedAt: post.PubDate,
				FeedID:      next.ID,
			},
		); err != nil {
			return err
		}
	}

	return nil
}
