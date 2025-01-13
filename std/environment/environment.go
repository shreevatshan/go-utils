package environment

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/shreevatshan/go-utils/std/communication"
)

const (
	cloudTypeAWS        = "AWS"
	cloudTypeAZURE      = "AZURE"
	cloudTypeGCP        = "GCP"
	cloudTypeGeneric    = "CLOUD"
	urlAWSInstanceID    = "http://169.254.169.254/latest/meta-data/instance-id"
	urlAWSToken         = "http://169.254.169.254/latest/api/token"
	urlAWSAutoScaling   = "http://169.254.169.254/latest/meta-data/autoscaling/target-lifecycle-state"
	urlAZUREInstanceId  = "http://169.254.169.254/metadata/instance/compute/vmId?api-version=2017-08-01&format=text"
	urlAZUREAutoScaling = "http://169.254.169.254/metadata/instance/compute/vmScaleSetName?api-version=2023-11-15&format=text"
	urlGCPInstanceID    = "http://169.254.169.254/computeMetadata/v1/instance/id"
	urlGenericCloud     = "http://169.254.169.254/"
)

type cloudInfo struct {
	isCloudinstance   bool
	cloudinstanceType string
	cloudinstanceID   string
	isAutoScaling     bool
}

type Environment struct {
	installPath        string
	operatingSystem    string
	systemArchitecture string
	cloudInfo          cloudInfo
}

func isAWSInstance() cloudInfo {

	cloudInfo := cloudInfo{cloudinstanceType: cloudTypeAWS, isAutoScaling: true}

	request1 := communication.HTTPRequest{
		RequestType: communication.RequestTypeGET,
		API:         urlAWSInstanceID,
		TimeOut:     5,
	}

	response1 := request1.Send()

	if response1.Err == nil {

		cloudInfo.isCloudinstance = true

		if len(response1.Body) > 0 {
			cloudInfo.cloudinstanceID = string(response1.Body)
		}

	} else {
		// verification for IMDSv2
		request2 := communication.HTTPRequest{
			RequestType: communication.RequestTypePUT,
			API:         urlAWSToken,
			Headers:     map[string]string{"X-aws-ec2-metadata-token-ttl-seconds": "21600"},
			TimeOut:     5,
		}

		response2 := request2.Send()

		if response2.Err == nil {

			request3 := communication.HTTPRequest{
				RequestType: communication.RequestTypeGET,
				API:         urlAWSInstanceID,
				Headers:     map[string]string{"X-aws-ec2-metadata-token": string(response2.Body)},
				TimeOut:     5,
			}
			response3 := request3.Send()

			if response3.Err == nil {

				cloudInfo.isCloudinstance = true

				if len(response3.Body) > 0 {
					cloudInfo.cloudinstanceID = string(response3.Body)
				}
			}
		}
	}

	if cloudInfo.isCloudinstance {
		// check for auto scaling
		request4 := communication.HTTPRequest{
			RequestType: communication.RequestTypeGET,
			API:         urlAWSAutoScaling,
			TimeOut:     5,
		}

		response4 := request4.Send()

		if response4.Err != nil {
			cloudInfo.isAutoScaling = false
		}
	}
	return cloudInfo
}

func isGCPInstance() cloudInfo {

	cloudInfo := cloudInfo{cloudinstanceType: cloudTypeGCP, isAutoScaling: true}

	request := communication.HTTPRequest{
		RequestType: communication.RequestTypeGET,
		API:         urlGCPInstanceID,
		Headers:     map[string]string{"Metadata-Flavor": "Google"},
		TimeOut:     5,
	}

	response := request.Send()

	if response.Err == nil {

		cloudInfo.isCloudinstance = true

		if len(response.Body) > 0 {
			cloudInfo.cloudinstanceID = string(response.Body)
		}
	}

	return cloudInfo
}

func isAZUREInstance() cloudInfo {

	cloudInfo := cloudInfo{cloudinstanceType: cloudTypeAZURE, isAutoScaling: true}

	request := communication.HTTPRequest{
		RequestType: communication.RequestTypeGET,
		API:         urlAZUREInstanceId,
		Headers:     map[string]string{"Metadata": "true"},
		TimeOut:     5,
	}

	response := request.Send()

	if response.Err == nil {

		cloudInfo.isCloudinstance = true

		if len(response.Body) > 0 {
			cloudInfo.cloudinstanceID = string(response.Body)
		}
	}

	if cloudInfo.isCloudinstance {
		// check for auto scaling
		request4 := communication.HTTPRequest{
			RequestType: communication.RequestTypeGET,
			API:         urlAZUREAutoScaling,
			Headers:     map[string]string{"Metadata": "true"},
			TimeOut:     5,
		}

		response4 := request4.Send()

		if response4.Err != nil || len(response4.Body) == 0 {
			cloudInfo.isAutoScaling = false
		}
	}

	return cloudInfo
}

func isOtherCloudInstance() cloudInfo {

	cloudInfo := cloudInfo{cloudinstanceType: cloudTypeGeneric, isAutoScaling: true}

	request := communication.HTTPRequest{
		RequestType: communication.RequestTypeGET,
		API:         urlGenericCloud,
		TimeOut:     5,
	}

	response := request.Send()

	if response.Err == nil {
		cloudInfo.isCloudinstance = true
	}

	return cloudInfo
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

	wg.Add(1)
	go func() {
		info := isAWSInstance()
		if info.isCloudinstance {
			environmentDetails.cloudInfo = info
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		info := isGCPInstance()
		if info.isCloudinstance {
			environmentDetails.cloudInfo = info
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		info := isAZUREInstance()
		if info.isCloudinstance {
			environmentDetails.cloudInfo = info
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		info := isOtherCloudInstance()
		if info.isCloudinstance {
			environmentDetails.cloudInfo = info
		}
		wg.Done()
	}()

	wg.Wait()

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

func (environmentDetails *Environment) GetIsAutoScaling() bool {
	return environmentDetails.cloudInfo.isAutoScaling
}
