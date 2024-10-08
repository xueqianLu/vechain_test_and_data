// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: hackcenter/center.proto

package hackcenter

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CenterServiceClient is the client API for CenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CenterServiceClient interface {
	SubmitBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*SubmitBlockResponse, error)
	SubscribeBlock(ctx context.Context, in *SubscribeBlockRequest, opts ...grpc.CallOption) (CenterService_SubscribeBlockClient, error)
	SubBroadcastTask(ctx context.Context, in *SubBroadcastTaskRequest, opts ...grpc.CallOption) (CenterService_SubBroadcastTaskClient, error)
	//    rpc BeginToHack(BeginToHackRequest) returns (BeginToHackResponse);
	RegisterNode(ctx context.Context, in *NodeRegisterInfo, opts ...grpc.CallOption) (*NodeRegisterResponse, error)
	FetchNode(ctx context.Context, in *FetchNodeRequest, opts ...grpc.CallOption) (*FetchNodeResponse, error)
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	SubscribeMinedBlock(ctx context.Context, in *SubscribeBlockRequest, opts ...grpc.CallOption) (CenterService_SubscribeMinedBlockClient, error)
	BroadcastBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*SubmitBlockResponse, error)
	UpdateHack(ctx context.Context, in *UpdateHackRequest, opts ...grpc.CallOption) (*Empty, error)
}

type centerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCenterServiceClient(cc grpc.ClientConnInterface) CenterServiceClient {
	return &centerServiceClient{cc}
}

func (c *centerServiceClient) SubmitBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*SubmitBlockResponse, error) {
	out := new(SubmitBlockResponse)
	err := c.cc.Invoke(ctx, "/hackcenter.CenterService/SubmitBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) SubscribeBlock(ctx context.Context, in *SubscribeBlockRequest, opts ...grpc.CallOption) (CenterService_SubscribeBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &CenterService_ServiceDesc.Streams[0], "/hackcenter.CenterService/SubscribeBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &centerServiceSubscribeBlockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CenterService_SubscribeBlockClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type centerServiceSubscribeBlockClient struct {
	grpc.ClientStream
}

func (x *centerServiceSubscribeBlockClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *centerServiceClient) SubBroadcastTask(ctx context.Context, in *SubBroadcastTaskRequest, opts ...grpc.CallOption) (CenterService_SubBroadcastTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &CenterService_ServiceDesc.Streams[1], "/hackcenter.CenterService/SubBroadcastTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &centerServiceSubBroadcastTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CenterService_SubBroadcastTaskClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type centerServiceSubBroadcastTaskClient struct {
	grpc.ClientStream
}

func (x *centerServiceSubBroadcastTaskClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *centerServiceClient) RegisterNode(ctx context.Context, in *NodeRegisterInfo, opts ...grpc.CallOption) (*NodeRegisterResponse, error) {
	out := new(NodeRegisterResponse)
	err := c.cc.Invoke(ctx, "/hackcenter.CenterService/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) FetchNode(ctx context.Context, in *FetchNodeRequest, opts ...grpc.CallOption) (*FetchNodeResponse, error) {
	out := new(FetchNodeResponse)
	err := c.cc.Invoke(ctx, "/hackcenter.CenterService/FetchNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, "/hackcenter.CenterService/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) SubscribeMinedBlock(ctx context.Context, in *SubscribeBlockRequest, opts ...grpc.CallOption) (CenterService_SubscribeMinedBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &CenterService_ServiceDesc.Streams[2], "/hackcenter.CenterService/SubscribeMinedBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &centerServiceSubscribeMinedBlockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CenterService_SubscribeMinedBlockClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type centerServiceSubscribeMinedBlockClient struct {
	grpc.ClientStream
}

func (x *centerServiceSubscribeMinedBlockClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *centerServiceClient) BroadcastBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*SubmitBlockResponse, error) {
	out := new(SubmitBlockResponse)
	err := c.cc.Invoke(ctx, "/hackcenter.CenterService/BroadcastBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) UpdateHack(ctx context.Context, in *UpdateHackRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/hackcenter.CenterService/UpdateHack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CenterServiceServer is the server API for CenterService service.
// All implementations must embed UnimplementedCenterServiceServer
// for forward compatibility
type CenterServiceServer interface {
	SubmitBlock(context.Context, *Block) (*SubmitBlockResponse, error)
	SubscribeBlock(*SubscribeBlockRequest, CenterService_SubscribeBlockServer) error
	SubBroadcastTask(*SubBroadcastTaskRequest, CenterService_SubBroadcastTaskServer) error
	//    rpc BeginToHack(BeginToHackRequest) returns (BeginToHackResponse);
	RegisterNode(context.Context, *NodeRegisterInfo) (*NodeRegisterResponse, error)
	FetchNode(context.Context, *FetchNodeRequest) (*FetchNodeResponse, error)
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	SubscribeMinedBlock(*SubscribeBlockRequest, CenterService_SubscribeMinedBlockServer) error
	BroadcastBlock(context.Context, *Block) (*SubmitBlockResponse, error)
	UpdateHack(context.Context, *UpdateHackRequest) (*Empty, error)
	mustEmbedUnimplementedCenterServiceServer()
}

// UnimplementedCenterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCenterServiceServer struct {
}

func (UnimplementedCenterServiceServer) SubmitBlock(context.Context, *Block) (*SubmitBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBlock not implemented")
}
func (UnimplementedCenterServiceServer) SubscribeBlock(*SubscribeBlockRequest, CenterService_SubscribeBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBlock not implemented")
}
func (UnimplementedCenterServiceServer) SubBroadcastTask(*SubBroadcastTaskRequest, CenterService_SubBroadcastTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method SubBroadcastTask not implemented")
}
func (UnimplementedCenterServiceServer) RegisterNode(context.Context, *NodeRegisterInfo) (*NodeRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (UnimplementedCenterServiceServer) FetchNode(context.Context, *FetchNodeRequest) (*FetchNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchNode not implemented")
}
func (UnimplementedCenterServiceServer) Vote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedCenterServiceServer) SubscribeMinedBlock(*SubscribeBlockRequest, CenterService_SubscribeMinedBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeMinedBlock not implemented")
}
func (UnimplementedCenterServiceServer) BroadcastBlock(context.Context, *Block) (*SubmitBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastBlock not implemented")
}
func (UnimplementedCenterServiceServer) UpdateHack(context.Context, *UpdateHackRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHack not implemented")
}
func (UnimplementedCenterServiceServer) mustEmbedUnimplementedCenterServiceServer() {}

// UnsafeCenterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CenterServiceServer will
// result in compilation errors.
type UnsafeCenterServiceServer interface {
	mustEmbedUnimplementedCenterServiceServer()
}

func RegisterCenterServiceServer(s grpc.ServiceRegistrar, srv CenterServiceServer) {
	s.RegisterService(&CenterService_ServiceDesc, srv)
}

func _CenterService_SubmitBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).SubmitBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hackcenter.CenterService/SubmitBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).SubmitBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_SubscribeBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CenterServiceServer).SubscribeBlock(m, &centerServiceSubscribeBlockServer{stream})
}

