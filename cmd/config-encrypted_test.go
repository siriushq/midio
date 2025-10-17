package cmd

import (
	"bytes"
	"testing"

	"github.com/siriushq/midio/pkg/auth"
	"github.com/siriushq/midio/pkg/madmin"
)

func TestDecryptData(t *testing.T) {
	cred1 := auth.Credentials{
		AccessKey: "minio",
		SecretKey: "minio123",
	}

	cred2 := auth.Credentials{
		AccessKey: "minio",
		SecretKey: "minio1234",
	}

	data := []byte(`config data`)
	edata1, err := madmin.EncryptData(cred1.String(), data)
	if err != nil {
		t.Fatal(err)
	}

	edata2, err := madmin.EncryptData(cred2.String(), data)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		edata   []byte
		creds   []auth.Credentials
		success bool
	}{
		{edata1, []auth.Credentials{cred1, cred2}, true},
		{edata2, []auth.Credentials{cred1, cred2}, true},
		{data, []auth.Credentials{cred1, cred2}, false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			ddata, err := decryptData(test.edata, test.creds...)
			if err != nil && test.success {
				t.Errorf("Expected success, saw failure %v", err)
			}
			if err == nil && !test.success {
				t.Error("Expected failure, saw success")
			}
			if test.success {
				if !bytes.Equal(ddata, data) {
					t.Errorf("Expected %s, got %s", string(data), string(ddata))
				}
			}
		})
	}
}
