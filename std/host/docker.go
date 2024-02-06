package host

import (
	"bufio"
	"os"
	"strings"
)

func dockerCheckV1() (bool, string) {
	var containerID string
	var isDocker bool

	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return isDocker, containerID
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && (containerID == "") {
		line := scanner.Text()
		tokens := strings.Split(line, "/")
		for i := 0; i < len(tokens) && (containerID == ""); i++ {
			if tokens[i] == "docker" {
				if i+1 < len(tokens) {
					containerID = strings.TrimSuffix(tokens[i+1], "\n")
				}
				isDocker = true
			}
		}
	}
	return isDocker, containerID
}

func dockerCheckV2() (bool, string) {
	var containerID string
	var isDocker bool

	file, err := os.Open("/proc/1/mountinfo")
	if err != nil {
		return isDocker, containerID
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		containerIDSubstrPrefix := strings.Index(line, "/var/lib/docker/containers/")
		containerIDSubstrSuffix := strings.Index(line, "/hostname")
		if containerIDSubstrPrefix != -1 && containerIDSubstrSuffix != -1 {
			containerIDSubstrPrefix += len("/var/lib/docker/containers/")
			containerID = line[containerIDSubstrPrefix:containerIDSubstrSuffix]
			isDocker = true
		}
	}
	return isDocker, containerID
}

func getDockerInfo() (bool, string) {

	var isDocker bool
	var containerID string

	isDocker, containerID = dockerCheckV1()
	if !isDocker {
		isDocker, containerID = dockerCheckV2()
	}

	return isDocker, containerID
}
