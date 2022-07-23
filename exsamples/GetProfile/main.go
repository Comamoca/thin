package main

import (
	"fmt"
	"os"

	"github.com/Comamoca/thin"
	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
)

func getTweetList(jsn string) []string  {
  var tweets []string

  for _, tweet := range gjson.Get(jsn, "statuses").Array() {
    tweets = append(tweets, gjson.Get(tweet.String(), "text").String())
  }
  return tweets
}

func main() {
	godotenv.Load("../../.env")

	key := thin.ApiKeys{
		ConsumerKey:       os.Getenv("APIKEY"),
		ConsumerSecret:    os.Getenv("APIKEYSECRET"),
		AccessToken:       os.Getenv("ACTOKEN"),
		AccessTokenSecret: os.Getenv("ACTOKEN_SECRET"),
	}

  client := key.Auth()

  endp := "https://api.twitter.com/1.1/account/settings.json"
  resp, _ := client.Get(endp)

  name := gjson.Get(resp, "screen_name")
  lang := gjson.Get(resp, "language")

  fmt.Printf("User Id: %s\nLanguage: %s", name, lang)
}
