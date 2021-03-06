// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabric.proto

package fabric

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type of Fabric transaction: INVOKE or QUERY.
type TransactionType int32

const (
	TransactionType_INVOKE TransactionType = 0
	TransactionType_QUERY  TransactionType = 1
)

var TransactionType_name = map[int32]string{
	0: "INVOKE",
	1: "QUERY",
}

var TransactionType_value = map[string]int32{
	"INVOKE": 0,
	"QUERY":  1,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}

func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{0}
}

// Fabric transaction request.
type TransactionRequest struct {
	Data                 *TransactionData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{0}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetData() *TransactionData {
	if m != nil {
		return m.Data
	}
	return nil
}

// Request data for a Fabric chaincode transaction
type TransactionData struct {
	// Required. ID of a Fabric connection returned by a ConnectRequest
	ConnectionId uint64 `protobuf:"varint,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// Required. 0=INVOKE or 1=QUERY
	Type TransactionType `protobuf:"varint,2,opt,name=type,proto3,enum=dovetail.fabric.v1.TransactionType" json:"type,omitempty"`
	// Required. Name of the Fabric chaincode to be called.
	ChaincodeId string `protobuf:"bytes,3,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	// Request timeout in milliseconds. Default is 0, i.e., no timeout.
	Timeout int64 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Fabric connnection endpoints to send the request. Default is randomly chosen from all available endpoints.
	Endpoint []string `protobuf:"bytes,5,rep,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Required. Name of the Fabric chaincode transaction.
	Transaction string `protobuf:"bytes,6,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// parameters for the transaction
	Parameter []string `protobuf:"bytes,7,rep,name=parameter,proto3" json:"parameter,omitempty"`
	// transient map for the transaction as a JSON string.
	TransientMap         string   `protobuf:"bytes,8,opt,name=transient_map,json=transientMap,proto3" json:"transient_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionData) Reset()         { *m = TransactionData{} }
func (m *TransactionData) String() string { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()    {}
func (*TransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{1}
}

func (m *TransactionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionData.Unmarshal(m, b)
}
func (m *TransactionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionData.Marshal(b, m, deterministic)
}
func (m *TransactionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionData.Merge(m, src)
}
func (m *TransactionData) XXX_Size() int {
	return xxx_messageInfo_TransactionData.Size(m)
}
func (m *TransactionData) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionData.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionData proto.InternalMessageInfo

func (m *TransactionData) GetConnectionId() uint64 {
	if m != nil {
		return m.ConnectionId
	}
	return 0
}

func (m *TransactionData) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_INVOKE
}

func (m *TransactionData) GetChaincodeId() string {
	if m != nil {
		return m.ChaincodeId
	}
	return ""
}

func (m *TransactionData) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *TransactionData) GetEndpoint() []string {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *TransactionData) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *TransactionData) GetParameter() []string {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *TransactionData) GetTransientMap() string {
	if m != nil {
		return m.TransientMap
	}
	return ""
}

// Response from a Fabric transaction
type TransactionResponse struct {
	// Status code. 200=success
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Error messages if code is not 200
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Returned data as JSON string
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{2}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TransactionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TransactionResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Fabric connection request.
type ConnectionRequest struct {
	Data                 *ConnectionData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ConnectionRequest) Reset()         { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()    {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{3}
}

func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionRequest.Unmarshal(m, b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
}
func (m *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(m, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionRequest.Size(m)
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetData() *ConnectionData {
	if m != nil {
		return m.Data
	}
	return nil
}

// Request data for creating a Fabric connection in the gateway service
type ConnectionData struct {
	// Name of the Fabric channel to connect to. Default is pre-configured by gateway service.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Name of Fabric client user for establishing the connection. Default is pre-configured by gateway service.
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// Name of the Fabric org that created the client user. Default is pre-configured by gateway service.
	OrgName string `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// File name for Fabric network config. Default is pre-configured by gateway service.
	NetworkConfig string `protobuf:"bytes,4,opt,name=network_config,json=networkConfig,proto3" json:"network_config,omitempty"`
	// File name for endpoint patten matchers. Default is pre-configured by gateway service.
	PattenMatchers       string   `protobuf:"bytes,5,opt,name=patten_matchers,json=pattenMatchers,proto3" json:"patten_matchers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionData) Reset()         { *m = ConnectionData{} }
func (m *ConnectionData) String() string { return proto.CompactTextString(m) }
func (*ConnectionData) ProtoMessage()    {}
func (*ConnectionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{4}
}

func (m *ConnectionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionData.Unmarshal(m, b)
}
func (m *ConnectionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionData.Marshal(b, m, deterministic)
}
func (m *ConnectionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionData.Merge(m, src)
}
func (m *ConnectionData) XXX_Size() int {
	return xxx_messageInfo_ConnectionData.Size(m)
}
func (m *ConnectionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionData.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionData proto.InternalMessageInfo

func (m *ConnectionData) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ConnectionData) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ConnectionData) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *ConnectionData) GetNetworkConfig() string {
	if m != nil {
		return m.NetworkConfig
	}
	return ""
}

func (m *ConnectionData) GetPattenMatchers() string {
	if m != nil {
		return m.PattenMatchers
	}
	return ""
}

// Response from a Fabric connection request
type ConnectionResponse struct {
	// Status code. 200=success
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Error messages if code is not 200
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// ID of the resulting Fabric connection
	ConnectionId         uint64   `protobuf:"varint,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionResponse) Reset()         { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()    {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{5}
}

func (m *ConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionResponse.Unmarshal(m, b)
}
func (m *ConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionResponse.Marshal(b, m, deterministic)
}
func (m *ConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResponse.Merge(m, src)
}
func (m *ConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionResponse.Size(m)
}
func (m *ConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResponse proto.InternalMessageInfo

func (m *ConnectionResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ConnectionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ConnectionResponse) GetConnectionId() uint64 {
	if m != nil {
		return m.ConnectionId
	}
	return 0
}

func init() {
	proto.RegisterEnum("dovetail.fabric.v1.TransactionType", TransactionType_name, TransactionType_value)
	proto.RegisterType((*TransactionRequest)(nil), "dovetail.fabric.v1.TransactionRequest")
	proto.RegisterType((*TransactionData)(nil), "dovetail.fabric.v1.TransactionData")
	proto.RegisterType((*TransactionResponse)(nil), "dovetail.fabric.v1.TransactionResponse")
	proto.RegisterType((*ConnectionRequest)(nil), "dovetail.fabric.v1.ConnectionRequest")
	proto.RegisterType((*ConnectionData)(nil), "dovetail.fabric.v1.ConnectionData")
	proto.RegisterType((*ConnectionResponse)(nil), "dovetail.fabric.v1.ConnectionResponse")
}

func init() { proto.RegisterFile("fabric.proto", fileDescriptor_284efff686d8e9bf) }

var fileDescriptor_284efff686d8e9bf = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xcd, 0x8f, 0xe3, 0x69, 0x9a, 0x96, 0x45, 0x02, 0x53, 0x5a, 0xc9, 0xb8, 0x6a, 0x6b,
	0x71, 0x48, 0xd4, 0x22, 0x81, 0xc4, 0x91, 0x52, 0xa1, 0xa8, 0x6a, 0x11, 0xab, 0x02, 0x82, 0x4b,
	0x34, 0xb5, 0xa7, 0xae, 0x45, 0xbd, 0x6b, 0xec, 0x4d, 0x4a, 0xae, 0xbc, 0x02, 0xcf, 0xc1, 0x99,
	0x07, 0xe1, 0x15, 0x78, 0x0c, 0x0e, 0xc8, 0x6b, 0x3b, 0x3f, 0x04, 0x11, 0x89, 0x5b, 0xf6, 0x9b,
	0x9f, 0x6f, 0xe6, 0xfb, 0xc6, 0x81, 0xf6, 0x25, 0x5e, 0xa4, 0x91, 0xdf, 0x4d, 0x52, 0xa9, 0x24,
	0x63, 0x81, 0x1c, 0x91, 0xc2, 0xe8, 0xba, 0x5b, 0xc2, 0xa3, 0x83, 0xcd, 0xad, 0x50, 0xca, 0xf0,
	0x9a, 0x7a, 0x98, 0x44, 0x3d, 0x14, 0x42, 0x2a, 0x54, 0x91, 0x14, 0x59, 0x51, 0xe1, 0x9e, 0x02,
	0x3b, 0x4f, 0x51, 0x64, 0xe8, 0xe7, 0x28, 0xa7, 0x4f, 0x43, 0xca, 0x14, 0x7b, 0x0a, 0xf5, 0x00,
	0x15, 0xda, 0x86, 0x63, 0x78, 0xab, 0x87, 0x3b, 0xdd, 0xc5, 0xb6, 0xdd, 0x99, 0xaa, 0x17, 0xa8,
	0x90, 0xeb, 0x02, 0xf7, 0xdb, 0x0a, 0xac, 0xff, 0x11, 0x61, 0x3b, 0xb0, 0xe6, 0x4b, 0x21, 0x48,
	0x23, 0x83, 0x28, 0xd0, 0x5d, 0xeb, 0xbc, 0x3d, 0x05, 0xfb, 0x41, 0xce, 0xa8, 0xc6, 0x09, 0xd9,
	0x2b, 0x8e, 0xe1, 0x75, 0x96, 0x32, 0x9e, 0x8f, 0x13, 0xe2, 0xba, 0x80, 0x3d, 0x84, 0xb6, 0x7f,
	0x85, 0x91, 0xf0, 0x65, 0x40, 0x79, 0xf3, 0x9a, 0x63, 0x78, 0x16, 0x5f, 0x9d, 0x60, 0xfd, 0x80,
	0xd9, 0x60, 0xaa, 0x28, 0x26, 0x39, 0x54, 0x76, 0xdd, 0x31, 0xbc, 0x1a, 0xaf, 0x9e, 0x6c, 0x13,
	0x5a, 0x24, 0x82, 0x44, 0x46, 0x42, 0xd9, 0x0d, 0xa7, 0xe6, 0x59, 0x7c, 0xf2, 0x66, 0x0e, 0xac,
	0xaa, 0x29, 0xa3, 0xdd, 0x2c, 0xfa, 0xce, 0x40, 0x6c, 0x0b, 0xac, 0x04, 0x53, 0x8c, 0x49, 0x51,
	0x6a, 0x9b, 0xba, 0x7c, 0x0a, 0xe4, 0x6b, 0xeb, 0xe4, 0x88, 0x84, 0x1a, 0xc4, 0x98, 0xd8, 0x2d,
	0xdd, 0xa1, 0x3d, 0x01, 0x4f, 0x31, 0x71, 0xdf, 0xc1, 0x9d, 0x39, 0xf9, 0xb3, 0x44, 0x8a, 0x8c,
	0x18, 0x83, 0x7a, 0x3e, 0xbb, 0x56, 0xaa, 0xc1, 0xf5, 0xef, 0x7c, 0x8b, 0x98, 0xb2, 0x0c, 0xc3,
	0x42, 0x24, 0x8b, 0x57, 0xcf, 0x3c, 0x5b, 0xbb, 0x55, 0xac, 0x5e, 0x18, 0x71, 0x02, 0xb7, 0x8f,
	0x26, 0xfa, 0x56, 0xb6, 0x3e, 0x99, 0xb3, 0xd5, 0xfd, 0x9b, 0xc8, 0xd3, 0xa2, 0x19, 0x57, 0xbf,
	0x1b, 0xd0, 0x99, 0x0f, 0xb0, 0x6d, 0x00, 0xff, 0x0a, 0x85, 0xa0, 0xeb, 0xca, 0x51, 0x8b, 0x5b,
	0x25, 0xd2, 0x0f, 0xd8, 0x03, 0xb0, 0x86, 0x19, 0xa5, 0x03, 0x81, 0x71, 0x35, 0x6e, 0x2b, 0x07,
	0xce, 0x30, 0x26, 0x76, 0x1f, 0x5a, 0x32, 0x0d, 0x8b, 0x58, 0x31, 0xb3, 0x29, 0xd3, 0x50, 0x87,
	0x76, 0xa1, 0x23, 0x48, 0xdd, 0xc8, 0xf4, 0xe3, 0xc0, 0x97, 0xe2, 0x32, 0x0a, 0xb5, 0x63, 0x16,
	0x5f, 0x2b, 0xd1, 0x23, 0x0d, 0xb2, 0x7d, 0x58, 0x4f, 0x50, 0x29, 0x12, 0x83, 0x18, 0x95, 0x7f,
	0x45, 0x69, 0x66, 0x37, 0x74, 0x5e, 0xa7, 0x80, 0x4f, 0x4b, 0xd4, 0x0d, 0x81, 0xcd, 0xca, 0xf0,
	0x5f, 0xf2, 0x2e, 0xdc, 0x6f, 0x6d, 0xf1, 0x7e, 0x1f, 0x79, 0x73, 0x77, 0x9f, 0xdf, 0x27, 0x03,
	0x68, 0xf6, 0xcf, 0xde, 0xbe, 0x3a, 0x39, 0xde, 0xb8, 0xc5, 0x2c, 0x68, 0xbc, 0x7e, 0x73, 0xcc,
	0xdf, 0x6f, 0x18, 0x87, 0xbf, 0x0c, 0x30, 0x5f, 0xa2, 0xa2, 0x1b, 0x1c, 0xb3, 0x11, 0x98, 0xe5,
	0x78, 0x6c, 0xf7, 0xdf, 0x6e, 0x94, 0x16, 0x6e, 0xee, 0x2d, 0x4b, 0x2b, 0x56, 0x74, 0xb7, 0xbe,
	0xfc, 0xf8, 0xf9, 0x75, 0xe5, 0xae, 0xdb, 0xe9, 0x8d, 0x0e, 0x7a, 0xd3, 0x49, 0x9f, 0x69, 0x43,
	0xd9, 0x18, 0xcc, 0xe3, 0xcf, 0xe4, 0x0f, 0x15, 0xb1, 0xbd, 0x25, 0x9f, 0x5a, 0x45, 0xbc, 0xbf,
	0x34, 0xaf, 0x64, 0xde, 0xd6, 0xcc, 0xf7, 0xdc, 0xf5, 0x9c, 0x79, 0xe6, 0x73, 0x29, 0xa8, 0x9f,
	0xb7, 0x3e, 0x34, 0x8b, 0xfa, 0x8b, 0xa6, 0xfe, 0x07, 0x7a, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x09, 0x58, 0x7b, 0x33, 0xc3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	// Connect: create or find a Fabric connection.
	// Returns the corresponding connection ID.
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	// Execute: execute a transaction on a Fabric endpoint.
	// Returns the result of the transaction.
	Execute(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/dovetail.fabric.v1.Gateway/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Execute(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/dovetail.fabric.v1.Gateway/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	// Connect: create or find a Fabric connection.
	// Returns the corresponding connection ID.
	Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	// Execute: execute a transaction on a Fabric endpoint.
	// Returns the result of the transaction.
	Execute(context.Context, *TransactionRequest) (*TransactionResponse, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) Connect(ctx context.Context, req *ConnectionRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedGatewayServer) Execute(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dovetail.fabric.v1.Gateway/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dovetail.fabric.v1.Gateway/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Execute(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dovetail.fabric.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Gateway_Connect_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Gateway_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fabric.proto",
}
