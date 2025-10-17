package crypto

import (
	"net/http"
	"testing"
)

func TestS3String(t *testing.T) {
	const Domain = "SSE-S3"
	if domain := S3.String(); domain != Domain {
		t.Errorf("S3's string method returns wrong domain: got '%s' - want '%s'", domain, Domain)
	}
}

func TestSSECString(t *testing.T) {
	const Domain = "SSE-C"
	if domain := SSEC.String(); domain != Domain {
		t.Errorf("SSEC's string method returns wrong domain: got '%s' - want '%s'", domain, Domain)
	}
}

var ssecUnsealObjectKeyTests = []struct {
	Headers        http.Header
	Bucket, Object string
	Metadata       map[string]string

	ExpectedErr error
}{
	{ // 0 - Valid HTTP headers and valid metadata entries for bucket/object
		Headers: http.Header{
			"X-Amz-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
			"X-Amz-Server-Side-Encryption-Customer-Key-Md5":   []string{"7PpPLAK26ONlVUGOWlusfg=="},
		},
		Bucket: "bucket",
		Object: "object",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key":     "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Seal-Algorithm": "DAREv2-HMAC-SHA256",
			"X-Minio-Internal-Server-Side-Encryption-Iv":             "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: nil,
	},
	{ // 1 - Valid HTTP headers but invalid metadata entries for bucket/object2
		Headers: http.Header{
			"X-Amz-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
			"X-Amz-Server-Side-Encryption-Customer-Key-Md5":   []string{"7PpPLAK26ONlVUGOWlusfg=="},
		},
		Bucket: "bucket",
		Object: "object2",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key":     "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Seal-Algorithm": "DAREv2-HMAC-SHA256",
			"X-Minio-Internal-Server-Side-Encryption-Iv":             "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: ErrSecretKeyMismatch,
	},
	{ // 2 - Valid HTTP headers but invalid metadata entries for bucket/object
		Headers: http.Header{
			"X-Amz-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
			"X-Amz-Server-Side-Encryption-Customer-Key-Md5":   []string{"7PpPLAK26ONlVUGOWlusfg=="},
		},
		Bucket: "bucket",
		Object: "object",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key": "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Iv":         "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: errMissingInternalSealAlgorithm,
	},
	{ // 3 - Invalid HTTP headers for valid metadata entries for bucket/object
		Headers: http.Header{
			"X-Amz-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
		},
		Bucket: "bucket",
		Object: "object",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key":     "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Seal-Algorithm": "DAREv2-HMAC-SHA256",
			"X-Minio-Internal-Server-Side-Encryption-Iv":             "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: ErrMissingCustomerKeyMD5,
	},
}

func TestSSECUnsealObjectKey(t *testing.T) {
	for i, test := range ssecUnsealObjectKeyTests {
		if _, err := SSEC.UnsealObjectKey(test.Headers, test.Metadata, test.Bucket, test.Object); err != test.ExpectedErr {
			t.Errorf("Test %d: got: %v - want: %v", i, err, test.ExpectedErr)
		}
	}
}

var sseCopyUnsealObjectKeyTests = []struct {
	Headers        http.Header
	Bucket, Object string
	Metadata       map[string]string

	ExpectedErr error
}{
	{ // 0 - Valid HTTP headers and valid metadata entries for bucket/object
		Headers: http.Header{
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5":   []string{"7PpPLAK26ONlVUGOWlusfg=="},
		},
		Bucket: "bucket",
		Object: "object",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key":     "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Seal-Algorithm": "DAREv2-HMAC-SHA256",
			"X-Minio-Internal-Server-Side-Encryption-Iv":             "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: nil,
	},
	{ // 1 - Valid HTTP headers but invalid metadata entries for bucket/object2
		Headers: http.Header{
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5":   []string{"7PpPLAK26ONlVUGOWlusfg=="},
		},
		Bucket: "bucket",
		Object: "object2",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key":     "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Seal-Algorithm": "DAREv2-HMAC-SHA256",
			"X-Minio-Internal-Server-Side-Encryption-Iv":             "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: ErrSecretKeyMismatch,
	},
	{ // 2 - Valid HTTP headers but invalid metadata entries for bucket/object
		Headers: http.Header{
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5":   []string{"7PpPLAK26ONlVUGOWlusfg=="},
		},
		Bucket: "bucket",
		Object: "object",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key": "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Iv":         "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: errMissingInternalSealAlgorithm,
	},
	{ // 3 - Invalid HTTP headers for valid metadata entries for bucket/object
		Headers: http.Header{
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm": []string{"AES256"},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key":       []string{"MzJieXRlc2xvbmdzZWNyZXRrZXltdXN0cHJvdmlkZWQ="},
		},
		Bucket: "bucket",
		Object: "object",
		Metadata: map[string]string{
			"X-Minio-Internal-Server-Side-Encryption-Sealed-Key":     "IAAfAMBdYor5tf/UlVaQvwYlw5yKbPBeQqfygqsfHqhu1wHD9KDAP4bw38AhL12prFTS23JbbR9Re5Qv26ZnlQ==",
			"X-Minio-Internal-Server-Side-Encryption-Seal-Algorithm": "DAREv2-HMAC-SHA256",
			"X-Minio-Internal-Server-Side-Encryption-Iv":             "coVfGS3I/CTrqexX5vUN+PQPoP9aUFiPYYrSzqTWfBA=",
		},
		ExpectedErr: ErrMissingCustomerKeyMD5,
	},
}

func TestSSECopyUnsealObjectKey(t *testing.T) {
	for i, test := range sseCopyUnsealObjectKeyTests {
		if _, err := SSECopy.UnsealObjectKey(test.Headers, test.Metadata, test.Bucket, test.Object); err != test.ExpectedErr {
			t.Errorf("Test %d: got: %v - want: %v", i, err, test.ExpectedErr)
		}
	}
}
