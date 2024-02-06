package host

import "runtime"

func getHostOS() string {
	return runtime.GOOS
}
