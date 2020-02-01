// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SumRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	SumResult            int32    `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	PrimeFactor          int64    `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetPrimeFactor() int64 {
	if m != nil {
		return m.PrimeFactor
	}
	return 0
}

type AverageRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRequest) Reset()         { *m = AverageRequest{} }
func (m *AverageRequest) String() string { return proto.CompactTextString(m) }
func (*AverageRequest) ProtoMessage()    {}
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *AverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRequest.Unmarshal(m, b)
}
func (m *AverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRequest.Marshal(b, m, deterministic)
}
func (m *AverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRequest.Merge(m, src)
}
func (m *AverageRequest) XXX_Size() int {
	return xxx_messageInfo_AverageRequest.Size(m)
}
func (m *AverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRequest proto.InternalMessageInfo

func (m *AverageRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (m *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(m, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximum              int32    `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*AverageRequest)(nil), "calculator.AverageRequest")
	proto.RegisterType((*AverageResponse)(nil), "calculator.AverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x5f, 0x21, 0x40, 0xde, 0x6d, 0x1f, 0x2f, 0x6f, 0x5e, 0x82, 0xa4, 0x46, 0x81, 0xba,
	0x69, 0x02, 0x01, 0xa2, 0x1b, 0x5d, 0xfa, 0x8f, 0x9d, 0xc6, 0x14, 0x56, 0x6e, 0x48, 0x29, 0x83,
	0x69, 0xd2, 0xe9, 0xd4, 0xf9, 0x43, 0xf4, 0x0b, 0xfb, 0x39, 0x8c, 0xd3, 0x0e, 0x1d, 0x82, 0x04,
	0x77, 0x73, 0xcf, 0xfc, 0xee, 0xb9, 0x93, 0x7b, 0x32, 0xe0, 0x47, 0x61, 0x12, 0xc9, 0x24, 0x14,
	0x94, 0x8d, 0xca, 0x63, 0xb6, 0x30, 0x8a, 0x61, 0xc6, 0xa8, 0xa0, 0x08, 0x4a, 0xc5, 0x9b, 0x01,
	0x4c, 0x25, 0x09, 0xf0, 0xab, 0xc4, 0x5c, 0xa0, 0x1e, 0x38, 0xab, 0x98, 0x71, 0x31, 0x4f, 0x25,
	0x59, 0x60, 0xd6, 0xb6, 0xba, 0x96, 0x5f, 0x0b, 0x6c, 0xa5, 0x3d, 0x2a, 0x09, 0x9d, 0xc1, 0x1f,
	0x8e, 0x23, 0x9a, 0x2e, 0x35, 0x53, 0x51, 0x8c, 0x93, 0x8b, 0x39, 0xe4, 0x0d, 0xc0, 0x56, 0xae,
	0x3c, 0xa3, 0x29, 0xc7, 0xe8, 0x04, 0x80, 0x4b, 0x32, 0x67, 0x98, 0xcb, 0x44, 0x14, 0xa6, 0xbf,
	0xb9, 0x02, 0x64, 0x22, 0xbc, 0x2b, 0xe8, 0x3c, 0xb1, 0x98, 0xe0, 0xbc, 0xf9, 0x0e, 0x47, 0x94,
	0x64, 0x94, 0xc7, 0x22, 0xa6, 0xa9, 0x7e, 0x58, 0x0b, 0xea, 0xc6, 0x93, 0xaa, 0x41, 0x51, 0x79,
	0xf7, 0xd0, 0xdd, 0xdf, 0x5a, 0x4c, 0xef, 0x81, 0x93, 0x7d, 0x31, 0xf3, 0x55, 0x18, 0x09, 0xaa,
	0x1d, 0x6c, 0xa5, 0x4d, 0x94, 0xe4, 0xf9, 0xd0, 0xbc, 0x5e, 0x63, 0x16, 0xbe, 0xe0, 0x43, 0x03,
	0xfb, 0xf0, 0x77, 0x43, 0x16, 0xfe, 0x6d, 0x68, 0x84, 0xb9, 0xa4, 0x58, 0x2b, 0xd0, 0xa5, 0x37,
	0x00, 0x34, 0x89, 0xd3, 0xe5, 0x43, 0xf8, 0x16, 0x93, 0x72, 0xc9, 0xdb, 0xd6, 0xb5, 0x8d, 0xf5,
	0x08, 0xfe, 0x6f, 0xd1, 0xa5, 0x3d, 0xc9, 0xa5, 0x82, 0xd7, 0xe5, 0xf9, 0x47, 0x05, 0xfe, 0xdd,
	0x6e, 0xa2, 0x9c, 0x62, 0xb6, 0x8e, 0x23, 0x8c, 0x2e, 0xa1, 0x3a, 0x95, 0x04, 0xb5, 0x86, 0x46,
	0xee, 0x65, 0xc4, 0xee, 0xd1, 0x8e, 0x9e, 0xcf, 0xf1, 0x7e, 0xa1, 0x77, 0x68, 0xef, 0x5b, 0x26,
	0xea, 0x9b, 0x6d, 0x07, 0xd2, 0x72, 0x07, 0x3f, 0x83, 0xf5, 0xe0, 0xb1, 0x85, 0x26, 0xd0, 0x28,
	0xd6, 0x8a, 0x5c, 0xb3, 0x79, 0x3b, 0x15, 0xf7, 0xf8, 0xdb, 0x3b, 0xed, 0xe3, 0x5b, 0x68, 0x06,
	0xb6, 0xb1, 0x43, 0x74, 0x6a, 0xf2, 0xbb, 0x51, 0xb8, 0x9d, 0xbd, 0xf7, 0xa5, 0xe7, 0xd8, 0xba,
	0x69, 0x3e, 0x3b, 0xe6, 0x8f, 0x5a, 0xd4, 0xd5, 0x3f, 0xba, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x46, 0x9d, 0x01, 0x95, 0x73, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
	Average(CalculatorService_AverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(req *PrimeNumberDecompositionRequest, srv CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) Average(srv CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(srv CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
