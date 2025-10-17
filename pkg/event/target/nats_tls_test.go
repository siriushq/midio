package target

import (
	"path"
	"path/filepath"
	"testing"

	natsserver "github.com/nats-io/nats-server/v2/test"
	xnet "github.com/siriushq/midio/pkg/net"
)

func TestNatsConnTLSCustomCA(t *testing.T) {
	s, opts := natsserver.RunServerWithConfig(filepath.Join("testdata", "nats_tls.conf"))
	defer s.Shutdown()

	clientConfig := &NATSArgs{
		Enable: true,
		Address: xnet.Host{Name: "localhost",
			Port:      (xnet.Port(opts.Port)),
			IsPortSet: true},
		Subject:       "test",
		Secure:        true,
		CertAuthority: path.Join("testdata", "certs", "root_ca_cert.pem"),
	}

	con, err := clientConfig.connectNats()
	if err != nil {
		t.Errorf("Could not connect to nats: %v", err)
	}
	defer con.Close()
}

func TestNatsConnTLSClientAuthorization(t *testing.T) {
	s, opts := natsserver.RunServerWithConfig(filepath.Join("testdata", "nats_tls_client_cert.conf"))
	defer s.Shutdown()

	clientConfig := &NATSArgs{
		Enable: true,
		Address: xnet.Host{Name: "localhost",
			Port:      (xnet.Port(opts.Port)),
			IsPortSet: true},
		Subject:       "test",
		Secure:        true,
		CertAuthority: path.Join("testdata", "certs", "root_ca_cert.pem"),
		ClientCert:    path.Join("testdata", "certs", "nats_client_cert.pem"),
		ClientKey:     path.Join("testdata", "certs", "nats_client_key.pem"),
	}

	con, err := clientConfig.connectNats()
	if err != nil {
		t.Errorf("Could not connect to nats: %v", err)
	}
	defer con.Close()
}
