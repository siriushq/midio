//go:build linux
// +build linux

/*
 * This file has been adopted and then modified from Daniel Swarbrick's smart
 * project residing at https://github.com/dswarbrick/smart
 *
 */

package smart

import (
	"fmt"
	"math/big"
	"unsafe"

	"github.com/dswarbrick/smart/ioctl"
	"golang.org/x/sys/unix"
)

// NVMe admin disk query constants
const (
	NvmeAdminGetLogPage = 0x02
	NvmeAdminIdentify   = 0x06
)

var (
	nvmeIoctlAdminCmd = ioctl.Iowr('N', 0x41, unsafe.Sizeof(nvmePassthruCommand{}))
)

// NewNVMeDevice creates a new NVMeDevice struct with name
func NewNVMeDevice(name string) *NVMeDevice {
	return &NVMeDevice{name, -1}
}

// Open - open device file to find kernel info
func (d *NVMeDevice) Open() (err error) {
	d.fd, err = unix.Open(d.Name, unix.O_RDWR, 0600)
	return err
}

// Close - closes device file
func (d *NVMeDevice) Close() error {
	return unix.Close(d.fd)
}

func (d *NVMeDevice) readLogPage(logID uint8, buf *[]byte) error {
	bufLen := len(*buf)

	if (bufLen < 4) || (bufLen > 0x4000) || (bufLen%4 != 0) {
		return fmt.Errorf("Invalid buffer size")
	}

	cmd := nvmePassthruCommand{
		opcode:  NvmeAdminGetLogPage,
		nsid:    0xffffffff, // FIXME
		addr:    uint64(uintptr(unsafe.Pointer(&(*buf)[0]))),
		dataLen: uint32(bufLen),
		cdw10:   uint32(logID) | (((uint32(bufLen) / 4) - 1) << 16),
	}

	return ioctl.Ioctl(uintptr(d.fd), nvmeIoctlAdminCmd, uintptr(unsafe.Pointer(&cmd)))
}

// le128ToBigInt takes a little-endian 16-byte slice and returns a *big.Int representing it.
func le128ToBigInt(buf [16]byte) *big.Int {
	// Int.SetBytes() expects big-endian input, so reverse the bytes locally first
	rev := make([]byte, 16)
	for x := 0; x < 16; x++ {
		rev[x] = buf[16-x-1]
	}

	return new(big.Int).SetBytes(rev)
}
