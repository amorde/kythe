// Code generated by protoc-gen-go.
// source: kythe/proto/analysis.proto
// DO NOT EDIT!

/*
Package analysis_proto is a generated protocol buffer package.

It is generated from these files:
	kythe/proto/analysis.proto

It has these top-level messages:
	AnalysisRequest
	AnalysisOutput
	CompilationUnit
	FileInfo
	FileData
*/
package analysis_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "kythe.io/third_party/proto/any_proto"
import kythe_proto "kythe.io/kythe/proto/storage_proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An AnalysisRequest instructs an analyzer to perform an analysis on a single
// CompilationUnit.
type AnalysisRequest struct {
	// The compilation to analyze.
	Compilation *CompilationUnit `protobuf:"bytes,1,opt,name=compilation" json:"compilation,omitempty"`
	// The address of a file data service to use.  If this is provided, it should
	// be used in preference to any other file data service the analyzer may know
	// about for this compilation.
	FileDataService string `protobuf:"bytes,2,opt,name=file_data_service" json:"file_data_service,omitempty"`
}

func (m *AnalysisRequest) Reset()         { *m = AnalysisRequest{} }
func (m *AnalysisRequest) String() string { return proto.CompactTextString(m) }
func (*AnalysisRequest) ProtoMessage()    {}

func (m *AnalysisRequest) GetCompilation() *CompilationUnit {
	if m != nil {
		return m.Compilation
	}
	return nil
}

// AnalysisOutput contains an output artifact for the current analysis taking
// place.  A given analysis may not produce any outputs.  It is okay for an
// indexer to send an empty AnalysisOutput message if needed to keep the RPC
// channel alive; the driver must correctly handle this.
type AnalysisOutput struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *AnalysisOutput) Reset()         { *m = AnalysisOutput{} }
func (m *AnalysisOutput) String() string { return proto.CompactTextString(m) }
func (*AnalysisOutput) ProtoMessage()    {}

// Describes a single unit of compilation.
type CompilationUnit struct {
	// The base VName for the compilation and any generated VNames from its
	// analysis. Generally, the `language` component designates the language of
	// the compilation's sources.
	VName *kythe_proto.VName `protobuf:"bytes,1,opt,name=v_name" json:"v_name,omitempty"`
	// The revision of the compilation.
	Revision string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	// All files that might be touched in the course of this compilation.
	RequiredInput []*CompilationUnit_FileInput `protobuf:"bytes,3,rep,name=required_input" json:"required_input,omitempty"`
	// Set by the extractor to indicate that the original input had compile
	// errors. This is used to check validity of the sharded analysis.
	HasCompileErrors bool `protobuf:"varint,4,opt,name=has_compile_errors" json:"has_compile_errors,omitempty"`
	// The arguments to pass to a compiler tool for this compilation unit,
	// including the compiler executable, flags, and input files.
	Argument []string `protobuf:"bytes,5,rep,name=argument" json:"argument,omitempty"`
	// Of those files in `required_input`, the ones that this CompilationUnit
	// is intended to analyze. This is necessary to support languages like Go,
	// where a single translation unit may contain many source files that must all
	// be processed at once (while excluding source files that belong to other
	// CUs/packages, if any).
	SourceFile []string `protobuf:"bytes,6,rep,name=source_file" json:"source_file,omitempty"`
	// The output key of the CompilationUnit; for example, the object file that
	// it writes.
	// TODO(zarko): should this be a VName? Are there uniqueness requirements?
	// How is this used in the pipeline?
	OutputKey string `protobuf:"bytes,7,opt,name=output_key" json:"output_key,omitempty"`
	// The absolute path of the current working directory where the build tool
	// was invoked.  During analysis, a file whose path has working_directory
	// plus a path separator as an exact prefix is considered accessible from
	// that same path without said prefix.  It is only necessary to set this
	// field if the build tool requires it.
	WorkingDirectory string `protobuf:"bytes,8,opt,name=working_directory" json:"working_directory,omitempty"`
	// For languages that make use of resource contexts (like C++), the context
	// that should be initially entered.
	// TODO(zarko): What is a "resource context"? Needs a clear definition and/or
	// a link to one.
	EntryContext string `protobuf:"bytes,9,opt,name=entry_context" json:"entry_context,omitempty"`
	// A collection of environment variables that the build environment expects
	// to be set.  As a rule, we only record variables here that must be set to
	// specific values for the build to work.
	// TODO(fromberger): When we move to NWP, use a map instead.
	Environment []*CompilationUnit_Env `protobuf:"bytes,10,rep,name=environment" json:"environment,omitempty"`
	// Per-language or -tool details.
	Details []*google_protobuf.Any `protobuf:"bytes,11,rep,name=details" json:"details,omitempty"`
}

