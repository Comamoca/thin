package thin

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
)

func auth() Client {
  godotenv.Load()

	keys := ApiKeys{os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET")}
	return keys.Auth()
}

func gen() string {
	v := url.Values{}
	v.Set("q", "#golang")
	url := "https://api.twitter.com/1.1/search/tweets.json?"
	tu := ThinUrl{url, v}

	url, err := GenerateUrl(tu)
	if err != nil {
		log.Fatal(err)
	}

	return url
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Notthing .env")
	}
}

func TestAuth(t *testing.T) {
	_ = auth()
}

func TestGenerateUrl(t *testing.T) {
	_ = gen()
}

func TestGet(t *testing.T) {
  url := "https://api.twitter.com/1.1/account/settings.json"
	client := auth()
	// log.Print(url)
	jsn, err := client.Get(url)
	if err != nil {
		t.Errorf("Client not accessed Twitter API\n %v", err)
	}

  if gjson.Get(jsn, "screen_name").String() != "Comamoca_" {
    t.Error("Incorrect value identified")
  }
}
