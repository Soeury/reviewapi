// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: api/operation/v1/operation.proto

package v1

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

// OperationClient is the client API for Operation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationClient interface {
	// 运营审核评价
	AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...grpc.CallOption) (*AuditReviewReply, error)
	// 运营审核商家申诉
	AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...grpc.CallOption) (*AuditAppealReply, error)
}

type operationClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationClient(cc grpc.ClientConnInterface) OperationClient {
	return &operationClient{cc}
}

func (c *operationClient) AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...grpc.CallOption) (*AuditReviewReply, error) {
	out := new(AuditReviewReply)
	err := c.cc.Invoke(ctx, "/api.operation.v1.Operation/AuditReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...grpc.CallOption) (*AuditAppealReply, error) {
	out := new(AuditAppealReply)
	err := c.cc.Invoke(ctx, "/api.operation.v1.Operation/AuditAppeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServer is the server API for Operation service.
// All implementations must embed UnimplementedOperationServer
// for forward compatibility
type OperationServer interface {
	// 运营审核评价
	AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error)
	// 运营审核商家申诉
	AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error)
	mustEmbedUnimplementedOperationServer()
}

// UnimplementedOperationServer must be embedded to have forward compatible implementations.
type UnimplementedOperationServer struct {
}

func (UnimplementedOperationServer) AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditReview not implemented")
}
func (UnimplementedOperationServer) AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditAppeal not implemented")
}
func (UnimplementedOperationServer) mustEmbedUnimplementedOperationServer() {}

// UnsafeOperationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServer will
// result in compilation errors.
type UnsafeOperationServer interface {
	mustEmbedUnimplementedOperationServer()
}

func RegisterOperationServer(s grpc.ServiceRegistrar, srv OperationServer) {
	s.RegisterService(&Operation_ServiceDesc, srv)
}

func _Operation_AuditReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).AuditReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.operation.v1.Operation/AuditReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).AuditReview(ctx, req.(*AuditReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_AuditAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).AuditAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.operation.v1.Operation/AuditAppeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).AuditAppeal(ctx, req.(*AuditAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operation_ServiceDesc is the grpc.ServiceDesc for Operation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.operation.v1.Operation",
	HandlerType: (*OperationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuditReview",
			Handler:    _Operation_AuditReview_Handler,
		},
		{
			MethodName: "AuditAppeal",
			Handler:    _Operation_AuditAppeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/operation/v1/operation.proto",
}
