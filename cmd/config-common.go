package cmd

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/siriushq/midio/pkg/hash"
)

var errConfigNotFound = errors.New("config file not found")

func readConfig(ctx context.Context, objAPI ObjectLayer, configFile string) ([]byte, error) {
	// Read entire content by setting size to -1
	r, err := objAPI.GetObjectNInfo(ctx, minioMetaBucket, configFile, nil, http.Header{}, readLock, ObjectOptions{})
	if err != nil {
		// Treat object not found as config not found.
		if isErrObjectNotFound(err) {
			return nil, errConfigNotFound
		}

		return nil, err
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, errConfigNotFound
	}
	return buf, nil
}

type objectDeleter interface {
	DeleteObject(ctx context.Context, bucket, object string, opts ObjectOptions) (ObjectInfo, error)
}

func deleteConfig(ctx context.Context, objAPI objectDeleter, configFile string) error {
	_, err := objAPI.DeleteObject(ctx, minioMetaBucket, configFile, ObjectOptions{})
	if err != nil && isErrObjectNotFound(err) {
		return errConfigNotFound
	}
	return err
}

func saveConfig(ctx context.Context, objAPI ObjectLayer, configFile string, data []byte) error {
	hashReader, err := hash.NewReader(bytes.NewReader(data), int64(len(data)), "", getSHA256Hash(data), int64(len(data)))
	if err != nil {
		return err
	}

	_, err = objAPI.PutObject(ctx, minioMetaBucket, configFile, NewPutObjReader(hashReader), ObjectOptions{MaxParity: true})
	return err
}

func checkConfig(ctx context.Context, objAPI ObjectLayer, configFile string) error {
	if _, err := objAPI.GetObjectInfo(ctx, minioMetaBucket, configFile, ObjectOptions{}); err != nil {
		// Treat object not found as config not found.
		if isErrObjectNotFound(err) {
			return errConfigNotFound
		}

		return err
	}
	return nil
}
