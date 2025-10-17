/*
 * Below main package has canonical imports for 'go get' and 'go build'
 * to work with all other clones of github.com/minio/minio repository. For
 * more information refer https://golang.org/doc/go1.4#canonicalimports
 */

//go:generate go run main_build.go

package main

import (
	"os"
	"path/filepath"

	midio "github.com/siriushq/midio/cmd"

	// Import gateway
	_ "github.com/siriushq/midio/cmd/gateway"
)

func main() {
	args := os.Args
	appName := filepath.Base(args[0])

	app := midio.NewApp(appName)
	if err := app.Run(args); err != nil {
		os.Exit(1)
	}
}
