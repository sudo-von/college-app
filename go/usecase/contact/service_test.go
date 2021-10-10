package contact

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
	"reflect"
	"testing"
	"time"
)

func TestService_GetContactByUserID(t *testing.T) {
	type fields struct {
		userRepository    user.UserReaderMock
		contactRepository ContactRepositoryMock
	}
	type args struct {
		userID          string
		requestedUserID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Contact
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
			},
			want: &entity.Contact{
				ID:            "615c09f7309d7ded48c7a0c7",
				UserID:        "615c09f7309d7ded48c7a053",
				ContactName:   "Fake contact name",
				ContactNumber: "Fake contact number",
				Message:       "Fake message",
				CreationDate:  time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "invalid permissions",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a054",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				userRepository:    tt.fields.userRepository,
				contactRepository: tt.fields.contactRepository,
			}
			got, err := s.GetContactByUserID(tt.args.userID, tt.args.requestedUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetContactByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetContactByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateContact(t *testing.T) {
	type fields struct {
		userRepository    user.UserReaderMock
		contactRepository ContactRepositoryMock
	}
	type args struct {
		userID          string
		requestedUserID string
		newContact      entity.ContactPayload
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
				userID:          "615c09f7309d7ded48c7a054",
				requestedUserID: "615c09f7309d7ded48c7a054",
				newContact: entity.ContactPayload{
					UserID:        "615c09f7309d7ded48c7a054",
					ContactName:   "Fake contact name",
					ContactNumber: "528661234567",
					Message:       "Fake message",
					CreationDate:  time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name: "user not found",
			args: args{
				userID:          "615c09f7309d7ded48c7a056",
				requestedUserID: "615c09f7309d7ded48c7a056",
			},
			wantErr: true,
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
			name: "invalid contact number",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newContact: entity.ContactPayload{
					UserID:        "615c09f7309d7ded48c7a053",
					ContactName:   "Fake contact name",
					ContactNumber: "abcdefghjklm",
					Message:       "Fake message",
					CreationDate:  time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
			wantErr: true,
		},
		{
			name: "contact already registered",
			args: args{
				userID:          "615c09f7309d7ded48c7a053",
				requestedUserID: "615c09f7309d7ded48c7a053",
				newContact: entity.ContactPayload{
					UserID:        "615c09f7309d7ded48c7a053",
					ContactName:   "Fake contact name",
					ContactNumber: "528661234567",
					Message:       "Fake message",
					CreationDate:  time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				userRepository:    tt.fields.userRepository,
				contactRepository: tt.fields.contactRepository,
			}
			if err := s.CreateContact(tt.args.userID, tt.args.requestedUserID, tt.args.newContact); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UpdateContactByID(t *testing.T) {
	type fields struct {
		userRepository    user.UserReaderMock
		contactRepository ContactRepositoryMock
	}
	type args struct {
		userID     string
		contactID  string
		newContact entity.UpdateContactPayload
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
				userID:    "615c09f7309d7ded48c7a053",
				contactID: "615c09f7309d7ded48c7a0c7",
				newContact: entity.UpdateContactPayload{
					ContactName:   "Fake contact name",
					ContactNumber: "528661234567",
					Message:       "Fake message",
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
		{
			name: "contact not found",
			args: args{
				userID:    "615c09f7309d7ded48c7a054",
				contactID: "615c09f7309d7ded48c7a0c9",
			},
			wantErr: true,
		},
		{
			name: "insufficient permissions",
			args: args{
				userID:    "615c09f7309d7ded48c7a054",
				contactID: "615c09f7309d7ded48c7a0c7",
			},
			wantErr: true,
		},
		{
			name: "invalid contact number",
			args: args{
				userID:    "615c09f7309d7ded48c7a053",
				contactID: "615c09f7309d7ded48c7a0c7",
				newContact: entity.UpdateContactPayload{
					ContactName:   "Fake contact name",
					ContactNumber: "abcdefghjklm",
					Message:       "Fake message",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				userRepository:    tt.fields.userRepository,
				contactRepository: tt.fields.contactRepository,
			}
			if err := s.UpdateContactByID(tt.args.userID, tt.args.contactID, tt.args.newContact); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateContactByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
