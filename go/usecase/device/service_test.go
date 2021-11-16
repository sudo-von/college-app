package device

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
	"testing"
)

func TestService_CreateDevice(t *testing.T) {
	type fields struct {
		deviceRepository DeviceRepositoryMock
		userRepository   user.UserRepositoryMock
	}
	type args struct {
		device entity.DevicePayload
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
				device: entity.DevicePayload{
					UserID: "615c09f7309d7ded48c7a053",
				},
			},
		},
		{
			name: "user not found",
			args: args{
				device: entity.DevicePayload{
					UserID: "615c09f7309d7ded48c7a051",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				deviceRepository: tt.fields.deviceRepository,
				userRepository:   tt.fields.userRepository,
			}
			if err := s.CreateDevice(tt.args.device); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