func (m *CompilationUnit) Reset()         { *m = CompilationUnit{} }
func (m *CompilationUnit) String() string { return proto.CompactTextString(m) }
func (*CompilationUnit) ProtoMessage()    {}

func (m *CompilationUnit) GetVName() *kythe_proto.VName {
	if m != nil {
		return m.VName
	}
	return nil
}

func (m *CompilationUnit) GetRequiredInput() []*CompilationUnit_FileInput {
	if m != nil {
		return m.RequiredInput
	}
	return nil
}

func (m *CompilationUnit) GetEnvironment() []*CompilationUnit_Env {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *CompilationUnit) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// ContextDependentVersionColumn and ContextDependentVersionRow
// define a table that relates input contexts (keyed by a single
// source context per row) to tuples of (byte offset * linked context).
// When a FileInput F being processed in context C refers to another
// FileInput F' at offset O (perhaps because F has an #include directive at O)
// the context in which F' should be processed is the linked context derived
// from this table.
type CompilationUnit_ContextDependentVersionColumn struct {
	// The byte offset into the file resource.
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	// The signature for the resulting context.
	LinkedContext string `protobuf:"bytes,2,opt,name=linked_context" json:"linked_context,omitempty"`
}

func (m *CompilationUnit_ContextDependentVersionColumn) Reset() {
	*m = CompilationUnit_ContextDependentVersionColumn{}
}
func (m *CompilationUnit_ContextDependentVersionColumn) String() string {
	return proto.CompactTextString(m)
}
func (*CompilationUnit_ContextDependentVersionColumn) ProtoMessage() {}

// See ContextDependentVersionColumn for details.
// It is valid for a ContextDependentVersionRow to have no columns. In this
// case, the associated FileInput was seen to exist in some context C, but
// did not refer to any other FileInputs while in that context.
type CompilationUnit_ContextDependentVersionRow struct {
	// The context to be applied to all columns.
	SourceContext string `protobuf:"bytes,1,opt,name=source_context" json:"source_context,omitempty"`
	// A map from byte offsets to linked contexts.
	Column []*CompilationUnit_ContextDependentVersionColumn `protobuf:"bytes,2,rep,name=column" json:"column,omitempty"`
	// If true, this version should always be processed regardless of any
	// claiming.
	AlwaysProcess bool `protobuf:"varint,3,opt,name=always_process" json:"always_process,omitempty"`
}

func (m *CompilationUnit_ContextDependentVersionRow) Reset() {
	*m = CompilationUnit_ContextDependentVersionRow{}
}
func (m *CompilationUnit_ContextDependentVersionRow) String() string {
	return proto.CompactTextString(m)
}
func (*CompilationUnit_ContextDependentVersionRow) ProtoMessage() {}

func (m *CompilationUnit_ContextDependentVersionRow) GetColumn() []*CompilationUnit_ContextDependentVersionColumn {
	if m != nil {
		return m.Column
	}
	return nil
}

type CompilationUnit_FileInput struct {
	// If set, overrides the `v_name` in the `CompilationUnit` for deriving
	// VNames during analysis.
	VName *kythe_proto.VName `protobuf:"bytes,1,opt,name=v_name" json:"v_name,omitempty"`
	// The file's metadata. It is invalid to provide a FileInput without both
	// the file's path and digest.
	Info *FileInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	// The file's context-dependent versions.
	Context []*CompilationUnit_ContextDependentVersionRow `protobuf:"bytes,3,rep,name=context" json:"context,omitempty"`
}

