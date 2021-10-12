package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserMoodController_Show(t *testing.T) {
	type fields struct {
		UserMoodService UserMoodServiceMock
		AuthService     func(http.Handler) http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		userID string
	}{
		{
			name:   "ok",
			userID: "615c09f7309d7ded48c7a053",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserMoodController{
				UserMoodService: tt.fields.UserMoodService,
				AuthService:     tt.fields.AuthService,
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/users/{id}", nil)
			ctx := context.WithValue(r.Context(), ContextKeyUserID, tt.userID)
			r = r.WithContext(ctx)
			c.Show(w, r)
		})
	}
}
