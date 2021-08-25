package entity

type TinyUniversity struct {
	ID             string
	Name           string
	ProfilePicture string
}

type University struct {
	ID             string
	Name           string
	ProfilePicture string
	Classrooms     []Classroom
}

type UniversityPayload struct {
	ID             string
	Name           string
	ProfilePicture string
}
