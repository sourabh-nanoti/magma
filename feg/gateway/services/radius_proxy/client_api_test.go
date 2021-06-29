package radius_proxy_test

import (
	"log"
	"testing"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/services/radius_proxy"
	"magma/feg/gateway/services/radius_proxy/test_init"

	"github.com/stretchr/testify/assert"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	ServerSecret = "123456"
)

func TestRadiusProxyClient(t *testing.T) {
	// run both s8 and pgw
	mAaa, err := test_init.StartRadiusProxyAndAaaService(t)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer mAaa.Close()
	log.Printf("Started the server")

	//------------------------
	//---- Create Session ----
	reqPacket := radius.New(radius.CodeAccessRequest, []byte(`123456`))
	rfc2865.UserName_SetString(reqPacket, "tim")
	rfc2865.UserPassword_SetString(reqPacket, "123456")

	_reqPacket, err := reqPacket.Encode()
	if err != nil {
		t.Fatalf("Cannot Create Radius Packet %v ", err)
		return
	}

	Req := &protos.AaaRequest{
		Packet: _reqPacket,
	}

	Rsp, err := radius_proxy.ProxyPacket(Req)
	if err != nil {
		t.Fatalf("Radius proxy ProxyPacket Error: %v", err)
		return
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, Rsp)

	_rspPacket := Rsp.GetPacket()

	rspPacket, err := radius.Parse(_rspPacket, []byte(ServerSecret))
	if err != nil {
		t.Fatalf("Radius proxy unable to parse packet: %v", err)
	}

	// check fteid was received properly
	assert.Equal(t, radius.CodeAccessAccept, rspPacket.Code)
}

func TestRadiusProxyClientReject(t *testing.T) {
	// run both s8 and pgw
	mAaa, err := test_init.StartRadiusProxyAndAaaService(t)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer mAaa.Close()
	log.Printf("Started the server")

	//------------------------
	//---- Create Session ----
	reqPacket := radius.New(radius.CodeAccessRequest, []byte(`12345`))
	rfc2865.UserName_SetString(reqPacket, "tim")
	rfc2865.UserPassword_SetString(reqPacket, "123456789")

	_reqPacket, err := reqPacket.Encode()
	if err != nil {
		t.Fatalf("Cannot Create Radius Packet %v ", err)
		return
	}

	Req := &protos.AaaRequest{
		Packet: _reqPacket,
	}

	Rsp, err := radius_proxy.ProxyPacket(Req)
	if err != nil {
		t.Fatalf("Radius proxy ProxyPacket Error: %v", err)
		return
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, Rsp)

	_rspPacket := Rsp.GetPacket()

	rspPacket, err := radius.Parse(_rspPacket, []byte(ServerSecret))
	if err != nil {
		t.Fatalf("Radius proxy unable to parse packet: %v", err)
	}

	// check fteid was received properly
	assert.Equal(t, radius.CodeAccessReject, rspPacket.Code)
}
