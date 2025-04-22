package core

// Error codes for user-related errors.
var (
	// Error codes for user authentication errors.
	ErrUserInvalid = NewErrCode(400, 4000, "User invalid")
	// Error codes for user registration errors.
	ErrUserNameNeed = NewErrCode(400, 4001, "User name required")
	// Error codes for user registration errors.
	ErrUserNameExist = NewErrCode(400, 4002, "User name already exists")
	// Error codes for user login errors.
	ErrUserVerify = NewErrCode(400, 4003, "username or password invalid")
)

// Error codes for user authentication errors.
var (
	// Error codes for user authentication errors.
	ErrNoAuthorization = NewErrCode(401, 3000, "No authorization")
	// Error codes for user login errors.
	ErrLoginNeed = NewErrCode(401, 3003, "Login required")
	// Error codes for user token errors.
	ErrAdminNeed = NewErrCode(403, 3004, "Admin required")
	// Error codes for user token errors.
	ErrTokenInvalid = NewErrCode(401, 3005, "Token invalid")
	// Error codes for user token errors.
	ErrTokenCreateFailed = NewErrCode(500, 3006, "Token create failed")

	ErrTokenUnsported = NewErrCode(500, 3007, "Token unsupported")

	ErrUnknownOperator = NewErrCode(500, 3008, "Unknown operator")
)
