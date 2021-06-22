package test_init

import (
	"context"
	"fmt"
	"testing"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/registry"
	"magma/feg/gateway/services/radius_proxy/servicers"
	"magma/feg/gateway/services/radius_proxy/servicers/mock_aaa"
	"magma/gateway/mconfig"
	"magma/orc8r/cloud/go/test_utils"
)

// StartRadiusProxyAndAaaService start both Radius proxy service and AAA for testing
func StartRadiusProxyAndAaaService(t *testing.T) (*mock_aaa.MockAaa, error) {
	// Start pgw and get the server address and real port
	mockAaa, err := mock_aaa.NewStarted(context.Background())
	if err != nil {
		return nil, err
	}

	// create config string with its proper values
	fegConfigFmt := `{
		"configsByKey": {
			"radius_proxy": {
				"@type": "type.googleapis.com/magma.mconfig.RadiusProxyConfig",
				"logLevel": "INFO",
				"local_address": "%s",
				"aaa_address": "%s"
			}
		}
	}`
	configStr := fmt.Sprintf(fegConfigFmt, "localhost", "localhost:1812")

	// load mconfig
	err = mconfig.CreateLoadTempConfig(configStr)
	if err != nil {
		return nil, err
	}
	config := servicers.GetRadiusProxyConfig()

	// create and launch Radius Proxy
	radiusproxyservice, err := servicers.NewRadiusProxy(config)
	if err != nil {
		return nil, err
	}
	srv, lis := test_utils.NewTestService(t, registry.ModuleName, registry.RADIUS_PROXY)
	protos.RegisterRadiusProxyServer(srv.GrpcServer, radiusproxyservice)
	go srv.RunTest(lis)
	return mockAaa, nil
}
