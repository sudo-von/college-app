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

func TestService_CreateUser(t *testing.T) {
	type fields struct {
		userRepository       UserRepositoryMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		newUser entity.UserPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				newUser: entity.UserPayload{
					Name:               "Von",
					BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "16190770",
					Password:           "123456",
					UniversityID:       "615c09f7309d7ded48c7a049",
					Status:             entity.ActiveStatus,
					Role:               entity.StudentRole,
					CreationDate:       time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name: "invalid email",
			args: args{
				newUser: entity.UserPayload{
					Email: "sudo.von.contact",
				},
			},
			wantErr: true,
		},
		{
			name: "registration number accepts only numbers",
			args: args{
				newUser: entity.UserPayload{
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "abcdefgh",
				},
			},
			wantErr: true,
		},
		{
			name: "exceeds the length of the registration number ",
			args: args{
				newUser: entity.UserPayload{
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "161907707",
				},
			},
			wantErr: true,
		},
		{
			name: "email already in use",
			args: args{
				newUser: entity.UserPayload{
					Email:              "sudo.von.contact@gmail.com",
					RegistrationNumber: "16190770",
				},
			},
			wantErr: true,
		},
		{
			name: "registration number already in use",
			args: args{
				newUser: entity.UserPayload{
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "16190775",
				},
			},
			wantErr: true,
		},
		{
			name: "university id not found",
			args: args{
				newUser: entity.UserPayload{
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "16190770",
					UniversityID:       "615c09f7309d7ded48c7a048",
				},
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
			if err := s.CreateUser(tt.args.newUser); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UpdateTinyUser(t *testing.T) {
	type fields struct {
		userRepository       UserRepositoryMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID          string
		requestedUserID string
		newUser         entity.UpdateUserPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newUser: entity.UpdateUserPayload{
					Name:               "Von",
					BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
					Email:              "sudo.von.contact@gmail.com",
					RegistrationNumber: "16190770",
				},
			},
		},
		{
			name: "invalid email",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newUser: entity.UpdateUserPayload{
					Email: "sudo.von.contact",
				},
			},
			wantErr: true,
		},
		{
			name: "registration number accepts only numbers",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newUser: entity.UpdateUserPayload{
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "abcdefgh",
				},
			},
			wantErr: true,
		},
		{
			name: "exceeds the length of the registration number ",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newUser: entity.UpdateUserPayload{
					Email:              "sudovoncontact@gmail.com",
					RegistrationNumber: "161907707",
				},
			},
			wantErr: true,
		},
		{
			name: "email already in use",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newUser: entity.UpdateUserPayload{
					Email:              "sudovonctf@gmail.com",
					RegistrationNumber: "16190770",
				},
			},
			wantErr: true,
		},
		{
			name: "registration number already in use",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newUser: entity.UpdateUserPayload{
					Email:              "sudo.von.contact@gmail.com",
					RegistrationNumber: "16190776",
				},
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
			if err := s.UpdateTinyUser(tt.args.userID, tt.args.requestedUserID, tt.args.newUser); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateTinyUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_AuthenticateUser(t *testing.T) {
	type fields struct {
		userRepository       UserRepositoryMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "invalid credentials",
			args: args{
				email:    "sudo.von.contact@gmail.com",
				password: "123abc",
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
			got, err := s.AuthenticateUser(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AuthenticateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	type args struct {
		userRepository       UserRepository
		universityRepository UniversityRepository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "ok",
			args: args{
				userRepository:       UserRepositoryMock{},
				universityRepository: university.UniversityReaderMock{},
			},
			want: &Service{
				userRepository:       UserRepositoryMock{},
				universityRepository: university.UniversityReaderMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.userRepository, tt.args.universityRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