func (m *CompilationUnit_FileInput) Reset()         { *m = CompilationUnit_FileInput{} }
func (m *CompilationUnit_FileInput) String() string { return proto.CompactTextString(m) }
func (*CompilationUnit_FileInput) ProtoMessage()    {}

func (m *CompilationUnit_FileInput) GetVName() *kythe_proto.VName {
	if m != nil {
		return m.VName
	}
	return nil
}

func (m *CompilationUnit_FileInput) GetInfo() *FileInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *CompilationUnit_FileInput) GetContext() []*CompilationUnit_ContextDependentVersionRow {
	if m != nil {
		return m.Context
	}
	return nil
}

// An Env message represents the name and value of a single environment
// variable in the build environment.
type CompilationUnit_Env struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *CompilationUnit_Env) Reset()         { *m = CompilationUnit_Env{} }
func (m *CompilationUnit_Env) String() string { return proto.CompactTextString(m) }
func (*CompilationUnit_Env) ProtoMessage()    {}

// A FileInfo specifies metadata for a file under analysis.
type FileInfo struct {
	// This path should be relative to the working directory of the compilation
	// command -- typically the root of the build.
	// i.e. file/base/file.cc or ../../base/atomic_ref_count.h
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The lowercase ascii hex SHA-256 digest of the file contents.
	Digest string `protobuf:"bytes,2,opt,name=digest" json:"digest,omitempty"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}

// A FileData describes the content of a single file.  A server responding to a
// FileDataRequest must populate the same path and digest in the reply that were
// provided by the client in the request, if any.  If either field was empty in
// the request, the server may leave it empty or fill in its value.
type FileData struct {
	Content []byte    `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Info    *FileInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *FileData) Reset()         { *m = FileData{} }
func (m *FileData) String() string { return proto.CompactTextString(m) }
func (*FileData) ProtoMessage()    {}

func (m *FileData) GetInfo() *FileInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for CompilationAnalyzer service

type CompilationAnalyzerClient interface {
	// Analyze is the main entry point for the analysis driver to send work to the
	// analyzer.  The analysis may produce many outputs which will be streamed as
	// framed AnalysisOutput messages.
	//
	// A driver may choose to retry analyses that return RPC errors.  It should
	// not retry analyses that are reported as finished unless it is necessary to
	// recover from an external production issue.
	Analyze(ctx context.Context, in *AnalysisRequest, opts ...grpc.CallOption) (CompilationAnalyzer_AnalyzeClient, error)
}

type compilationAnalyzerClient struct {
	cc *grpc.ClientConn
}

func NewCompilationAnalyzerClient(cc *grpc.ClientConn) CompilationAnalyzerClient {
	return &compilationAnalyzerClient{cc}
}

func (c *compilationAnalyzerClient) Analyze(ctx context.Context, in *AnalysisRequest, opts ...grpc.CallOption) (CompilationAnalyzer_AnalyzeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CompilationAnalyzer_serviceDesc.Streams[0], c.cc, "/kythe.proto.CompilationAnalyzer/Analyze", opts...)
	if err != nil {
		return nil, err
	}
	x := &compilationAnalyzerAnalyzeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompilationAnalyzer_AnalyzeClient interface {
	Recv() (*AnalysisOutput, error)
	grpc.ClientStream
}

type compilationAnalyzerAnalyzeClient struct {
	grpc.ClientStream
}

func (x *compilationAnalyzerAnalyzeClient) Recv() (*AnalysisOutput, error) {
	m := new(AnalysisOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CompilationAnalyzer service

type CompilationAnalyzerServer interface {
	// Analyze is the main entry point for the analysis driver to send work to the
	// analyzer.  The analysis may produce many outputs which will be streamed as
	// framed AnalysisOutput messages.
	//
	// A driver may choose to retry analyses that return RPC errors.  It should
	// not retry analyses that are reported as finished unless it is necessary to
	// recover from an external production issue.
	Analyze(*AnalysisRequest, CompilationAnalyzer_AnalyzeServer) error
}

func RegisterCompilationAnalyzerServer(s *grpc.Server, srv CompilationAnalyzerServer) {
	s.RegisterService(&_CompilationAnalyzer_serviceDesc, srv)
}

func _CompilationAnalyzer_Analyze_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AnalysisRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompilationAnalyzerServer).Analyze(m, &compilationAnalyzerAnalyzeServer{stream})
}

