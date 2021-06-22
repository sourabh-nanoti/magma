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
	"context"
	"log"
	"magma/feg/cloud/go/protos"

	"github.com/golang/glog"
	"layeh.com/radius"
)

// RadiusProxy structure
type RadiusProxy struct {
	config *RadiusProxyConfig
	client radius.Client
}

// RadiusProxyConfig Radius proxy Configuration
type RadiusProxyConfig struct {
	ClientAddr   string
	ServerAddr   string
	ServerSecret string
}

// NewRadiusProxy creates an Radius proy
func NewRadiusProxy(config *RadiusProxyConfig) (*RadiusProxy, error) {
	/*gtpCli, err := gtp.NewRunningClient(
		context.Background(), config.ClientAddr,
		gtp.SGWControlPlaneIfType, config.GtpTimeout)
	if err != nil {
		return nil, fmt.Errorf("Error creating S8_Proxy: %s", err)
	} */
	return newRadiusProxyImp(config)
}

func newRadiusProxyImp(config *RadiusProxyConfig) (*RadiusProxy, error) {
	// TODO: validate config
	radiusproxy := &RadiusProxy{
		config: config,
		client: radius.Client{InsecureSkipVerify: true},
	}
	return radiusproxy, nil
}

// ProxyPacket : Proxies the packet to the AAA server
func (s *RadiusProxy) ProxyPacket(ctx context.Context, req *protos.AaaRequest) (*protos.AaaResponse, error) {
	packet := req.GetPacket()
	log.Printf("InProxyPacket")
	_packet, err := radius.Parse(packet, []byte(s.config.ServerSecret))
	if err != nil {
		glog.Errorf("Failed to parse Radius packet %s", err)
		return nil, err
	}
	response, err := s.client.Exchange(context.Background(), _packet, s.config.ServerAddr)
	if err != nil {
		glog.Errorf("Failed to exchange Radius packet with error %s ", err)
		return nil, err
	}
	_response, err := response.Encode()
	if err != nil {
		glog.Errorf("Failed to encode Radius packet with error %s ", err)
		return nil, err
	}
	aaaResponse := &protos.AaaResponse{
		Packet: _response,
	}

	glog.V(2).Info("Succssfully proxied the Radius Request-Response")
	return aaaResponse, nil
}
