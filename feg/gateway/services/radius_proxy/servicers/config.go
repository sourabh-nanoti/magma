/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package servicers

import (
	"log"
	"os"

	mcfgprotos "magma/feg/cloud/go/protos/mconfig"
	utils "magma/feg/gateway/diameter"
	"magma/gateway/mconfig"

	"github.com/golang/glog"
)

// Constants
const (
	RadiusProxyServiceName = "radius_proxy"
	ClientAddrEnv          = "RADIUS_CLIENT_ADDRESS"
	ServerAddrEnv          = "RADIUS_SERVER_ADDRESS"
	ServerSecretEnv        = "RADIUS_SERVER_SECRET"
)

// GetRadiusProxyConfig Gets the Radius Proxy Config
func GetRadiusProxyConfig() *RadiusProxyConfig {
	configPtr := &mcfgprotos.RadiusProxyConfig{}
	conf := &RadiusProxyConfig{}
	err := mconfig.GetServiceConfigs(RadiusProxyServiceName, configPtr)
	if err != nil {
		glog.V(2).Infof("%s Managed Configs Load Error: %v Using EnvVars", RadiusProxyServiceName, err)
		conf.ClientAddr = os.Getenv(ClientAddrEnv)
		conf.ServerAddr = os.Getenv(ServerAddrEnv)
		conf.ServerSecret = os.Getenv(ServerSecretEnv)

	} else {
		log.Printf("Configuration")
		conf.ClientAddr = utils.GetValueOrEnv("", ClientAddrEnv, configPtr.LocalAddress)
		conf.ServerAddr = utils.GetValueOrEnv("", ServerAddrEnv, configPtr.AaaAddress)
		conf.ServerSecret = utils.GetValueOrEnv("", ServerSecretEnv, configPtr.ServerSecret)
	}
	glog.V(2).Infof("Loaded configs: %+v", conf)
	return conf
}
