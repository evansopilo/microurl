package service

import (
	"github.com/evansopilo/Micro-URL/pkg/db"
	"github.com/google/uuid"
)

type URLSRVC struct {
	*db.URLModel
}

func (u *URLSRVC) InsertURL(url db.URL) (db.URL, error) {

	url.ID = uuid.New().String()

	url, err := u.URLModel.Insert(url)

	if err != nil {
		return db.URL{}, err
	}

	return url, nil
}

func (u *URLSRVC) ReadURL(id string) (db.URL, error) {

	url, err := u.URLModel.ReadByID(id)

	if err != nil {
		return db.URL{}, err
	}

	return url, nil
}

func (u *URLSRVC) ReadAllURLs() ([]db.URL, error) {

	urls, err := u.URLModel.ReadAll()

	if err != nil {
		return []db.URL{}, err
	}

	return urls, nil
}

func (u *URLSRVC) UpdateURL(id string, url db.URL) (db.URL, error) {

	url, err := u.URLModel.Update(id, url)

	if err != nil {
		return db.URL{}, err
	}

	return url, nil
}

func (u *URLSRVC) DeleteURL(id string) error {

	err := u.URLModel.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
