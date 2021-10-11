package entity

import (
	"testing"
	"time"
)

func TestAdvice_ValidateRequestedAdvice(t *testing.T) {
	type fields struct {
		ID                 string
		User               TinyUser
		Classroom          Classroom
		UniversityID       string
		Subject            string
		AdviceDate         time.Time
		StudentsWillAttend []string
		Status             string
		CreationDate       time.Time
	}
	type args struct {
		userID string
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
				User: TinyUser{
					ID: "615c09f7309d7ded48c7a053",
				},
			},
			args: args{
				userID: "615c09f7309d7ded48c7a053",
			},
		},
		{
			name: "insufficient permissions",
			fields: fields{
				User: TinyUser{
					ID: "615c09f7309d7ded48c7a053",
				},
			},
			args: args{
				userID: "615c09f7309d7ded48c7a054",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Advice{
				ID:                 tt.fields.ID,
				User:               tt.fields.User,
				Classroom:          tt.fields.Classroom,
				UniversityID:       tt.fields.UniversityID,
				Subject:            tt.fields.Subject,
				AdviceDate:         tt.fields.AdviceDate,
				StudentsWillAttend: tt.fields.StudentsWillAttend,
				Status:             tt.fields.Status,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := a.ValidateRequestedAdvice(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Advice.ValidateRequestedAdvice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdvice_ValidateUniversity(t *testing.T) {
	type fields struct {
		ID                 string
		User               TinyUser
		Classroom          Classroom
		UniversityID       string
		Subject            string
		AdviceDate         time.Time
		StudentsWillAttend []string
		Status             string
		CreationDate       time.Time
	}
	type args struct {
		universityID string
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
				UniversityID: "615c09f7309d7ded48c7a049",
			},
			args: args{
				universityID: "615c09f7309d7ded48c7a049",
			},
		},
		{
			name: "insufficient permissions",
			fields: fields{
				UniversityID: "615c09f7309d7ded48c7a049",
			},
			args: args{
				universityID: "615c09f7309d7ded48c7a050",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Advice{
				ID:                 tt.fields.ID,
				User:               tt.fields.User,
				Classroom:          tt.fields.Classroom,
				UniversityID:       tt.fields.UniversityID,
				Subject:            tt.fields.Subject,
				AdviceDate:         tt.fields.AdviceDate,
				StudentsWillAttend: tt.fields.StudentsWillAttend,
				Status:             tt.fields.Status,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := a.ValidateUniversity(tt.args.universityID); (err != nil) != tt.wantErr {
				t.Errorf("Advice.ValidateUniversity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdvicePayload_ValidateDate(t *testing.T) {
	type fields struct {
		ID                 string
		UserID             string
		ClassroomID        string
		UniversityID       string
		Subject            string
		AdviceDate         time.Time
		StudentsWillAttend []string
		Status             string
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
				AdviceDate: time.Now().Add(time.Minute * 15),
			},
		},
		{
			name: "invalid date",
			fields: fields{
				AdviceDate: time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &AdvicePayload{
				ID:                 tt.fields.ID,
				UserID:             tt.fields.UserID,
				ClassroomID:        tt.fields.ClassroomID,
				UniversityID:       tt.fields.UniversityID,
				Subject:            tt.fields.Subject,
				AdviceDate:         tt.fields.AdviceDate,
				StudentsWillAttend: tt.fields.StudentsWillAttend,
				Status:             tt.fields.Status,
				CreationDate:       tt.fields.CreationDate,
			}
			if err := ap.ValidateDate(); (err != nil) != tt.wantErr {
				t.Errorf("AdvicePayload.ValidateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateAdvicePayload_ValidateDate(t *testing.T) {
	type fields struct {
		Subject     string
		ClassroomID string
		AdviceDate  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				AdviceDate: time.Now().Add(time.Minute * 15),
			},
		},
		{
			name: "invalid date",
			fields: fields{
				AdviceDate: time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uap := &UpdateAdvicePayload{
				Subject:     tt.fields.Subject,
				ClassroomID: tt.fields.ClassroomID,
				AdviceDate:  tt.fields.AdviceDate,
			}
			if err := uap.ValidateDate(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdvicePayload.ValidateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
