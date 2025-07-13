package auth

import (
	"encoding/base64"
	"encoding/json"
)

type loginState struct {
	ReturnTo string `json:"returnTo"`
}

func newLoginState(query *loginQueryParams) (string, error) {
	state := &loginState{ReturnTo: query.ReturnTo}

	bytes, err := json.Marshal(state)
	if err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(bytes)
	return token, nil
}

func parseLoginState(stateStr string) (*loginState, error) {
	bytes, err := base64.StdEncoding.DecodeString(stateStr)
	if err != nil {
		return nil, err
	}

	var state loginState
	err = json.Unmarshal(bytes, &state)
	return &state, err
}
