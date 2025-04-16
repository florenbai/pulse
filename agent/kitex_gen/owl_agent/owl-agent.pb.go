// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: owl-agent.proto

package owl_agent

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op       string       `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Filename string       `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Groups   []*RuleGroup `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *RuleRequest) Reset() {
	*x = RuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_owl_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleRequest) ProtoMessage() {}

func (x *RuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_owl_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleRequest.ProtoReflect.Descriptor instead.
func (*RuleRequest) Descriptor() ([]byte, []int) {
	return file_owl_agent_proto_rawDescGZIP(), []int{0}
}

func (x *RuleRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *RuleRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *RuleRequest) GetGroups() []*RuleGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

type RuleGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rules []*Rule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *RuleGroup) Reset() {
	*x = RuleGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_owl_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroup) ProtoMessage() {}

func (x *RuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_owl_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleGroup.ProtoReflect.Descriptor instead.
func (*RuleGroup) Descriptor() ([]byte, []int) {
	return file_owl_agent_proto_rawDescGZIP(), []int{1}
}

func (x *RuleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RuleGroup) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert       string            `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
	Expr        string            `protobuf:"bytes,2,opt,name=expr,proto3" json:"expr,omitempty"`
	For         string            `protobuf:"bytes,3,opt,name=for,proto3" json:"for,omitempty"`
	Labels      map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations map[string]string `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_owl_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_owl_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_owl_agent_proto_rawDescGZIP(), []int{2}
}

func (x *Rule) GetAlert() string {
	if x != nil {
		return x.Alert
	}
	return ""
}

func (x *Rule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *Rule) GetFor() string {
	if x != nil {
		return x.For
	}
	return ""
}

func (x *Rule) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Rule) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_owl_agent_proto protoreflect.FileDescriptor

var file_owl_agent_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x77, 0x6c, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6f, 0x77, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0b, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x77, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x22, 0x46, 0x0a, 0x09, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x77, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb6, 0x02, 0x0a, 0x04, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x72, 0x12,
	0x33, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6f, 0x77, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x77, 0x6c, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x51, 0x0a, 0x0f, 0x4f, 0x77, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x6f, 0x77, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x69, 0x6e, 0x6f, 0x73, 0x2f,
	0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x77, 0x6c, 0x5f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_owl_agent_proto_rawDescOnce sync.Once
	file_owl_agent_proto_rawDescData = file_owl_agent_proto_rawDesc
)

func file_owl_agent_proto_rawDescGZIP() []byte {
	file_owl_agent_proto_rawDescOnce.Do(func() {
		file_owl_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_owl_agent_proto_rawDescData)
	})
	return file_owl_agent_proto_rawDescData
}

var file_owl_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_owl_agent_proto_goTypes = []interface{}{
	(*RuleRequest)(nil),   // 0: owl_agent.RuleRequest
	(*RuleGroup)(nil),     // 1: owl_agent.RuleGroup
	(*Rule)(nil),          // 2: owl_agent.Rule
	nil,                   // 3: owl_agent.Rule.LabelsEntry
	nil,                   // 4: owl_agent.Rule.AnnotationsEntry
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_owl_agent_proto_depIdxs = []int32{
	1, // 0: owl_agent.RuleRequest.groups:type_name -> owl_agent.RuleGroup
	2, // 1: owl_agent.RuleGroup.rules:type_name -> owl_agent.Rule
	3, // 2: owl_agent.Rule.labels:type_name -> owl_agent.Rule.LabelsEntry
	4, // 3: owl_agent.Rule.annotations:type_name -> owl_agent.Rule.AnnotationsEntry
	0, // 4: owl_agent.OwlAgentService.UpdateRule:input_type -> owl_agent.RuleRequest
	5, // 5: owl_agent.OwlAgentService.UpdateRule:output_type -> google.protobuf.Empty
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_owl_agent_proto_init() }
func file_owl_agent_proto_init() {
	if File_owl_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_owl_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleRequest); i {
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
		file_owl_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleGroup); i {
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
		file_owl_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
			RawDescriptor: file_owl_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_owl_agent_proto_goTypes,
		DependencyIndexes: file_owl_agent_proto_depIdxs,
		MessageInfos:      file_owl_agent_proto_msgTypes,
	}.Build()
	File_owl_agent_proto = out.File
	file_owl_agent_proto_rawDesc = nil
	file_owl_agent_proto_goTypes = nil
	file_owl_agent_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type OwlAgentService interface {
	UpdateRule(ctx context.Context, req *RuleRequest) (res *emptypb.Empty, err error)
}
