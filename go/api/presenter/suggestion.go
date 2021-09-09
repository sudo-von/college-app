package presenter

import (
	"errors"
	"net/http"
	"strings"
)

type SuggestionPayload struct {
	Suggestion string `json:"suggestion"`
}

func (sp *SuggestionPayload) validate() (err error) {
	if len(strings.TrimSpace(sp.Suggestion)) == 0 {
		err = mergeErrors(err, errors.New("missing field suggestion"))
	}
	return
}

func (sp *SuggestionPayload) Bind(r *http.Request) error {
	if err := sp.validate(); err != nil {
		return err
	}
	return nil
}
