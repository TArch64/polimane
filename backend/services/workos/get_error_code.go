package workos

import (
	"encoding/json"

	"github.com/workos/workos-go/v4/pkg/workos_errors"
)

const (
	CodeEmailVerificationCodeExpired = "email_verification_code_expired"
	CodeUnknownError                 = "unknown"
)

type httpErrorRawBody struct {
	Code string `json:"code"`
}

func GetErrorCode(httpErr workos_errors.HTTPError) string {
	var rawBody httpErrorRawBody
	if err := json.Unmarshal([]byte(httpErr.RawBody), &rawBody); err != nil {
		return CodeUnknownError
	}
	if rawBody.Code == "" {
		return CodeUnknownError
	}
	return rawBody.Code
}
