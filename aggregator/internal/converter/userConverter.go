package converter

import (
	"aggregatorProject/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

type UserConvert struct {
}

func NewUserConvert() *UserConvert {
	return &UserConvert{}
}

func (c *UserConvert) ConvertBytesToUser(req *http.Request) (*model.User, error) {
	dataBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	user := &model.User{}

	err = json.Unmarshal(dataBytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserConvert) ConvertUserToBytes(user *model.User) ([]byte, error) {
	dataBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return []byte{}, err
	}
	return dataBytes, nil
}
