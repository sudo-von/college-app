package entity

import (
	"testing"
	"time"
)

func TestContactPayload_ValidateNumber(t *testing.T) {
	type fields struct {
		ID            string
		UserID        string
		ContactName   string
		ContactNumber string
		Message       string
		CreationDate  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ContactNumber: "528661234567",
			},
		},
		{
			name: "invalid contact number",
			fields: fields{
				ContactNumber: "8661234567",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ContactPayload{
				ID:            tt.fields.ID,
				UserID:        tt.fields.UserID,
				ContactName:   tt.fields.ContactName,
				ContactNumber: tt.fields.ContactNumber,
				Message:       tt.fields.Message,
				CreationDate:  tt.fields.CreationDate,
			}
			if err := cp.ValidateNumber(); (err != nil) != tt.wantErr {
				t.Errorf("ContactPayload.ValidateNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateContactPayload_ValidateNumber(t *testing.T) {
	type fields struct {
		ContactName   string
		ContactNumber string
		Message       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ContactNumber: "528661234567",
			},
		},
		{
			name: "invalid contact number",
			fields: fields{
				ContactNumber: "8661234567",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucp := &UpdateContactPayload{
				ContactName:   tt.fields.ContactName,
				ContactNumber: tt.fields.ContactNumber,
				Message:       tt.fields.Message,
			}
			if err := ucp.ValidateNumber(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateContactPayload.ValidateNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
