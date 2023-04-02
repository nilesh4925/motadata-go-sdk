package sdkutil

import (
	"github.com/nilesh4925/motadata-go-sdk/logger"
	. "github.com/nilesh4925/motadata-go-sdk/motadatatypes"
	"github.com/nilesh4925/motadata-go-sdk/sdkconstant"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"unicode"
)

var loggerObj = logger.New("GO-Polling/SDK Util")

func GetLogLevel() MotadataUINT8 {

	var logLevel MotadataUINT8 = sdkconstant.LogLevelInfo

	configDetails := ReadYAMLFile(sdkconstant.CurrentDir + sdkconstant.PathSeparator + sdkconstant.ConfigFolder + sdkconstant.PathSeparator + sdkconstant.SystemConfigFile)

	configLogLevel := getConfigParam(configDetails, sdkconstant.ParamLogLevel)

	if len(configLogLevel) > 0 {

		logLevel = configLogLevel.ToUINT8()
	}

	return logLevel
}

func ReadYAMLFile(filePath string) MotadataMap {

	_, err := os.Stat(filePath)

	if err == nil {

		streamFile, err := ioutil.ReadFile(filePath)

		if err == nil {

			data := make(MotadataMap)

			err = yaml.Unmarshal(streamFile, &data)

			if err == nil {

				return data

			} else {

				loggerObj.Error(sdkconstant.BlankString, err, sdkconstant.ModuleSDK)

			}
		} else {

			loggerObj.Error(sdkconstant.BlankString, err, sdkconstant.ModuleSDK)

		}
	} else {

		loggerObj.Warning(sdkconstant.BlankString, MotadataString("File does not exist "+filePath), sdkconstant.ModuleSDK)

	}

	return nil
}

func getConfigParam(configDetails MotadataMap, paramName string) MotadataString {

	if configDetails != nil && configDetails.Contains(paramName) {

		return MotadataString(ToString(configDetails[paramName]))
	}
	return ""
}

func IsASCII(value string) bool {

	for _, char := range value {

		if char > unicode.MaxASCII {

			return false
		}
	}
	return true
}

func IsFileExists(filePath string) bool {

	_, err := os.Stat(filePath)

	if err == nil {

		return true
	}

	return false
}

func CreateFile(filePath string) bool {

	file, err := os.Create(filePath)

	if err == nil {

		loggerObj.Warning(sdkconstant.BlankString, MotadataString("Failed to create file"+filePath), sdkconstant.ModuleSDK)

		return true
	}

	defer func(file *os.File) {

		err := file.Close()

		if err != nil {

			loggerObj.Error(sdkconstant.BlankString, err, sdkconstant.ModuleSDK)
		}

	}(file)

	return false
}

func CreateYAMLFile(host MotadataString, fileName string, context MotadataMap) {

	yamlData, err := yaml.Marshal(&context)

	if err == nil {

		err := ioutil.WriteFile(fileName, yamlData, 0)

		if err != nil {

			loggerObj.Warning(host, MotadataString("Failed to create file"+fileName), sdkconstant.ModuleSDK)
		}

	} else {

		loggerObj.Warning(host, MotadataString("Failed to marshal file data"+fileName), sdkconstant.ModuleSDK)

	}

}

func GetPingRetryCount() MotadataUINT16 {

	var retryCount MotadataUINT16 = 3

	configDetails := ReadYAMLFile(sdkconstant.CurrentDir + sdkconstant.PathSeparator + sdkconstant.ConfigFolder + sdkconstant.PathSeparator + sdkconstant.SystemConfigFile)

	configRetryCount := getConfigParam(configDetails, sdkconstant.ParamPingRetryCount)

	if len(configRetryCount) > 0 {

		retryCount = configRetryCount.ToUINT16()
	}

	return retryCount
}

func GetPingTimeout() MotadataUINT16 {

	var timeout MotadataUINT16 = 10

	configDetails := ReadYAMLFile(sdkconstant.CurrentDir + sdkconstant.PathSeparator + sdkconstant.ConfigFolder + sdkconstant.PathSeparator + sdkconstant.SystemConfigFile)

	configTimeout := getConfigParam(configDetails, sdkconstant.ParamPingTimeout)

	if len(configTimeout) > 0 {

		timeout = configTimeout.ToUINT16()
	}

	return timeout
}

func GetPingPacketSize() MotadataUINT16 {

	var packetSize MotadataUINT16 = 56

	configDetails := ReadYAMLFile(sdkconstant.CurrentDir + sdkconstant.PathSeparator + sdkconstant.ConfigFolder + sdkconstant.PathSeparator + sdkconstant.SystemConfigFile)

	configPacketSize := getConfigParam(configDetails, sdkconstant.ParamPingPacketSize)

	if len(configPacketSize) > 0 {

		packetSize = configPacketSize.ToUINT16()
	}

	return packetSize
}
