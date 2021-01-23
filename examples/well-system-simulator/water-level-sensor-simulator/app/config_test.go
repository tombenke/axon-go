package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/messenger"
	"testing"
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
			Name:           "well-water-upper-level-sensor-simulator",
			Type:           "water-level-sensor-simulator",
			ConfigFileName: "config.yml",
			LogLevel:       "debug",
			LogFormat:      "json",
			Ports: config.Ports{
				Configure: config.Configure{
					Extend: true,
					Modify: true,
				},
				Inputs: config.Inputs{
					config.In{
						IO: config.IO{
							Name:           "reference-water-level",
							Type:           "base/Float64",
							Representation: "application/json",
							Channel:        "",
						},
						Default: "{ \"Body\": { \"Data\": 0.75 } }",
					},
					config.In{
						IO: config.IO{
							Name:           "water-level",
							Type:           "base/Float64",
							Representation: "application/json",
							Channel:        "well-water-level",
						},
						Default: "{ \"Body\": { \"Data\": 0.0 } }",
					},
				},
				Outputs: config.Outputs{
					config.Out{
						IO: config.IO{
							Name:           "water-level-state",
							Type:           "base/Bool",
							Representation: "application/json",
							Channel:        "well-water-upper-level-state",
						},
					},
				},
			},
			Orchestration: config.Orchestration{
				Presence:        true,
				Synchronization: true,
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
	}

	config, err := readConfigFromFile(defaultConfig, "./config.yml")
	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestReadConfigFromFile_Partial(t *testing.T) {
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
			Name:           "well-water-upper-level-sensor-simulator",
			Type:           "untyped",
			ConfigFileName: "config.yml",
			LogLevel:       "info",
			LogFormat:      "text",
			Ports: config.Ports{
				Configure: config.Configure{
					Extend: true,
					Modify: true,
				},
				Inputs: config.Inputs{
					config.In{
						IO: config.IO{
							Name:           "reference-water-level",
							Type:           "base/Float64",
							Representation: "application/json",
							Channel:        "",
						},
						Default: "{ \"Body\": { \"Data\": 0.75 } }",
					},
					config.In{
						IO: config.IO{
							Name:           "water-level",
							Type:           "base/Float64",
							Representation: "application/json",
							Channel:        "well-water-level",
						},
						Default: "{ \"Body\": { \"Data\": 0.0 } }",
					},
				},
				Outputs: config.Outputs{
					config.Out{
						IO: config.IO{
							Name:           "water-level-state",
							Type:           "base/Bool",
							Representation: "application/json",
							Channel:        "well-water-upper-level-state",
						},
					},
				},
			},
			Orchestration: config.Orchestration{
				Presence:        true,
				Synchronization: true,
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
	}

	config, err := readConfigFromFile(defaultConfig, "./config-partial.yml")
	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestGetConfigFileName_Default(t *testing.T) {
	assert.Equal(t, "config.yml", getConfigFileName([]string{}))
}

func TestGetConfigFileName_ByArgs(t *testing.T) {
	expectedConfigFileName := "config-by-args.yml"
	var args []string
	args = append(args, "-config")
	args = append(args, expectedConfigFileName)

	assert.Equal(t, expectedConfigFileName, getConfigFileName(args))
}
