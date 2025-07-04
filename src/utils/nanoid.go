package utils

import (
	"fmt"

	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateNanoId(len int) (string, error) {
	id, err := nanoid.New(len)
	if err != nil {
		return "", fmt.Errorf("Error while generating nanoid: %v", err)
	}
	return id, nil
}
