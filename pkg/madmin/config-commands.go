package madmin

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

// GetConfig - returns the config.json of a minio setup, incoming data is encrypted.
func (adm *AdminClient) GetConfig(ctx context.Context) ([]byte, error) {
	// Execute GET on /minio/admin/v3/config to get config of a setup.
	resp, err := adm.executeMethod(ctx,
		http.MethodGet,
		requestData{relPath: adminAPIPrefix + "/config"})
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	return DecryptData(adm.getSecretKey(), resp.Body)
}

// SetConfig - set config supplied as config.json for the setup.
func (adm *AdminClient) SetConfig(ctx context.Context, config io.Reader) (err error) {
	const maxConfigJSONSize = 256 * 1024 // 256KiB

	// Read configuration bytes
	configBuf := make([]byte, maxConfigJSONSize+1)
	n, err := io.ReadFull(config, configBuf)
	if err == nil {
		return bytes.ErrTooLarge
	}
	if err != io.ErrUnexpectedEOF {
		return err
	}
	configBytes := configBuf[:n]
	econfigBytes, err := EncryptData(adm.getSecretKey(), configBytes)
	if err != nil {
		return err
	}

	reqData := requestData{
		relPath: adminAPIPrefix + "/config",
		content: econfigBytes,
	}

	// Execute PUT on /minio/admin/v3/config to set config.
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)

	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}

	return nil
}
