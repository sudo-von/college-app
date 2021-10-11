package search

import "errors"

func FindInSlice(id string, list []string) error {

	found := false
	for _, l := range list {
		if id == l {
			found = true
		}
	}
	if !found {
		return errors.New("not found")
	}
	return nil

}
