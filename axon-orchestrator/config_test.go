package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go-common/config"
	"github.com/tombenke/axon-go-common/messenger"
	"testing"
)

func TestReadConfigFromFile_NoFile(t *testing.T) {
	defaultConfig := Config{
		PrintConfig: false,
	}

	config, err := readConfigFromFile(defaultConfig, "./non-existing-config.yml")
	assert.NotNil(t, err)
	assert.Equal(t, "open ./non-existing-config.yml: no such file or directory", err.Error())
	assert.Equal(t, defaultConfig, config)
}

func TestReadConfigFromFile_Ok(t *testing.T) {
	defaultConfig := Config{
		PrintConfig: false,
	}
	expectedConfig := Config{
		Heartbeat:        defaultHeartbeat,
		MaxResponseTime:  defaultMaxResponseTime,
		EPNStatusChannel: defaultEPNStatusChannel,
		Name:             appName,
		ConfigFileName:   defaultConfigFileName,
		LogLevel:         "debug",
		LogFormat:        "text",
		PrintConfig:      false,
		Messenger: messenger.Config{
			Urls:      "localhost:4222",
			UserCreds: "",
		},
		Orchestration: config.Orchestration{
			Presence:        true,
			Synchronization: false,
			Channels: config.Channels{
				StatusRequest:       "status-request",
				StatusReport:        "status-report",
				SendResults:         "send-results",
				SendingCompleted:    "sending-completed",
				ReceiveAndProcess:   "receive-and-process",
				ProcessingCompleted: "processing-completed",
			},
		},
	}

	config, err := readConfigFromFile(defaultConfig, "./config.yml")
	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestGetConfigFileName_Default(t *testing.T) {
	dfltConfig := Config{}
	assert.Equal(t, defaultConfigFileName, getConfigFileName("", dfltConfig, []string{}))
}

func TestGetConfigFileName_ByArgs(t *testing.T) {
	expectedConfigFileName := "config-by-args.yml"
	var args []string
	args = append(args, "-config")
	args = append(args, expectedConfigFileName)

	assert.Equal(t, expectedConfigFileName, getConfigFileName("", Config{}, args))
}
