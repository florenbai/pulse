// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.22.2
// source: integration.proto

package integration

import (
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

type IntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description" form:"description" query:"description"`
	LevelField  string `protobuf:"bytes,3,opt,name=levelField,proto3" json:"levelField" form:"levelField" query:"levelField"`
	Token       string `protobuf:"bytes,4,opt,name=token,proto3" json:"token" form:"token" query:"token"`
}

func (x *IntegrationRequest) Reset() {
	*x = IntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationRequest) ProtoMessage() {}

func (x *IntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationRequest.ProtoReflect.Descriptor instead.
func (*IntegrationRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{0}
}

func (x *IntegrationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IntegrationRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IntegrationRequest) GetLevelField() string {
	if x != nil {
		return x.LevelField
	}
	return ""
}

func (x *IntegrationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IntegrationRouterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces []uint64   `protobuf:"varint,1,rep,packed,name=workspaces,proto3" json:"workspaces" form:"workspaces" query:"workspaces"`
	Next       int64      `protobuf:"varint,2,opt,name=next,proto3" json:"next" form:"next" query:"next"`
	Filters    []*Filters `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters" form:"filters" query:"filters"`
	Sort       int64      `protobuf:"varint,4,opt,name=sort,proto3" json:"sort" form:"sort" query:"sort"`
}

func (x *IntegrationRouterRequest) Reset() {
	*x = IntegrationRouterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationRouterRequest) ProtoMessage() {}

func (x *IntegrationRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationRouterRequest.ProtoReflect.Descriptor instead.
func (*IntegrationRouterRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{1}
}

func (x *IntegrationRouterRequest) GetWorkspaces() []uint64 {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *IntegrationRouterRequest) GetNext() int64 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *IntegrationRouterRequest) GetFilters() []*Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *IntegrationRouterRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Filter `protobuf:"bytes,1,rep,name=values,proto3" json:"values" form:"values" query:"values"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{2}
}

func (x *Filters) GetValues() []*Filter {
	if x != nil {
		return x.Values
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag       string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag" form:"tag" query:"tag"`
	Operation string   `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation" form:"operation" query:"operation"`
	Value     []string `protobuf:"bytes,3,rep,name=value,proto3" json:"value" form:"value" query:"value"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Filter) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Filter) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type TagRewriteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldTag     *string  `protobuf:"bytes,1,opt,name=oldTag,proto3,oneof" json:"oldTag" form:"oldTag" query:"oldTag"`
	Expression *string  `protobuf:"bytes,2,opt,name=expression,proto3,oneof" json:"expression" form:"expression" query:"expression"`
	NewTag     *string  `protobuf:"bytes,3,opt,name=newTag,proto3,oneof" json:"newTag" form:"newTag" query:"newTag"`
	Value      *string  `protobuf:"bytes,4,opt,name=value,proto3,oneof" json:"value" form:"value" query:"value"`
	DeleteTag  []string `protobuf:"bytes,5,rep,name=deleteTag,proto3" json:"deleteTag" form:"deleteTag" query:"deleteTag"`
}

func (x *TagRewriteItem) Reset() {
	*x = TagRewriteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagRewriteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagRewriteItem) ProtoMessage() {}

func (x *TagRewriteItem) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagRewriteItem.ProtoReflect.Descriptor instead.
func (*TagRewriteItem) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{4}
}

func (x *TagRewriteItem) GetOldTag() string {
	if x != nil && x.OldTag != nil {
		return *x.OldTag
	}
	return ""
}

func (x *TagRewriteItem) GetExpression() string {
	if x != nil && x.Expression != nil {
		return *x.Expression
	}
	return ""
}

func (x *TagRewriteItem) GetNewTag() string {
	if x != nil && x.NewTag != nil {
		return *x.NewTag
	}
	return ""
}

func (x *TagRewriteItem) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *TagRewriteItem) GetDeleteTag() []string {
	if x != nil {
		return x.DeleteTag
	}
	return nil
}

type TagRewriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewriteType string          `protobuf:"bytes,1,opt,name=rewriteType,proto3" json:"rewriteType" form:"rewriteType" query:"rewriteType"`
	RewriteItem *TagRewriteItem `protobuf:"bytes,4,opt,name=rewriteItem,proto3" json:"rewriteItem" form:"rewriteItem" query:"rewriteItem"`
	Filters     []*Filter       `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters" form:"filters" query:"filters"`
}

func (x *TagRewriteRequest) Reset() {
	*x = TagRewriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagRewriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagRewriteRequest) ProtoMessage() {}

func (x *TagRewriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagRewriteRequest.ProtoReflect.Descriptor instead.
func (*TagRewriteRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{5}
}

func (x *TagRewriteRequest) GetRewriteType() string {
	if x != nil {
		return x.RewriteType
	}
	return ""
}

func (x *TagRewriteRequest) GetRewriteItem() *TagRewriteItem {
	if x != nil {
		return x.RewriteItem
	}
	return nil
}

func (x *TagRewriteRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

var File_integration_proto protoreflect.FileDescriptor

var file_integration_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x80, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x36, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x4e, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xd7, 0x01, 0x0a, 0x0e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x6f, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6e, 0x65, 0x77, 0x54, 0x61, 0x67,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x42, 0x1d, 0x5a, 0x1b, 0x6d, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integration_proto_rawDescOnce sync.Once
	file_integration_proto_rawDescData = file_integration_proto_rawDesc
)

func file_integration_proto_rawDescGZIP() []byte {
	file_integration_proto_rawDescOnce.Do(func() {
		file_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_integration_proto_rawDescData)
	})
	return file_integration_proto_rawDescData
}

var file_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_integration_proto_goTypes = []interface{}{
	(*IntegrationRequest)(nil),       // 0: integration.IntegrationRequest
	(*IntegrationRouterRequest)(nil), // 1: integration.IntegrationRouterRequest
	(*Filters)(nil),                  // 2: integration.Filters
	(*Filter)(nil),                   // 3: integration.Filter
	(*TagRewriteItem)(nil),           // 4: integration.TagRewriteItem
	(*TagRewriteRequest)(nil),        // 5: integration.TagRewriteRequest
}
var file_integration_proto_depIdxs = []int32{
	2, // 0: integration.IntegrationRouterRequest.filters:type_name -> integration.Filters
	3, // 1: integration.Filters.values:type_name -> integration.Filter
	4, // 2: integration.TagRewriteRequest.rewriteItem:type_name -> integration.TagRewriteItem
	3, // 3: integration.TagRewriteRequest.filters:type_name -> integration.Filter
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_integration_proto_init() }
func file_integration_proto_init() {
	if File_integration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationRequest); i {
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
		file_integration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationRouterRequest); i {
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
		file_integration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_integration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_integration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagRewriteItem); i {
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
		file_integration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagRewriteRequest); i {
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
	file_integration_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integration_proto_goTypes,
		DependencyIndexes: file_integration_proto_depIdxs,
		MessageInfos:      file_integration_proto_msgTypes,
	}.Build()
	File_integration_proto = out.File
	file_integration_proto_rawDesc = nil
	file_integration_proto_goTypes = nil
	file_integration_proto_depIdxs = nil
}
