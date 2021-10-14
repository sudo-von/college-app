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
		{
			name: "user not found",
			args: args{
				userID: "615c09f7309d7ded48c7a052",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				departmentRepository: tt.fields.departmentRepository,
				userRepository:       tt.fields.userRepository,
			}
			got, total, err := s.GetDepartments(tt.args.userID, tt.args.departmentFilters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetDepartments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.totalWant = total
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetDepartments() got = %v, want %v", got, tt.want)
			}
			if total != tt.totalWant {
				t.Errorf("Service.GetDepartments() total got = %v, want %v", total, tt.totalWant)
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
				userID: "615c09f7309d7ded48c7a053",
				newDepartment: entity.DepartmentPayload{
					Cost: 1,
				},
			},
		},
		{
			name: "user not found",
			args: args{
				userID: "615c09f7309d7ded48c7a052",
				newDepartment: entity.DepartmentPayload{
					Cost: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid cost",
			args: args{
				userID: "615c09f7309d7ded48c7a052",
				newDepartment: entity.DepartmentPayload{
					Cost: 0,
				},
			},
			wantErr: true,
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

func TestService_GetDepartmentByID(t *testing.T) {
	type fields struct {
		departmentRepository DepartmentRepositoryMock
		userRepository       user.UserReaderMock
	}
	type args struct {
		userID       string
		departmentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Department
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userID:       "615c09f7309d7ded48c7a053",
				departmentID: "6167698bd0187e53d4789212",
			},
			want: &entity.Department{
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
		{
			name: "user not found",
			args: args{
				userID:       "615c09f7309d7ded48c7a052",
				departmentID: "6167698bd0187e53d4789212",
			},
			wantErr: true,
		},
		{
			name: "department not found",
			args: args{
				userID:       "615c09f7309d7ded48c7a053",
				departmentID: "6167698bd0187e53d4789211",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				departmentRepository: tt.fields.departmentRepository,
				userRepository:       tt.fields.userRepository,
			}
			got, err := s.GetDepartmentByID(tt.args.userID, tt.args.departmentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetDepartmentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetDepartmentByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	type args struct {
		departmentRepository DepartmentRepositoryMock
		userRepository       user.UserReaderMock
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "ok",
			want: &Service{
				departmentRepository: DepartmentRepositoryMock{},
				userRepository:       user.UserReaderMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.departmentRepository, tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateDepartment(t *testing.T) {
	type fields struct {
		departmentRepository DepartmentRepositoryMock
		userRepository       user.UserReaderMock
	}
	type args struct {
		userID            string
		departmentID      string
		departmentPayload entity.UpdateDepartmentPayload
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
				userID:       "615c09f7309d7ded48c7a053",
				departmentID: "6167698bd0187e53d4789212",
				departmentPayload: entity.UpdateDepartmentPayload{
					Description:  "updated description",
					Street:       "updated stret",
					Neighborhood: "updated neighborhood",
					Cost:         1,
					Available:    true,
				},
			},
		},
		{
			name: "user not found",
			args: args{
				userID:       "615c09f7309d7ded48c7a052",
				departmentID: "6167698bd0187e53d4789212",
				departmentPayload: entity.UpdateDepartmentPayload{
					Description:  "updated description",
					Street:       "updated stret",
					Neighborhood: "updated neighborhood",
					Cost:         1,
					Available:    true,
				},
			},
			wantErr: true,
		},
		{
			name: "department not found",
			args: args{
				userID:       "615c09f7309d7ded48c7a052",
				departmentID: "6167698bd0187e53d4789211",
				departmentPayload: entity.UpdateDepartmentPayload{
					Description:  "updated description",
					Street:       "updated stret",
					Neighborhood: "updated neighborhood",
					Cost:         1,
					Available:    true,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid cost",
			args: args{
				userID:       "615c09f7309d7ded48c7a052",
				departmentID: "6167698bd0187e53d4789212",
				departmentPayload: entity.UpdateDepartmentPayload{
					Description:  "updated description",
					Street:       "updated stret",
					Neighborhood: "updated neighborhood",
					Cost:         0,
					Available:    true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				departmentRepository: tt.fields.departmentRepository,
				userRepository:       tt.fields.userRepository,
			}
			if err := s.UpdateDepartment(tt.args.userID, tt.args.departmentID, tt.args.departmentPayload); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateDepartment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
