package db

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestUserModel_Insert(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserModel{
				DB: tt.fields.DB,
			}
			got, err := u.Insert(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserModel.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserModel_ReadByID(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserModel{
				DB: tt.fields.DB,
			}
			got, err := u.ReadByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.ReadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserModel.ReadByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserModel_ReadAll(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	tests := []struct {
		name    string
		fields  fields
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserModel{
				DB: tt.fields.DB,
			}
			got, err := u.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserModel.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserModel_Update(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	type args struct {
		id   string
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserModel{
				DB: tt.fields.DB,
			}
			got, err := u.Update(tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserModel.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserModel_Delete(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserModel{
				DB: tt.fields.DB,
			}
			if err := u.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
