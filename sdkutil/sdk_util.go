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
