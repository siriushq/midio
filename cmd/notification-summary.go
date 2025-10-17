package cmd

import (
	"github.com/siriushq/midio/pkg/madmin"
)

// GetTotalCapacity gets the total capacity in the cluster.
func GetTotalCapacity(diskInfo []madmin.Disk) (capacity uint64) {

	for _, disk := range diskInfo {
		capacity += disk.TotalSpace
	}
	return
}

// GetTotalUsableCapacity gets the total usable capacity in the cluster.
// This value is not an accurate representation of total usable in a multi-tenant deployment.
func GetTotalUsableCapacity(diskInfo []madmin.Disk, s StorageInfo) (capacity float64) {
	raw := GetTotalCapacity(diskInfo)
	var approxDataBlocks float64
	var actualDisks float64
	for _, scData := range s.Backend.StandardSCData {
		approxDataBlocks += float64(scData)
		actualDisks += float64(scData + s.Backend.StandardSCParity)
	}
	ratio := approxDataBlocks / actualDisks
	return float64(raw) * ratio
}

// GetTotalCapacityFree gets the total capacity free in the cluster.
func GetTotalCapacityFree(diskInfo []madmin.Disk) (capacity uint64) {
	for _, d := range diskInfo {
		capacity += d.AvailableSpace
	}
	return
}

// GetTotalUsableCapacityFree gets the total usable capacity free in the cluster.
// This value is not an accurate representation of total free in a multi-tenant deployment.
func GetTotalUsableCapacityFree(diskInfo []madmin.Disk, s StorageInfo) (capacity float64) {
	raw := GetTotalCapacityFree(diskInfo)
	var approxDataBlocks float64
	var actualDisks float64
	for _, scData := range s.Backend.StandardSCData {
		approxDataBlocks += float64(scData)
		actualDisks += float64(scData + s.Backend.StandardSCParity)
	}
	ratio := approxDataBlocks / actualDisks
	return float64(raw) * ratio
}
