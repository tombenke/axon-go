package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/messenger"
	"testing"
)

func TestGetConfig(t *testing.T) {
	c := GetConfig()
	assert.Equal(t, c.LogLevel, defaultLogLevel)
	assert.Equal(t, c.LogFormat, defaultLogFormat)
	assert.Equal(t, c.Urls, messenger.DefaultNatsURL())
	assert.Equal(t, c.UserCreds, defaultNatsUserCreds)
}
