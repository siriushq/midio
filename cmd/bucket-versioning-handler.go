package cmd

import (
	"encoding/xml"
	"io"
	"net/http"

	humanize "github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	"github.com/siriushq/midio/cmd/logger"
	"github.com/siriushq/midio/pkg/bucket/policy"
	"github.com/siriushq/midio/pkg/bucket/versioning"
)

const (
	bucketVersioningConfig = "versioning.xml"

	// Maximum size of bucket versioning configuration payload sent to the PutBucketVersioningHandler.
	maxBucketVersioningConfigSize = 1 * humanize.MiByte
)

// PutBucketVersioningHandler - PUT Bucket Versioning.
// ----------
func (api objectAPIHandlers) PutBucketVersioningHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "PutBucketVersioning")

	defer logger.AuditLog(ctx, w, r, mustGetClaimsFromToken(r))

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	if s3Error := checkRequestAuthType(ctx, r, policy.PutBucketVersioningAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL, guessIsBrowserReq(r))
		return
	}

	v, err := versioning.ParseConfig(io.LimitReader(r.Body, maxBucketVersioningConfigSize))
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	if rcfg, _ := globalBucketObjectLockSys.Get(bucket); rcfg.LockEnabled && v.Suspended() {
		writeErrorResponse(ctx, w, APIError{
			Code:           "InvalidBucketState",
			Description:    "An Object Lock configuration is present on this bucket, so the versioning state cannot be changed.",
			HTTPStatusCode: http.StatusConflict,
		}, r.URL, guessIsBrowserReq(r))
		return
	}
	if _, err := getReplicationConfig(ctx, bucket); err == nil && v.Suspended() {
		writeErrorResponse(ctx, w, APIError{
			Code:           "InvalidBucketState",
			Description:    "A replication configuration is present on this bucket, so the versioning state cannot be changed.",
			HTTPStatusCode: http.StatusConflict,
		}, r.URL, guessIsBrowserReq(r))
		return
	}

	configData, err := xml.Marshal(v)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	if err = globalBucketMetadataSys.Update(bucket, bucketVersioningConfig, configData); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	writeSuccessResponseHeadersOnly(w)
}

// GetBucketVersioningHandler - GET Bucket Versioning.
// ----------
func (api objectAPIHandlers) GetBucketVersioningHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "GetBucketVersioning")

	defer logger.AuditLog(ctx, w, r, mustGetClaimsFromToken(r))

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	if s3Error := checkRequestAuthType(ctx, r, policy.GetBucketVersioningAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL, guessIsBrowserReq(r))
		return
	}

	// Check if bucket exists.
	if _, err := objectAPI.GetBucketInfo(ctx, bucket); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	config, err := globalBucketVersioningSys.Get(bucket)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	configData, err := xml.Marshal(config)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Write bucket versioning configuration to client
	writeSuccessResponseXML(w, configData)

}
