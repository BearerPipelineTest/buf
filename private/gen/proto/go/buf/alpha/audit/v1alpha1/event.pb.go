// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: buf/alpha/audit/v1alpha1/event.proto

package auditv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// ActorType is the type of actor that caused the audited event.
type ActorType int32

const (
	ActorType_ACTOR_TYPE_UNSPECIFIED ActorType = 0
	ActorType_ACTOR_TYPE_USER        ActorType = 1
)

// Enum value maps for ActorType.
var (
	ActorType_name = map[int32]string{
		0: "ACTOR_TYPE_UNSPECIFIED",
		1: "ACTOR_TYPE_USER",
	}
	ActorType_value = map[string]int32{
		"ACTOR_TYPE_UNSPECIFIED": 0,
		"ACTOR_TYPE_USER":        1,
	}
)

func (x ActorType) Enum() *ActorType {
	p := new(ActorType)
	*p = x
	return p
}

func (x ActorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActorType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_audit_v1alpha1_event_proto_enumTypes[0].Descriptor()
}

func (ActorType) Type() protoreflect.EnumType {
	return &file_buf_alpha_audit_v1alpha1_event_proto_enumTypes[0]
}

func (x ActorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActorType.Descriptor instead.
func (ActorType) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{0}
}

// ResourceType is the type of the resource that was affected by the audited
// event.
type ResourceType int32

const (
	ResourceType_RESOURCE_TYPE_UNSPECIFIED         ResourceType = 0
	ResourceType_RESOURCE_TYPE_USER                ResourceType = 1
	ResourceType_RESOURCE_TYPE_ORGANIZATION        ResourceType = 2
	ResourceType_RESOURCE_TYPE_ORGANIZATION_MEMBER ResourceType = 3
	ResourceType_RESOURCE_TYPE_REPOSITORY          ResourceType = 4
	ResourceType_RESOURCE_TYPE_REPOSITORY_MEMBER   ResourceType = 5
	ResourceType_RESOURCE_TYPE_REPOSITORY_COMMIT   ResourceType = 6
	ResourceType_RESOURCE_TYPE_PLUGIN              ResourceType = 7
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RESOURCE_TYPE_UNSPECIFIED",
		1: "RESOURCE_TYPE_USER",
		2: "RESOURCE_TYPE_ORGANIZATION",
		3: "RESOURCE_TYPE_ORGANIZATION_MEMBER",
		4: "RESOURCE_TYPE_REPOSITORY",
		5: "RESOURCE_TYPE_REPOSITORY_MEMBER",
		6: "RESOURCE_TYPE_REPOSITORY_COMMIT",
		7: "RESOURCE_TYPE_PLUGIN",
	}
	ResourceType_value = map[string]int32{
		"RESOURCE_TYPE_UNSPECIFIED":         0,
		"RESOURCE_TYPE_USER":                1,
		"RESOURCE_TYPE_ORGANIZATION":        2,
		"RESOURCE_TYPE_ORGANIZATION_MEMBER": 3,
		"RESOURCE_TYPE_REPOSITORY":          4,
		"RESOURCE_TYPE_REPOSITORY_MEMBER":   5,
		"RESOURCE_TYPE_REPOSITORY_COMMIT":   6,
		"RESOURCE_TYPE_PLUGIN":              7,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_audit_v1alpha1_event_proto_enumTypes[1].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_buf_alpha_audit_v1alpha1_event_proto_enumTypes[1]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{1}
}

// EventType is the type of audited event.
type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED          EventType = 0
	EventType_EVENT_TYPE_ORGANIZATION_CREATED EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "EVENT_TYPE_ORGANIZATION_CREATED",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED":          0,
		"EVENT_TYPE_ORGANIZATION_CREATED": 1,
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
	return file_buf_alpha_audit_v1alpha1_event_proto_enumTypes[2].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_buf_alpha_audit_v1alpha1_event_proto_enumTypes[2]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{2}
}

// Actor is the actor who caused the audited event.
type Actor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of actor who caused the audited event.
	Type ActorType `protobuf:"varint,1,opt,name=type,proto3,enum=buf.alpha.audit.v1alpha1.ActorType" json:"type,omitempty"`
	// Id of the actor who caused the audited event.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the actor who caused the audited event.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Actor) Reset() {
	*x = Actor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actor) ProtoMessage() {}

func (x *Actor) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actor.ProtoReflect.Descriptor instead.
func (*Actor) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Actor) GetType() ActorType {
	if x != nil {
		return x.Type
	}
	return ActorType_ACTOR_TYPE_UNSPECIFIED
}

