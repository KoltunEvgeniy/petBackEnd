package modelErrors

import "errors"

var (
	ErrName            = errors.New("Name_empty")
	ErrPhone           = errors.New("Phone_empty")
	ErrEmail           = errors.New("Email_empty")
	ErrRole            = errors.New("Role_empty")
	ErrAllEmpty        = errors.New("all_field_empty")
	ErrAccountNotFound = errors.New("Account_not_found")
	ErrSmsCode         = errors.New("Invalid_or_expired_code")
	ErrRefreshToken    = errors.New("Invalid_Refresh_Token")
	ErrTokenExp        = errors.New("Token_Expired")
	ErrMaster          = errors.New("Master_not_found")
)
