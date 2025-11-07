/*
 * Below main package has canonical imports for 'go get' and 'go build'
 * to work with all other clones of github.com/siriushq/midio repository. For
 * more information refer https://golang.org/doc/go1.4#canonicalimports
 */

//go:generate go run main_build.go

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package main

// #include <stdlib.h>
//
// typedef void* midio_builder;
// typedef void* midio_server;
// typedef void (*midio_thread_closure)();
//
// static inline void __midio_call_closure(midio_thread_closure closure) {
//  	closure();
// }
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

// MidioFlags
// A holder object for flags provided when using the below building functions.
// Some are unused depending on whether a `server`, `gateway s3`, `gateway aws`, etc., is built.
type MidioFlags struct {
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
func midio_create() *C.midio_builder {
	builder := &MidioFlags{}
	return (*C.midio_builder)(unsafe.Pointer(builder))
}

//export midio_volume
func midio_volume(builder *C.midio_builder, path *C.char) {
	flags := (*MidioFlags)(unsafe.Pointer(builder))
	flags.Volumes = append(flags.Volumes, C.GoString(path))
}

//export midio_certificates
func midio_certificates(builder *C.midio_builder, directory *C.char) {
	flags := (*MidioFlags)(unsafe.Pointer(builder))
	flags.CertsDir = C.GoString(directory)
}

//export midio_address
func midio_address(builder *C.midio_builder, address *C.char) {
	flags := (*MidioFlags)(unsafe.Pointer(builder))
	flags.Address = C.GoString(address)
}

//export midio_quiet
func midio_quiet(builder *C.midio_builder) {
	flags := (*MidioFlags)(unsafe.Pointer(builder))
	flags.Quiet = true
}

//export midio_anonymous
func midio_anonymous(builder *C.midio_builder) {
	flags := (*MidioFlags)(unsafe.Pointer(builder))
	flags.Anonymous = true
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

// Run the provided function pointer in a goroutine.
// This is a simple, managed and cross-platform alternative to managing threads
// directly, which may be favourable depending on your usage environment.
//
//export midio_thread
func midio_thread(closure *C.midio_thread_closure) {
	go C.__midio_call_closure(*closure)
}

// Entry point for `midio server` in shared-library build, returning a closeable resource.
//
//export midio_create_server
func midio_create_server(builder *C.midio_builder) *C.midio_server {
	flags := (*MidioFlags)(unsafe.Pointer(builder))
	arguments := []string{"server"}

	logger.EnableShared(MidioChannel)
	if flags.Address != "" {
		arguments = append(arguments, "--address", flags.Address)
	}
	if flags.CertsDir != "" {
		arguments = append(arguments, "--certs-dir", flags.CertsDir)
	}
	if flags.Quiet {
		logger.EnableQuiet()
	}
	if flags.Anonymous {
		logger.EnableAnonymous()
	}

	for _, volume := range flags.Volumes {
		arguments = append(arguments, volume)
	}

	app := midio.NewApp("midio")
	go func() {
		if err := app.Run(arguments); err != nil {
			panic(err)
		}
	}()

	server := &MidioServer{App: app}
	return (*C.midio_server)(unsafe.Pointer(server))
}

//export midio_close
func midio_close(builder *C.midio_server) {
	server := (*MidioServer)(unsafe.Pointer(builder))
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
