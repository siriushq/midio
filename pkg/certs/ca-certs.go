package certs

import (
	"crypto/x509"
	"io/ioutil"
	"os"
	"path"
)

// GetRootCAs - returns all the root CAs into certPool
// at the input certsCADir
func GetRootCAs(certsCAsDir string) (*x509.CertPool, error) {
	rootCAs, _ := loadSystemRoots()
	if rootCAs == nil {
		// In some systems system cert pool is not supported
		// or no certificates are present on the
		// system - so we create a new cert pool.
		rootCAs = x509.NewCertPool()
	}

	fis, err := ioutil.ReadDir(certsCAsDir)
	if err != nil {
		if os.IsNotExist(err) || os.IsPermission(err) {
			// Return success if CA's directory is missing or permission denied.
			return rootCAs, nil
		}
		return rootCAs, err
	}

	// Load all custom CA files.
	for _, fi := range fis {
		caCert, err := ioutil.ReadFile(path.Join(certsCAsDir, fi.Name()))
		if err == nil {
			rootCAs.AppendCertsFromPEM(caCert)
		}
		// ignore files which are not readable.
	}

	return rootCAs, nil
}
