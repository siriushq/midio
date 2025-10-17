//go:build openbsd
// +build openbsd

package disk

import (
	"fmt"
	"syscall"
)

// GetInfo returns total and free bytes available in a directory, e.g. `/`.
func GetInfo(path string) (info Info, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return Info{}, err
	}
	reservedBlocks := uint64(s.F_bfree) - uint64(s.F_bavail)
	info = Info{
		Total:  uint64(s.F_bsize) * (uint64(s.F_blocks) - reservedBlocks),
		Free:   uint64(s.F_bsize) * uint64(s.F_bavail),
		Files:  uint64(s.F_files),
		Ffree:  uint64(s.F_ffree),
		FSType: getFSType(s.F_fstypename[:]),
	}
	if info.Free > info.Total {
		return info, fmt.Errorf("detected free space (%d) > total disk space (%d), fs corruption at (%s). please run 'fsck'", info.Free, info.Total, path)
	}
	info.Used = info.Total - info.Free
	return info, nil
}
