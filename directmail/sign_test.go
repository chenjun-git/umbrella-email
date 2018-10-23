package directmail

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chenjun-git/umbrella-email/common"
)

const (
	configPath = "../conf/config.email.toml"
)

func init() {
	common.InitConfig(configPath)
}

func TestSign(t *testing.T) {
	assert := assert.New(t)

	signedStr := signSample()
	assert.Equal("llJfXJjBW3OacrVgxxsITgYaYm0=", signedStr)
}