func (x *Actor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Actor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Resource is the affected resource by the audited event.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of resource that was affected by the audited event.
	Type ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=buf.alpha.audit.v1alpha1.ResourceType" json:"type,omitempty"`
	// Id of the affected resource by the audited event.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the affected resource by the audited event.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetType() ResourceType {
	if x != nil {
		return x.Type
	}
	return ResourceType_RESOURCE_TYPE_UNSPECIFIED
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Event is an audited action that happened in the BSR, with the information of
// what happened, when it happened, who did it, which resource was affected, and
// more contextual information on the event.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id of the audited event.
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Type of the audited event. It specifies "what" happened.
	Type EventType `protobuf:"varint,2,opt,name=type,proto3,enum=buf.alpha.audit.v1alpha1.EventType" json:"type,omitempty"`
	// Actor of the audited event. It specifies "who" did it.
	Actor *Actor `protobuf:"bytes,3,opt,name=actor,proto3" json:"actor,omitempty"`
	// Resource of the audited event. It specifies "which resource" was affected.
	Resource *Resource `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	// Time of the audited event. It specifies "when" it happened.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Payload of the audited event. It specifies additional context on the event.
	//
	// Types that are assignable to Payload:
	//	*Event_OrganizationCreated
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[2]
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
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (x *Event) GetActor() *Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *Event) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Event) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetOrganizationCreated() *PayloadOrganizationCreated {
	if x, ok := x.GetPayload().(*Event_OrganizationCreated); ok {
		return x.OrganizationCreated
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_OrganizationCreated struct {
	OrganizationCreated *PayloadOrganizationCreated `protobuf:"bytes,6,opt,name=organization_created,json=organizationCreated,proto3,oneof"`
}

func (*Event_OrganizationCreated) isEvent_Payload() {}

// PayloadOrganizationCreated is the payload for an ORGANIZATION_CREATED event.
type PayloadOrganizationCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// created_time is when the organization was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *PayloadOrganizationCreated) Reset() {
	*x = PayloadOrganizationCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadOrganizationCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadOrganizationCreated) ProtoMessage() {}

func (x *PayloadOrganizationCreated) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadOrganizationCreated.ProtoReflect.Descriptor instead.
func (*PayloadOrganizationCreated) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadOrganizationCreated) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

var File_buf_alpha_audit_v1alpha1_event_proto protoreflect.FileDescriptor

var file_buf_alpha_audit_v1alpha1_event_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x64, 0x0a, 0x05, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x83, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x69, 0x0a, 0x14, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x13, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x09,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5b, 0x0a, 0x1a, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x3c, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x01, 0x2a, 0x8e, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52,
	0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52,
	0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x59, 0x10,
	0x04, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x45,
	0x4d, 0x42, 0x45, 0x52, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f,
	0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x55,
	0x47, 0x49, 0x4e, 0x10, 0x07, 0x2a, 0x4c, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23,
	0x0a, 0x1f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x42, 0x82, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x41, 0xaa, 0x02, 0x18,
	0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18, 0x42, 0x75, 0x66, 0x5c, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x42, 0x75, 0x66,
	0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_audit_v1alpha1_event_proto_rawDescOnce sync.Once
	file_buf_alpha_audit_v1alpha1_event_proto_rawDescData = file_buf_alpha_audit_v1alpha1_event_proto_rawDesc
)

func file_buf_alpha_audit_v1alpha1_event_proto_rawDescGZIP() []byte {
	file_buf_alpha_audit_v1alpha1_event_proto_rawDescOnce.Do(func() {
		file_buf_alpha_audit_v1alpha1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_audit_v1alpha1_event_proto_rawDescData)
	})
	return file_buf_alpha_audit_v1alpha1_event_proto_rawDescData
}

var file_buf_alpha_audit_v1alpha1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_buf_alpha_audit_v1alpha1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_alpha_audit_v1alpha1_event_proto_goTypes = []interface{}{
	(ActorType)(0),                     // 0: buf.alpha.audit.v1alpha1.ActorType
	(ResourceType)(0),                  // 1: buf.alpha.audit.v1alpha1.ResourceType
	(EventType)(0),                     // 2: buf.alpha.audit.v1alpha1.EventType
	(*Actor)(nil),                      // 3: buf.alpha.audit.v1alpha1.Actor
	(*Resource)(nil),                   // 4: buf.alpha.audit.v1alpha1.Resource
	(*Event)(nil),                      // 5: buf.alpha.audit.v1alpha1.Event
	(*PayloadOrganizationCreated)(nil), // 6: buf.alpha.audit.v1alpha1.PayloadOrganizationCreated
	(*timestamppb.Timestamp)(nil),      // 7: google.protobuf.Timestamp
}
var file_buf_alpha_audit_v1alpha1_event_proto_depIdxs = []int32{
	0, // 0: buf.alpha.audit.v1alpha1.Actor.type:type_name -> buf.alpha.audit.v1alpha1.ActorType
	1, // 1: buf.alpha.audit.v1alpha1.Resource.type:type_name -> buf.alpha.audit.v1alpha1.ResourceType
	2, // 2: buf.alpha.audit.v1alpha1.Event.type:type_name -> buf.alpha.audit.v1alpha1.EventType
	3, // 3: buf.alpha.audit.v1alpha1.Event.actor:type_name -> buf.alpha.audit.v1alpha1.Actor
	4, // 4: buf.alpha.audit.v1alpha1.Event.resource:type_name -> buf.alpha.audit.v1alpha1.Resource
	7, // 5: buf.alpha.audit.v1alpha1.Event.event_time:type_name -> google.protobuf.Timestamp
	6, // 6: buf.alpha.audit.v1alpha1.Event.organization_created:type_name -> buf.alpha.audit.v1alpha1.PayloadOrganizationCreated
	7, // 7: buf.alpha.audit.v1alpha1.PayloadOrganizationCreated.created_time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_buf_alpha_audit_v1alpha1_event_proto_init() }
func file_buf_alpha_audit_v1alpha1_event_proto_init() {
	if File_buf_alpha_audit_v1alpha1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actor); i {
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
		file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadOrganizationCreated); i {
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
	file_buf_alpha_audit_v1alpha1_event_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Event_OrganizationCreated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_audit_v1alpha1_event_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_audit_v1alpha1_event_proto_goTypes,
		DependencyIndexes: file_buf_alpha_audit_v1alpha1_event_proto_depIdxs,
		EnumInfos:         file_buf_alpha_audit_v1alpha1_event_proto_enumTypes,
		MessageInfos:      file_buf_alpha_audit_v1alpha1_event_proto_msgTypes,
	}.Build()
	File_buf_alpha_audit_v1alpha1_event_proto = out.File
	file_buf_alpha_audit_v1alpha1_event_proto_rawDesc = nil
	file_buf_alpha_audit_v1alpha1_event_proto_goTypes = nil
	file_buf_alpha_audit_v1alpha1_event_proto_depIdxs = nil
}
