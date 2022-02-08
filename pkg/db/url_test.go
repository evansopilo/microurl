package db

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestURLModel_Insert(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	type args struct {
		url URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLModel{
				DB: tt.fields.DB,
			}
			got, err := u.Insert(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLModel.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLModel.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLModel_ReadByID(t *testing.T) {
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
		want    URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLModel{
				DB: tt.fields.DB,
			}
			got, err := u.ReadByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLModel.ReadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLModel.ReadByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLModel_ReadAll(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	tests := []struct {
		name    string
		fields  fields
		want    []URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLModel{
				DB: tt.fields.DB,
			}
			got, err := u.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("URLModel.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLModel.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLModel_Update(t *testing.T) {
	type fields struct {
		DB *mongo.Collection
	}
	type args struct {
		id  string
		url URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLModel{
				DB: tt.fields.DB,
			}
			got, err := u.Update(tt.args.id, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLModel.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLModel.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLModel_Delete(t *testing.T) {
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
			u := &URLModel{
				DB: tt.fields.DB,
			}
			if err := u.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("URLModel.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
