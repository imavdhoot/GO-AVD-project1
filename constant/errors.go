package constant

import "errors"

var (
	//varErrRequestBody          = errors.New("Error in parsing request body")

	ErrRecNotFound = "record not found"

	ErrMerchantAlreadyExist = errors.New("merchant already exists with same name")
	ErrMerchantNotFound     = errors.New("merchant not found for that ID")
	ErrMerchantIdInvalid    = errors.New("invalid/empty merchant ID for update")
	ErrMemberIdInvalid      = errors.New("invalid/empty merchant ID for update")
	ErrInvalidEmailAddress  = errors.New("invalid email address")
	ErrMerchantIdEmpty      = errors.New("invalid/empty merchant ID")
	ErrInvalidPageNo        = errors.New("invalid page number in query")
)
