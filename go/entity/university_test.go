package entity

import "testing"

func TestUniversity_ValidateClassroom(t *testing.T) {
	type fields struct {
		ID             string
		Name           string
		ProfilePicture string
		Classrooms     []Classroom
	}
	type args struct {
		classroomID string
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
				Classrooms: []Classroom{
					{
						ID: "615c09f7309d7def48c7a010",
					},
				},
			},
			args: args{
				classroomID: "615c09f7309d7def48c7a010",
			},
		},
		{
			name: "invalid classroom id",
			fields: fields{
				Classrooms: []Classroom{
					{
						ID: "615c09f7309d7def48c7a010",
					},
				},
			},
			args: args{
				classroomID: "615c09f7309d7def48c7a009",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &University{
				ID:             tt.fields.ID,
				Name:           tt.fields.Name,
				ProfilePicture: tt.fields.ProfilePicture,
				Classrooms:     tt.fields.Classrooms,
			}
			if err := u.ValidateClassroom(tt.args.classroomID); (err != nil) != tt.wantErr {
				t.Errorf("University.ValidateClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
