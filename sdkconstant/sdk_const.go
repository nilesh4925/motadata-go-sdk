package sdkconstant

import "os"

const (
	PathSeparator = string(os.PathSeparator)

	LogLevelTypeTrace = "TRACE"

	LogLevelTypeDebug = "DEBUG"

	LogLevelTypeInfo = "INFO"

	LogLevelTypeWarning = "WARNING"

	LogLevelTypeError = "ERROR"

	LogLevelTypeFatal = "FATAL"

	LogLevelTrace = 0

	LogLevelDebug = 1

	LogLevelInfo = 2

	LogLevelWarning = 3

	LogLevelError = 4

	LogDirectory = "log"

	LogFile = "@@@-GO-POLLING-ENGINE-###.log"

	LogNCMFile = "@@@-GO-NCM-ENGINE-###.log"

	LogSDKFile = "@@@-GO-SDK-###.log"

	LogFileDateFormat = "02-January-2006"

	LogFileTimeFormat = "03:04:05 PM"

	TimeFormat = "2006-01-02 03:04:05"

	SpaceSeparator = " "

	BlankString = ""

	NewLineSeparator = "\n"

	ConfigFolder = "config"

	MotadataConfigFile = "motadata-conf.yml"

	SystemConfigFile = "system-conf.yml"

	ModulePolling = "polling"

	ModuleNCM = "ncm"

	ModuleSDK = "sdk"

	ParamLogLevel = "log-level"

	NCMContextDir = "ncm-context"

	Host = "host"

	Port = "port"

	PluginContext = "plugin-context"

	ResultSeparator = "_|@#|_"

	Contexts = "work-context"

	EventType = "event-type"

	EventTypeDiscover = "discover"

	EventTypeCollect = "collect"

	YmlFileExtension = ".yml"

	ErrorCodeTimedOut = "Timed out"

	PluginPingCheckRequired = "ping-check-required"

	PluginNMapPing = "nmap-ping"

	ErrorCode = "error-code"

	StatusCode = "Status"

	StatusCodeUnreachable = 4

	StatusCodeReachable = 1

	ParamPingRetryCount = "max-ping-check-retry-count"

	ParamPingTimeout = "max-ping-check-timeout"

	ParamProcessRequestTimeout = "process-request-timeout"

	ParamPingPacketSize = "ping-packet-size"
)

var (
	CurrentDir, _ = os.Getwd()
)
