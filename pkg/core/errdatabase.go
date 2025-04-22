package core

// Error codes for database-related errors.
var (
	// Error codes for database connection errors.
	ErrDatabaseUpdate = NewErrCode(500, 2000, "Database update error")
	// Error codes for database deletion errors.
	ErrDatabaseDelete = NewErrCode(500, 2001, "Database delete error")
	// Error codes for database retrieval errors.
	ErrDatabaseGet = NewErrCode(500, 2002, "Database get error")
	// Error codes for database creation errors.
	ErrDatabaseCreate = NewErrCode(500, 2003, "Database create error")
	// Error codes for database query errors.
	ErrDatabaseQuery = NewErrCode(500, 2004, "Database query error")

	// Error codes for database not found errors.
	ErrDatabaseNotFound = NewErrCode(404, 2005, "Database not found")
	// Error codes for database duplicate entry errors.
	ErrDatabaseDuplicate = NewErrCode(409, 2006, "Database duplicate entry")
)
