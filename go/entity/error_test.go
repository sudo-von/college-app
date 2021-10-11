package entity

import (
	"errors"
	"testing"
)

func TestErrorBadRequest_Error(t *testing.T) {
	type fields struct {
		Message error
		Code    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Message: errors.New("fake error"),
				Code:    "fake code",
			},
			want: "fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorBadRequest{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorBadRequest.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewErrorBadRequest(t *testing.T) {
	type args struct {
		message error
		code    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				message: errors.New("fake error"),
				code:    "fake code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewErrorBadRequest(tt.args.message, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("NewErrorBadRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrorUnauthorized_Error(t *testing.T) {
	type fields struct {
		Message error
		Code    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Message: errors.New("fake error"),
				Code:    "fake code",
			},
			want: "fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorUnauthorized{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorUnauthorized.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorForbidden_Error(t *testing.T) {
	type fields struct {
		Message error
		Code    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Message: errors.New("fake error"),
				Code:    "fake code",
			},
			want: "fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorForbidden{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorForbidden.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorNotFound_Error(t *testing.T) {
	type fields struct {
		Message error
		Code    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Message: errors.New("fake error"),
				Code:    "fake code",
			},
			want: "fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorNotFound{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorNotFound.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorConflict_Error(t *testing.T) {
	type fields struct {
		Message error
		Code    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Message: errors.New("fake error"),
				Code:    "fake code",
			},
			want: "fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorConflict{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorConflict.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorInternalServer_Error(t *testing.T) {
	type fields struct {
		Message error
		Code    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Message: errors.New("fake error"),
				Code:    "fake code",
			},
			want: "fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorInternalServer{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorInternalServer.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewErrorUnauthorized(t *testing.T) {
	type args struct {
		message error
		code    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				message: errors.New("fake error"),
				code:    "fake code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewErrorUnauthorized(tt.args.message, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("NewErrorUnauthorized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewErrorForbidden(t *testing.T) {
	type args struct {
		message error
		code    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				message: errors.New("fake error"),
				code:    "fake code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewErrorForbidden(tt.args.message, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("NewErrorForbidden() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewErrorNotFound(t *testing.T) {
	type args struct {
		message error
		code    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				message: errors.New("fake error"),
				code:    "fake code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewErrorNotFound(tt.args.message, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("NewErrorNotFound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewErrorConflict(t *testing.T) {
	type args struct {
		message error
		code    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				message: errors.New("fake error"),
				code:    "fake code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewErrorConflict(tt.args.message, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("NewErrorConflict() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewErrorInternalServer(t *testing.T) {
	type args struct {
		message error
		code    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				message: errors.New("fake error"),
				code:    "fake code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewErrorInternalServer(tt.args.message, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("NewErrorInternalServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
