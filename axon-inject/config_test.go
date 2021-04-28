package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go-common/config"
	"github.com/tombenke/axon-go-common/messenger"
	"testing"
	"time"
)

func TestReadConfigFromFile_NoFile(t *testing.T) {
	defaultConfig := Config{
		Node:        config.GetDefaultNode(),
		PrintConfig: false,
	}

	config, err := readConfigFromFile(defaultConfig, "./non-existing-config.yml")
	assert.NotNil(t, err)
	assert.Equal(t, "open ./non-existing-config.yml: no such file or directory", err.Error())
	assert.Equal(t, defaultConfig, config)
}

func TestReadConfigFromFile_Ok(t *testing.T) {
	defaultConfig := Config{
		Node:        config.GetDefaultNode(),
		PrintConfig: false,
	}
	expectedConfig := Config{
		Node: config.Node{
			Messenger: messenger.Config{
				Urls:      "localhost:4222",
				UserCreds: "",
			},
			Name:           actorName,
			Type:           "untyped",
			ConfigFileName: defaultConfigFileName,
			LogLevel:       "debug",
			LogFormat:      "json",
			Ports: config.Ports{
				Configure: config.Configure{
					Extend: false,
					Modify: true,
				},
				Inputs: config.Inputs{
					config.In{
						IO: config.IO{
							Name:           "inject",
							Type:           "base/Bytes",
							Representation: "text/plain",
							Channel:        "",
						},
						Default: "",
					},
				},
				Outputs: config.Outputs{
					config.Out{
						IO: config.IO{
							Name:           "output",
							Type:           "base/Bytes",
							Representation: "text/plain",
							Channel:        "axon-inject.output",
						},
					},
				},
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
		},
		PrintConfig: false,
		Precision:   "ns",
		Message:     "A plain text message",
		Repeat:      1,
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

func TestParseDelay(t *testing.T) {
	okTestCases := []struct {
		DelayStr string
		Expected time.Duration
	}{
		{DelayStr: "1s", Expected: 1 * time.Second},
		{DelayStr: "1ms", Expected: 1 * time.Millisecond},
		{DelayStr: "42ns", Expected: 42 * time.Nanosecond},
		{DelayStr: "424us", Expected: 424 * time.Microsecond},
	}

	for _, testCase := range okTestCases {
		actual, err := parseDelay(testCase.DelayStr)
		assert.Nil(t, err)
		assert.Equal(t, testCase.Expected, actual)
	}

	wrongTestCases := []string{
		"", "999", "1xs", "ns", "ms23",
	}
	for _, wrongTestCase := range wrongTestCases {
		_, err := parseDelay(wrongTestCase)
		fmt.Printf("err: %s\n", err)
		assert.NotNil(t, err)
	}
}
