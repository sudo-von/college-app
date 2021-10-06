package university

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/search"
)

var fakeUniversityIDS = []string{"615c09f7309d7ded48c7a049", "615c09f7309d7ded48c7a050"}

type UniversityReaderMock struct{}

func (u UniversityReaderMock) GetTinyUniversities() ([]entity.TinyUniversity, *int, error) {
	tinyUniversities := []entity.TinyUniversity{
		{
			ID:             "615c09f7309d7ded48c7a050",
			Name:           "Fake name",
			ProfilePicture: "Fake profile picture",
		},
	}
	total := len(tinyUniversities)
	return tinyUniversities, &total, nil
}

func (u UniversityReaderMock) GetUniversityByID(universityID string) (*entity.University, error) {

	err := search.FindInSlice(universityID, fakeUniversityIDS)
	if err != nil {
		return nil, err
	}

	university := &entity.University{
		ID:             universityID,
		Name:           "Fake name",
		ProfilePicture: "Fake profile picture",
		Classrooms: []entity.Classroom{
			{
				ID:   "615c09f7309d7def48c7a010",
				Name: "Fake name",
			},
		},
	}
	return university, nil
}

type UniversityRepositoryMock struct {
	UniversityReaderMock
}
