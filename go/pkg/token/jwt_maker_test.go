package token

import (
	"freelancer/college-app/go/entity"
	"reflect"
	"testing"
	"time"
)

func TestNewService(t *testing.T) {
	type args struct {
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				secretKey: "67eec82ed526e288e022ee743e254defeef869ff1af05b39943f0c46f9ed83e8",
			},
			want: &Service{
				secretKey: "67eec82ed526e288e022ee743e254defeef869ff1af05b39943f0c46f9ed83e8",
			},
		},
		{
			name: "invalid secret key",
			args: args{
				secretKey: "67eec82",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateToken(t *testing.T) {
	type fields struct {
		secretKey string
	}
	type args struct {
		user     *entity.User
		duration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				secretKey: "67eec82ed526e288e022ee743e254defeef869ff1af05b39943f0c46f9ed83e8",
			},
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				secretKey: tt.fields.secretKey,
			}
			got, err := s.CreateToken(tt.args.user, tt.args.duration)
			tt.want = got
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
