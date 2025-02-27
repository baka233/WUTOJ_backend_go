package judger

import (
	"context"
	"fmt"
	"math"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defintions of all judge result types.
type Report_ResultType int32

const (
	Report_Accepted            Report_ResultType = 0
	Report_WrongAnswer         Report_ResultType = 1
	Report_TimeLimitExceeded   Report_ResultType = 2
	Report_MemoryLimitExceeded Report_ResultType = 3
	Report_RuntimeError        Report_ResultType = 4
	Report_CompileError        Report_ResultType = 5
	Report_SystemError         Report_ResultType = 6
)

var Report_ResultType_name = map[int32]string{
	0: "AC",
	1: "WA",
	2: "TLE",
	3: "MLE",
	4: "RE",
	5: "CE",
	6: "SE",
	7: "UE",
}

var Report_ResultType_value = map[string]int32{
	"AC":  0,
	"WA":  1,
	"TLE": 2,
	"MLE": 3,
	"RE":  4,
	"CE":  5,
	"SE":  6,
	"UE":  7,
}

func (x Report_ResultType) String() string {
	return proto.EnumName(Report_ResultType_name, int32(x))
}

func (Report_ResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2, 0}
}

type Workspace struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workspace.Unmarshal(m, b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
}
func (m *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(m, src)
}
func (m *Workspace) XXX_Size() int {
	return xxx_messageInfo_Workspace.Size(m)
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Workspace) GetPath() *wrappers.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

type Request struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

