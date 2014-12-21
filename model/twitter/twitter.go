package twitter

import (
	"github.com/satoshun/goutils/model/mgodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// for Tweet struct
const (
	TimeLayout = "Mon Jan 2 15:04:05 +0000 2006"
)

// Tweet fields
type Tweet struct {
	ID            int64    `bson:"_id" json:"id"`
	Text          string   `bson:"text" json:"text"`
	URL           string   `bson:"url" json:"url"`
	Hashtags      []string `bson:"hashtags" json:"hashtags"`
	RetweetCount  int      `bson:"retweet_count" json:"retweet_count"`
	FavoriteCount int      `bson:"favorite_count" json:"favorite_count"`
	CreatedAt     int64    `bson:"created_at" json:"created_at"`
}

// Update tweet
func (tweet *Tweet) Update() (err error) {
	err = CTweet().Update(bson.M{"_id": tweet.ID}, tweet)
	return
}

// GetTweets specified start, end time
func GetTweets(start int64, end int64) (tweets []Tweet) {
	CTweet().Find(bson.M{"created_at": bson.M{"$gt": start, "$lt": end}}).All(&tweets)
	return
}

// CTweet return Collection of tweet
func CTweet() *mgo.Collection {
	return mgodb.GetCollection("tweet")
}

// GetNotShortenURLTweets is already Longer URL
func GetNotShortenURLTweets() (tweets []Tweet, err error) {
	err = CTweet().Find(bson.M{"is_shorten_url": false}).Limit(3000).All(&tweets)
	return
}

// BulkTweet is return Bulk Session
func BulkTweet() *mgo.Bulk {
	return CTweet().Bulk()
}

// Insert tweet
func Insert(tweet Tweet) error {
	return CTweet().Insert(&tweet)
}
