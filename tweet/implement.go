package tweet

import (
	"botqepo/file"
	"botqepo/zenquotes"
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

func SearchTweets(client *twitter.Client) (twitter.Tweet, error) {
	readID := file.ReadLatestID()

	tweetSearch, _, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{
		Count:   1,
		SinceID: readID,
	})
	if err != nil {
		log.Printf("error search tweet: %s", err)
		return twitter.Tweet{}, err
	}

	if len(tweetSearch) == 0 {
		return twitter.Tweet{}, err
	}

	return tweetSearch[0], nil
}

func PostOpenaiTweet(client *twitter.Client, status string, tweet twitter.Tweet) error {
	_, _, err := client.Statuses.Update(fmt.Sprintf("@%s %s", tweet.User.ScreenName, status), &twitter.StatusUpdateParams{
		InReplyToStatusID: tweet.ID,
	})
	if err != nil {
		log.Println("Error PostOpenaiTweet: ", err)
		return err
	}

	return nil
}

func PostZenquotesTweet(client *twitter.Client, quote zenquotes.Quote) error {
	_, _, err := client.Statuses.Update(fmt.Sprintf("Motivational / Inspirational Quote Today:\n\n(Author) %s: %s", quote.A, quote.Q), &twitter.StatusUpdateParams{})
	if err != nil {
		log.Println("Error PostZenquotesTweet: ", err)
		return err
	}

	return nil
}
