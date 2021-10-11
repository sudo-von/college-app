package entity

import (
	"testing"
	"time"
)

func TestUser_ValidateRequestedUser(t *testing.T) {
	type fields struct {
		ID                 string
		Name               string
		BirthDate          time.Time
		Email              string
		Password           string
		RegistrationNumber string
		Status             string
		Role               string
		University         University
		CreationDate       time.Time
	}
	type args struct {
		requestedUserID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ID: "615c09f7309d7ded48c7a053",
			},
			args: args{
				requestedUserID: "615c09f7309d7ded48c7a053",
			},
		},
		{
			name: "insufficient permissions",
			fields: fields{
				ID: "615c09f7309d7ded48c7a053",
			},
			args: args{
				requestedUserID: "615c09f7309d7ded48c7a054",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:                 tt.fields.ID,
				Name:               tt.fields.Name,
				BirthDate:          tt.fields.BirthDate,
				Email:              tt.fields.Email,
				Password:           tt.fields.Password,
				RegistrationNumber: tt.fields.RegistrationNumber,
				Status:             tt.fields.Status,
				Role:               tt.fields.Role,
				University:         tt.fields.University,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := u.ValidateRequestedUser(tt.args.requestedUserID); (err != nil) != tt.wantErr {
				t.Errorf("User.ValidateRequestedUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ValidateRequestedUniversity(t *testing.T) {
	type fields struct {
		ID                 string
		Name               string
		BirthDate          time.Time
		Email              string
		Password           string
		RegistrationNumber string
		Status             string
		Role               string
		University         University
		CreationDate       time.Time
	}
	type args struct {
		requestedUniversityID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				University: University{
					ID: "615c09f7309d7ded48c7a049",
				},
			},
			args: args{
				requestedUniversityID: "615c09f7309d7ded48c7a049",
			},
		},
		{
			name: "insufficient permissions",
			fields: fields{
				University: University{
					ID: "615c09f7309d7ded48c7a049",
				},
			},
			args: args{
				requestedUniversityID: "615c09f7309d7ded48c7a048",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:                 tt.fields.ID,
				Name:               tt.fields.Name,
				BirthDate:          tt.fields.BirthDate,
				Email:              tt.fields.Email,
				Password:           tt.fields.Password,
				RegistrationNumber: tt.fields.RegistrationNumber,
				Status:             tt.fields.Status,
				Role:               tt.fields.Role,
				University:         tt.fields.University,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := u.ValidateRequestedUniversity(tt.args.requestedUniversityID); (err != nil) != tt.wantErr {
				t.Errorf("User.ValidateRequestedUniversity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserPayload_ValidateRegistrationNumberDigits(t *testing.T) {
	type fields struct {
		ID                 string
		Name               string
		BirthDate          time.Time
		Email              string
		RegistrationNumber string
		Password           string
		UniversityID       string
		Status             string
		Role               string
		CreationDate       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				RegistrationNumber: "12345678",
			},
		},
		{
			name: "invalid registration number",
			fields: fields{
				RegistrationNumber: "abcdefgh",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &UserPayload{
				ID:                 tt.fields.ID,
				Name:               tt.fields.Name,
				BirthDate:          tt.fields.BirthDate,
				Email:              tt.fields.Email,
				RegistrationNumber: tt.fields.RegistrationNumber,
				Password:           tt.fields.Password,
				UniversityID:       tt.fields.UniversityID,
				Status:             tt.fields.Status,
				Role:               tt.fields.Role,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := up.ValidateRegistrationNumberDigits(); (err != nil) != tt.wantErr {
				t.Errorf("UserPayload.ValidateRegistrationNumberDigits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserPayload_ValidateRegistrationNumberLength(t *testing.T) {
	type fields struct {
		ID                 string
		Name               string
		BirthDate          time.Time
		Email              string
		RegistrationNumber string
		Password           string
		UniversityID       string
		Status             string
		Role               string
		CreationDate       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				RegistrationNumber: "12345678",
			},
		},
		{
			name: "invalid registration number length",
			fields: fields{
				RegistrationNumber: "123456789",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &UserPayload{
				ID:                 tt.fields.ID,
				Name:               tt.fields.Name,
				BirthDate:          tt.fields.BirthDate,
				Email:              tt.fields.Email,
				RegistrationNumber: tt.fields.RegistrationNumber,
				Password:           tt.fields.Password,
				UniversityID:       tt.fields.UniversityID,
				Status:             tt.fields.Status,
				Role:               tt.fields.Role,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := up.ValidateRegistrationNumberLength(); (err != nil) != tt.wantErr {
				t.Errorf("UserPayload.ValidateRegistrationNumberLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUserPayload_ValidateRegistrationNumberDigits(t *testing.T) {
	type fields struct {
		Name               string
		BirthDate          time.Time
		Email              string
		RegistrationNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				RegistrationNumber: "12345678",
			},
		},
		{
			name: "invalid registration number",
			fields: fields{
				RegistrationNumber: "abcdefgh",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &UpdateUserPayload{
				Name:               tt.fields.Name,
				BirthDate:          tt.fields.BirthDate,
				Email:              tt.fields.Email,
				RegistrationNumber: tt.fields.RegistrationNumber,
			}
			if err := up.ValidateRegistrationNumberDigits(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserPayload.ValidateRegistrationNumberDigits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUserPayload_ValidateRegistrationNumberLength(t *testing.T) {
	type fields struct {
		Name               string
		BirthDate          time.Time
		Email              string
		RegistrationNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				RegistrationNumber: "12345678",
			},
		},
		{
			name: "invalid registration number length",
			fields: fields{
				RegistrationNumber: "123456789",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &UpdateUserPayload{
				Name:               tt.fields.Name,
				BirthDate:          tt.fields.BirthDate,
				Email:              tt.fields.Email,
				RegistrationNumber: tt.fields.RegistrationNumber,
			}
			if err := up.ValidateRegistrationNumberLength(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserPayload.ValidateRegistrationNumberLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
