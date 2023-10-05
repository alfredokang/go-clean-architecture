package shared

import (
	"github.com/google/uuid"
	"log"
)

func GetUuid() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Error", err)
	}

	return id
}

func GetUuidByString(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func GetUuidEmpty() uuid.UUID {
	return uuid.Nil
}
