package service

import (
	"reflect"
	"testing"

	"github.com/evansopilo/Micro-URL/pkg/db"
)

func TestUSERSRVC_InsertUser(t *testing.T) {
	type fields struct {
		UserModel db.UserModel
	}
	type args struct {
		user db.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &USERSRVC{
				UserModel: tt.fields.UserModel,
			}
			got, err := u.InsertUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("USERSRVC.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USERSRVC.InsertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSERSRVC_ReadUser(t *testing.T) {
	type fields struct {
		UserModel db.UserModel
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &USERSRVC{
				UserModel: tt.fields.UserModel,
			}
			got, err := u.ReadUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("USERSRVC.ReadUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USERSRVC.ReadUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSERSRVC_ReadAllUsers(t *testing.T) {
	type fields struct {
		UserModel db.UserModel
	}
	tests := []struct {
		name    string
		fields  fields
		want    []db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &USERSRVC{
				UserModel: tt.fields.UserModel,
			}
			got, err := u.ReadAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("USERSRVC.ReadAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USERSRVC.ReadAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSERSRVC_UpdateUser(t *testing.T) {
	type fields struct {
		UserModel db.UserModel
	}
	type args struct {
		id   string
		user db.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &USERSRVC{
				UserModel: tt.fields.UserModel,
			}
			got, err := u.UpdateUser(tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("USERSRVC.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USERSRVC.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSERSRVC_DeleteUser(t *testing.T) {
	type fields struct {
		UserModel db.UserModel
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
			u := &USERSRVC{
				UserModel: tt.fields.UserModel,
			}
			if err := u.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("USERSRVC.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
