package token

import (
	"freelancer/college-app/go/entity"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewPayload(t *testing.T) {
	type args struct {
		user     *entity.User
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *Payload
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				user: &entity.User{
					ID:   "615c09f7309d7ded48c7a053",
					Name: "Fake user name",
					University: entity.University{
						ID:             "615c09f7309d7ded48c7a049",
						Name:           "Fake university name",
						ProfilePicture: "Fake university profile picture",
					},
				},
				duration: 15,
			},
			want: &Payload{
				UserID:                   "615c09f7309d7ded48c7a053",
				UserName:                 "Fake user name",
				UniversityID:             "615c09f7309d7ded48c7a049",
				UniversityName:           "Fake university name",
				UniversityProfilePicture: "Fake university profile picture",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPayload(tt.args.user, tt.args.duration)
			tt.want.ID = got.ID
			tt.want.IssuedAt = got.IssuedAt
			tt.want.ExpiredAt = got.ExpiredAt
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayload_Valid(t *testing.T) {
	type fields struct {
		ID                       uuid.UUID
		UserID                   string
		UserName                 string
		UniversityID             string
		UniversityName           string
		UniversityProfilePicture string
		IssuedAt                 time.Time
		ExpiredAt                time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ID:                       uuid.New(),
				UserID:                   "615c09f7309d7ded48c7a053",
				UserName:                 "Fake user name",
				UniversityID:             "615c09f7309d7ded48c7a049",
				UniversityName:           "Fake university name",
				UniversityProfilePicture: "Fake university profile picture",
				IssuedAt:                 time.Now(),
				ExpiredAt:                time.Now().Add(time.Minute * 15),
			},
		},
		{
			name: "expired token",
			fields: fields{
				ID:                       uuid.New(),
				UserID:                   "615c09f7309d7ded48c7a053",
				UserName:                 "Fake user name",
				UniversityID:             "615c09f7309d7ded48c7a049",
				UniversityName:           "Fake university name",
				UniversityProfilePicture: "Fake university profile picture",
				IssuedAt:                 time.Now(),
				ExpiredAt:                time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := &Payload{
				ID:                       tt.fields.ID,
				UserID:                   tt.fields.UserID,
				UserName:                 tt.fields.UserName,
				UniversityID:             tt.fields.UniversityID,
				UniversityName:           tt.fields.UniversityName,
				UniversityProfilePicture: tt.fields.UniversityProfilePicture,
				IssuedAt:                 tt.fields.IssuedAt,
				ExpiredAt:                tt.fields.ExpiredAt,
			}
			if err := payload.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Payload.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
