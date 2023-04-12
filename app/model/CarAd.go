package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// CarAd structure for our blog
type CarAd struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Km       string    `json:"km"`
	Location string    `json:"location"`
	Date     time.Time `bson:"date,omitempty"`
	Price    string    `json:"price"`
	Url      string    `json:"url"`
	Img_data []byte    `json:"img_data"`
}

func GetAllPosts() ([]CarAd, error) {

	var posts []CarAd

	session, err := db.StartSession()
	if err != nil {
		return posts, err
	}

	defer session.EndSession(context.Background())

	cursor, err := session.Client().Database("db_name").Collection("fbMarketplace").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var post CarAd
		if err := cursor.Decode(&post); err != nil {
			fmt.Print("ERROR", err.Error())
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}

// func GetPost(id uint64) (CarAd, error) {

// }

// func CreatePost(post CarAd) error {
// }

// func UpdatePost(post CarAd) error {

// }

// func DeletePost(id uint64) error {

// }
