package sql

import "fmt"

type s3Error struct {
	code       string
	message    string
	statusCode int
	cause      error
}

func (err *s3Error) Cause() error {
	return err.cause
}

func (err *s3Error) ErrorCode() string {
	return err.code
}

func (err *s3Error) ErrorMessage() string {
	return err.message
}

func (err *s3Error) HTTPStatusCode() int {
	return err.statusCode
}

func (err *s3Error) Error() string {
	return err.message
}

func errInvalidDataType(err error) *s3Error {
	return &s3Error{
		code:       "InvalidDataType",
		message:    "The SQL expression contains an invalid data type.",
		statusCode: 400,
		cause:      err,
	}
}

func errIncorrectSQLFunctionArgumentType(err error) *s3Error {
	return &s3Error{
		code:       "IncorrectSqlFunctionArgumentType",
		message:    "Incorrect type of arguments in function call.",
		statusCode: 400,
		cause:      err,
	}
}

func errLikeInvalidInputs(err error) *s3Error {
	return &s3Error{
		code:       "LikeInvalidInputs",
		message:    "Invalid argument given to the LIKE clause in the SQL expression.",
		statusCode: 400,
		cause:      err,
	}
}

func errQueryParseFailure(err error) *s3Error {
	return &s3Error{
		code:       "ParseSelectFailure",
		message:    err.Error(),
		statusCode: 400,
		cause:      err,
	}
}

func errQueryAnalysisFailure(err error) *s3Error {
	return &s3Error{
		code:       "InvalidQuery",
		message:    err.Error(),
		statusCode: 400,
		cause:      err,
	}
}

func errBadTableName(err error) *s3Error {
	return &s3Error{
		code:       "BadTableName",
		message:    fmt.Sprintf("The table name is not supported: %v", err),
		statusCode: 400,
		cause:      err,
	}
}

func errDataSource(err error) *s3Error {
	return &s3Error{
		code:       "DataSourcePathUnsupported",
		message:    fmt.Sprintf("Data source: %v", err),
		statusCode: 400,
		cause:      err,
	}
}
