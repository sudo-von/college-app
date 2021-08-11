package entity

const (
	FacultadSistemas = "60a1b132cc1fa10008fa31ad"
)

type University struct {
	ID          string
	Name        string
	Description string
	Picture     string
}

type UniversityPayload struct {
	ID          string
	Name        string
	Description string
	Picture     string
}
