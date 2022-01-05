package config

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

// AppConfig application configuration instance
var AppConfig Config

// Config application configuration
type Config struct {
	HostIPAddress   string 			`json:"address"`
	HTTPPort        int				`json:"port"` 
	TLS             bool			`json:"tls"` 
	LogLevel        string			`json:"loglevel"`
}


var errEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string) (string, error) {
    v := os.Getenv(key)
    if v == "" {
        return v, errEnvVarEmpty
    }
    return v, nil
}

func getenvInt(key string) (int, error) {
    s, err := getenvStr(key)
    if err != nil {
        return 0, err
    }
    v, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    return v, nil
}

func getenvBool(key string) (bool, error) {
    s, err := getenvStr(key)
    if err != nil {
        return false, err
    }
    v, err := strconv.ParseBool(s)
    if err != nil {
        return false, err
    }
    return v, nil
}

func init() {
	var err error
	const (
		defaultHostIPAddress = "127.0.0.1"
		defaultHostPort = 8888
		defaultHostTLS = true
		defaultLogLevel = "trace"
		HostIPAddressKey="IPADDRESS"
		hostPortKey="PORT"
		hostTLSKey="TLS"
		logLevelKey="LOG_LEVEL"
	)
	AppConfig = Config{defaultHostIPAddress, defaultHostPort, defaultHostTLS, defaultLogLevel}
	ipAddress, err := getenvStr(HostIPAddressKey)
	if err == nil {
		AppConfig.HostIPAddress = ipAddress
	}

	port, err := getenvInt(hostPortKey)
	if err == nil {
		AppConfig.HTTPPort = port
	}

	isTLS, err := getenvBool(hostTLSKey)
	if err == nil {
		AppConfig.TLS = isTLS
	}

	logLevel, err := getenvStr(logLevelKey)
	if err == nil {
		AppConfig.LogLevel = logLevel
	}
}

// String return application configuration to string
func (c Config) String() string {
	s, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		return string(s)
	}
	return ""
}
