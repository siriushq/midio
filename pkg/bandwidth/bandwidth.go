package bandwidth

// Details for the measured bandwidth
type Details struct {
	LimitInBytesPerSecond            int64   `json:"limitInBits"`
	CurrentBandwidthInBytesPerSecond float64 `json:"currentBandwidth"`
}

// Report captures the details for all buckets.
type Report struct {
	BucketStats map[string]Details `json:"bucketStats,omitempty"`
}
