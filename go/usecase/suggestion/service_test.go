package suggestion

import (
	"freelancer/college-app/go/entity"
	"reflect"
	"testing"
	"time"
)

func TestService_CreateSuggestion(t *testing.T) {
	type fields struct {
		suggestionRepository SuggestionRepositoryMock
	}
	type args struct {
		newSuggestion entity.SuggestionPayload
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
				newSuggestion: entity.SuggestionPayload{
					UserID:       "615c09f7309d7ded48c7a053",
					Suggestion:   "Fake suggestion",
					CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				suggestionRepository: tt.fields.suggestionRepository,
			}
			if err := s.CreateSuggestion(tt.args.newSuggestion); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateSuggestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	type args struct {
		suggestionRepository SuggestionRepositoryMock
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "ok",
			args: args{
				suggestionRepository: SuggestionRepositoryMock{},
			},
			want: &Service{
				suggestionRepository: SuggestionRepositoryMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.suggestionRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
