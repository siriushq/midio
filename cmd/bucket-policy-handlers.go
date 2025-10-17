package cmd

import (
	"encoding/json"
	"io"
	"net/http"

	humanize "github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	"github.com/siriushq/midio/cmd/logger"
	"github.com/siriushq/midio/pkg/bucket/policy"
)

const (
	// As per AWS S3 specification, 20KiB policy JSON data is allowed.
	maxBucketPolicySize = 20 * humanize.KiByte

	// Policy configuration file.
	bucketPolicyConfig = "policy.json"
)

// PutBucketPolicyHandler - This HTTP handler stores given bucket policy configuration as per
// https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html
func (api objectAPIHandlers) PutBucketPolicyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "PutBucketPolicy")

	defer logger.AuditLog(ctx, w, r, mustGetClaimsFromToken(r))

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	if s3Error := checkRequestAuthType(ctx, r, policy.PutBucketPolicyAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL, guessIsBrowserReq(r))
		return
	}

	// Check if bucket exists.
	if _, err := objAPI.GetBucketInfo(ctx, bucket); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Error out if Content-Length is missing.
	// PutBucketPolicy always needs Content-Length.
	if r.ContentLength <= 0 {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrMissingContentLength), r.URL, guessIsBrowserReq(r))
		return
	}

	// Error out if Content-Length is beyond allowed size.
	if r.ContentLength > maxBucketPolicySize {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrPolicyTooLarge), r.URL, guessIsBrowserReq(r))
		return
	}

	bucketPolicy, err := policy.ParseConfig(io.LimitReader(r.Body, r.ContentLength), bucket)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Version in policy must not be empty
	if bucketPolicy.Version == "" {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrMalformedPolicy), r.URL, guessIsBrowserReq(r))
		return
	}

	configData, err := json.Marshal(bucketPolicy)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	if err = globalBucketMetadataSys.Update(bucket, bucketPolicyConfig, configData); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Success.
	writeSuccessNoContent(w)
}

// DeleteBucketPolicyHandler - This HTTP handler removes bucket policy configuration.
func (api objectAPIHandlers) DeleteBucketPolicyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "DeleteBucketPolicy")

	defer logger.AuditLog(ctx, w, r, mustGetClaimsFromToken(r))

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	if s3Error := checkRequestAuthType(ctx, r, policy.DeleteBucketPolicyAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL, guessIsBrowserReq(r))
		return
	}

	// Check if bucket exists.
	if _, err := objAPI.GetBucketInfo(ctx, bucket); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	if err := globalBucketMetadataSys.Update(bucket, bucketPolicyConfig, nil); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Success.
	writeSuccessNoContent(w)
}

// GetBucketPolicyHandler - This HTTP handler returns bucket policy configuration.
func (api objectAPIHandlers) GetBucketPolicyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "GetBucketPolicy")

	defer logger.AuditLog(ctx, w, r, mustGetClaimsFromToken(r))

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	if s3Error := checkRequestAuthType(ctx, r, policy.GetBucketPolicyAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL, guessIsBrowserReq(r))
		return
	}

	// Check if bucket exists.
	if _, err := objAPI.GetBucketInfo(ctx, bucket); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Read bucket access policy.
	config, err := globalPolicySys.Get(bucket)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	configData, err := json.Marshal(config)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Write to client.
	writeSuccessResponseJSON(w, configData)
}
