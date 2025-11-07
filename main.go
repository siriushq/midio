/*
 * Below main package has canonical imports for 'go get' and 'go build'
 * to work with all other clones of github.com/siriushq/midio repository. For
 * more information refer https://golang.org/doc/go1.4#canonicalimports
 */

//go:generate go run main_build.go

package main

// #include <stdlib.h>
import "C"

import (
	"os"
	"path/filepath"
	"unsafe"

	midio "github.com/siriushq/midio/cmd"
	_ "github.com/siriushq/midio/cmd/gateway"
	"github.com/siriushq/midio/cmd/logger"

	"github.com/minio/cli"
)

// Entry-point for regular runnable binary build
func main() {
	arguments := os.Args
	appName := filepath.Base(arguments[0])

	app := midio.NewApp(appName)
	if err := app.Run(arguments); err != nil {
		os.Exit(1)
	}
}

// MidioBuilder
// A holder object for flags provided when using the below building functions.
// Some are unused depending on whether a `server`, `gateway s3`, `gateway aws`, etc., is built.
type MidioBuilder struct {
	Volumes []string

	Address   string
	CertsDir  string
	Quiet     bool
	Anonymous bool
}

// MidioServer
// A holder object for server instances, allowing for them to be later closed and shut down.
type MidioServer struct {
	App *cli.App
}

// MidioChannel
// Go channel where log messages are published if running as a shared
// library using the below building functions.
var MidioChannel = make(chan string, 16)

// Returns version currently being used, the result must have `free()` called
// against it later as `malloc()` is called, even though this is a constant.
//
//export midio_version
func midio_version() *C.char {
	return C.CString(midio.Version)
}

//export midio_create
func midio_create() *C.void {
	builder := &MidioBuilder{}
	return (*C.void)(unsafe.Pointer(builder))
}

//export midio_volume
func midio_volume(pointer *C.void, path *C.char) {
	builder := (*MidioBuilder)(unsafe.Pointer(pointer))
	builder.Volumes = append(builder.Volumes, C.GoString(path))
}

//export midio_certificates
func midio_certificates(pointer *C.void, directory *C.char) {
	builder := (*MidioBuilder)(unsafe.Pointer(pointer))
	builder.CertsDir = C.GoString(directory)
}

//export midio_address
func midio_address(pointer *C.void, address *C.char) {
	builder := (*MidioBuilder)(unsafe.Pointer(pointer))
	builder.Address = C.GoString(address)
}

//export midio_quiet
func midio_quiet(pointer *C.void) {
	builder := (*MidioBuilder)(unsafe.Pointer(pointer))
	builder.Quiet = true
}

//export midio_anonymous
func midio_anonymous(pointer *C.void) {
	builder := (*MidioBuilder)(unsafe.Pointer(pointer))
	builder.Anonymous = true
}

// Consume a message, block until another one is available.
// This method must be called consistently or the service will deadlock.
// The returned string must have `free()` called after consumption.
//
//export midio_message
func midio_message() *C.char {
	message := <-MidioChannel
	return C.CString(message)
}

// Entry point for `midio server` in shared-library build, returning a closeable resource.
//
//export midio_server
func midio_server(pointer *C.void) *C.void {
	builder := (*MidioBuilder)(unsafe.Pointer(pointer))
	arguments := []string{"server"}

	logger.EnableShared(MidioChannel)
	if builder.Address != "" {
		arguments = append(arguments, "--address", builder.Address)
	}
	if builder.CertsDir != "" {
		arguments = append(arguments, "--certs-dir", builder.CertsDir)
	}
	if builder.Quiet {
		arguments = append(arguments, "--quiet")
	}
	if builder.Anonymous {
		arguments = append(arguments, "--anonymous")
	}

	for _, volume := range builder.Volumes {
		arguments = append(arguments, volume)
	}

	app := midio.NewApp("midio")
	go func() {
		if err := app.Run(arguments); err != nil {
			panic(err)
		}
	}()

	server := &MidioServer{App: app}
	return (*C.void)(unsafe.Pointer(server))
}

//export midio_close
func midio_close(pointer *C.void) {
	server := (*MidioServer)(unsafe.Pointer(pointer))
	if server.App != nil {
		// TODO - use App in the future, for now just send interrupt
		// on Windows, we catch and ignore SIGINT error, then throw SIGKILL
		// midio cannot close without closing the current process, right now

		pid := os.Getpid()
		process, err := os.FindProcess(pid)
		if err != nil {
			panic(err)
		}

		if err := process.Signal(os.Interrupt); err == nil {
			return
		}
		if err := process.Signal(os.Kill); err != nil {
			panic(err)
		}
	}
}
