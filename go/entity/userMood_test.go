package entity

import (
	"testing"
	"time"
)

func TestUserMoodPayload_ValidateMood(t *testing.T) {
	type fields struct {
		ID           string
		UserID       string
		Mood         float64
		CreationDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ID:           "61639d228e3ca087c7a5fdad",
				UserID:       "615c09f7309d7ded48c7a053",
				Mood:         5,
				CreationDate: time.Now(),
			},
		},
		{
			name: "invalid mood",
			fields: fields{
				ID:           "61639d228e3ca087c7a5fdad",
				UserID:       "615c09f7309d7ded48c7a053",
				Mood:         11,
				CreationDate: time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ump := &UserMoodPayload{
				ID:           tt.fields.ID,
				UserID:       tt.fields.UserID,
				Mood:         tt.fields.Mood,
				CreationDate: tt.fields.CreationDate,
			}
			if err := ump.ValidateMood(); (err != nil) != tt.wantErr {
				t.Errorf("UserMoodPayload.ValidateMood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
