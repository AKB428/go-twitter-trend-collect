package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	api := getTwitterApi()
	v := url.Values{}
	trendResp, err := api.GetTrendsByPlace(23424856, v)

	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()

	fmt.Println(trendResp.AsOf)
	fmt.Println(trendResp.CreatedAt)
	fmt.Println(trendResp.Locations)
	fmt.Println(len(trendResp.Trends))
	fmt.Println(t)

	// https://pkg.go.dev/github.com/chimeracoder/anaconda#TrendResponse
	// https://pkg.go.dev/github.com/chimeracoder/anaconda#Trend
	for i, v := range trendResp.Trends {
		ranking := i + 1
		fmt.Printf("%d %s\n", ranking, v.Name)
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
