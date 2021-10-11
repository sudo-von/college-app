package search

import "testing"

func TestFindInSlice(t *testing.T) {
	type args struct {
		id   string
		list []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				id:   "615c09f7309d7ded48c7a053",
				list: []string{"615c09f7309d7ded48c7a053", "615c09f7309d7ded48c7a054"},
			},
		},
		{
			name: "not found",
			args: args{
				id:   "615c09f7309d7ded48c7a052",
				list: []string{"615c09f7309d7ded48c7a053", "615c09f7309d7ded48c7a054"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FindInSlice(tt.args.id, tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("FindInSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
