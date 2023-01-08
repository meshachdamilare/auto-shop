package shops

import "github.com/google/uuid"

// Generate a new Id for each product added to the store
func GenerateID() string {
	Id := uuid.New().String()
	return Id
}
