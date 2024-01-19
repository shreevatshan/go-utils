package environment

import (
	"dataexporter/pkg/std"
	"dataexporter/pkg/std/communication"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	cloudTypeAWS       = "AWS"
	cloudTypeAZURE     = "AZURE"
	cloudTypeGCP       = "GCP"
	cloudTypeGeneric   = "CLOUD"
	urlAWSInstanceID   = "http://169.254.169.254/latest/meta-data/instance-id"
	urlAWSToken        = "http://169.254.169.254/latest/api/token"
	urlAWSMetaData     = "http://169.254.169.254/latest/meta-data/"
	urlAZUREInstanceId = "http://169.254.169.254/metadata/instance/compute/vmId?api-version=2017-08-01&format=text"
	urlGCPInstanceID   = "http://169.254.169.254/computeMetadata/v1/instance/id"
	urlGenericCloud    = "http://169.254.169.254/"
)

type cloudInfo struct {
	isCloudinstance   bool
	cloudinstanceType string
	cloudinstanceID   string
}

type Environment struct {
	installPath        string
	operatingSystem    string
	systemArchitecture string
	cloudInfo          cloudInfo
}

func isAWSInstance() (bool, string) {

	request1 := communication.HTTPRequest{
		RequestType: communication.RequestTypeGet,
		API:         urlAWSInstanceID,
		TimeOut:     5,
	}

	response1 := request1.Send()
	if response1.Err == nil {

		if len(response1.Body) > 0 {
			return true, string(response1.Body)
		}

		return true, std.EmptyString

	} else {

		request2 := communication.HTTPRequest{
			RequestType: communication.RequestTypePut,
			API:         urlAWSToken,
			Headers:     map[string]string{"X-aws-ec2-metadata-token-ttl-seconds": "21600"},
			TimeOut:     5,
		}
		response2 := request2.Send()

		if response2.Err == nil {

			request3 := communication.HTTPRequest{
				RequestType: communication.RequestTypeGet,
				API:         urlAWSMetaData,
				Headers:     map[string]string{"X-aws-ec2-metadata-token": string(response2.Body)},
				TimeOut:     5,
			}
			response3 := request3.Send()

			if response3.Err == nil {
				if len(response3.Body) > 0 {
					return true, string(response3.Body)
				}
				return true, std.EmptyString
			}
		}
	}

	return false, std.EmptyString
}

func isGCPInstance() (bool, string) {

	request := communication.HTTPRequest{
		RequestType: communication.RequestTypeGet,
		API:         urlGCPInstanceID,
		Headers:     map[string]string{"Metadata-Flavor": "Google"},
		TimeOut:     5,
	}

	response := request.Send()
	if response.Err != nil {
		return false, std.EmptyString
	}

	if len(response.Body) > 0 {
		return true, string(response.Body)
	}

	return true, std.EmptyString
}

func isAZUREInstance() (bool, string) {

	request := communication.HTTPRequest{
		RequestType: communication.RequestTypeGet,
		API:         urlAZUREInstanceId,
		Headers:     map[string]string{"Metadata": "true"},
		TimeOut:     5,
	}

	response := request.Send()
	if response.Err != nil {
		return false, std.EmptyString
	}

	if len(response.Body) > 0 {
		return true, string(response.Body)
	}

	return true, std.EmptyString
}

func isOtherCloudInstance() bool {

	request := communication.HTTPRequest{
		RequestType: communication.RequestTypeGet,
		API:         urlGenericCloud,
		TimeOut:     5,
	}

	response := request.Send()

	return response.Err == nil
}

func Init() *Environment {
	environmentDetails := &Environment{}

	return environmentDetails
}

func (environmentDetails *Environment) FetchDetails() {
	environmentDetails.setOperatingSystem()
	environmentDetails.setSystemArchitecture()
	environmentDetails.setInstallPath()
	environmentDetails.setCloudInfo()
}

func (environmentDetails *Environment) setOperatingSystem() {
	environmentDetails.operatingSystem = runtime.GOOS
}

func (environmentDetails *Environment) setSystemArchitecture() {
	environmentDetails.systemArchitecture = runtime.GOARCH
}

func (environmentDetails *Environment) setInstallPath() {
	var installPath = ""
	currentExecutable, err := os.Executable()
	if err == nil {
		installPath = filepath.Dir(currentExecutable)
	}
	environmentDetails.installPath = installPath
}

func (environmentDetails *Environment) setCloudInfo() {

	var wg sync.WaitGroup
	var isGCP, isAWS, isAZURE, isOtherCloud bool
	var idGCP, idAWS, idAZURE string

	wg.Add(1)
	go func() {
		isAWS, idAWS = isAWSInstance()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		isGCP, idGCP = isGCPInstance()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		isAZURE, idAZURE = isAZUREInstance()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		isOtherCloud = isOtherCloudInstance()
		wg.Done()
	}()

	wg.Wait()

	if isAWS || isGCP || isAZURE || isOtherCloud {

		environmentDetails.cloudInfo.isCloudinstance = true

		if isAWS {
			environmentDetails.cloudInfo.cloudinstanceType = cloudTypeAWS
			environmentDetails.cloudInfo.cloudinstanceID = idAWS
		} else if isGCP {
			environmentDetails.cloudInfo.cloudinstanceType = cloudTypeGCP
			environmentDetails.cloudInfo.cloudinstanceID = idGCP
		} else if isAZURE {
			environmentDetails.cloudInfo.cloudinstanceType = cloudTypeAZURE
			environmentDetails.cloudInfo.cloudinstanceID = idAZURE
		} else {
			environmentDetails.cloudInfo.cloudinstanceType = cloudTypeGeneric
		}

	}
}

func (environmentDetails *Environment) GetOperatingSystem() string {
	return environmentDetails.operatingSystem
}

func (environmentDetails *Environment) GetSystemArchitecture() string {
	return environmentDetails.systemArchitecture
}

func (environmentDetails *Environment) GetInstallPath() string {
	return environmentDetails.installPath
}

func (environmentDetails *Environment) GetIsCloudInstance() bool {
	return environmentDetails.cloudInfo.isCloudinstance
}

func (environmentDetails *Environment) GetCloudInstanceType() string {
	return environmentDetails.cloudInfo.cloudinstanceType
}

func (environmentDetails *Environment) GetCloudInstanceID() string {
	return environmentDetails.cloudInfo.cloudinstanceID
}
