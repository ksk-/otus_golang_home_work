// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: calendar/calendar.proto

package calendarpb

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

// CalendarApiClient is the client API for CalendarApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarApiClient interface {
	// Create a new event
	CreateEventV1(ctx context.Context, in *CreateEventV1Request, opts ...grpc.CallOption) (*CreateEventV1Response, error)
	// Update event
	UpdateEventV1(ctx context.Context, in *UpdateEventV1Request, opts ...grpc.CallOption) (*UpdateEventV1Response, error)
	// Delete event
	DeleteEventV1(ctx context.Context, in *DeleteEventV1Request, opts ...grpc.CallOption) (*DeleteEventV1Response, error)
	// Get event for the specified date
	GetEventsOfDayV1(ctx context.Context, in *GetEventsV1Request, opts ...grpc.CallOption) (*GetEventsV1Response, error)
	// Get event for the specified week
	GetEventsOfWeekV1(ctx context.Context, in *GetEventsV1Request, opts ...grpc.CallOption) (*GetEventsV1Response, error)
	// Get event for the specified month
	GetEventsOfMonthV1(ctx context.Context, in *GetEventsV1Request, opts ...grpc.CallOption) (*GetEventsV1Response, error)
}

type calendarApiClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarApiClient(cc grpc.ClientConnInterface) CalendarApiClient {
	return &calendarApiClient{cc}
}

func (c *calendarApiClient) CreateEventV1(ctx context.Context, in *CreateEventV1Request, opts ...grpc.CallOption) (*CreateEventV1Response, error) {
	out := new(CreateEventV1Response)
	err := c.cc.Invoke(ctx, "/calendar.CalendarApi/CreateEventV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) UpdateEventV1(ctx context.Context, in *UpdateEventV1Request, opts ...grpc.CallOption) (*UpdateEventV1Response, error) {
	out := new(UpdateEventV1Response)
	err := c.cc.Invoke(ctx, "/calendar.CalendarApi/UpdateEventV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) DeleteEventV1(ctx context.Context, in *DeleteEventV1Request, opts ...grpc.CallOption) (*DeleteEventV1Response, error) {
	out := new(DeleteEventV1Response)
	err := c.cc.Invoke(ctx, "/calendar.CalendarApi/DeleteEventV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) GetEventsOfDayV1(ctx context.Context, in *GetEventsV1Request, opts ...grpc.CallOption) (*GetEventsV1Response, error) {
	out := new(GetEventsV1Response)
	err := c.cc.Invoke(ctx, "/calendar.CalendarApi/GetEventsOfDayV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) GetEventsOfWeekV1(ctx context.Context, in *GetEventsV1Request, opts ...grpc.CallOption) (*GetEventsV1Response, error) {
	out := new(GetEventsV1Response)
	err := c.cc.Invoke(ctx, "/calendar.CalendarApi/GetEventsOfWeekV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) GetEventsOfMonthV1(ctx context.Context, in *GetEventsV1Request, opts ...grpc.CallOption) (*GetEventsV1Response, error) {
	out := new(GetEventsV1Response)
	err := c.cc.Invoke(ctx, "/calendar.CalendarApi/GetEventsOfMonthV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarApiServer is the server API for CalendarApi service.
// All implementations must embed UnimplementedCalendarApiServer
// for forward compatibility
type CalendarApiServer interface {
	// Create a new event
	CreateEventV1(context.Context, *CreateEventV1Request) (*CreateEventV1Response, error)
	// Update event
	UpdateEventV1(context.Context, *UpdateEventV1Request) (*UpdateEventV1Response, error)
	// Delete event
	DeleteEventV1(context.Context, *DeleteEventV1Request) (*DeleteEventV1Response, error)
	// Get event for the specified date
	GetEventsOfDayV1(context.Context, *GetEventsV1Request) (*GetEventsV1Response, error)
	// Get event for the specified week
	GetEventsOfWeekV1(context.Context, *GetEventsV1Request) (*GetEventsV1Response, error)
	// Get event for the specified month
	GetEventsOfMonthV1(context.Context, *GetEventsV1Request) (*GetEventsV1Response, error)
	mustEmbedUnimplementedCalendarApiServer()
}

// UnimplementedCalendarApiServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarApiServer struct {
}

func (UnimplementedCalendarApiServer) CreateEventV1(context.Context, *CreateEventV1Request) (*CreateEventV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventV1 not implemented")
}
func (UnimplementedCalendarApiServer) UpdateEventV1(context.Context, *UpdateEventV1Request) (*UpdateEventV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEventV1 not implemented")
}
func (UnimplementedCalendarApiServer) DeleteEventV1(context.Context, *DeleteEventV1Request) (*DeleteEventV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEventV1 not implemented")
}
func (UnimplementedCalendarApiServer) GetEventsOfDayV1(context.Context, *GetEventsV1Request) (*GetEventsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsOfDayV1 not implemented")
}
func (UnimplementedCalendarApiServer) GetEventsOfWeekV1(context.Context, *GetEventsV1Request) (*GetEventsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsOfWeekV1 not implemented")
}
func (UnimplementedCalendarApiServer) GetEventsOfMonthV1(context.Context, *GetEventsV1Request) (*GetEventsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsOfMonthV1 not implemented")
}
func (UnimplementedCalendarApiServer) mustEmbedUnimplementedCalendarApiServer() {}

// UnsafeCalendarApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarApiServer will
// result in compilation errors.
type UnsafeCalendarApiServer interface {
	mustEmbedUnimplementedCalendarApiServer()
}

func RegisterCalendarApiServer(s grpc.ServiceRegistrar, srv CalendarApiServer) {
	s.RegisterService(&CalendarApi_ServiceDesc, srv)
}

func _CalendarApi_CreateEventV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).CreateEventV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarApi/CreateEventV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).CreateEventV1(ctx, req.(*CreateEventV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_UpdateEventV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).UpdateEventV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarApi/UpdateEventV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).UpdateEventV1(ctx, req.(*UpdateEventV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_DeleteEventV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).DeleteEventV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarApi/DeleteEventV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).DeleteEventV1(ctx, req.(*DeleteEventV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_GetEventsOfDayV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).GetEventsOfDayV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarApi/GetEventsOfDayV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).GetEventsOfDayV1(ctx, req.(*GetEventsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_GetEventsOfWeekV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).GetEventsOfWeekV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarApi/GetEventsOfWeekV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).GetEventsOfWeekV1(ctx, req.(*GetEventsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_GetEventsOfMonthV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).GetEventsOfMonthV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarApi/GetEventsOfMonthV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).GetEventsOfMonthV1(ctx, req.(*GetEventsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarApi_ServiceDesc is the grpc.ServiceDesc for CalendarApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.CalendarApi",
	HandlerType: (*CalendarApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEventV1",
			Handler:    _CalendarApi_CreateEventV1_Handler,
		},
		{
			MethodName: "UpdateEventV1",
			Handler:    _CalendarApi_UpdateEventV1_Handler,
		},
		{
			MethodName: "DeleteEventV1",
			Handler:    _CalendarApi_DeleteEventV1_Handler,
		},
		{
			MethodName: "GetEventsOfDayV1",
			Handler:    _CalendarApi_GetEventsOfDayV1_Handler,
		},
		{
			MethodName: "GetEventsOfWeekV1",
			Handler:    _CalendarApi_GetEventsOfWeekV1_Handler,
		},
		{
			MethodName: "GetEventsOfMonthV1",
			Handler:    _CalendarApi_GetEventsOfMonthV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar/calendar.proto",
}
