package constant

import "errors"

var (
	//varErrRequestBody          = errors.New("Error in parsing request body")
	ErrMerchantAlreadyExist = errors.New("merchant already exists with same name")
)
