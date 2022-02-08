package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	LongURL  string `json:"long_url,omitempty" bson:"long_url,omitempty"`
	ShortURL string `json:"short_url,omitempty" bson:"short_url,omitempty"`
	UserID   string `json:"user_id,omitempty"`
}

type URLModel struct {
	DB *mongo.Collection
}

func (u *URLModel) Insert(url URL) (URL, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if _, err := u.DB.InsertOne(ctx, &url); err != nil {
		return URL{}, err
	}

	url, err := u.ReadByID(url.ID)

	if err != nil {
		return URL{}, err
	}

	return url, nil
}

func (u *URLModel) ReadByID(id string) (URL, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	var url URL

	if err := u.DB.FindOne(ctx, filter).Decode(&url); err != nil {
		return URL{}, err
	}

	return url, nil
}

func (u *URLModel) ReadAll() ([]URL, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{}

	filterCursor, err := u.DB.Find(ctx, filter)

	if err != nil {
		return []URL{}, err
	}

	var urls []URL

	for filterCursor.Next(ctx) {

		var url URL
		if err := filterCursor.Decode(&url); err != nil {
			return []URL{}, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func (u *URLModel) Update(id string, url URL) (URL, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.D{{"$set", url}}

	if _, err := u.DB.UpdateOne(ctx, filter, update); err != nil {
		return URL{}, err
	}

	url, err := u.ReadByID(id)

	if err != nil {
		return URL{}, err
	}

	return url, nil
}

func (u *URLModel) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	if _, err := u.DB.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
