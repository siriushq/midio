// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package disk

import "runtime"

// IsRootDisk returns if diskPath belongs to root-disk, i.e the disk mounted at "/"
func IsRootDisk(diskPath string, rootDisk string) (bool, error) {
	if runtime.GOOS == "windows" {
		// On windows this function is not implemented.
		return false, nil
	}
	return SameDisk(diskPath, rootDisk)
}
