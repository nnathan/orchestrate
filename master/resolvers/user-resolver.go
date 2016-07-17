package resolvers

import (
	"encoding/json"
	"io"

	"github.com/vjftw/orchestrate/master/models"
)

type IUserResolver interface {
	FromRequest(*models.User, io.ReadCloser) error
}

type UserResolver struct {
}

func (uR UserResolver) FromRequest(u *models.User, b io.ReadCloser) error {

	var rJSON map[string]interface{}

	err := json.NewDecoder(b).Decode(&rJSON)
	if err != nil {
		return err
	}

	u.EmailAddress = rJSON["emailAddress"].(string)
	u.Password = rJSON["password"].(string)

	if val, ok := rJSON["firstName"]; ok {
		u.FirstName = val.(string)
	}

	if val, ok := rJSON["lastName"]; ok {
		u.LastName = val.(string)
	}

	return nil
}
