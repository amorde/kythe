// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: kythe/proto/extraction_config.proto

package extraction_config_go_proto

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

type ExtractionConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiredImage []*ExtractionConfiguration_Image      `protobuf:"bytes,1,rep,name=required_image,json=requiredImage,proto3" json:"required_image,omitempty"`
	RunCommand    []*ExtractionConfiguration_RunCommand `protobuf:"bytes,2,rep,name=run_command,json=runCommand,proto3" json:"run_command,omitempty"`
	EntryPoint    []string                              `protobuf:"bytes,3,rep,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
}

func (x *ExtractionConfiguration) Reset() {
	*x = ExtractionConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_extraction_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionConfiguration) ProtoMessage() {}

func (x *ExtractionConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_extraction_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionConfiguration.ProtoReflect.Descriptor instead.
func (*ExtractionConfiguration) Descriptor() ([]byte, []int) {
	return file_kythe_proto_extraction_config_proto_rawDescGZIP(), []int{0}
}

func (x *ExtractionConfiguration) GetRequiredImage() []*ExtractionConfiguration_Image {
	if x != nil {
		return x.RequiredImage
	}
	return nil
}

func (x *ExtractionConfiguration) GetRunCommand() []*ExtractionConfiguration_RunCommand {
	if x != nil {
		return x.RunCommand
	}
	return nil
}

func (x *ExtractionConfiguration) GetEntryPoint() []string {
	if x != nil {
		return x.EntryPoint
	}
	return nil
}

type ExtractionConfiguration_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri      string                              `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Name     string                              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CopySpec []*ExtractionConfiguration_CopySpec `protobuf:"bytes,3,rep,name=copy_spec,json=copySpec,proto3" json:"copy_spec,omitempty"`
	EnvVar   []*ExtractionConfiguration_EnvVar   `protobuf:"bytes,4,rep,name=env_var,json=envVar,proto3" json:"env_var,omitempty"`
}

func (x *ExtractionConfiguration_Image) Reset() {
	*x = ExtractionConfiguration_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_extraction_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionConfiguration_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionConfiguration_Image) ProtoMessage() {}

func (x *ExtractionConfiguration_Image) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_extraction_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionConfiguration_Image.ProtoReflect.Descriptor instead.
func (*ExtractionConfiguration_Image) Descriptor() ([]byte, []int) {
	return file_kythe_proto_extraction_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExtractionConfiguration_Image) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *ExtractionConfiguration_Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExtractionConfiguration_Image) GetCopySpec() []*ExtractionConfiguration_CopySpec {
	if x != nil {
		return x.CopySpec
	}
	return nil
}

func (x *ExtractionConfiguration_Image) GetEnvVar() []*ExtractionConfiguration_EnvVar {
	if x != nil {
		return x.EnvVar
	}
	return nil
}

type ExtractionConfiguration_CopySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *ExtractionConfiguration_CopySpec) Reset() {
	*x = ExtractionConfiguration_CopySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_extraction_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionConfiguration_CopySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionConfiguration_CopySpec) ProtoMessage() {}

func (x *ExtractionConfiguration_CopySpec) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_extraction_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionConfiguration_CopySpec.ProtoReflect.Descriptor instead.
func (*ExtractionConfiguration_CopySpec) Descriptor() ([]byte, []int) {
	return file_kythe_proto_extraction_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ExtractionConfiguration_CopySpec) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ExtractionConfiguration_CopySpec) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type ExtractionConfiguration_EnvVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExtractionConfiguration_EnvVar) Reset() {
	*x = ExtractionConfiguration_EnvVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_extraction_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionConfiguration_EnvVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionConfiguration_EnvVar) ProtoMessage() {}

func (x *ExtractionConfiguration_EnvVar) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_extraction_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionConfiguration_EnvVar.ProtoReflect.Descriptor instead.
func (*ExtractionConfiguration_EnvVar) Descriptor() ([]byte, []int) {
	return file_kythe_proto_extraction_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ExtractionConfiguration_EnvVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExtractionConfiguration_EnvVar) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ExtractionConfiguration_RunCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Arg     []string `protobuf:"bytes,2,rep,name=arg,proto3" json:"arg,omitempty"`
}