type CompilationAnalyzer_AnalyzeServer interface {
	Send(*AnalysisOutput) error
	grpc.ServerStream
}

type compilationAnalyzerAnalyzeServer struct {
	grpc.ServerStream
}

func (x *compilationAnalyzerAnalyzeServer) Send(m *AnalysisOutput) error {
	return x.ServerStream.SendMsg(m)
}

var _CompilationAnalyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.CompilationAnalyzer",
	HandlerType: (*CompilationAnalyzerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Analyze",
			Handler:       _CompilationAnalyzer_Analyze_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for FileDataService service

type FileDataServiceClient interface {
	// Get returns the contents of one or more files needed for analysis.  It is
	// the server implementation's responsibility to do any caching that might be
	// necessary to make this perform well so that an analyzer does not need to
	// implement its own caches unless it is doing something unusual.
	//
	// Except in case of error, the server is required to return at least one of
	// the requested files; however, the server may also limit how much data is
	// returned for a single request.  If this occurs, any files it does return
	// must be complete.  The client is responsible for issuing additional
	// requests as necessary to obtain any missing files.  For example, the
	// following is a valid interaction:
	//
	//   C: Get(a, b, c)
	//   S: FileData(a)
	//   S: FileData(c)
	//   C: Get(b)
	//   S: FileData(b)
	//
	Get(ctx context.Context, opts ...grpc.CallOption) (FileDataService_GetClient, error)
}

type fileDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileDataServiceClient(cc *grpc.ClientConn) FileDataServiceClient {
	return &fileDataServiceClient{cc}
}

func (c *fileDataServiceClient) Get(ctx context.Context, opts ...grpc.CallOption) (FileDataService_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FileDataService_serviceDesc.Streams[0], c.cc, "/kythe.proto.FileDataService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileDataServiceGetClient{stream}
	return x, nil
}

type FileDataService_GetClient interface {
	Send(*FileInfo) error
	Recv() (*FileData, error)
	grpc.ClientStream
}

type fileDataServiceGetClient struct {
	grpc.ClientStream
}

func (x *fileDataServiceGetClient) Send(m *FileInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileDataServiceGetClient) Recv() (*FileData, error) {
	m := new(FileData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FileDataService service

type FileDataServiceServer interface {
	// Get returns the contents of one or more files needed for analysis.  It is
	// the server implementation's responsibility to do any caching that might be
	// necessary to make this perform well so that an analyzer does not need to
	// implement its own caches unless it is doing something unusual.
	//
	// Except in case of error, the server is required to return at least one of
	// the requested files; however, the server may also limit how much data is
	// returned for a single request.  If this occurs, any files it does return
	// must be complete.  The client is responsible for issuing additional
	// requests as necessary to obtain any missing files.  For example, the
	// following is a valid interaction:
	//
	//   C: Get(a, b, c)
	//   S: FileData(a)
	//   S: FileData(c)
	//   C: Get(b)
	//   S: FileData(b)
	//
	Get(FileDataService_GetServer) error
}

func RegisterFileDataServiceServer(s *grpc.Server, srv FileDataServiceServer) {
	s.RegisterService(&_FileDataService_serviceDesc, srv)
}

func _FileDataService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileDataServiceServer).Get(&fileDataServiceGetServer{stream})
}

type FileDataService_GetServer interface {
	Send(*FileData) error
	Recv() (*FileInfo, error)
	grpc.ServerStream
}

type fileDataServiceGetServer struct {
	grpc.ServerStream
}

func (x *fileDataServiceGetServer) Send(m *FileData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileDataServiceGetServer) Recv() (*FileInfo, error) {
	m := new(FileInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.FileDataService",
	HandlerType: (*FileDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _FileDataService_Get_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
