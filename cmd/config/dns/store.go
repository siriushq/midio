package dns

// Error - DNS related errors error.
type Error struct {
	Bucket string
	Err    error
}

// ErrInvalidBucketName for buckets with invalid name
type ErrInvalidBucketName Error

func (e ErrInvalidBucketName) Error() string {
	return e.Bucket + " invalid bucket name error: " + e.Err.Error()
}

func (e Error) Error() string {
	return "dns related error: " + e.Err.Error()
}

// ErrBucketConflict for buckets that already exist
type ErrBucketConflict Error

func (e ErrBucketConflict) Error() string {
	return e.Bucket + " bucket conflict error: " + e.Err.Error()
}

// Store dns record store
type Store interface {
	Put(bucket string) error
	Get(bucket string) ([]SrvRecord, error)
	Delete(bucket string) error
	List() (map[string][]SrvRecord, error)
	DeleteRecord(record SrvRecord) error
	Close() error
	String() string
}
