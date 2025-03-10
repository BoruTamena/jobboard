package errors

import "github.com/joomcode/errorx"

var (
	InternalError = errorx.NewNamespace("Internal Server:error")
	InputError    = errorx.NewNamespace("InputValiation:error")
	DbError       = errorx.NewNamespace("DataBase:error")
)

var (

	// Internal server errors
	BadInput     = errorx.NewType(InputError, "Bad user input:error")
	MarshalErr   = errorx.NewType(InternalError, "unable to marshal:error")
	UnMarshalErr = errorx.NewType(InternalError, "unable to unmarshal:error")
	DublicateErr = errorx.NewType(InternalError, "Duplicate::error")
	AuthErr      = errorx.NewType(InternalError, "Unauthorized user :: error")

	// Database errors
	WriteErr  = errorx.NewType(DbError, "db write :: error ")
	NullObjId = errorx.NewType(DbError, "Null Object Id Returned :: error")
	DbReadErr = errorx.NewType(DbError, "db read::error")
)

var (
	ErrorCode = errorx.RegisterProperty("ERRCODE")
)
