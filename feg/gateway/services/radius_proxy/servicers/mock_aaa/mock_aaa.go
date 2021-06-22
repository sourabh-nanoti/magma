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

package mock_aaa

import (
	"context"
	"net"

	"github.com/golang/glog"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

// MockAaa is just a wrapper around radius PacketServer
type MockAaa struct {
	Addr     string
	Server   *radius.PacketServer
	l        net.PacketConn
	serveErr error
}

// NewStarted MockAaa
func NewStarted(ctx context.Context) (*MockAaa, error) {
	handler := radius.HandlerFunc(func(w radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)
		password := rfc2865.UserPassword_GetString(r.Packet)

		var code radius.Code
		if username == "tim" && password == "12345" {
			code = radius.CodeAccessAccept
		} else {
			code = radius.CodeAccessReject
		}
		glog.V(2).Infof("Writing %s to %s", code, r.RemoteAddr)
		w.Write(r.Response(code))
	})

	mAaa := New(handler)
	err := mAaa.Start(ctx)
	if err != nil {
		return nil, err
	}
	return mAaa, nil
}

// New MockAaa
func New(handler radius.Handler) *MockAaa {
	addr, err := net.ResolveUDPAddr("udp", "localhost:1812")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	return &MockAaa{
		Addr: conn.LocalAddr().String(),
		Server: &radius.PacketServer{
			Handler:      handler,
			SecretSource: radius.StaticSecretSource([]byte(`12345`)),
		},
		l: conn,
	}
}

// Start Start the MockAaa server
func (mAaa *MockAaa) Start(ctx context.Context) error {

	go func() {
		mAaa.serveErr = mAaa.Server.Serve(mAaa.l)
	}()

	return nil
}

// Close  Close the Mock AAA Server
func (mAaa *MockAaa) Close() error {
	return mAaa.l.Close()
}
