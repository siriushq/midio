package csv

import "errors"

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

func errCSVParsingError(err error) *s3Error {
	return &s3Error{
		code:       "CSVParsingError",
		message:    "Encountered an error parsing the CSV file. Check the file and try again.",
		statusCode: 400,
		cause:      err,
	}
}

func errInvalidTextEncodingError() *s3Error {
	return &s3Error{
		code:       "InvalidTextEncoding",
		message:    "UTF-8 encoding is required.",
		statusCode: 400,
		cause:      errors.New("invalid utf8 encoding"),
	}
}
