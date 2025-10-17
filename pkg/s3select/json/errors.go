package json

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

func errInvalidJSONType(err error) *s3Error {
	return &s3Error{
		code:       "InvalidJsonType",
		message:    "The JsonType is invalid. Only DOCUMENT and LINES are supported.",
		statusCode: 400,
		cause:      err,
	}
}

func errJSONParsingError(err error) *s3Error {
	return &s3Error{
		code:       "JSONParsingError",
		message:    "Encountered an error parsing the JSON file. Check the file and try again.",
		statusCode: 400,
		cause:      err,
	}
}
