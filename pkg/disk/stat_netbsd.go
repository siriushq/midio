//go:build netbsd
// +build netbsd

package disk

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// GetInfo returns total and free bytes available in a directory, e.g. `/`.
func GetInfo(path string) (info Info, err error) {
	s := unix.Statvfs_t{}
	if err = unix.Statvfs(path, &s); err != nil {
		return Info{}, err
	}
	reservedBlocks := uint64(s.Bfree) - uint64(s.Bavail)
	info = Info{
		Total:  uint64(s.Frsize) * (uint64(s.Blocks) - reservedBlocks),
		Free:   uint64(s.Frsize) * uint64(s.Bavail),
		Files:  uint64(s.Files),
		Ffree:  uint64(s.Ffree),
		FSType: string(s.Fstypename[:]),
	}
	if info.Free > info.Total {
		return info, fmt.Errorf("detected free space (%d) > total disk space (%d), fs corruption at (%s). please run 'fsck'", info.Free, info.Total, path)
	}
	info.Used = info.Total - info.Free
	return info, nil
}
