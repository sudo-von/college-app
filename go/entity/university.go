package entity

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

type Classroom struct {
	ID   string
	Name string
}
