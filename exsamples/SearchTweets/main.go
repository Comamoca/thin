package exsample

import (
	"fmt"
	"net/url"
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

  query := url.Values{}
  query.Set("q", "#Twitter")

  endp := "https://api.twitter.com/1.1/search/tweets.json?"
  tu := thin.ThinUrl{RawUrl: endp, Value: query}
  url, _ := thin.GenerateUrl(tu)

  resp, _ := client.Get(url)
  for idx, txt := range getTweetList(resp) {
    fmt.Printf("%d: %s\n", idx, txt)
  }
}
