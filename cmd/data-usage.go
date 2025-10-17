package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/siriushq/midio/cmd/logger"
	"github.com/siriushq/midio/pkg/hash"
	"github.com/siriushq/midio/pkg/madmin"
)

const (
	dataUsageRoot   = SlashSeparator
	dataUsageBucket = minioMetaBucket + SlashSeparator + bucketMetaPrefix

	dataUsageObjName   = ".usage.json"
	dataUsageCacheName = ".usage-cache.bin"
	dataUsageBloomName = ".bloomcycle.bin"
)

// storeDataUsageInBackend will store all objects sent on the gui channel until closed.
func storeDataUsageInBackend(ctx context.Context, objAPI ObjectLayer, dui <-chan madmin.DataUsageInfo) {
	for dataUsageInfo := range dui {
		dataUsageJSON, err := json.Marshal(dataUsageInfo)
		if err != nil {
			logger.LogIf(ctx, err)
			continue
		}
		size := int64(len(dataUsageJSON))
		r, err := hash.NewReader(bytes.NewReader(dataUsageJSON), size, "", "", size)
		if err != nil {
			logger.LogIf(ctx, err)
			continue
		}
		_, err = objAPI.PutObject(ctx, dataUsageBucket, dataUsageObjName, NewPutObjReader(r), ObjectOptions{})
		if !isErrBucketNotFound(err) {
			logger.LogIf(ctx, err)
		}
	}
}

func loadDataUsageFromBackend(ctx context.Context, objAPI ObjectLayer) (madmin.DataUsageInfo, error) {
	r, err := objAPI.GetObjectNInfo(ctx, dataUsageBucket, dataUsageObjName, nil, http.Header{}, readLock, ObjectOptions{})
	if err != nil {
		if isErrObjectNotFound(err) || isErrBucketNotFound(err) {
			return madmin.DataUsageInfo{}, nil
		}
		return madmin.DataUsageInfo{}, toObjectErr(err, dataUsageBucket, dataUsageObjName)
	}
	defer r.Close()

	var dataUsageInfo madmin.DataUsageInfo
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.NewDecoder(r).Decode(&dataUsageInfo); err != nil {
		return madmin.DataUsageInfo{}, err
	}

	// For forward compatibility reasons, we need to add this code.
	if len(dataUsageInfo.BucketsUsage) == 0 {
		dataUsageInfo.BucketsUsage = make(map[string]madmin.BucketUsageInfo, len(dataUsageInfo.BucketSizes))
		for bucket, size := range dataUsageInfo.BucketSizes {
			dataUsageInfo.BucketsUsage[bucket] = madmin.BucketUsageInfo{Size: size}
		}
	}

	// For backward compatibility reasons, we need to add this code.
	if len(dataUsageInfo.BucketSizes) == 0 {
		dataUsageInfo.BucketSizes = make(map[string]uint64, len(dataUsageInfo.BucketsUsage))
		for bucket, bui := range dataUsageInfo.BucketsUsage {
			dataUsageInfo.BucketSizes[bucket] = bui.Size
		}
	}

	return dataUsageInfo, nil
}