// // The judge report definition.
// // Contains the judge result and resource usage and an optional system
// message.
type Report struct {
	Result               Report_ResultType     `protobuf:"varint,1,opt,name=result,proto3,enum=ana_rpc.Report_ResultType" json:"result,omitempty"`
	Usage                *Resource             `protobuf:"bytes,2,opt,name=usage,proto3" json:"usage,omitempty"`
	Message              *wrappers.StringValue `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetResult() Report_ResultType {
	if m != nil {
		return m.Result
	}
	return Report_Accepted
}

func (m *Report) GetUsage() *Resource {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Report) GetMessage() *wrappers.StringValue {
	if m != nil {
		return m.Message
	}
	return nil
}

type Resource struct {
	RealTime             *duration.Duration    `protobuf:"bytes,1,opt,name=real_time,json=realTime,proto3" json:"real_time,omitempty"`
	CpuTime              *duration.Duration    `protobuf:"bytes,2,opt,name=cpu_time,json=cpuTime,proto3" json:"cpu_time,omitempty"`
	Memory               *wrappers.UInt64Value `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetRealTime() *duration.Duration {
	if m != nil {
		return m.RealTime
	}
	return nil
}

func (m *Resource) GetCpuTime() *duration.Duration {
	if m != nil {
		return m.CpuTime
	}
	return nil
}

func (m *Resource) GetMemory() *wrappers.UInt64Value {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("ana_rpc.Report_ResultType", Report_ResultType_name, Report_ResultType_value)
	proto.RegisterType((*Workspace)(nil), "ana_rpc.Workspace")
	proto.RegisterType((*Request)(nil), "ana_rpc.Request")
	proto.RegisterType((*Report)(nil), "ana_rpc.Report")
	proto.RegisterType((*Resource)(nil), "ana_rpc.Resource")
}

func init() {
	proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1)
}

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0xd3, 0x2d, 0x6d, 0xcf, 0xa6, 0x2d, 0x33, 0x02, 0x46, 0x40, 0x08, 0xe5, 0x06, 0x2e,
	0x50, 0x36, 0x95, 0xa9, 0xdc, 0x21, 0x55, 0x50, 0x21, 0x10, 0xdc, 0x64, 0x83, 0x5d, 0x4e, 0x5e,
	0x72, 0x08, 0xd6, 0xe2, 0xd8, 0x1c, 0xdb, 0x2a, 0x7d, 0x0b, 0x9e, 0x04, 0xf1, 0x88, 0x28, 0x89,
	0xbb, 0x8d, 0x55, 0x88, 0x71, 0xfb, 0xfd, 0x9c, 0xdf, 0x0f, 0xc6, 0xa4, 0xf3, 0x54, 0x93, 0xb2,
	0x8a, 0x0d, 0x79, 0xcd, 0xcf, 0x48, 0xe7, 0xf1, 0xe3, 0x52, 0xa9, 0xb2, 0xc2, 0x83, 0x16, 0x3e,
	0x77, 0x5f, 0x0e, 0x0a, 0x47, 0xdc, 0x0a, 0x55, 0x77, 0xc2, 0xf8, 0xe1, 0x4d, 0x1e, 0xa5, 0xb6,
	0x4b, 0x4f, 0xae, 0x99, 0x17, 0xc4, 0xb5, 0x46, 0x32, 0x1d, 0x9f, 0x5c, 0xc0, 0xf8, 0x54, 0xd1,
	0x85, 0xd1, 0x3c, 0x47, 0xf6, 0x1c, 0x02, 0x51, 0xec, 0xf7, 0x9f, 0xf4, 0x9f, 0x6d, 0x4d, 0x1e,
	0xa5, 0x9d, 0x33, 0x5d, 0x39, 0xd3, 0x63, 0x4b, 0xa2, 0x2e, 0x3f, 0xf3, 0xca, 0x61, 0x16, 0x88,
	0x82, 0x1d, 0xc2, 0x86, 0xe6, 0xf6, 0xeb, 0x7e, 0x70, 0x0b, 0x7d, 0xab, 0x4c, 0x5e, 0xc2, 0x30,
	0xc3, 0x6f, 0x0e, 0x8d, 0xfd, 0xbf, 0x56, 0xc9, 0xaf, 0x00, 0xc2, 0x0c, 0xb5, 0x22, 0xcb, 0x26,
	0x10, 0x12, 0x1a, 0x57, 0xd9, 0xd6, 0xbc, 0x33, 0x89, 0x53, 0x7f, 0xa7, 0xb4, 0x13, 0xa4, 0x59,
	0xcb, 0x9e, 0x2c, 0x35, 0x66, 0x5e, 0xc9, 0x9e, 0xc2, 0xa6, 0x33, 0xbc, 0x44, 0x3f, 0xea, 0xde,
	0x35, 0x8b, 0x51, 0x8e, 0x72, 0xcc, 0x3a, 0x9e, 0x4d, 0x61, 0x28, 0xd1, 0xb4, 0xd2, 0xc1, 0x2d,
	0x46, 0x5b, 0x89, 0x93, 0x1f, 0x7d, 0x80, 0xab, 0xbe, 0x6c, 0x1b, 0x46, 0xb3, 0x3c, 0x47, 0x6d,
	0xb1, 0x88, 0x7a, 0x6c, 0x17, 0xb6, 0x4e, 0x49, 0xd5, 0xe5, 0xac, 0x36, 0x0b, 0xa4, 0xa8, 0xcf,
	0xee, 0xc2, 0xde, 0x89, 0x90, 0xf8, 0x41, 0x48, 0x61, 0xe7, 0xdf, 0x73, 0xc4, 0x02, 0x8b, 0x28,
	0x60, 0xf7, 0xe1, 0xce, 0x47, 0x94, 0x8a, 0x96, 0x7f, 0x12, 0x03, 0x16, 0xc1, 0x76, 0xe6, 0x6a,
	0x2b, 0x24, 0xce, 0x89, 0x14, 0x45, 0x1b, 0x0d, 0xf2, 0x5a, 0x49, 0x2d, 0x2a, 0x8f, 0x6c, 0x36,
	0x4d, 0x8e, 0x97, 0xc6, 0xa2, 0xec, 0x80, 0x30, 0xf9, 0xd9, 0x87, 0xd1, 0x6a, 0x3d, 0x36, 0x85,
	0x31, 0x21, 0xaf, 0xce, 0x9a, 0x1a, 0xfe, 0xe8, 0x0f, 0xd6, 0x36, 0x7b, 0xe3, 0x63, 0x95, 0x8d,
	0x1a, 0x6d, 0x33, 0x21, 0x3b, 0x82, 0x51, 0xae, 0x5d, 0x67, 0x0b, 0xfe, 0x65, 0x1b, 0xe6, 0xda,
	0x79, 0x57, 0x28, 0xdb, 0x45, 0xfe, 0x7a, 0xc4, 0x4f, 0xef, 0x6a, 0x3b, 0x3d, 0xea, 0x8e, 0xe8,
	0xb5, 0x93, 0x05, 0x0c, 0x66, 0x35, 0x67, 0xaf, 0x60, 0xe7, 0xbd, 0x2b, 0x4a, 0xbc, 0x4a, 0x25,
	0xbb, 0x7c, 0xd7, 0x25, 0x16, 0xdf, 0x5b, 0x2b, 0x39, 0x6f, 0x42, 0x9f, 0xf4, 0xd8, 0x21, 0x8c,
	0xdf, 0xa2, 0xf5, 0x61, 0x89, 0xae, 0x7d, 0xba, 0xcd, 0x5d, 0xbc, 0x7b, 0x23, 0x2e, 0x49, 0xef,
	0x3c, 0x6c, 0x6b, 0xbc, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x05, 0xe5, 0x50, 0x7c, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnaClient is the client API for Ana service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnaClient interface {
	// rpc Judge(Task) returns (stream Report) {}
	// rpc Cache(Problem) returns (google.protobuf.Empty) {}
	JudgeWorkspace(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (*empty.Empty, error)
	GetReport(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Report, error)
}

type anaClient struct {
	cc grpc.ClientConnInterface
}

func NewAnaClient(cc grpc.ClientConnInterface) AnaClient {
	return &anaClient{cc}
}

func (c *anaClient) JudgeWorkspace(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ana_rpc.Ana/JudgeWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anaClient) GetReport(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Report, error) {
	out := new(Report)
	err := c.cc.Invoke(ctx, "/ana_rpc.Ana/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnaServer is the server API for Ana service.
type AnaServer interface {
	// rpc Judge(Task) returns (stream Report) {}
	// rpc Cache(Problem) returns (google.protobuf.Empty) {}
	JudgeWorkspace(context.Context, *Workspace) (*empty.Empty, error)
	GetReport(context.Context, *Request) (*Report, error)
}

// UnimplementedAnaServer can be embedded to have forward compatible implementations.
type UnimplementedAnaServer struct {
}

func (*UnimplementedAnaServer) JudgeWorkspace(ctx context.Context, req *Workspace) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JudgeWorkspace not implemented")
}
func (*UnimplementedAnaServer) GetReport(ctx context.Context, req *Request) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}

func RegisterAnaServer(s *grpc.Server, srv AnaServer) {
	s.RegisterService(&_Ana_serviceDesc, srv)
}

func _Ana_JudgeWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Workspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnaServer).JudgeWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ana_rpc.Ana/JudgeWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnaServer).JudgeWorkspace(ctx, req.(*Workspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ana_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnaServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ana_rpc.Ana/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnaServer).GetReport(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ana_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ana_rpc.Ana",
	HandlerType: (*AnaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JudgeWorkspace",
			Handler:    _Ana_JudgeWorkspace_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _Ana_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}