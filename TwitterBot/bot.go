package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Credentials struct necessary to authenticate the account.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

/* Helper function that returns a twitter client.
   It recieves a pointer with the necessary credentials to authenticate the account
   and returns the pointer to the twitter client and an error. */
func getClient(creds *Credentials) (*twitter.Client, error) {

	//Consumer key (API key) and consumerSecret (API secret).
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	//AccessToken and AccessTokenSecret.
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	//Verifies credentials.
	verifyCreds := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyCreds)
	if err != nil {
		return nil, err
	}
	if user != nil {
		fmt.Println("Worked Correctly")
		fmt.Println()
	}

	return client, nil

}

//Tweets in the account provided the credentials have been approved.
func tweet(twit string, client *twitter.Client) {

	tweet, resp, err := client.Statuses.Update(twit, nil)

	if err != nil {
		log.Println(err)
	}
	if resp != nil {
		fmt.Println("Correctly tweeted")
	}

	fmt.Println(tweet.Text)

}

//Searches tweets containing the searchWord.
func search(searchWord string, client *twitter.Client) {

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: searchWord,
	})

	if err != nil {
		log.Print(err)
	}
	if resp != nil {
		fmt.Println("Correctly searched")
		fmt.Println()
	}

	tweets := search.Statuses
	for _, tweet := range tweets {
		fmt.Println(tweet.Text)
		fmt.Println()
	}

}

func main() {

	fmt.Println("TwitterBot V1.1")

	//Env keys must be set previously to work.
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}
	client, err := getClient(&creds)
	if err != nil {
		log.Println(err)
		log.Println("Error finding twitter client")
	}

	tweet("Testing Bot", client)
	search("Elon Musk", client)

}
