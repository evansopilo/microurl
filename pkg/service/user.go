package service

import (
	"github.com/evansopilo/Micro-URL/pkg/db"
	"github.com/google/uuid"
)

type USERSRVC struct {
	db.UserModel
}

func (u *USERSRVC) InsertUser(user db.User) (db.User, error) {

	user.ID = uuid.New().String()

	user, err := u.UserModel.Insert(user)

	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *USERSRVC) ReadUser(id string) (db.User, error) {

	user, err := u.UserModel.ReadByID(id)

	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *USERSRVC) ReadAllUsers() ([]db.User, error) {

	users, err := u.UserModel.ReadAll()

	if err != nil {
		return []db.User{}, err
	}

	return users, nil
}

func (u *USERSRVC) UpdateUser(id string, user db.User) (db.User, error) {

	user, err := u.UserModel.Update(id, user)

	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *USERSRVC) DeleteUser(id string) error {

	err := u.UserModel.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
