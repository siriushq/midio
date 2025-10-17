package target

import "github.com/google/uuid"

func getNewUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
