package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/AKB428/go-twitter-trend-collect/model"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := getTwitterApi()
	v := url.Values{}
	trendResp, err := api.GetTrendsByPlace(23424856, v)

	if err != nil {
		log.Fatal(err)
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	now := time.Now().In(jst)

	fmt.Println(trendResp.AsOf)
	fmt.Println(trendResp.CreatedAt)
	fmt.Println(trendResp.Locations)
	fmt.Println(len(trendResp.Trends))
	fmt.Println(now)

	// https://pkg.go.dev/github.com/chimeracoder/anaconda#TrendResponse
	// https://pkg.go.dev/github.com/chimeracoder/anaconda#Trend

	var twitterTrends []model.TwitterTrend

	for i, v := range trendResp.Trends {
		ranking := i + 1
		fmt.Printf("%d %s\n", ranking, v.Name)

		t := model.TwitterTrend{Name: v.Name, Rank: int32(ranking), CreatedAt: now, UpdatedAt: now}
		twitterTrends = append(twitterTrends, t)
	}

	db := gormConnect()
	//https://gorm.io/docs/create.html
	db.Create(&twitterTrends)

}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func gormConnect() *gorm.DB {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	if len(dbHost) == 0 {
		dbUser = "root"
	}

	if len(dbUser) > 0 {
		dbPass = ":" + dbPass
	}

	if len(dbHost) == 0 {
		dbHost = "localhost"
	}

	db, err := gorm.Open(mysql.Open(dbUser + dbPass + "@" + "tcp(" + dbHost + ")/" + dbName + "?parseTime=true"))
	if err != nil {
		panic(err.Error())
	}

	return db
}
