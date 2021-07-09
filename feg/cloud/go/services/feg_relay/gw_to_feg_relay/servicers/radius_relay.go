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

	"magma/feg/cloud/go/protos"
)

// RadiusProxyServer implementation
//
// RadiusProxy sends radius aaa request, waits (blocks)
// for response then returns its RPC representation
func (s *RelayRouter) ProxyPacket(
	c context.Context, r *protos.AaaRequest) (*protos.AaaResponse, error) {

	client, ctx, cancel, err := s.getRadiusClient(c, r.GetSubscriberId())
	if err != nil {
		return nil, err
	}
	defer cancel()
	ret, err := client.ProxyPacket(ctx, r)
	return ret, err
}

func (s *RelayRouter) getRadiusClient(
	c context.Context, imsi string) (protos.RadiusProxyClient, context.Context, context.CancelFunc, error) {

	conn, ctx, cancel, err := s.GetFegServiceConnection(c, imsi, FegRadiusProxy)
	if err != nil {
		return nil, nil, nil, err
	}
	return protos.NewRadiusProxyClient(conn), ctx, cancel, nil
}
