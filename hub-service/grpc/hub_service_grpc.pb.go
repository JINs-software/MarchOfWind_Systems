// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proto/hub_service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HubService_CreateMatchRoom_FullMethodName = "/proto.HubService/CreateMatchRoom"
	HubService_GetMatchRooms_FullMethodName   = "/proto.HubService/GetMatchRooms"
)

// HubServiceClient is the client API for HubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubServiceClient interface {
	CreateMatchRoom(ctx context.Context, in *CreateMatchRoomRequest, opts ...grpc.CallOption) (*CreateMatchRoomResponse, error)
	GetMatchRooms(ctx context.Context, in *GetMatchRoomsRequest, opts ...grpc.CallOption) (*GetMatchRoomsResponse, error)
}

type hubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHubServiceClient(cc grpc.ClientConnInterface) HubServiceClient {
	return &hubServiceClient{cc}
}

func (c *hubServiceClient) CreateMatchRoom(ctx context.Context, in *CreateMatchRoomRequest, opts ...grpc.CallOption) (*CreateMatchRoomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMatchRoomResponse)
	err := c.cc.Invoke(ctx, HubService_CreateMatchRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubServiceClient) GetMatchRooms(ctx context.Context, in *GetMatchRoomsRequest, opts ...grpc.CallOption) (*GetMatchRoomsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMatchRoomsResponse)
	err := c.cc.Invoke(ctx, HubService_GetMatchRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubServiceServer is the server API for HubService service.
// All implementations must embed UnimplementedHubServiceServer
// for forward compatibility.
type HubServiceServer interface {
	CreateMatchRoom(context.Context, *CreateMatchRoomRequest) (*CreateMatchRoomResponse, error)
	GetMatchRooms(context.Context, *GetMatchRoomsRequest) (*GetMatchRoomsResponse, error)
	mustEmbedUnimplementedHubServiceServer()
}

// UnimplementedHubServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHubServiceServer struct{}

func (UnimplementedHubServiceServer) CreateMatchRoom(context.Context, *CreateMatchRoomRequest) (*CreateMatchRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatchRoom not implemented")
}
func (UnimplementedHubServiceServer) GetMatchRooms(context.Context, *GetMatchRoomsRequest) (*GetMatchRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchRooms not implemented")
}
func (UnimplementedHubServiceServer) mustEmbedUnimplementedHubServiceServer() {}
func (UnimplementedHubServiceServer) testEmbeddedByValue()                    {}

// UnsafeHubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServiceServer will
// result in compilation errors.
type UnsafeHubServiceServer interface {
	mustEmbedUnimplementedHubServiceServer()
}

func RegisterHubServiceServer(s grpc.ServiceRegistrar, srv HubServiceServer) {
	// If the following call pancis, it indicates UnimplementedHubServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HubService_ServiceDesc, srv)
}

func _HubService_CreateMatchRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServiceServer).CreateMatchRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HubService_CreateMatchRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServiceServer).CreateMatchRoom(ctx, req.(*CreateMatchRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubService_GetMatchRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServiceServer).GetMatchRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HubService_GetMatchRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServiceServer).GetMatchRooms(ctx, req.(*GetMatchRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HubService_ServiceDesc is the grpc.ServiceDesc for HubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HubService",
	HandlerType: (*HubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMatchRoom",
			Handler:    _HubService_CreateMatchRoom_Handler,
		},
		{
			MethodName: "GetMatchRooms",
			Handler:    _HubService_GetMatchRooms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hub_service.proto",
}
