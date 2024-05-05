package constant

type ErrorCode int

const (
	TokenIsExpired                  ErrorCode = 1
	TokenIsInvalid                  ErrorCode = 2
	TokenMustBeRefresh              ErrorCode = 3
	UserIsNotAuthorized             ErrorCode = 4
	TokenMustBeAccess               ErrorCode = 5
	UserIsNotActive                 ErrorCode = 6
	SuccessLogin                    ErrorCode = 7
	SuccessDeleted                  ErrorCode = 8
	SuccessUpdated                  ErrorCode = 9
	SuccessCreated                  ErrorCode = 10
	SuccessSend                     ErrorCode = 11
	SuccessLogout                   ErrorCode = 12
	MailFailed                      ErrorCode = 13
	BadRequest                      ErrorCode = 400
	UserIsNotAuthenticatErrorCodeed ErrorCode = 401
	FORBIDDEN                       ErrorCode = 403
	NotFound                        ErrorCode = 404
	MethodNotAllowed                ErrorCode = 405
	TooManyRequests                 ErrorCode = 429
	InternalServerError             ErrorCode = 500
	REQUIRED                        ErrorCode = 1001
	MIN                             ErrorCode = 1002
	MAX                             ErrorCode = 1003
	BOOLEAN                         ErrorCode = 1004
	NUMERIC                         ErrorCode = 1005
	TooLong                         ErrorCode = 1006
	AlreadyTaken                    ErrorCode = 1007
	NotExists                       ErrorCode = 1008
	INVALID                         ErrorCode = 1009
	NotMatch                        ErrorCode = 1010
	FAILED                          ErrorCode = 1011
	NotIn                           ErrorCode = 1012
	DateFormat                      ErrorCode = 1013
	FileType                        ErrorCode = 1014
	FileMimes                       ErrorCode = 1015
	NotSelected                     ErrorCode = 1016
	DUPLICATE                       ErrorCode = 1017
	STRING                          ErrorCode = 1018
	INTEGER                         ErrorCode = 1019
	ARRAY                           ErrorCode = 1020
	EMAIL                           ErrorCode = 1021
	LENGTH                          ErrorCode = 1022
	DIGITS                          ErrorCode = 1023
	NUMBER                          ErrorCode = 1024
	NotLastChildren                 ErrorCode = 1090
	NotColor                        ErrorCode = 2001
)
