package main

import (
	"botqepo/config"
	"botqepo/file"
	"botqepo/openai"
	"botqepo/tweet"
	"botqepo/zenquotes"
	"log"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	cfg := *config.Init()
	config := oauth1.NewConfig(cfg.TwitterConsumerKey, cfg.TwitterConsumerSecret)
	token := oauth1.NewToken(cfg.TwitterAccessToken, cfg.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	openaiCfg := openai.CfgOpenai{
		BaseUrl: cfg.OpenaiBaseUrl,
		Model:   cfg.OpenaiModel,
		Token:   cfg.OpenaiToken,
	}
	zenquotesCfg := zenquotes.CfgZenquotes{
		BaseUrl: cfg.ZenquotesBaseUrl,
	}
	for {
		tweets, err := tweet.SearchTweets(client)
		if err != nil {
			if err.Error() == "twitter: 88 Rate limit exceeded" {
				log.Println("Rate limit exceeded, waiting for 15 minutes")
				time.Sleep(time.Second * 60 * 15)
			}
		}

		if tweets.ID != 0 {
			var answers string
			if cfg.IsOpenAIEnable {
				openaiResp := openai.Completions(openaiCfg, strings.Replace(tweets.Text, "@botqepo", "", -1))
				answers = strings.Replace(openaiResp.Choices[0].Text, "\n", " ", -1)
				answers = strings.Replace(answers, "A:", "", -1)
			}
			err = tweet.PostOpenaiTweet(client, answers, tweets)
			if err != nil {
				log.Println("Error PostTweet: ", err)
			} else {
				file.WriteLatestID(tweets.ID)
			}
		}

		if cfg.IsZenQuotesEnable {
			latestQuote := file.ReadLatestQuote()

			// 21600 is 6 hours in seconds
			if latestQuote == 0 || time.Now().Unix()-latestQuote > 21600 {
				quote, err := zenquotes.GetQuote(zenquotesCfg)
				if err != nil {
					log.Println("Error GetQuote: ", err)
				}

				if quote.Q != "" {
					err = tweet.PostZenquotesTweet(client, quote)
					if err != nil {
						log.Println("Error PostTweet: ", err)
					}

					file.WriteLatestQuote(time.Now().Unix())
				}
			}
		}

		time.Sleep(time.Second * 40)
	}
}
