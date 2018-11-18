// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplane.proto

package dataplane

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	crossconnect "github.com/ligato/networkservicemesh/controlplane/pkg/apis/crossconnect"
	connection1 "github.com/ligato/networkservicemesh/controlplane/pkg/apis/local/connection"
	connection "github.com/ligato/networkservicemesh/controlplane/pkg/apis/remote/connection"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message sent by dataplane module informing NSM of any changes in its
// operations parameters or constraints
type MechanismUpdate struct {
	RemoteMechanisms     []*connection.Mechanism  `protobuf:"bytes,1,rep,name=remote_mechanisms,json=remoteMechanisms,proto3" json:"remote_mechanisms,omitempty"`
	LocalMechanisms      []*connection1.Mechanism `protobuf:"bytes,2,rep,name=local_mechanisms,json=localMechanisms,proto3" json:"local_mechanisms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MechanismUpdate) Reset()         { *m = MechanismUpdate{} }
func (m *MechanismUpdate) String() string { return proto.CompactTextString(m) }
func (*MechanismUpdate) ProtoMessage()    {}
func (*MechanismUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_617387e490a04ffa, []int{0}
}

func (m *MechanismUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MechanismUpdate.Unmarshal(m, b)
}
func (m *MechanismUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MechanismUpdate.Marshal(b, m, deterministic)
}
func (m *MechanismUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MechanismUpdate.Merge(m, src)
}
func (m *MechanismUpdate) XXX_Size() int {
	return xxx_messageInfo_MechanismUpdate.Size(m)
}
func (m *MechanismUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MechanismUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MechanismUpdate proto.InternalMessageInfo

func (m *MechanismUpdate) GetRemoteMechanisms() []*connection.Mechanism {
	if m != nil {
		return m.RemoteMechanisms
	}
	return nil
}

func (m *MechanismUpdate) GetLocalMechanisms() []*connection1.Mechanism {
	if m != nil {
		return m.LocalMechanisms
	}
	return nil
}

func init() {
	proto.RegisterType((*MechanismUpdate)(nil), "dataplane.MechanismUpdate")
}

func init() { proto.RegisterFile("dataplane.proto", fileDescriptor_617387e490a04ffa) }

var fileDescriptor_617387e490a04ffa = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4e, 0x23, 0x31,
	0x10, 0x87, 0xb5, 0x77, 0xba, 0x3b, 0xc5, 0x57, 0x24, 0x6c, 0x81, 0xd0, 0x42, 0x81, 0xa8, 0xa8,
	0x6c, 0x94, 0x94, 0x54, 0x28, 0x80, 0x44, 0x91, 0x26, 0x12, 0x0d, 0x02, 0x21, 0xaf, 0x33, 0x71,
	0xac, 0x78, 0x3d, 0xc6, 0xf6, 0x82, 0xf2, 0x3e, 0x3c, 0x13, 0xcf, 0x83, 0xd6, 0xce, 0x9f, 0x0d,
	0x82, 0x34, 0x69, 0x56, 0x3b, 0xfe, 0x3c, 0x9f, 0x35, 0x3f, 0x0d, 0xe9, 0x4e, 0x78, 0xe0, 0x56,
	0x73, 0x03, 0xd4, 0x3a, 0x0c, 0x98, 0x77, 0xd6, 0x07, 0xc5, 0xa3, 0x54, 0x61, 0x56, 0x97, 0x54,
	0x60, 0xc5, 0xb4, 0x92, 0x3c, 0x20, 0x33, 0x10, 0xde, 0xd0, 0xcd, 0x3d, 0xb8, 0x57, 0x25, 0xa0,
	0x02, 0x3f, 0x63, 0x02, 0x4d, 0x70, 0xa8, 0x63, 0x0b, 0xb3, 0x73, 0xc9, 0xb8, 0x55, 0x9e, 0x69,
	0x14, 0x5c, 0x37, 0xcc, 0x80, 0x08, 0x0a, 0x4d, 0xeb, 0x37, 0x3d, 0x54, 0x3c, 0xed, 0x61, 0x77,
	0x50, 0x61, 0x80, 0x9d, 0xfa, 0x87, 0x3d, 0xf4, 0xc2, 0xa1, 0xf7, 0x4b, 0xe3, 0x56, 0xb1, 0x74,
	0x0f, 0x5a, 0x6e, 0x89, 0x9a, 0x1b, 0xc9, 0x22, 0x28, 0xeb, 0x29, 0xb3, 0x61, 0x61, 0xc1, 0x33,
	0xa8, 0x6c, 0x58, 0xa4, 0x6f, 0x6a, 0x3a, 0x7b, 0xcf, 0x48, 0x77, 0x04, 0x62, 0xc6, 0x8d, 0xf2,
	0xd5, 0xbd, 0x9d, 0xf0, 0x00, 0xf9, 0x1d, 0x39, 0x48, 0xa3, 0x3c, 0x57, 0x2b, 0xe2, 0x8f, 0xb2,
	0xd3, 0xdf, 0xe7, 0xff, 0xfb, 0x27, 0x34, 0x11, 0xda, 0x9a, 0x6c, 0xdd, 0x3e, 0xee, 0x25, 0xb8,
	0x3e, 0xf0, 0xf9, 0x2d, 0xe9, 0xc5, 0xcc, 0xdb, 0xa6, 0x5f, 0xd1, 0x74, 0x4c, 0x23, 0xf8, 0x5e,
	0xd4, 0x8d, 0x6c, 0xe3, 0xe9, 0x7f, 0x64, 0xa4, 0x73, 0xbd, 0x5a, 0x81, 0xfc, 0x8a, 0xfc, 0x1b,
	0xc3, 0x4b, 0x0d, 0x3e, 0xe4, 0x05, 0xdd, 0x4a, 0x62, 0xd8, 0x14, 0xc3, 0x54, 0x14, 0x3b, 0x58,
	0x7e, 0x49, 0xfe, 0x0c, 0x35, 0x7a, 0xd8, 0x29, 0x38, 0xa4, 0x12, 0x51, 0xea, 0xe5, 0x12, 0x96,
	0xf5, 0x94, 0xde, 0x34, 0xd1, 0x35, 0x01, 0x8d, 0xd0, 0xa8, 0x80, 0xae, 0x35, 0xea, 0x0f, 0x97,
	0x8b, 0x82, 0x6e, 0x96, 0xf9, 0x4b, 0xd2, 0x17, 0x59, 0xf9, 0x37, 0xde, 0x1e, 0x7c, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0xf5, 0x5c, 0x35, 0xf2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneClient is the client API for Dataplane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneClient interface {
	Request(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*crossconnect.CrossConnect, error)
	Close(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*empty.Empty, error)
	MonitorMechanisms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Dataplane_MonitorMechanismsClient, error)
}

type dataplaneClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneClient(cc *grpc.ClientConn) DataplaneClient {
	return &dataplaneClient{cc}
}

func (c *dataplaneClient) Request(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*crossconnect.CrossConnect, error) {
	out := new(crossconnect.CrossConnect)
	err := c.cc.Invoke(ctx, "/dataplane.Dataplane/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneClient) Close(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dataplane.Dataplane/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneClient) MonitorMechanisms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Dataplane_MonitorMechanismsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dataplane_serviceDesc.Streams[0], "/dataplane.Dataplane/MonitorMechanisms", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneMonitorMechanismsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dataplane_MonitorMechanismsClient interface {
	Recv() (*MechanismUpdate, error)
	grpc.ClientStream
}

type dataplaneMonitorMechanismsClient struct {
	grpc.ClientStream
}

func (x *dataplaneMonitorMechanismsClient) Recv() (*MechanismUpdate, error) {
	m := new(MechanismUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneServer is the server API for Dataplane service.
type DataplaneServer interface {
	Request(context.Context, *crossconnect.CrossConnect) (*crossconnect.CrossConnect, error)
	Close(context.Context, *crossconnect.CrossConnect) (*empty.Empty, error)
	MonitorMechanisms(*empty.Empty, Dataplane_MonitorMechanismsServer) error
}

func RegisterDataplaneServer(s *grpc.Server, srv DataplaneServer) {
	s.RegisterService(&_Dataplane_serviceDesc, srv)
}

func _Dataplane_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(crossconnect.CrossConnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplane.Dataplane/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneServer).Request(ctx, req.(*crossconnect.CrossConnect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataplane_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(crossconnect.CrossConnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplane.Dataplane/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneServer).Close(ctx, req.(*crossconnect.CrossConnect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataplane_MonitorMechanisms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataplaneServer).MonitorMechanisms(m, &dataplaneMonitorMechanismsServer{stream})
}

type Dataplane_MonitorMechanismsServer interface {
	Send(*MechanismUpdate) error
	grpc.ServerStream
}

type dataplaneMonitorMechanismsServer struct {
	grpc.ServerStream
}

func (x *dataplaneMonitorMechanismsServer) Send(m *MechanismUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _Dataplane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplane.Dataplane",
	HandlerType: (*DataplaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Dataplane_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Dataplane_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorMechanisms",
			Handler:       _Dataplane_MonitorMechanisms_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dataplane.proto",
}