type CenterService_SubscribeBlockServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type centerServiceSubscribeBlockServer struct {
	grpc.ServerStream
}

func (x *centerServiceSubscribeBlockServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _CenterService_SubBroadcastTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubBroadcastTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CenterServiceServer).SubBroadcastTask(m, &centerServiceSubBroadcastTaskServer{stream})
}

type CenterService_SubBroadcastTaskServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type centerServiceSubBroadcastTaskServer struct {
	grpc.ServerStream
}

func (x *centerServiceSubBroadcastTaskServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _CenterService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRegisterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hackcenter.CenterService/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).RegisterNode(ctx, req.(*NodeRegisterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_FetchNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).FetchNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hackcenter.CenterService/FetchNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).FetchNode(ctx, req.(*FetchNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hackcenter.CenterService/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_SubscribeMinedBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CenterServiceServer).SubscribeMinedBlock(m, &centerServiceSubscribeMinedBlockServer{stream})
}

type CenterService_SubscribeMinedBlockServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type centerServiceSubscribeMinedBlockServer struct {
	grpc.ServerStream
}

func (x *centerServiceSubscribeMinedBlockServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _CenterService_BroadcastBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).BroadcastBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hackcenter.CenterService/BroadcastBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).BroadcastBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_UpdateHack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).UpdateHack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hackcenter.CenterService/UpdateHack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).UpdateHack(ctx, req.(*UpdateHackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CenterService_ServiceDesc is the grpc.ServiceDesc for CenterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CenterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hackcenter.CenterService",
	HandlerType: (*CenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitBlock",
			Handler:    _CenterService_SubmitBlock_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _CenterService_RegisterNode_Handler,
		},
		{
			MethodName: "FetchNode",
			Handler:    _CenterService_FetchNode_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _CenterService_Vote_Handler,
		},
		{
			MethodName: "BroadcastBlock",
			Handler:    _CenterService_BroadcastBlock_Handler,
		},
		{
			MethodName: "UpdateHack",
			Handler:    _CenterService_UpdateHack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeBlock",
			Handler:       _CenterService_SubscribeBlock_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubBroadcastTask",
			Handler:       _CenterService_SubBroadcastTask_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeMinedBlock",
			Handler:       _CenterService_SubscribeMinedBlock_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hackcenter/center.proto",
}