func (x *ExtractionConfiguration_RunCommand) Reset() {
	*x = ExtractionConfiguration_RunCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_extraction_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionConfiguration_RunCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionConfiguration_RunCommand) ProtoMessage() {}

func (x *ExtractionConfiguration_RunCommand) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_extraction_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionConfiguration_RunCommand.ProtoReflect.Descriptor instead.
func (*ExtractionConfiguration_RunCommand) Descriptor() ([]byte, []int) {
	return file_kythe_proto_extraction_config_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ExtractionConfiguration_RunCommand) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ExtractionConfiguration_RunCommand) GetArg() []string {
	if x != nil {
		return x.Arg
	}
	return nil
}

var File_kythe_proto_extraction_config_proto protoreflect.FileDescriptor

var file_kythe_proto_extraction_config_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd5, 0x04, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51,
	0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x50, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x1a, 0xbf, 0x01, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x70, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x63, 0x6f, 0x70, 0x79, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x44, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x52, 0x06,
	0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x1a, 0x44, 0x0a, 0x08, 0x43, 0x6f, 0x70, 0x79, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x32, 0x0a, 0x06,
	0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x38, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x67, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x67, 0x42, 0x31, 0x5a, 0x2f, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kythe_proto_extraction_config_proto_rawDescOnce sync.Once
	file_kythe_proto_extraction_config_proto_rawDescData = file_kythe_proto_extraction_config_proto_rawDesc
)

func file_kythe_proto_extraction_config_proto_rawDescGZIP() []byte {
	file_kythe_proto_extraction_config_proto_rawDescOnce.Do(func() {
		file_kythe_proto_extraction_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_extraction_config_proto_rawDescData)
	})
	return file_kythe_proto_extraction_config_proto_rawDescData
}

var file_kythe_proto_extraction_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kythe_proto_extraction_config_proto_goTypes = []interface{}{
	(*ExtractionConfiguration)(nil),            // 0: kythe.proto.ExtractionConfiguration
	(*ExtractionConfiguration_Image)(nil),      // 1: kythe.proto.ExtractionConfiguration.Image
	(*ExtractionConfiguration_CopySpec)(nil),   // 2: kythe.proto.ExtractionConfiguration.CopySpec
	(*ExtractionConfiguration_EnvVar)(nil),     // 3: kythe.proto.ExtractionConfiguration.EnvVar
	(*ExtractionConfiguration_RunCommand)(nil), // 4: kythe.proto.ExtractionConfiguration.RunCommand
}
var file_kythe_proto_extraction_config_proto_depIdxs = []int32{
	1, // 0: kythe.proto.ExtractionConfiguration.required_image:type_name -> kythe.proto.ExtractionConfiguration.Image
	4, // 1: kythe.proto.ExtractionConfiguration.run_command:type_name -> kythe.proto.ExtractionConfiguration.RunCommand
	2, // 2: kythe.proto.ExtractionConfiguration.Image.copy_spec:type_name -> kythe.proto.ExtractionConfiguration.CopySpec
	3, // 3: kythe.proto.ExtractionConfiguration.Image.env_var:type_name -> kythe.proto.ExtractionConfiguration.EnvVar
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kythe_proto_extraction_config_proto_init() }
func file_kythe_proto_extraction_config_proto_init() {
	if File_kythe_proto_extraction_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_extraction_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionConfiguration); i {
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
		file_kythe_proto_extraction_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionConfiguration_Image); i {
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
		file_kythe_proto_extraction_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionConfiguration_CopySpec); i {
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
		file_kythe_proto_extraction_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionConfiguration_EnvVar); i {
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
		file_kythe_proto_extraction_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionConfiguration_RunCommand); i {
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
			RawDescriptor: file_kythe_proto_extraction_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kythe_proto_extraction_config_proto_goTypes,
		DependencyIndexes: file_kythe_proto_extraction_config_proto_depIdxs,
		MessageInfos:      file_kythe_proto_extraction_config_proto_msgTypes,
	}.Build()
	File_kythe_proto_extraction_config_proto = out.File
	file_kythe_proto_extraction_config_proto_rawDesc = nil
	file_kythe_proto_extraction_config_proto_goTypes = nil
	file_kythe_proto_extraction_config_proto_depIdxs = nil
}
