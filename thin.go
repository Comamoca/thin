// Very thin Twitter API client
package thin

import (
	"io/ioutil"

	"github.com/dghubble/oauth1"

	"net/http"
	"net/url"
)

type Client struct {
  client *http.Client
}

type ApiKeys struct {
	ConsumerKey    string
	ConsumerSecret string

	AccessToken       string
	AccessTokenSecret string
}

type ThinUrl struct {
	RawUrl string
	Value  url.Values
}

func GenerateUrl(tu ThinUrl) (string, error) {
	url, err := url.Parse(tu.RawUrl + tu.Value.Encode())
	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (clt Client) Get (url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	response, err := clt.client.Do(request)
	if err != nil {
		return "", err
	}

	barray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	response.Body.Close()

	return string(barray), nil
}

func (keys ApiKeys) Auth() Client {
	config := oauth1.NewConfig(keys.ConsumerKey, keys.ConsumerSecret)
	token := oauth1.NewToken(keys.AccessToken, keys.AccessTokenSecret)

	clt := config.Client(oauth1.NoContext, token)

  client := Client{clt}

	return client
}
