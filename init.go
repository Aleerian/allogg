package exloggo

import (
	"errors"
)

const (
	levelInfo          = "INFO"
	levelWarning       = "WARNING"
	levelError         = "ERROR"
	levelDebug         = "DEBUG"
	levelFatal         = "FATAL"
	levelRequestResult = "REQUEST_RESULT"
)

const (
	ReleaseMode     = "RELEASE"
	DevelopmentMode = "DEVELOPMENT"
)

const (
	defaultServerVersion   = "dev"
	defaultOutputDirectory = "/logs"
	defaultMode            = DevelopmentMode
)

// Цветовые коды
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

type Parameters struct {
	Mode          string
	ServerVersion string
	Directory     string
}

var loggerMode string
var prefixPath string
var serverVersion string
var logsDirectoryPath string

func setMode(mode string) error {
	if mode != "" {
		if mode != DevelopmentMode && mode != ReleaseMode {
			return errors.New(errorInvalidMode)
		}
	} else {
		loggerMode = DevelopmentMode
		return nil
	}
	loggerMode = mode
	return nil
}

// Возвращает цвет для уровня логирования
func getColor(level string) string {
	switch level {
	case "INFO":
		return Green
	case "ERROR":
		return Red
	case "WARN":
		return Yellow
	default:
		return Reset
	}
}

func setServerVersion(version string) {
	if version == "" {
		serverVersion = defaultServerVersion
		return
	}
	serverVersion = version
}

func setOutputDirectory(path string) {
	if path == "" {
		logsDirectoryPath = defaultOutputDirectory
		return
	}
	logsDirectoryPath = path
}
