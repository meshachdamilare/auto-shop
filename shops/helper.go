package shops

import "github.com/google/uuid"

func GenerateID() string {
	Id := uuid.New().String()
	return Id
}
