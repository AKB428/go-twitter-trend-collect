package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	api := getTwitterApi()
	v := url.Values{}
	trendResp, err := api.GetTrendsByPlace(23424856, v)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(trendResp.AsOf)
	fmt.Println(trendResp.CreatedAt)
	fmt.Println(trendResp.Locations)
	fmt.Println(len(trendResp.Trends))

	// https://pkg.go.dev/github.com/chimeracoder/anaconda#TrendResponse
	// https://pkg.go.dev/github.com/chimeracoder/anaconda#Trend
	for _, v := range trendResp.Trends {
		fmt.Printf("%s %s %s\n", v.Name, v.Query, v.Url)
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
