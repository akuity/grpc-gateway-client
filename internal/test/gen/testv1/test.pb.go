// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: testv1/test.proto

package testv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_EVENT_TYPE_UNDEFINED EventType = 0
	EventType_EVENT_TYPE_SEEN      EventType = 1
	EventType_EVENT_TYPE_ACCEPTED  EventType = 2
	EventType_EVENT_TYPE_REJECTED  EventType = 3
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNDEFINED",
		1: "EVENT_TYPE_SEEN",
		2: "EVENT_TYPE_ACCEPTED",
		3: "EVENT_TYPE_REJECTED",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNDEFINED": 0,
		"EVENT_TYPE_SEEN":      1,
		"EVENT_TYPE_ACCEPTED":  2,
		"EVENT_TYPE_REJECTED":  3,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_testv1_test_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_testv1_test_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{0}
}

type InvitationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw map[string]string `protobuf:"bytes,1,rep,name=raw,proto3" json:"raw,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InvitationMetadata) Reset() {
	*x = InvitationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationMetadata) ProtoMessage() {}

func (x *InvitationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationMetadata.ProtoReflect.Descriptor instead.
func (*InvitationMetadata) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{0}
}

func (x *InvitationMetadata) GetRaw() map[string]string {
	if x != nil {
		return x.Raw
	}
	return nil
}

type Invitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{1}
}

func (x *Invitation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invitation) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type ListInvitationsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListInvitationsQuery) Reset() {
	*x = ListInvitationsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitationsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitationsQuery) ProtoMessage() {}

func (x *ListInvitationsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitationsQuery.ProtoReflect.Descriptor instead.
func (*ListInvitationsQuery) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{2}
}

func (x *ListInvitationsQuery) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type ListInvitationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *ListInvitationsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ListInvitationsRequest) Reset() {
	*x = ListInvitationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitationsRequest) ProtoMessage() {}

func (x *ListInvitationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitationsRequest.ProtoReflect.Descriptor instead.
func (*ListInvitationsRequest) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{3}
}

func (x *ListInvitationsRequest) GetQuery() *ListInvitationsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type ListInvitationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invitations []*Invitation `protobuf:"bytes,1,rep,name=invitations,proto3" json:"invitations,omitempty"`
}

func (x *ListInvitationsResponse) Reset() {
	*x = ListInvitationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitationsResponse) ProtoMessage() {}

func (x *ListInvitationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitationsResponse.ProtoReflect.Descriptor instead.
func (*ListInvitationsResponse) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{4}
}

func (x *ListInvitationsResponse) GetInvitations() []*Invitation {
	if x != nil {
		return x.Invitations
	}
	return nil
}

type SendInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendInvitationRequest) Reset() {
	*x = SendInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendInvitationRequest) ProtoMessage() {}

func (x *SendInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendInvitationRequest.ProtoReflect.Descriptor instead.
func (*SendInvitationRequest) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{5}
}

func (x *SendInvitationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendInvitationResponse) Reset() {
	*x = SendInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendInvitationResponse) ProtoMessage() {}

func (x *SendInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendInvitationResponse.ProtoReflect.Descriptor instead.
func (*SendInvitationResponse) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{6}
}

func (x *SendInvitationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TrackInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type *EventType `protobuf:"varint,2,opt,name=type,proto3,enum=io.akuity.test.v1.EventType,oneof" json:"type,omitempty"`
}

func (x *TrackInvitationRequest) Reset() {
	*x = TrackInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackInvitationRequest) ProtoMessage() {}

func (x *TrackInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackInvitationRequest.ProtoReflect.Descriptor instead.
func (*TrackInvitationRequest) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{7}
}

func (x *TrackInvitationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackInvitationRequest) GetType() EventType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return EventType_EVENT_TYPE_UNDEFINED
}

type TrackInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    EventType `protobuf:"varint,1,opt,name=type,proto3,enum=io.akuity.test.v1.EventType" json:"type,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TrackInvitationResponse) Reset() {
	*x = TrackInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackInvitationResponse) ProtoMessage() {}

