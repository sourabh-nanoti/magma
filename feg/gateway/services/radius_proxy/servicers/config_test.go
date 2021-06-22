package servicers_test

import (
	"os"
	"testing"

	"magma/feg/gateway/services/radius_proxy/servicers"
	"magma/gateway/mconfig"

	"github.com/stretchr/testify/assert"
)

func TestGetRadiusProxyConfig(t *testing.T) {
	conf := generateRadiusMconfig(t, config)
	assert.Equal(t, ":2222", conf.ClientAddr)
	assert.Equal(t, "10.0.0.1:9999", conf.ServerAddr)
}

func TestGetRadiusProxyConfig_EnvVars(t *testing.T) {
	os.Setenv(servicers.ClientAddrEnv, ":4444")
	os.Setenv(servicers.ServerAddrEnv, "9.9.9.9:99")
	conf := generateRadiusMconfig(t, config)
	assert.Equal(t, ":4444", conf.ClientAddr)
	assert.Equal(t, "9.9.9.9:99", conf.ServerAddr)
}

func generateRadiusMconfig(t *testing.T, configString string) *servicers.RadiusProxyConfig {
	err := mconfig.CreateLoadTempConfig(configString)
	assert.NoError(t, err)
	return servicers.GetRadiusProxyConfig()
}

var (
	config = `{
		"configsByKey": {
			"radius_proxy": {
				"@type": "type.googleapis.com/magma.mconfig.RadiusProxyConfig",
				"logLevel": "INFO",
				"localAddress": ":2222",
				"aaaAddress": "10.0.0.1:9999"
			}
		}
	}`

	config_noAAA = `{
		"configsByKey": {
			"radius_proxy": {
				"@type": "type.googleapis.com/magma.mconfig.RadiusProxyConfig",
				"logLevel": "INFO",
				"localAddress": ":3333",
				"aaaAddress": ""
			}
		}
	}`
)
