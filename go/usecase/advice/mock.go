package advice

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/search"
	"time"
)

var (
	fakeAdviceIDS     = []string{"615e7802e13d8485f4f5e35e", "615e7802e13d8485f4f5e35f"}
	fakeUniversityIDS = []string{"615c09f7309d7ded48c7a049", "615c0794590f4315693633a6"}
)

type AdviceReaderMock struct{}

func (a AdviceReaderMock) GetAdviceByID(adviceID string) (*entity.Advice, error) {

	err := search.FindInSlice(adviceID, fakeAdviceIDS)
	if err != nil {
		return nil, err
	}

	advice := &entity.Advice{
		ID: adviceID,
		User: entity.TinyUser{
			ID: "615c09f7309d7ded48c7a053",
		},
		Classroom: entity.Classroom{
			ID:   "615c09f7309d7def48c7a010",
			Name: "Fake name",
		},
		UniversityID:       "615c09f7309d7ded48c7a049",
		Subject:            "Fake subject",
		AdviceDate:         time.Date(2021, 01, 02, 0, 0, 0, 0, time.Local),
		StudentsWillAttend: []string{"615c09f7309d7ded48c7a053"},
		CreationDate:       time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
	return advice, nil
}

func (a AdviceReaderMock) GetAdvices(universityID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error) {

	err := search.FindInSlice(universityID, fakeUniversityIDS)
	if err != nil {
		return nil, nil, err
	}

	advices := []entity.Advice{
		{
			ID: "615e7802e13d8485f4f5e35e",
			User: entity.TinyUser{
				ID: "615c09f7309d7ded48c7a053",
			},
			Classroom: entity.Classroom{
				ID:   "615c09f7309d7def48c7a010",
				Name: "Fake name",
			},
			UniversityID:       universityID,
			Subject:            "Fake subject",
			AdviceDate:         time.Date(2021, 01, 02, 0, 0, 0, 0, time.Local),
			StudentsWillAttend: []string{"615c09f7309d7ded48c7a053"},
			CreationDate:       time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
		},
	}
	total := len(advices)
	return advices, &total, nil
}

type AdviceWriterMock struct{}

func (a AdviceWriterMock) CreateAdvice(newAdvice entity.AdvicePayload) error {
	return nil
}

func (a AdviceWriterMock) UpdateAdvice(updatedAdvice entity.AdvicePayload) error {
	return nil
}

type AdviceRepositoryMock struct {
	AdviceReaderMock
	AdviceWriterMock
}
