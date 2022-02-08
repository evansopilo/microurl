package service

import (
	"reflect"
	"testing"

	"github.com/evansopilo/Micro-URL/pkg/db"
)

func TestURLSRVC_InsertURL(t *testing.T) {
	type fields struct {
		URLModel *db.URLModel
	}
	type args struct {
		url db.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLSRVC{
				URLModel: tt.fields.URLModel,
			}
			got, err := u.InsertURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLSRVC.InsertURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLSRVC.InsertURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLSRVC_ReadURL(t *testing.T) {
	type fields struct {
		URLModel *db.URLModel
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLSRVC{
				URLModel: tt.fields.URLModel,
			}
			got, err := u.ReadURL(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLSRVC.ReadURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLSRVC.ReadURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLSRVC_ReadAllURLs(t *testing.T) {
	type fields struct {
		URLModel *db.URLModel
	}
	tests := []struct {
		name    string
		fields  fields
		want    []db.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLSRVC{
				URLModel: tt.fields.URLModel,
			}
			got, err := u.ReadAllURLs()
			if (err != nil) != tt.wantErr {
				t.Errorf("URLSRVC.ReadAllURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLSRVC.ReadAllURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLSRVC_UpdateURL(t *testing.T) {
	type fields struct {
		URLModel *db.URLModel
	}
	type args struct {
		id  string
		url db.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLSRVC{
				URLModel: tt.fields.URLModel,
			}
			got, err := u.UpdateURL(tt.args.id, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLSRVC.UpdateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLSRVC.UpdateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLSRVC_DeleteURL(t *testing.T) {
	type fields struct {
		URLModel *db.URLModel
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
			u := &URLSRVC{
				URLModel: tt.fields.URLModel,
			}
			if err := u.DeleteURL(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("URLSRVC.DeleteURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
