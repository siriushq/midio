//go:build linux
// +build linux

package sys

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// GetMaxThreads returns the maximum number of threads that the system can create.
func GetMaxThreads() (int, error) {
	sysMaxThreadsStr, err := ioutil.ReadFile("/proc/sys/kernel/threads-max")
	if err != nil {
		return 0, err
	}
	sysMaxThreads, err := strconv.Atoi(strings.TrimSpace(string(sysMaxThreadsStr)))
	if err != nil {
		return 0, err
	}
	return sysMaxThreads, nil
}
