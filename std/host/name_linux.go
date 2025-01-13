package host

import (
	"os"
	"syscall"
)

func getHostName() string {

	var err error
	var name string

	name, err = os.Hostname()
	if err != nil {
		var uts syscall.Utsname
		err = syscall.Uname(&uts)
		if err == nil {
			buf := make([]byte, 0, len(uts.Nodename))
			for _, c := range uts.Nodename {
				if c == 0 {
					break
				}
				buf = append(buf, byte(c))
			}
			name = string(buf)
		}
	}

	return name
}