func (x *TrackInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackInvitationResponse.ProtoReflect.Descriptor instead.
func (*TrackInvitationResponse) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{8}
}

func (x *TrackInvitationResponse) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_TYPE_UNDEFINED
}

func (x *TrackInvitationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DownloadInvitationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *EventType `protobuf:"varint,1,opt,name=type,proto3,enum=io.akuity.test.v1.EventType,oneof" json:"type,omitempty"`
}

func (x *DownloadInvitationsRequest) Reset() {
	*x = DownloadInvitationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadInvitationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadInvitationsRequest) ProtoMessage() {}

func (x *DownloadInvitationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadInvitationsRequest.ProtoReflect.Descriptor instead.
func (*DownloadInvitationsRequest) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{9}
}

func (x *DownloadInvitationsRequest) GetType() EventType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return EventType_EVENT_TYPE_UNDEFINED
}

type DownloadLargeFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DownloadLargeFileRequest) Reset() {
	*x = DownloadLargeFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testv1_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadLargeFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadLargeFileRequest) ProtoMessage() {}

func (x *DownloadLargeFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testv1_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadLargeFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadLargeFileRequest) Descriptor() ([]byte, []int) {
	return file_testv1_test_proto_rawDescGZIP(), []int{10}
}

var File_testv1_test_proto protoreflect.FileDescriptor

var file_testv1_test_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8e, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x61, 0x77, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x03, 0x72, 0x61, 0x77, 0x1a, 0x36, 0x0a, 0x08, 0x52, 0x61, 0x77, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x9a, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x41, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9e, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69,
	0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75,
	0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x5a, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75,
	0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x16,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x65, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a,
	0x1a, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x6c, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x45,
	0x4e, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x32, 0x88, 0x05, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b,
	0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7d, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75,
	0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b,
	0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x30, 0x01, 0x12, 0x7b, 0x0a, 0x13, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12,
	0x15, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x30, 0x01, 0x12, 0x76, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x2d, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x30, 0x01,
	0x42, 0xd0, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x6b, 0x75, 0x69,
	0x74, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x49, 0x41, 0x54, 0xaa, 0x02, 0x11, 0x49, 0x6f, 0x2e, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x49, 0x6f, 0x5c, 0x41, 0x6b,
	0x75, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x49,
	0x6f, 0x5c, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x49,
	0x6f, 0x3a, 0x3a, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testv1_test_proto_rawDescOnce sync.Once
	file_testv1_test_proto_rawDescData = file_testv1_test_proto_rawDesc
)

func file_testv1_test_proto_rawDescGZIP() []byte {
	file_testv1_test_proto_rawDescOnce.Do(func() {
		file_testv1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_testv1_test_proto_rawDescData)
	})
	return file_testv1_test_proto_rawDescData
}

