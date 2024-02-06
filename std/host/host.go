package host

const (
	HostTypeDocker = "docker"
)

type Docker struct {
	isDocker    bool
	containerID string
}

type Host struct {
	name   string
	os     string
	fqdn   string
	docker Docker
}

var (
	host Host
)

func Name() string {
	return host.name
}

func OS() string {
	return host.os
}

func FQDN() string {
	return host.fqdn
}

func ContainerID() string {
	return host.docker.containerID
}

func IsDocker() bool {
	return host.docker.isDocker
}

func init() {
	host.name = getHostName()
	host.os = getHostOS()
	host.fqdn = getHostFQDN()
	host.docker.isDocker, host.docker.containerID = getDockerInfo()
}
