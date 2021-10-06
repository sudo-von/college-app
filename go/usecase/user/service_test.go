package user

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/university"
	"reflect"
	"testing"
	"time"
)

func TestService_GetTinyUserByID(t *testing.T) {
	type fields struct {
		userRepository       UserRepositoryMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID          string
		requestedUserID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.TinyUser
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
			},
			want: &entity.TinyUser{
				ID:                 "615c09f7309d7ded48c7a053",
				Name:               "Von",
				BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
				Email:              "sudo.von.contact@gmail.com",
				RegistrationNumber: "16190775",
			},
		},
		{
			name: "insufficient permissions",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a054",
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				userID:          "615c09f7309d7ded48c7a055",
				requestedUserID: "615c09f7309d7ded48c7a055",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			got, err := s.GetTinyUserByID(tt.args.userID, tt.args.requestedUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetTinyUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetTinyUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
