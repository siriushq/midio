package target

import (
	"testing"

	natsserver "github.com/nats-io/nats-server/v2/test"
	xnet "github.com/siriushq/midio/pkg/net"
)

func TestNatsConnPlain(t *testing.T) {
	opts := natsserver.DefaultTestOptions
	opts.Port = 14222
	s := natsserver.RunServer(&opts)
	defer s.Shutdown()

	clientConfig := &NATSArgs{
		Enable: true,
		Address: xnet.Host{Name: "localhost",
			Port:      (xnet.Port(opts.Port)),
			IsPortSet: true},
		Subject: "test",
	}
	con, err := clientConfig.connectNats()
	if err != nil {
		t.Errorf("Could not connect to nats: %v", err)
	}
	defer con.Close()
}

func TestNatsConnUserPass(t *testing.T) {
	opts := natsserver.DefaultTestOptions
	opts.Port = 14223
	opts.Username = "testminio"
	opts.Password = "miniotest"
	s := natsserver.RunServer(&opts)
	defer s.Shutdown()

	clientConfig := &NATSArgs{
		Enable: true,
		Address: xnet.Host{Name: "localhost",
			Port:      (xnet.Port(opts.Port)),
			IsPortSet: true},
		Subject:  "test",
		Username: opts.Username,
		Password: opts.Password,
	}

	con, err := clientConfig.connectNats()
	if err != nil {
		t.Errorf("Could not connect to nats: %v", err)
	}
	defer con.Close()
}

func TestNatsConnToken(t *testing.T) {
	opts := natsserver.DefaultTestOptions
	opts.Port = 14223
	opts.Authorization = "s3cr3t"
	s := natsserver.RunServer(&opts)
	defer s.Shutdown()

	clientConfig := &NATSArgs{
		Enable: true,
		Address: xnet.Host{Name: "localhost",
			Port:      (xnet.Port(opts.Port)),
			IsPortSet: true},
		Subject: "test",
		Token:   opts.Authorization,
	}

	con, err := clientConfig.connectNats()
	if err != nil {
		t.Errorf("Could not connect to nats: %v", err)
	}
	defer con.Close()
}
