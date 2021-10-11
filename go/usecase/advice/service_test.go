package advice

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
	"reflect"
	"testing"
	"time"
)

func TestService_GetAdvices(t *testing.T) {
	type fields struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserReaderMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID        string
		adviceFilters entity.AdviceFilters
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []entity.Advice
		totalWant *int
		wantErr   bool
	}{
		{
			name: "ok",
			args: args{
				userID: "615c09f7309d7ded48c7a053",
				adviceFilters: entity.AdviceFilters{
					AdviceDate: &time.Time{},
				},
			},
			want: []entity.Advice{
				{
					ID: "615e7802e13d8485f4f5e35e",
					User: entity.TinyUser{
						ID: "615c09f7309d7ded48c7a053",
					},
					Classroom: entity.Classroom{
						ID:   "615c09f7309d7def48c7a010",
						Name: "Fake name",
					},
					UniversityID:       "615c09f7309d7ded48c7a049",
					Subject:            "Fake subject",
					AdviceDate:         time.Date(2021, 01, 02, 0, 0, 0, 0, time.Local),
					StudentsWillAttend: []string{"615c09f7309d7ded48c7a053"},
					CreationDate:       time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name: "user not found",
			args: args{
				userID:        "615c09f7309d7ded48c7a052",
				adviceFilters: entity.AdviceFilters{},
			},
			wantErr: true,
		},
		{
			name: "insufficient permissions",
			args: args{
				userID: "615c09f7309d7ded48c7a053",
				adviceFilters: entity.AdviceFilters{
					AdviceDate:     &time.Time{},
					UserWillAttend: "615c09f7309d7ded48c7a054",
				},
			},
			wantErr: true,
		},
		{
			name: "insufficient permissions",
			args: args{
				userID: "615c09f7309d7ded48c7a053",
				adviceFilters: entity.AdviceFilters{
					AdviceDate:    &time.Time{},
					UserWillTeach: "615c09f7309d7ded48c7a054",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				adviceRepository:     tt.fields.adviceRepository,
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			got, totalGot, err := s.GetAdvices(tt.args.userID, tt.args.adviceFilters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAdvices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.totalWant = totalGot
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAdvices() got = %v, want %v", got, tt.want)
			}
			if totalGot != tt.totalWant {
				t.Errorf("Service.GetAdvices() totalGot = %v, want %v", totalGot, tt.totalWant)
			}
		})
	}
}

func TestService_GetAdviceByID(t *testing.T) {
	type fields struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserReaderMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID   string
		adviceID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Advice
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35e",
			},
			want: &entity.Advice{
				ID: "615e7802e13d8485f4f5e35e",
				User: entity.TinyUser{
					ID: "615c09f7309d7ded48c7a053",
				},
				Classroom: entity.Classroom{
					ID:   "615c09f7309d7def48c7a010",
					Name: "Fake name",
				},
				UniversityID:       "615c09f7309d7ded48c7a049",
				Subject:            "Fake subject",
				AdviceDate:         time.Date(2021, 01, 02, 0, 0, 0, 0, time.Local),
				StudentsWillAttend: []string{"615c09f7309d7ded48c7a053"},
				CreationDate:       time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "user not found",
			args: args{
				userID:   "615c09f7309d7ded48c7a052",
				adviceID: "615e7802e13d8485f4f5e35e",
			},
			wantErr: true,
		},
		{
			name: "advice not found",
			args: args{
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35d",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				adviceRepository:     tt.fields.adviceRepository,
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			got, err := s.GetAdviceByID(tt.args.userID, tt.args.adviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAdviceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAdviceByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateAdvice(t *testing.T) {
	type fields struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserReaderMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		newAdvice entity.AdvicePayload
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
				newAdvice: entity.AdvicePayload{
					UserID:             "615c09f7309d7ded48c7a053",
					ClassroomID:        "615c09f7309d7def48c7a010",
					UniversityID:       "615c09f7309d7ded48c7a049",
					Subject:            "Fake subject",
					AdviceDate:         time.Now().In(time.Local).Add(time.Hour * time.Duration(1)),
					StudentsWillAttend: []string{"615c09f7309d7ded48c7a053"},
					Status:             entity.ActiveStatus,
					CreationDate:       time.Now().In(time.Local),
				},
			},
		},
		{
			name: "invalid advice date",
			args: args{
				newAdvice: entity.AdvicePayload{
					UserID:             "615c09f7309d7ded48c7a053",
					ClassroomID:        "615c09f7309d7def48c7a010",
					UniversityID:       "615c09f7309d7ded48c7a049",
					Subject:            "Fake subject",
					AdviceDate:         time.Now().In(time.Local),
					StudentsWillAttend: []string{"615c09f7309d7ded48c7a053"},
					Status:             entity.ActiveStatus,
					CreationDate:       time.Now().In(time.Local),
				},
			},
			wantErr: true,
		},
		{
			name: "user not found",
			args: args{
				newAdvice: entity.AdvicePayload{
					UserID:             "615c09f7309d7ded48c7a052",
					ClassroomID:        "615c09f7309d7def48c7a010",
					UniversityID:       "615c09f7309d7ded48c7a049",
					Subject:            "Fake subject",
					AdviceDate:         time.Now().In(time.Local).Add(time.Hour * time.Duration(1)),
					StudentsWillAttend: []string{"615c09f7309d7ded48c7a052"},
					Status:             entity.ActiveStatus,
					CreationDate:       time.Now().In(time.Local),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				adviceRepository:     tt.fields.adviceRepository,
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			if err := s.CreateAdvice(tt.args.newAdvice); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateAdvice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UpdateAdvice(t *testing.T) {
	type fields struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserReaderMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID    string
		adviceID  string
		newAdvice entity.UpdateAdvicePayload
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
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35e",
				newAdvice: entity.UpdateAdvicePayload{
					ClassroomID: "615c09f7309d7def48c7a010",
					Subject:     "Fake subject",
					AdviceDate:  time.Now().In(time.Local).Add(time.Hour * time.Duration(1)),
				},
			},
		},
		{
			name: "insufficient permissions",
			args: args{
				userID:   "615c09f7309d7ded48c7a054",
				adviceID: "615e7802e13d8485f4f5e35e",
				newAdvice: entity.UpdateAdvicePayload{
					ClassroomID: "615c09f7309d7def48c7a010",
					Subject:     "Fake subject",
					AdviceDate:  time.Now().In(time.Local).Add(time.Hour * time.Duration(1)),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid advice date",
			args: args{
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35e",
				newAdvice: entity.UpdateAdvicePayload{
					ClassroomID: "615c09f7309d7def48c7a010",
					Subject:     "Fake subject",
					AdviceDate:  time.Now().In(time.Local),
				},
			},
			wantErr: true,
		},
		{
			name: "user not found",
			args: args{
				userID:   "615c09f7309d7ded48c7a052",
				adviceID: "615e7802e13d8485f4f5e35e",
				newAdvice: entity.UpdateAdvicePayload{
					ClassroomID: "615c09f7309d7def48c7a010",
					Subject:     "Fake subject",
					AdviceDate:  time.Now().In(time.Local).Add(time.Hour * time.Duration(1)),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid classroom id",
			args: args{
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35e",
				newAdvice: entity.UpdateAdvicePayload{
					ClassroomID: "615c09f7309d7def48c7a009",
					Subject:     "Fake subject",
					AdviceDate:  time.Now().In(time.Local).Add(time.Hour * time.Duration(1)),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				adviceRepository:     tt.fields.adviceRepository,
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			if err := s.UpdateAdvice(tt.args.userID, tt.args.adviceID, tt.args.newAdvice); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateAdvice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteAdvice(t *testing.T) {
	type fields struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserReaderMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID   string
		adviceID string
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
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35e",
			},
		},
		{
			name: "insufficient permissions",
			args: args{
				userID:   "615c09f7309d7ded48c7a054",
				adviceID: "615e7802e13d8485f4f5e35e",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				adviceRepository:     tt.fields.adviceRepository,
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			if err := s.DeleteAdvice(tt.args.userID, tt.args.adviceID); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteAdvice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UpdateAdviceStudentsNumber(t *testing.T) {
	type fields struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserReaderMock
		universityRepository university.UniversityRepositoryMock
	}
	type args struct {
		userID   string
		adviceID string
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
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35e",
			},
		},
		{
			name: "advice not found",
			args: args{
				userID:   "615c09f7309d7ded48c7a053",
				adviceID: "615e7802e13d8485f4f5e35d",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				adviceRepository:     tt.fields.adviceRepository,
				userRepository:       tt.fields.userRepository,
				universityRepository: tt.fields.universityRepository,
			}
			if err := s.UpdateAdviceStudentsNumber(tt.args.userID, tt.args.adviceID); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateAdviceStudentsNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	type args struct {
		adviceRepository     AdviceRepositoryMock
		userRepository       user.UserRepositoryMock
		universityRepository university.UniversityRepositoryMock
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "ok",
			args: args{
				adviceRepository:     AdviceRepositoryMock{},
				userRepository:       user.UserRepositoryMock{},
				universityRepository: university.UniversityRepositoryMock{},
			},
			want: &Service{
				adviceRepository:     AdviceRepositoryMock{},
				userRepository:       user.UserRepositoryMock{},
				universityRepository: university.UniversityRepositoryMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.adviceRepository, tt.args.userRepository, tt.args.universityRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
