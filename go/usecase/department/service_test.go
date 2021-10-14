package department

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
	"reflect"
	"testing"
	"time"
)

func TestService_GetDepartments(t *testing.T) {
	type fields struct {
		departmentRepository DepartmentRepositoryMock
		userRepository       user.UserReaderMock
	}
	type args struct {
		userID            string
		departmentFilters entity.DepartmentFilters
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []entity.Department
		totalWant *int
		wantErr   bool
	}{
		{
			name: "ok",
			args: args{
				userID: "615c09f7309d7ded48c7a053",
			},
			want: []entity.Department{
				{
					ID: "6167698bd0187e53d4789212",
					User: entity.BasicUser{
						ID: "615c09f7309d7ded48c7a053",
					},
					UniversityID: "615c09f7309d7ded48c7a049",
					Description:  "fake description",
					Street:       "fake street",
					Neighborhood: "fake neighborhood",
					Cost:         1000,
					Available:    true,
					CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				departmentRepository: tt.fields.departmentRepository,
				userRepository:       tt.fields.userRepository,
			}
			got, got1, err := s.GetDepartments(tt.args.userID, tt.args.departmentFilters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetDepartments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			total := len(tt.want)
			tt.totalWant = &total
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetDepartments() got = %v, want %v", got, tt.want)
			}
			if *got1 != *tt.totalWant {
				t.Errorf("Service.GetDepartments() total got = %v, want %v", got1, tt.totalWant)
			}
		})
	}
}

func TestService_CreateDepartment(t *testing.T) {
	type fields struct {
		departmentRepository DepartmentRepositoryMock
		userRepository       user.UserReaderMock
	}
	type args struct {
		userID        string
		newDepartment entity.DepartmentPayload
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
				userID:        "615c09f7309d7ded48c7a053",
				newDepartment: entity.DepartmentPayload{},
			},
		},
		{
			name: "user not found",
			args: args{
				userID:        "615c09f7309d7ded48c7a052",
				newDepartment: entity.DepartmentPayload{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				departmentRepository: tt.fields.departmentRepository,
				userRepository:       tt.fields.userRepository,
			}
			if err := s.CreateDepartment(tt.args.userID, tt.args.newDepartment); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateDepartment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
