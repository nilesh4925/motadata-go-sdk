package logger

import (
	. "github.com/nilesh4925/motadata-go-sdk/motadatatypes"
	"github.com/nilesh4925/motadata-go-sdk/sdkconstant"
	"github.com/pkg/errors"
	"os"
)

var logLevel = sdkconstant.LogLevelInfo

type Logger struct {
	component MotadataString
}

func traceEnabled() bool {

	return sdkconstant.LogLevelTrace >= logLevel
}

func debugEnabled() bool {

	return sdkconstant.LogLevelDebug >= logLevel
}

func infoEnabled() bool {

	return sdkconstant.LogLevelInfo >= logLevel
}

func warningEnabled() bool {

	return sdkconstant.LogLevelWarning >= logLevel
}

func errorEnabled() bool {

	return sdkconstant.LogLevelError >= logLevel
}

func SetLogLevel(value MotadataUINT8) {

	logLevel = int(value)
}

func New(component MotadataString) Logger {
	return Logger{component: component}
}

func (logger *Logger) Trace(host MotadataString, message MotadataString, module MotadataString) {

	if traceEnabled() {

		currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

		currentTime := MotadataTimeString(sdkconstant.LogFileTimeFormat).Format()

		message := currentDate + sdkconstant.SpaceSeparator + currentTime + sdkconstant.SpaceSeparator + "TRACE [" + logger.component + "]:" +
			"[ " + host + " ]:" + message

		logger.write(message.ToString(), sdkconstant.LogLevelTypeTrace, module)
	}
}

func (logger *Logger) Debug(host MotadataString, message MotadataString, module MotadataString) {

	if debugEnabled() {

		currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

		currentTime := MotadataTimeString(sdkconstant.LogFileTimeFormat).Format()

		message := currentDate + sdkconstant.SpaceSeparator + currentTime + sdkconstant.SpaceSeparator + "Debug [" + logger.component + "]:" +
			"[ " + host + " ]:" + message

		logger.write(message.ToString(), sdkconstant.LogLevelTypeDebug, module)
	}
}

func (logger *Logger) Info(host MotadataString, message MotadataString, module MotadataString) {

	if infoEnabled() {

		currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

		currentTime := MotadataTimeString(sdkconstant.LogFileTimeFormat).Format()

		message := currentDate + sdkconstant.SpaceSeparator + currentTime + sdkconstant.SpaceSeparator + "Info [" + logger.component + "]:" +
			"[ " + host + " ]:" + message

		logger.write(message.ToString(), sdkconstant.LogLevelTypeInfo, module)
	}
}

func (logger *Logger) Warning(host MotadataString, message MotadataString, module MotadataString) {

	if warningEnabled() {

		currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

		currentTime := MotadataTimeString(sdkconstant.LogFileTimeFormat).Format()

		message := currentDate + sdkconstant.SpaceSeparator + currentTime + sdkconstant.SpaceSeparator + "Warning [" + logger.component + "]:" +
			"[ " + host + " ]:" + message

		logger.write(message.ToString(), sdkconstant.LogLevelTypeWarning, module)
	}
}

func (logger *Logger) Error(host MotadataString, err error, module MotadataString) {

	if errorEnabled() {

		currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

		currentTime := MotadataTimeString(sdkconstant.LogFileTimeFormat).Format()

		err = errors.Errorf("%+v\n", errors.New(err.Error()))

		message := currentDate + sdkconstant.SpaceSeparator + currentTime + sdkconstant.SpaceSeparator + "Error [" + logger.component + "]: " +
			"[ " + host + " ]:" + MotadataString(err.Error())

		logger.write(message.ToString(), sdkconstant.LogLevelTypeError, module)
	}
}

func (logger *Logger) Fatal(host MotadataString, message MotadataString, module MotadataString) {

	currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

	currentTime := MotadataTimeString(sdkconstant.LogFileTimeFormat).Format()

	message = currentDate + sdkconstant.SpaceSeparator + currentTime + sdkconstant.SpaceSeparator + "Fatal [" + logger.component + "]:" +
		"[ " + host + " ]:" + message

	logger.write(message.ToString(), sdkconstant.LogLevelTypeFatal, module)
}

func (logger *Logger) write(message string, logType MotadataString, module MotadataString) {

	logDir := sdkconstant.CurrentDir + sdkconstant.PathSeparator + sdkconstant.LogDirectory + sdkconstant.PathSeparator

	_, err := os.Stat(logDir)

	if os.IsNotExist(err) {

		err := os.MkdirAll(logDir, 0755)

		if err != nil {

			panic(err)
		}

	}

	logFile := logDir + sdkconstant.PathSeparator

	currentDate := MotadataTimeString(sdkconstant.LogFileDateFormat).Format()

	logFile = logFile + MotadataString(getModuleFile(module)).ReplaceAll("@@@", currentDate).ReplaceAll("###", logType).ToString()

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {

		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	if _, err = file.WriteString(message + sdkconstant.NewLineSeparator); err != nil {

		panic(err)
	}
}

func getModuleFile(module MotadataString) MotadataString {

	file := sdkconstant.LogFile

	switch module {

	case sdkconstant.ModuleNCM:

		file = sdkconstant.LogNCMFile

		break

	case sdkconstant.ModuleSDK:

		file = sdkconstant.LogSDKFile

		break

	default:
		break
	}

	return MotadataString(file)
}
