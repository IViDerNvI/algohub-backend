package core

var (
	ErrFileTooLarge    = NewErrCode(400, 1002, "File too large")
	ErrInvalidFileType = NewErrCode(400, 1003, "Invalid file type")
)
