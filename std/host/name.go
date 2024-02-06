//go:build !linux
// +build !linux

package host

import (
	"os"
)

func getHostName() string {

	name, _ := os.Hostname()

	return name
}
