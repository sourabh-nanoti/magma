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

package radius_proxy

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/registry"
	"magma/orc8r/lib/go/util"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type radiusProxyClient struct {
	protos.RadiusProxyClient
}

func getRadiusProxyClient() (*radiusProxyClient, error) {
	var (
		conn *grpc.ClientConn
		err  error
	)
	if util.GetEnvBool("USE_REMOTE_RADIUS_PROXY") {
		conn, err = registry.Get().GetSharedCloudConnection(strings.ToLower(registry.RADIUS_PROXY))
	} else {
		conn, err = registry.GetConnection(registry.RADIUS_PROXY)
	}
	if err != nil {
		errMsg := fmt.Sprintf("Radius Proxy client initialization error: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return &radiusProxyClient{protos.NewRadiusProxyClient(conn)}, nil
}

func ProxyPacket(req *protos.AaaRequest) (*protos.AaaResponse, error) {
	if req == nil {
		return nil, errors.New("Invalid Request")
	}
	cli, err := getRadiusProxyClient()
	if err != nil {
		return nil, err
	}
	return cli.ProxyPacket(context.Background(), req)
}
