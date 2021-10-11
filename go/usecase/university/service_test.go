package university

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
	"reflect"
	"testing"
)

func TestService_GetTinyUniversities(t *testing.T) {
	type fields struct {
		userRepository       user.UserReaderMock
		universityRepository UniversityRepositoryMock
	}
	tests := []struct {
		name      string
		fields    fields
		want      []entity.TinyUniversity
		totalWant *int
		wantErr   bool
	}{
		{
			name: "ok",
			want: []entity.TinyUniversity{
				{
					ID:             "615c09f7309d7ded48c7a049",
					Name:           "Fake name",
					ProfilePicture: "Fake profile picture",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			got, totalGot, err := s.GetTinyUniversities()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetTinyUniversities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.totalWant = totalGot
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetTinyUniversities() got = %v, want %v", got, tt.want)
			}
			if totalGot != tt.totalWant {
				t.Errorf("Service.GetTinyUniversities() totalGot = %v, want %v", totalGot, tt.totalWant)
			}
		})
	}
}

func TestService_GetUniversityByID(t *testing.T) {
	type fields struct {
		userRepository       user.UserReaderMock
		universityRepository UniversityRepositoryMock
	}
	type args struct {
		userID       string
		universityID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.University
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userID:       "615c09f7309d7ded48c7a053",
				universityID: "615c09f7309d7ded48c7a049",
			},
			want: &entity.University{
				ID:             "615c09f7309d7ded48c7a049",
				Name:           "Fake name",
				ProfilePicture: "Fake profile picture",
				Classrooms: []entity.Classroom{
					{
						ID:   "615c09f7309d7def48c7a010",
						Name: "Fake name",
					},
				},
			},
		},
		{
			name: "insufficient permissions",
			args: args{
				userID:       "615c09f7309d7ded48c7a053",
				universityID: "615c09f7309d7ded48c7a050",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			got, err := s.GetUniversityByID(tt.args.userID, tt.args.universityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUniversityByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUniversityByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	type args struct {
		userRepository       user.UserReaderMock
		universityRepository UniversityRepositoryMock
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "ok",
			args: args{
				userRepository:       user.UserReaderMock{},
				universityRepository: UniversityRepositoryMock{},
			},
			want: &Service{
				userRepository:       user.UserReaderMock{},
				universityRepository: UniversityRepositoryMock{},
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
