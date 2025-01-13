package host

const (
	// HostTypeDocker specifies the type of host as Docker.
	HostTypeDocker = "docker"
)

// Host represents the host information.
type Host struct {
	name   string
	os     string
	fqdn   string
	docker Docker
}

var (
	host Host
)

// Name returns the host name.
func Name() string {
	return host.name
}

// OS returns the host OS.
func OS() string {
	return host.os
}

// FQDN returns the host FQDN.
func FQDN() string {
	return host.fqdn
}

// ContainerID returns the container ID if the host is a Docker.
func ContainerID() string {
	return host.docker.containerID
}

// IsDocker returns true if the host is a Docker.
func IsDocker() bool {
	return host.docker.isDocker
}

func init() {
	host.name = getHostName()
	host.os = getHostOS()
	host.fqdn = getHostFQDN()
	host.docker = getDockerInfo()
}
