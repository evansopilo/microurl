package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string
	UserName string
	Password string
}

type UserModel struct {
	DB *mongo.Collection
}

func (u *UserModel) Insert(user User) (User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if _, err := u.DB.InsertOne(ctx, &user); err != nil {
		return User{}, err
	}

	user, err := u.ReadByID(user.ID)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UserModel) ReadByID(id string) (User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	var user User

	if err := u.DB.FindOne(ctx, filter).Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UserModel) ReadAll() ([]User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{}

	filterCursor, err := u.DB.Find(ctx, filter)

	if err != nil {
		return []User{}, err
	}

	var users []User

	for filterCursor.Next(ctx) {

		var user User
		if err := filterCursor.Decode(&user); err != nil {
			return []User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserModel) Update(id string, user User) (User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.D{{"$set", user}}

	if _, err := u.DB.UpdateOne(ctx, filter, update); err != nil {
		return User{}, err
	}

	user, err := u.ReadByID(id)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UserModel) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	if _, err := u.DB.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
