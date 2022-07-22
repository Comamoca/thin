<div align="center">
<h1>🪶 thin</h1>

Very thin Twitter API client
</div>

thin is a Twitter library made by [Go]() that supports [Twitter API v1.1](). It is designed to have fewer functions than the library has.

Since thin does not have Json parsing function, you need to parse Json by yourself.
[gjson]() is recommended for Json parser.

## 🚀 Install
```sh
go get github.com/Comamoca/thin
```

## ❓ How to use

```go

package main

import (
	"github.com/Comamoca/thin"
	"github.com/tidwall/gjson"
)

func main(){
	// Set Keys
	keys := ApiKeys{os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET")}

	// get client
	client := thin.Auth(keys)

	// Set query
	v := url.Values{}
	v.Set("q", "#golang")

	// Twitter API endpoint
	endp := "https://api.twitter.com/1.1/search/tweets.json?"

	tu := ThinUrl{url, v}
	// Generate URL
	url, _ := GenerateUrl(url)

	tweets, _ := client.Get(url)
	for _, tweet := range gjson.Get(res, "statuses").Array() {
			fmt.Println(gjson.Get(tweet.String(), "text"))
			fmt.Println("---------------------")
	}

}
```

👀 See the [Exsamples](./_exsample/)

## 🍪 Types

## thin.Keys
```go

type ApiKeys struct {
	ConsumerKey    string
	ConsumerSecret string

	AccessToken       string
	AccessTokenSecret string
}

```

## thin.ThinUrl
```go
type ThinUrl struct {
	RawUrl string
	Value  url.Values
}
```

## thin.Client
```go
type Client struct {
  client *http.Client
}
```

## 🧩 Functions

### Auth

`func (keys ApiKeys) Auth() Client `

Authenticate based on `thin.ApiKey`.

### GenerateUrl

`func GenerateUrl(tu ThinUrl) (string, error)`

Generates a parsed URL based on `thin.ThinUrl`. This is useful when specifying parameters.

### Get

`func (clt Client) Get (url string) (string, error)`

Execute a Get request based on the URL.

## 💡 Source of ideas

[RubyでTwitter API v2を叩けるgemをつくりました](https://zenn.dev/yhara/articles/21e496263108ae#1.-api%E3%82%92%E8%AA%BF%E3%81%B9%E3%82%8B)