var file_testv1_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testv1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_testv1_test_proto_goTypes = []interface{}{
	(EventType)(0),                     // 0: io.akuity.test.v1.EventType
	(*InvitationMetadata)(nil),         // 1: io.akuity.test.v1.InvitationMetadata
	(*Invitation)(nil),                 // 2: io.akuity.test.v1.Invitation
	(*ListInvitationsQuery)(nil),       // 3: io.akuity.test.v1.ListInvitationsQuery
	(*ListInvitationsRequest)(nil),     // 4: io.akuity.test.v1.ListInvitationsRequest
	(*ListInvitationsResponse)(nil),    // 5: io.akuity.test.v1.ListInvitationsResponse
	(*SendInvitationRequest)(nil),      // 6: io.akuity.test.v1.SendInvitationRequest
	(*SendInvitationResponse)(nil),     // 7: io.akuity.test.v1.SendInvitationResponse
	(*TrackInvitationRequest)(nil),     // 8: io.akuity.test.v1.TrackInvitationRequest
	(*TrackInvitationResponse)(nil),    // 9: io.akuity.test.v1.TrackInvitationResponse
	(*DownloadInvitationsRequest)(nil), // 10: io.akuity.test.v1.DownloadInvitationsRequest
	(*DownloadLargeFileRequest)(nil),   // 11: io.akuity.test.v1.DownloadLargeFileRequest
	nil,                                // 12: io.akuity.test.v1.InvitationMetadata.RawEntry
	nil,                                // 13: io.akuity.test.v1.Invitation.LabelsEntry
	nil,                                // 14: io.akuity.test.v1.ListInvitationsQuery.LabelsEntry
	(*httpbody.HttpBody)(nil),          // 15: google.api.HttpBody
}
var file_testv1_test_proto_depIdxs = []int32{
	12, // 0: io.akuity.test.v1.InvitationMetadata.raw:type_name -> io.akuity.test.v1.InvitationMetadata.RawEntry
	13, // 1: io.akuity.test.v1.Invitation.labels:type_name -> io.akuity.test.v1.Invitation.LabelsEntry
	14, // 2: io.akuity.test.v1.ListInvitationsQuery.labels:type_name -> io.akuity.test.v1.ListInvitationsQuery.LabelsEntry
	3,  // 3: io.akuity.test.v1.ListInvitationsRequest.query:type_name -> io.akuity.test.v1.ListInvitationsQuery
	2,  // 4: io.akuity.test.v1.ListInvitationsResponse.invitations:type_name -> io.akuity.test.v1.Invitation
	0,  // 5: io.akuity.test.v1.TrackInvitationRequest.type:type_name -> io.akuity.test.v1.EventType
	0,  // 6: io.akuity.test.v1.TrackInvitationResponse.type:type_name -> io.akuity.test.v1.EventType
	0,  // 7: io.akuity.test.v1.DownloadInvitationsRequest.type:type_name -> io.akuity.test.v1.EventType
	4,  // 8: io.akuity.test.v1.TestService.ListInvitations:input_type -> io.akuity.test.v1.ListInvitationsRequest
	6,  // 9: io.akuity.test.v1.TestService.SendInvitation:input_type -> io.akuity.test.v1.SendInvitationRequest
	8,  // 10: io.akuity.test.v1.TestService.TrackInvitation:input_type -> io.akuity.test.v1.TrackInvitationRequest
	10, // 11: io.akuity.test.v1.TestService.DownloadInvitations:input_type -> io.akuity.test.v1.DownloadInvitationsRequest
	11, // 12: io.akuity.test.v1.TestService.DownloadLargeFile:input_type -> io.akuity.test.v1.DownloadLargeFileRequest
	5,  // 13: io.akuity.test.v1.TestService.ListInvitations:output_type -> io.akuity.test.v1.ListInvitationsResponse
	7,  // 14: io.akuity.test.v1.TestService.SendInvitation:output_type -> io.akuity.test.v1.SendInvitationResponse
	9,  // 15: io.akuity.test.v1.TestService.TrackInvitation:output_type -> io.akuity.test.v1.TrackInvitationResponse
	15, // 16: io.akuity.test.v1.TestService.DownloadInvitations:output_type -> google.api.HttpBody
	15, // 17: io.akuity.test.v1.TestService.DownloadLargeFile:output_type -> google.api.HttpBody
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_testv1_test_proto_init() }
func file_testv1_test_proto_init() {
	if File_testv1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testv1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationMetadata); i {
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
		file_testv1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invitation); i {
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
		file_testv1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitationsQuery); i {
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
		file_testv1_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitationsRequest); i {
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
		file_testv1_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitationsResponse); i {
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
		file_testv1_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendInvitationRequest); i {
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
		file_testv1_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendInvitationResponse); i {
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
		file_testv1_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackInvitationRequest); i {
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
		file_testv1_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackInvitationResponse); i {
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
		file_testv1_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadInvitationsRequest); i {
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
		file_testv1_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadLargeFileRequest); i {
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
	file_testv1_test_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_testv1_test_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testv1_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testv1_test_proto_goTypes,
		DependencyIndexes: file_testv1_test_proto_depIdxs,
		EnumInfos:         file_testv1_test_proto_enumTypes,
		MessageInfos:      file_testv1_test_proto_msgTypes,
	}.Build()
	File_testv1_test_proto = out.File
	file_testv1_test_proto_rawDesc = nil
	file_testv1_test_proto_goTypes = nil
	file_testv1_test_proto_depIdxs = nil
}
