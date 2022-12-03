// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: calendar/calendar.proto

package calendarpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Begin of event
	BeginTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	// End of event
	EndTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Description
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// User ID
	UserId string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Notify in
	NotifyIn *durationpb.Duration `protobuf:"bytes,7,opt,name=notify_in,json=notifyIn,proto3" json:"notify_in,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetBeginTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BeginTime
	}
	return nil
}

func (x *Event) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Event) GetNotifyIn() *durationpb.Duration {
	if x != nil {
		return x.NotifyIn
	}
	return nil
}

// Create event request
type CreateEventV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Title
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Begin of event
	BeginTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	// End of event
	EndTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Description
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// User ID
	UserId string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Notify in
	NotifyIn *durationpb.Duration `protobuf:"bytes,6,opt,name=notify_in,json=notifyIn,proto3" json:"notify_in,omitempty"`
}

func (x *CreateEventV1Request) Reset() {
	*x = CreateEventV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventV1Request) ProtoMessage() {}

func (x *CreateEventV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventV1Request.ProtoReflect.Descriptor instead.
func (*CreateEventV1Request) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventV1Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEventV1Request) GetBeginTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BeginTime
	}
	return nil
}

func (x *CreateEventV1Request) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *CreateEventV1Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEventV1Request) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateEventV1Request) GetNotifyIn() *durationpb.Duration {
	if x != nil {
		return x.NotifyIn
	}
	return nil
}

// Create event response
type CreateEventV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Created event ID
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *CreateEventV1Response) Reset() {
	*x = CreateEventV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventV1Response) ProtoMessage() {}

func (x *CreateEventV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventV1Response.ProtoReflect.Descriptor instead.
func (*CreateEventV1Response) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventV1Response) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

// Update event request
type UpdateEventV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *UpdateEventV1Request) Reset() {
	*x = UpdateEventV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventV1Request) ProtoMessage() {}

func (x *UpdateEventV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventV1Request.ProtoReflect.Descriptor instead.
func (*UpdateEventV1Request) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEventV1Request) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// Update event response
type UpdateEventV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEventV1Response) Reset() {
	*x = UpdateEventV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventV1Response) ProtoMessage() {}

func (x *UpdateEventV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventV1Response.ProtoReflect.Descriptor instead.
func (*UpdateEventV1Response) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{4}
}

// Delete event request
type DeleteEventV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event ID
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *DeleteEventV1Request) Reset() {
	*x = DeleteEventV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventV1Request) ProtoMessage() {}

func (x *DeleteEventV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventV1Request.ProtoReflect.Descriptor instead.
func (*DeleteEventV1Request) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEventV1Request) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

// Delete event response
type DeleteEventV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEventV1Response) Reset() {
	*x = DeleteEventV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventV1Response) ProtoMessage() {}

func (x *DeleteEventV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventV1Response.ProtoReflect.Descriptor instead.
func (*DeleteEventV1Response) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{6}
}

// Get events request
type GetEventsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The begin of requested period
	Since *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *GetEventsV1Request) Reset() {
	*x = GetEventsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsV1Request) ProtoMessage() {}

func (x *GetEventsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsV1Request.ProtoReflect.Descriptor instead.
func (*GetEventsV1Request) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventsV1Request) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

// Get events response
type GetEventsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested events
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsV1Response) Reset() {
	*x = GetEventsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_calendar_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsV1Response) ProtoMessage() {}

func (x *GetEventsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_calendar_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsV1Response.ProtoReflect.Descriptor instead.
func (*GetEventsV1Response) Descriptor() ([]byte, []int) {
	return file_calendar_calendar_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventsV1Response) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_calendar_calendar_proto protoreflect.FileDescriptor

var file_calendar_calendar_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x92, 0x41, 0x07, 0xa2, 0x02, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x92, 0x41, 0x07, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x5f, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x22,
	0x9d, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x92, 0x41, 0x07, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x5f, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x22,
	0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x92, 0x41, 0x07, 0xa2,
	0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x3d, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x17,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0x92, 0x41, 0x07, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xfd, 0x06, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x41, 0x70, 0x69, 0x12, 0xa5, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x37, 0x4a, 0x35, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12,
	0x2e, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x21, 0x1a, 0x1f,
	0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0xb0, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56,
	0x31, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x1a, 0x19, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x37, 0x4a, 0x35, 0x0a, 0x03, 0x32,
	0x30, 0x34, 0x12, 0x2e, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x21, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x56, 0x31, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x37, 0x4a, 0x35, 0x0a, 0x03, 0x32,
	0x30, 0x34, 0x12, 0x2e, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x21, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4f,
	0x66, 0x44, 0x61, 0x79, 0x56, 0x31, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x61, 0x79, 0x2f,
	0x7b, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x7d, 0x12, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x56, 0x31, 0x12, 0x1c, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x2f, 0x7b, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x7d, 0x12, 0x77,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x56, 0x31, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x2f,
	0x7b, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x7d, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x73, 0x6b, 0x2d, 0x2f, 0x6f, 0x74, 0x75, 0x73, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x68, 0x77, 0x31, 0x32, 0x5f, 0x31, 0x33, 0x5f, 0x31, 0x34, 0x5f, 0x31, 0x35, 0x5f, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calendar_calendar_proto_rawDescOnce sync.Once
	file_calendar_calendar_proto_rawDescData = file_calendar_calendar_proto_rawDesc
)

func file_calendar_calendar_proto_rawDescGZIP() []byte {
	file_calendar_calendar_proto_rawDescOnce.Do(func() {
		file_calendar_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_calendar_calendar_proto_rawDescData)
	})
	return file_calendar_calendar_proto_rawDescData
}

var file_calendar_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_calendar_calendar_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: calendar.Event
	(*CreateEventV1Request)(nil),  // 1: calendar.CreateEventV1Request
	(*CreateEventV1Response)(nil), // 2: calendar.CreateEventV1Response
	(*UpdateEventV1Request)(nil),  // 3: calendar.UpdateEventV1Request
	(*UpdateEventV1Response)(nil), // 4: calendar.UpdateEventV1Response
	(*DeleteEventV1Request)(nil),  // 5: calendar.DeleteEventV1Request
	(*DeleteEventV1Response)(nil), // 6: calendar.DeleteEventV1Response
	(*GetEventsV1Request)(nil),    // 7: calendar.GetEventsV1Request
	(*GetEventsV1Response)(nil),   // 8: calendar.GetEventsV1Response
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 10: google.protobuf.Duration
}
var file_calendar_calendar_proto_depIdxs = []int32{
	9,  // 0: calendar.Event.begin_time:type_name -> google.protobuf.Timestamp
	9,  // 1: calendar.Event.end_time:type_name -> google.protobuf.Timestamp
	10, // 2: calendar.Event.notify_in:type_name -> google.protobuf.Duration
	9,  // 3: calendar.CreateEventV1Request.begin_time:type_name -> google.protobuf.Timestamp
	9,  // 4: calendar.CreateEventV1Request.end_time:type_name -> google.protobuf.Timestamp
	10, // 5: calendar.CreateEventV1Request.notify_in:type_name -> google.protobuf.Duration
	0,  // 6: calendar.UpdateEventV1Request.event:type_name -> calendar.Event
	9,  // 7: calendar.GetEventsV1Request.since:type_name -> google.protobuf.Timestamp
	0,  // 8: calendar.GetEventsV1Response.events:type_name -> calendar.Event
	1,  // 9: calendar.CalendarApi.CreateEventV1:input_type -> calendar.CreateEventV1Request
	3,  // 10: calendar.CalendarApi.UpdateEventV1:input_type -> calendar.UpdateEventV1Request
	5,  // 11: calendar.CalendarApi.DeleteEventV1:input_type -> calendar.DeleteEventV1Request
	7,  // 12: calendar.CalendarApi.GetEventsOfDayV1:input_type -> calendar.GetEventsV1Request
	7,  // 13: calendar.CalendarApi.GetEventsOfWeekV1:input_type -> calendar.GetEventsV1Request
	7,  // 14: calendar.CalendarApi.GetEventsOfMonthV1:input_type -> calendar.GetEventsV1Request
	2,  // 15: calendar.CalendarApi.CreateEventV1:output_type -> calendar.CreateEventV1Response
	4,  // 16: calendar.CalendarApi.UpdateEventV1:output_type -> calendar.UpdateEventV1Response
	6,  // 17: calendar.CalendarApi.DeleteEventV1:output_type -> calendar.DeleteEventV1Response
	8,  // 18: calendar.CalendarApi.GetEventsOfDayV1:output_type -> calendar.GetEventsV1Response
	8,  // 19: calendar.CalendarApi.GetEventsOfWeekV1:output_type -> calendar.GetEventsV1Response
	8,  // 20: calendar.CalendarApi.GetEventsOfMonthV1:output_type -> calendar.GetEventsV1Response
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_calendar_calendar_proto_init() }
func file_calendar_calendar_proto_init() {
	if File_calendar_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calendar_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calendar_calendar_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calendar_calendar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calendar_calendar_proto_goTypes,
		DependencyIndexes: file_calendar_calendar_proto_depIdxs,
		MessageInfos:      file_calendar_calendar_proto_msgTypes,
	}.Build()
	File_calendar_calendar_proto = out.File
	file_calendar_calendar_proto_rawDesc = nil
	file_calendar_calendar_proto_goTypes = nil
	file_calendar_calendar_proto_depIdxs = nil
}
