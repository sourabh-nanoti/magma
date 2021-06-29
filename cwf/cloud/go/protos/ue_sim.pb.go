// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cwf/protos/ue_sim.proto

package protos // import "magma/cwf/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import protos "magma/orc8r/lib/go/protos"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthenticateRequestHssLess struct {
	// MSISDN
	Msisdn string `protobuf:"bytes,1,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	// APN
	Apn string `protobuf:"bytes,2,opt,name=apn,proto3" json:"apn,omitempty"`
	// RAT
	Rat                  uint32   `protobuf:"varint,3,opt,name=rat,proto3" json:"rat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequestHssLess) Reset()         { *m = AuthenticateRequestHssLess{} }
func (m *AuthenticateRequestHssLess) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequestHssLess) ProtoMessage()    {}
func (*AuthenticateRequestHssLess) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{0}
}
func (m *AuthenticateRequestHssLess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequestHssLess.Unmarshal(m, b)
}
func (m *AuthenticateRequestHssLess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequestHssLess.Marshal(b, m, deterministic)
}
func (dst *AuthenticateRequestHssLess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequestHssLess.Merge(dst, src)
}
func (m *AuthenticateRequestHssLess) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequestHssLess.Size(m)
}
func (m *AuthenticateRequestHssLess) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequestHssLess.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequestHssLess proto.InternalMessageInfo

func (m *AuthenticateRequestHssLess) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *AuthenticateRequestHssLess) GetApn() string {
	if m != nil {
		return m.Apn
	}
	return ""
}

func (m *AuthenticateRequestHssLess) GetRat() uint32 {
	if m != nil {
		return m.Rat
	}
	return 0
}

type UEConfig struct {
	// Unique identifier for the UE.
	Imsi string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// Authentication key (k).
	AuthKey []byte `protobuf:"bytes,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	// Operator configuration field (Op) signed with authentication key (k).
	AuthOpc []byte `protobuf:"bytes,3,opt,name=auth_opc,json=authOpc,proto3" json:"auth_opc,omitempty"`
	// Sequence Number (SEQ).
	Seq uint64 `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	// HSSLess Configuration
	HsslessCfg           *AuthenticateRequestHssLess `protobuf:"bytes,5,opt,name=hssless_cfg,json=hsslessCfg,proto3" json:"hssless_cfg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UEConfig) Reset()         { *m = UEConfig{} }
func (m *UEConfig) String() string { return proto.CompactTextString(m) }
func (*UEConfig) ProtoMessage()    {}
func (*UEConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{1}
}
func (m *UEConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UEConfig.Unmarshal(m, b)
}
func (m *UEConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UEConfig.Marshal(b, m, deterministic)
}
func (dst *UEConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UEConfig.Merge(dst, src)
}
func (m *UEConfig) XXX_Size() int {
	return xxx_messageInfo_UEConfig.Size(m)
}
func (m *UEConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UEConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UEConfig proto.InternalMessageInfo

func (m *UEConfig) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *UEConfig) GetAuthKey() []byte {
	if m != nil {
		return m.AuthKey
	}
	return nil
}

func (m *UEConfig) GetAuthOpc() []byte {
	if m != nil {
		return m.AuthOpc
	}
	return nil
}

func (m *UEConfig) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *UEConfig) GetHsslessCfg() *AuthenticateRequestHssLess {
	if m != nil {
		return m.HsslessCfg
	}
	return nil
}

type AuthenticateRequest struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	CalledStationID      string   `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{2}
}
func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(dst, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *AuthenticateRequest) GetCalledStationID() string {
	if m != nil {
		return m.CalledStationID
	}
	return ""
}

type AuthenticateResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	SessionId            string   `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{3}
}
func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (dst *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(dst, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

func (m *AuthenticateResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type DisconnectRequest struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	CalledStationID      string   `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{4}
}
func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (dst *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(dst, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *DisconnectRequest) GetCalledStationID() string {
	if m != nil {
		return m.CalledStationID
	}
	return ""
}

type DisconnectResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{5}
}
func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (dst *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(dst, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

func (m *DisconnectResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

type ProxyRadiusRequest struct {
	UeMac                string   `protobuf:"bytes,1,opt,name=ue_mac,json=ueMac,proto3" json:"ue_mac,omitempty"`
	CalledStationID      string   `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyRadiusRequest) Reset()         { *m = ProxyRadiusRequest{} }
func (m *ProxyRadiusRequest) String() string { return proto.CompactTextString(m) }
func (*ProxyRadiusRequest) ProtoMessage()    {}
func (*ProxyRadiusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{6}
}
func (m *ProxyRadiusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyRadiusRequest.Unmarshal(m, b)
}
func (m *ProxyRadiusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyRadiusRequest.Marshal(b, m, deterministic)
}
func (dst *ProxyRadiusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyRadiusRequest.Merge(dst, src)
}
func (m *ProxyRadiusRequest) XXX_Size() int {
	return xxx_messageInfo_ProxyRadiusRequest.Size(m)
}
func (m *ProxyRadiusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyRadiusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyRadiusRequest proto.InternalMessageInfo

func (m *ProxyRadiusRequest) GetUeMac() string {
	if m != nil {
		return m.UeMac
	}
	return ""
}

func (m *ProxyRadiusRequest) GetCalledStationID() string {
	if m != nil {
		return m.CalledStationID
	}
	return ""
}

type ProxyRadiusResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyRadiusResponse) Reset()         { *m = ProxyRadiusResponse{} }
func (m *ProxyRadiusResponse) String() string { return proto.CompactTextString(m) }
func (*ProxyRadiusResponse) ProtoMessage()    {}
func (*ProxyRadiusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{7}
}
func (m *ProxyRadiusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyRadiusResponse.Unmarshal(m, b)
}
func (m *ProxyRadiusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyRadiusResponse.Marshal(b, m, deterministic)
}
func (dst *ProxyRadiusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyRadiusResponse.Merge(dst, src)
}
func (m *ProxyRadiusResponse) XXX_Size() int {
	return xxx_messageInfo_ProxyRadiusResponse.Size(m)
}
func (m *ProxyRadiusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyRadiusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyRadiusResponse proto.InternalMessageInfo

func (m *ProxyRadiusResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

type GenTrafficRequest struct {
	Imsi                    string                `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Volume                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Bitrate                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	TimeInSecs              uint64                `protobuf:"varint,4,opt,name=timeInSecs,proto3" json:"timeInSecs,omitempty"`
	ReportingIntervalInSecs uint64                `protobuf:"varint,5,opt,name=reportingIntervalInSecs,proto3" json:"reportingIntervalInSecs,omitempty"`
	ReverseMode             bool                  `protobuf:"varint,6,opt,name=reverseMode,proto3" json:"reverseMode,omitempty"`
	// Configure a max timeout iperf client will run
	Timeout uint32 `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Enables/Disable a function to check if server is alive before sending any traffic
	DisableServerReachabilityCheck bool     `protobuf:"varint,8,opt,name=disableServerReachabilityCheck,proto3" json:"disableServerReachabilityCheck,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *GenTrafficRequest) Reset()         { *m = GenTrafficRequest{} }
func (m *GenTrafficRequest) String() string { return proto.CompactTextString(m) }
func (*GenTrafficRequest) ProtoMessage()    {}
func (*GenTrafficRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{8}
}
func (m *GenTrafficRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTrafficRequest.Unmarshal(m, b)
}
func (m *GenTrafficRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTrafficRequest.Marshal(b, m, deterministic)
}
func (dst *GenTrafficRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTrafficRequest.Merge(dst, src)
}
func (m *GenTrafficRequest) XXX_Size() int {
	return xxx_messageInfo_GenTrafficRequest.Size(m)
}
func (m *GenTrafficRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTrafficRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenTrafficRequest proto.InternalMessageInfo

func (m *GenTrafficRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *GenTrafficRequest) GetVolume() *wrappers.StringValue {
	if m != nil {
		return m.Volume
	}
	return nil
}

func (m *GenTrafficRequest) GetBitrate() *wrappers.StringValue {
	if m != nil {
		return m.Bitrate
	}
	return nil
}

func (m *GenTrafficRequest) GetTimeInSecs() uint64 {
	if m != nil {
		return m.TimeInSecs
	}
	return 0
}

func (m *GenTrafficRequest) GetReportingIntervalInSecs() uint64 {
	if m != nil {
		return m.ReportingIntervalInSecs
	}
	return 0
}

func (m *GenTrafficRequest) GetReverseMode() bool {
	if m != nil {
		return m.ReverseMode
	}
	return false
}

func (m *GenTrafficRequest) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *GenTrafficRequest) GetDisableServerReachabilityCheck() bool {
	if m != nil {
		return m.DisableServerReachabilityCheck
	}
	return false
}

type GenTrafficResponse struct {
	Output               []byte         `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	EndOutput            *TrafficOutput `protobuf:"bytes,2,opt,name=end_output,json=endOutput,proto3" json:"end_output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GenTrafficResponse) Reset()         { *m = GenTrafficResponse{} }
func (m *GenTrafficResponse) String() string { return proto.CompactTextString(m) }
func (*GenTrafficResponse) ProtoMessage()    {}
func (*GenTrafficResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{9}
}
func (m *GenTrafficResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTrafficResponse.Unmarshal(m, b)
}
func (m *GenTrafficResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTrafficResponse.Marshal(b, m, deterministic)
}
func (dst *GenTrafficResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTrafficResponse.Merge(dst, src)
}
func (m *GenTrafficResponse) XXX_Size() int {
	return xxx_messageInfo_GenTrafficResponse.Size(m)
}
func (m *GenTrafficResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTrafficResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenTrafficResponse proto.InternalMessageInfo

func (m *GenTrafficResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *GenTrafficResponse) GetEndOutput() *TrafficOutput {
	if m != nil {
		return m.EndOutput
	}
	return nil
}

type TrafficOutput struct {
	SumSent              *TrafficSummary `protobuf:"bytes,1,opt,name=sum_sent,json=sumSent,proto3" json:"sum_sent,omitempty"`
	SumReceived          *TrafficSummary `protobuf:"bytes,2,opt,name=sum_received,json=sumReceived,proto3" json:"sum_received,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TrafficOutput) Reset()         { *m = TrafficOutput{} }
func (m *TrafficOutput) String() string { return proto.CompactTextString(m) }
func (*TrafficOutput) ProtoMessage()    {}
func (*TrafficOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{10}
}
func (m *TrafficOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficOutput.Unmarshal(m, b)
}
func (m *TrafficOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficOutput.Marshal(b, m, deterministic)
}
func (dst *TrafficOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficOutput.Merge(dst, src)
}
func (m *TrafficOutput) XXX_Size() int {
	return xxx_messageInfo_TrafficOutput.Size(m)
}
func (m *TrafficOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficOutput proto.InternalMessageInfo

func (m *TrafficOutput) GetSumSent() *TrafficSummary {
	if m != nil {
		return m.SumSent
	}
	return nil
}

func (m *TrafficOutput) GetSumReceived() *TrafficSummary {
	if m != nil {
		return m.SumReceived
	}
	return nil
}

type TrafficSummary struct {
	Start                float64  `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	Seconds              float64  `protobuf:"fixed64,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Bytes                int32    `protobuf:"varint,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	BitsPerSecond        float64  `protobuf:"fixed64,5,opt,name=bits_per_second,json=bitsPerSecond,proto3" json:"bits_per_second,omitempty"`
	Retransmits          int32    `protobuf:"varint,6,opt,name=retransmits,proto3" json:"retransmits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficSummary) Reset()         { *m = TrafficSummary{} }
func (m *TrafficSummary) String() string { return proto.CompactTextString(m) }
func (*TrafficSummary) ProtoMessage()    {}
func (*TrafficSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_ue_sim_71973c8f5128f810, []int{11}
}
func (m *TrafficSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSummary.Unmarshal(m, b)
}
func (m *TrafficSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSummary.Marshal(b, m, deterministic)
}
func (dst *TrafficSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSummary.Merge(dst, src)
}
func (m *TrafficSummary) XXX_Size() int {
	return xxx_messageInfo_TrafficSummary.Size(m)
}
func (m *TrafficSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSummary.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSummary proto.InternalMessageInfo

func (m *TrafficSummary) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TrafficSummary) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *TrafficSummary) GetSeconds() float64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TrafficSummary) GetBytes() int32 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *TrafficSummary) GetBitsPerSecond() float64 {
	if m != nil {
		return m.BitsPerSecond
	}
	return 0
}

func (m *TrafficSummary) GetRetransmits() int32 {
	if m != nil {
		return m.Retransmits
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthenticateRequestHssLess)(nil), "magma.cwf.AuthenticateRequestHssLess")
	proto.RegisterType((*UEConfig)(nil), "magma.cwf.UEConfig")
	proto.RegisterType((*AuthenticateRequest)(nil), "magma.cwf.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "magma.cwf.AuthenticateResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "magma.cwf.DisconnectRequest")
	proto.RegisterType((*DisconnectResponse)(nil), "magma.cwf.DisconnectResponse")
	proto.RegisterType((*ProxyRadiusRequest)(nil), "magma.cwf.ProxyRadiusRequest")
	proto.RegisterType((*ProxyRadiusResponse)(nil), "magma.cwf.ProxyRadiusResponse")
	proto.RegisterType((*GenTrafficRequest)(nil), "magma.cwf.GenTrafficRequest")
	proto.RegisterType((*GenTrafficResponse)(nil), "magma.cwf.GenTrafficResponse")
	proto.RegisterType((*TrafficOutput)(nil), "magma.cwf.TrafficOutput")
	proto.RegisterType((*TrafficSummary)(nil), "magma.cwf.TrafficSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UESimClient is the client API for UESim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UESimClient interface {
	// Adds a new UE to the store.
	//
	AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*GenTrafficResponse, error)
	ProxyRadius(ctx context.Context, in *ProxyRadiusRequest, opts ...grpc.CallOption) (*ProxyRadiusResponse, error)
}

type uESimClient struct {
	cc *grpc.ClientConn
}

func NewUESimClient(cc *grpc.ClientConn) UESimClient {
	return &uESimClient{cc}
}

func (c *uESimClient) AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/AddUE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*GenTrafficResponse, error) {
	out := new(GenTrafficResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/GenTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) ProxyRadius(ctx context.Context, in *ProxyRadiusRequest, opts ...grpc.CallOption) (*ProxyRadiusResponse, error) {
	out := new(ProxyRadiusResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/ProxyRadius", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UESimServer is the server API for UESim service.
type UESimServer interface {
	// Adds a new UE to the store.
	//
	AddUE(context.Context, *UEConfig) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(context.Context, *GenTrafficRequest) (*GenTrafficResponse, error)
	ProxyRadius(context.Context, *ProxyRadiusRequest) (*ProxyRadiusResponse, error)
}

func RegisterUESimServer(s *grpc.Server, srv UESimServer) {
	s.RegisterService(&_UESim_serviceDesc, srv)
}

func _UESim_AddUE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).AddUE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/AddUE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).AddUE(ctx, req.(*UEConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_GenTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).GenTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/GenTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).GenTraffic(ctx, req.(*GenTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_ProxyRadius_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyRadiusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).ProxyRadius(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/ProxyRadius",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).ProxyRadius(ctx, req.(*ProxyRadiusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UESim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.cwf.UESim",
	HandlerType: (*UESimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUE",
			Handler:    _UESim_AddUE_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UESim_Authenticate_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _UESim_Disconnect_Handler,
		},
		{
			MethodName: "GenTraffic",
			Handler:    _UESim_GenTraffic_Handler,
		},
		{
			MethodName: "ProxyRadius",
			Handler:    _UESim_ProxyRadius_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cwf/protos/ue_sim.proto",
}

func init() { proto.RegisterFile("cwf/protos/ue_sim.proto", fileDescriptor_ue_sim_71973c8f5128f810) }

var fileDescriptor_ue_sim_71973c8f5128f810 = []byte{
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x73, 0x1b, 0x35,
	0x14, 0xce, 0xa6, 0xf1, 0x8f, 0x3c, 0x3b, 0x94, 0x28, 0xa5, 0xdd, 0x98, 0x60, 0x3c, 0x3b, 0x03,
	0xe3, 0x93, 0x3d, 0x84, 0x0e, 0x84, 0x19, 0x2e, 0x25, 0x4d, 0x21, 0x53, 0x4a, 0x53, 0x99, 0x74,
	0x3a, 0x5c, 0x3c, 0xb2, 0xf6, 0xd9, 0xd6, 0x64, 0x57, 0xda, 0x4a, 0x5a, 0x07, 0x9f, 0xf9, 0x7f,
	0x38, 0xf0, 0xaf, 0x71, 0xe0, 0xca, 0xac, 0x76, 0xb7, 0x59, 0xe3, 0x3a, 0xa5, 0x33, 0x3d, 0xf9,
	0xbd, 0xef, 0xfb, 0xf4, 0xed, 0xb3, 0xf4, 0xf4, 0x04, 0x0f, 0xf8, 0xf5, 0x74, 0x98, 0x68, 0x65,
	0x95, 0x19, 0xa6, 0x38, 0x36, 0x22, 0x1e, 0xb8, 0x8c, 0xec, 0xc6, 0x6c, 0x16, 0xb3, 0x01, 0xbf,
	0x9e, 0x76, 0x0e, 0x95, 0xe6, 0x27, 0xba, 0x54, 0x71, 0x15, 0xc7, 0x4a, 0xe6, 0xaa, 0x4e, 0x77,
	0xa6, 0xd4, 0x2c, 0xc2, 0x9c, 0x9b, 0xa4, 0xd3, 0xe1, 0xb5, 0x66, 0x49, 0x82, 0xda, 0xe4, 0x7c,
	0xf0, 0x0a, 0x3a, 0x8f, 0x52, 0x3b, 0x47, 0x69, 0x05, 0x67, 0x16, 0x29, 0xbe, 0x4e, 0xd1, 0xd8,
	0x9f, 0x8c, 0xf9, 0x19, 0x8d, 0x21, 0xf7, 0xa1, 0x1e, 0x1b, 0x61, 0x42, 0xe9, 0x7b, 0x3d, 0xaf,
	0xbf, 0x4b, 0x8b, 0x8c, 0x7c, 0x0c, 0x77, 0x58, 0x22, 0xfd, 0x6d, 0x07, 0x66, 0x61, 0x86, 0x68,
	0x66, 0xfd, 0x3b, 0x3d, 0xaf, 0xbf, 0x47, 0xb3, 0x30, 0xf8, 0xd3, 0x83, 0xe6, 0xe5, 0xd9, 0xa9,
	0x92, 0x53, 0x31, 0x23, 0x04, 0x76, 0x44, 0x6c, 0x44, 0x61, 0xe3, 0x62, 0x72, 0x08, 0x4d, 0x96,
	0xda, 0xf9, 0xf8, 0x0a, 0x97, 0xce, 0xa9, 0x4d, 0x1b, 0x59, 0xfe, 0x14, 0x97, 0x6f, 0x28, 0x95,
	0x70, 0x67, 0x59, 0x50, 0xcf, 0x13, 0x9e, 0x7d, 0xc8, 0xe0, 0x6b, 0x7f, 0xa7, 0xe7, 0xf5, 0x77,
	0x68, 0x16, 0x92, 0x27, 0xd0, 0x9a, 0x1b, 0x13, 0xa1, 0x31, 0x63, 0x3e, 0x9d, 0xf9, 0xb5, 0x9e,
	0xd7, 0x6f, 0x1d, 0x7f, 0x31, 0x78, 0xb3, 0x3d, 0x83, 0xcd, 0x7f, 0x90, 0x42, 0xb1, 0xf2, 0x74,
	0x3a, 0x0b, 0x46, 0x70, 0xf0, 0x16, 0xe5, 0x5b, 0x4b, 0xef, 0xc3, 0x5d, 0xce, 0xa2, 0x08, 0xc3,
	0x91, 0x65, 0x56, 0x28, 0x79, 0xfe, 0xb8, 0xd8, 0x8b, 0xff, 0xc2, 0xc1, 0x2b, 0xb8, 0xb7, 0x6a,
	0x6a, 0x12, 0x25, 0x0d, 0x92, 0x00, 0xda, 0x9a, 0x85, 0x22, 0x35, 0x17, 0x8c, 0x5f, 0xa1, 0x75,
	0xee, 0x6d, 0xba, 0x82, 0x91, 0x23, 0xd8, 0x35, 0x68, 0x4c, 0x66, 0x14, 0x16, 0xfe, 0x37, 0x40,
	0xf0, 0x02, 0xf6, 0x1f, 0x0b, 0xc3, 0x95, 0x94, 0xc8, 0xed, 0x87, 0x29, 0xf6, 0x04, 0x48, 0xd5,
	0xf2, 0xff, 0x97, 0x1a, 0x5c, 0x02, 0xb9, 0xd0, 0xea, 0xf7, 0x25, 0x75, 0x60, 0x59, 0xcd, 0x27,
	0x50, 0x4f, 0x71, 0x1c, 0x33, 0x5e, 0xd4, 0x53, 0x4b, 0xf1, 0x19, 0xe3, 0xef, 0x51, 0xd0, 0x77,
	0x70, 0xb0, 0x62, 0xfb, 0x1e, 0x15, 0xfd, 0xbd, 0x0d, 0xfb, 0x3f, 0xa2, 0xfc, 0x55, 0xb3, 0xe9,
	0x54, 0xf0, 0xdb, 0xf6, 0xe7, 0x21, 0xd4, 0x17, 0x2a, 0x4a, 0x63, 0x74, 0x55, 0xb4, 0x8e, 0x8f,
	0x06, 0xf9, 0x9d, 0x19, 0x94, 0x77, 0x66, 0x30, 0xb2, 0x5a, 0xc8, 0xd9, 0x4b, 0x16, 0xa5, 0x48,
	0x0b, 0x2d, 0xf9, 0x06, 0x1a, 0x13, 0x61, 0x35, 0xb3, 0xe8, 0x3a, 0xf4, 0x5d, 0xcb, 0x4a, 0x31,
	0xe9, 0x02, 0x58, 0x11, 0xe3, 0xb9, 0x1c, 0x21, 0x37, 0x45, 0x1b, 0x57, 0x10, 0x72, 0x02, 0x0f,
	0x34, 0x26, 0x4a, 0x5b, 0x21, 0x67, 0xe7, 0xd2, 0xa2, 0x5e, 0xb0, 0xa8, 0x10, 0xd7, 0x9c, 0x78,
	0x13, 0x4d, 0x7a, 0xd0, 0xd2, 0xb8, 0x40, 0x6d, 0xf0, 0x99, 0x0a, 0xd1, 0xaf, 0xf7, 0xbc, 0x7e,
	0x93, 0x56, 0x21, 0xe2, 0x43, 0x23, 0xfb, 0x92, 0x4a, 0xad, 0xdf, 0x70, 0x17, 0xb5, 0x4c, 0xc9,
	0x13, 0xe8, 0x86, 0xc2, 0xb0, 0x49, 0x84, 0x23, 0xd4, 0x0b, 0xd4, 0x14, 0x19, 0x9f, 0xb3, 0x89,
	0x88, 0x84, 0x5d, 0x9e, 0xce, 0x91, 0x5f, 0xf9, 0x4d, 0x67, 0xf7, 0x0e, 0x55, 0x80, 0x40, 0xaa,
	0x9b, 0x5e, 0x9c, 0xd7, 0x7d, 0xa8, 0xab, 0xd4, 0x26, 0x69, 0x79, 0x52, 0x45, 0x46, 0xbe, 0x05,
	0x40, 0x19, 0x8e, 0x0b, 0x2e, 0xdf, 0x7d, 0xbf, 0x72, 0x71, 0x0b, 0x9f, 0xe7, 0x8e, 0xa7, 0xbb,
	0x28, 0xc3, 0x3c, 0x0c, 0xfe, 0xf0, 0x60, 0x6f, 0x85, 0x24, 0x0f, 0xa1, 0x69, 0xd2, 0x78, 0x6c,
	0x50, 0xe6, 0x1f, 0x69, 0x1d, 0x1f, 0xae, 0x1b, 0x8d, 0xd2, 0x38, 0x66, 0x7a, 0x49, 0x1b, 0x26,
	0x8d, 0x47, 0x28, 0x2d, 0xf9, 0x1e, 0xda, 0xd9, 0x2a, 0x8d, 0x1c, 0xc5, 0x02, 0xc3, 0xa2, 0x84,
	0x5b, 0x56, 0xb6, 0x4c, 0x1a, 0xd3, 0x42, 0x1d, 0xfc, 0xe5, 0xc1, 0x47, 0xab, 0x3c, 0xb9, 0x07,
	0x35, 0x63, 0x99, 0xce, 0x6b, 0xf0, 0x68, 0x9e, 0x64, 0x33, 0x0b, 0x65, 0xee, 0xee, 0xd1, 0x2c,
	0xcc, 0x4e, 0xc2, 0x20, 0x57, 0x32, 0x34, 0xae, 0x7b, 0x3c, 0x5a, 0xa6, 0x99, 0xc3, 0x64, 0x69,
	0x31, 0x6f, 0x8d, 0x1a, 0xcd, 0x13, 0xf2, 0x25, 0xdc, 0x9d, 0x08, 0x6b, 0xc6, 0x09, 0xea, 0x71,
	0xae, 0x74, 0xdd, 0xe0, 0xd1, 0xbd, 0x0c, 0xbe, 0x40, 0x3d, 0x72, 0x60, 0xde, 0x03, 0x56, 0x33,
	0x69, 0x62, 0x61, 0x8d, 0xeb, 0x81, 0x1a, 0xad, 0x42, 0xc7, 0xff, 0x6c, 0x43, 0xed, 0xf2, 0x6c,
	0x24, 0x62, 0xf2, 0x15, 0xd4, 0x1e, 0x85, 0xe1, 0xe5, 0x19, 0x39, 0xa8, 0xfc, 0xdf, 0x72, 0x62,
	0x77, 0xf6, 0x0b, 0xd0, 0x3d, 0x2d, 0x83, 0x97, 0x4a, 0x84, 0xc1, 0x16, 0x79, 0x01, 0xed, 0xea,
	0x34, 0x23, 0xdd, 0xdb, 0xa7, 0x6c, 0xe7, 0xf3, 0x8d, 0x7c, 0xde, 0x19, 0xc1, 0x16, 0x79, 0x0a,
	0x70, 0x33, 0x73, 0xc8, 0x51, 0x65, 0xc1, 0xda, 0x74, 0xeb, 0x7c, 0xb6, 0x81, 0xad, 0x9a, 0xdd,
	0xb4, 0xdf, 0x8a, 0xd9, 0xda, 0x28, 0x58, 0x31, 0x5b, 0xef, 0xd9, 0x60, 0x8b, 0xfc, 0x02, 0xad,
	0xca, 0xf0, 0x21, 0x55, 0xfd, 0xfa, 0xac, 0xeb, 0x74, 0x37, 0xd1, 0xa5, 0xdf, 0x0f, 0x9f, 0xfe,
	0x76, 0xe8, 0x24, 0xc3, 0xec, 0x45, 0xe7, 0x91, 0x4a, 0xc3, 0xe1, 0x4c, 0x15, 0x8f, 0xf6, 0xa4,
	0xee, 0x7e, 0xbf, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xd0, 0xd5, 0x3e, 0xef, 0x07, 0x00,
	0x00,
}
