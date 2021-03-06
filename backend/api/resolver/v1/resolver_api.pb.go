// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resolver/v1/resolver_api.proto

package resolverv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type ResolveRequest struct {
	// The type URL of the desired result.
	Want string `protobuf:"bytes,1,opt,name=want,proto3" json:"want,omitempty"`
	// Filled in object schemas.
	Have *any.Any `protobuf:"bytes,2,opt,name=have,proto3" json:"have,omitempty"`
	// The maximum number of results to return.
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolveRequest) Reset()         { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()    {}
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6dc0c12b4e8ba83, []int{0}
}

func (m *ResolveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveRequest.Unmarshal(m, b)
}
func (m *ResolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveRequest.Marshal(b, m, deterministic)
}
func (m *ResolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveRequest.Merge(m, src)
}
func (m *ResolveRequest) XXX_Size() int {
	return xxx_messageInfo_ResolveRequest.Size(m)
}
func (m *ResolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveRequest proto.InternalMessageInfo

func (m *ResolveRequest) GetWant() string {
	if m != nil {
		return m.Want
	}
	return ""
}

func (m *ResolveRequest) GetHave() *any.Any {
	if m != nil {
		return m.Have
	}
	return nil
}

func (m *ResolveRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ResolveResponse struct {
	Results              []*any.Any       `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PartialFailures      []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ResolveResponse) Reset()         { *m = ResolveResponse{} }
func (m *ResolveResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveResponse) ProtoMessage()    {}
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6dc0c12b4e8ba83, []int{1}
}

func (m *ResolveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveResponse.Unmarshal(m, b)
}
func (m *ResolveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveResponse.Marshal(b, m, deterministic)
}
func (m *ResolveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveResponse.Merge(m, src)
}
func (m *ResolveResponse) XXX_Size() int {
	return xxx_messageInfo_ResolveResponse.Size(m)
}
func (m *ResolveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveResponse proto.InternalMessageInfo

func (m *ResolveResponse) GetResults() []*any.Any {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ResolveResponse) GetPartialFailures() []*status.Status {
	if m != nil {
		return m.PartialFailures
	}
	return nil
}

type SearchRequest struct {
	// The type URL of the desired result.
	Want string `protobuf:"bytes,1,opt,name=want,proto3" json:"want,omitempty"`
	// Free-form text query.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// The maximum number of results to return.
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6dc0c12b4e8ba83, []int{2}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetWant() string {
	if m != nil {
		return m.Want
	}
	return ""
}

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type SearchResponse struct {
	Results              []*any.Any       `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PartialFailures      []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6dc0c12b4e8ba83, []int{3}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetResults() []*any.Any {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SearchResponse) GetPartialFailures() []*status.Status {
	if m != nil {
		return m.PartialFailures
	}
	return nil
}

type GetObjectSchemasRequest struct {
	TypeUrl              string   `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectSchemasRequest) Reset()         { *m = GetObjectSchemasRequest{} }
func (m *GetObjectSchemasRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectSchemasRequest) ProtoMessage()    {}
func (*GetObjectSchemasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6dc0c12b4e8ba83, []int{4}
}

func (m *GetObjectSchemasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectSchemasRequest.Unmarshal(m, b)
}
func (m *GetObjectSchemasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectSchemasRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectSchemasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectSchemasRequest.Merge(m, src)
}
func (m *GetObjectSchemasRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectSchemasRequest.Size(m)
}
func (m *GetObjectSchemasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectSchemasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectSchemasRequest proto.InternalMessageInfo

func (m *GetObjectSchemasRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

type GetObjectSchemasResponse struct {
	TypeUrl              string    `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Schemas              []*Schema `protobuf:"bytes,2,rep,name=schemas,proto3" json:"schemas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetObjectSchemasResponse) Reset()         { *m = GetObjectSchemasResponse{} }
func (m *GetObjectSchemasResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectSchemasResponse) ProtoMessage()    {}
func (*GetObjectSchemasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6dc0c12b4e8ba83, []int{5}
}

func (m *GetObjectSchemasResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectSchemasResponse.Unmarshal(m, b)
}
func (m *GetObjectSchemasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectSchemasResponse.Marshal(b, m, deterministic)
}
func (m *GetObjectSchemasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectSchemasResponse.Merge(m, src)
}
func (m *GetObjectSchemasResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectSchemasResponse.Size(m)
}
func (m *GetObjectSchemasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectSchemasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectSchemasResponse proto.InternalMessageInfo

func (m *GetObjectSchemasResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *GetObjectSchemasResponse) GetSchemas() []*Schema {
	if m != nil {
		return m.Schemas
	}
	return nil
}

func init() {
	proto.RegisterType((*ResolveRequest)(nil), "clutch.resolver.v1.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "clutch.resolver.v1.ResolveResponse")
	proto.RegisterType((*SearchRequest)(nil), "clutch.resolver.v1.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "clutch.resolver.v1.SearchResponse")
	proto.RegisterType((*GetObjectSchemasRequest)(nil), "clutch.resolver.v1.GetObjectSchemasRequest")
	proto.RegisterType((*GetObjectSchemasResponse)(nil), "clutch.resolver.v1.GetObjectSchemasResponse")
}

func init() {
	proto.RegisterFile("resolver/v1/resolver_api.proto", fileDescriptor_b6dc0c12b4e8ba83)
}

var fileDescriptor_b6dc0c12b4e8ba83 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x95, 0xb3, 0x76, 0x2d, 0xbf, 0xb1, 0x3f, 0x32, 0x95, 0x9a, 0x86, 0x0d, 0x15, 0x03, 0x52,
	0x35, 0x90, 0xa3, 0x16, 0x4e, 0x93, 0x76, 0x58, 0x0f, 0x20, 0x4e, 0xa0, 0x54, 0x5c, 0xb8, 0x54,
	0x6e, 0xf0, 0xda, 0x40, 0x96, 0x64, 0xb6, 0x93, 0x29, 0x27, 0x10, 0x5f, 0x81, 0x33, 0x42, 0xe2,
	0xc8, 0xd7, 0xe1, 0x13, 0x20, 0xf1, 0x29, 0x76, 0x42, 0x4d, 0xec, 0x29, 0xed, 0xda, 0x69, 0x37,
	0x6e, 0x8e, 0xdf, 0x7b, 0x7e, 0xef, 0xc5, 0x3f, 0xc3, 0x03, 0xc1, 0x65, 0x1c, 0x66, 0x5c, 0xb8,
	0x59, 0xdf, 0x35, 0xeb, 0x31, 0x4b, 0x02, 0x9a, 0x88, 0x58, 0xc5, 0x18, 0xfb, 0x61, 0xaa, 0xfc,
	0x19, 0x35, 0x10, 0xcd, 0xfa, 0xce, 0xfe, 0x34, 0x8e, 0xa7, 0x21, 0x77, 0x59, 0x12, 0xb8, 0x2c,
	0x8a, 0x62, 0xc5, 0x54, 0x10, 0x47, 0xb2, 0x54, 0x38, 0x1d, 0x8d, 0x16, 0x5f, 0x93, 0xf4, 0xd4,
	0x65, 0x51, 0xae, 0xa1, 0xb6, 0x86, 0x44, 0xe2, 0xbb, 0x52, 0x31, 0x95, 0x1a, 0x4d, 0x3b, 0x63,
	0x61, 0xf0, 0x81, 0x29, 0xee, 0x9a, 0x85, 0x06, 0xec, 0xb9, 0x47, 0xd6, 0x5f, 0x61, 0x63, 0x57,
	0x83, 0x4b, 0x7f, 0xc6, 0xcf, 0x58, 0x89, 0x90, 0x0b, 0xd8, 0xf1, 0x4a, 0xcc, 0xe3, 0xe7, 0x29,
	0x97, 0x0a, 0xdf, 0x87, 0xda, 0x05, 0x8b, 0x94, 0x8d, 0xba, 0xa8, 0x77, 0x67, 0xd8, 0xb8, 0x1c,
	0xd6, 0x84, 0xd5, 0x45, 0x5e, 0xb1, 0x89, 0x07, 0x50, 0x9b, 0xb1, 0x8c, 0xdb, 0x56, 0x17, 0xf5,
	0xb6, 0x06, 0x2d, 0x5a, 0x66, 0xa4, 0x26, 0x3e, 0x3d, 0x89, 0xf2, 0x61, 0xf3, 0x72, 0x58, 0xff,
	0x89, 0xac, 0x26, 0xf2, 0x0a, 0x2e, 0x6e, 0x41, 0x3d, 0x0c, 0xce, 0x02, 0x65, 0x6f, 0x74, 0x51,
	0x6f, 0xdb, 0x2b, 0x3f, 0xc8, 0x17, 0x04, 0xbb, 0x57, 0xce, 0x32, 0x89, 0x23, 0xc9, 0x31, 0x85,
	0x86, 0xe0, 0x32, 0x0d, 0x95, 0xb4, 0x51, 0x77, 0x63, 0x9d, 0x81, 0x67, 0x48, 0xf8, 0x18, 0xf6,
	0x12, 0x26, 0x54, 0xc0, 0xc2, 0xf1, 0x29, 0x0b, 0xc2, 0x54, 0x70, 0x69, 0x5b, 0x85, 0x10, 0x1b,
	0xa1, 0x48, 0x7c, 0x3a, 0x2a, 0xfe, 0x9e, 0xb7, 0xab, 0xb9, 0x2f, 0x35, 0x95, 0x30, 0xd8, 0x1e,
	0x71, 0x26, 0xfc, 0xd9, 0xad, 0xaa, 0x1f, 0x40, 0xfd, 0x3c, 0xe5, 0x22, 0x2f, 0xba, 0x57, 0xd0,
	0x72, 0x77, 0x4d, 0xcb, 0xcf, 0xb0, 0x63, 0x2c, 0xfe, 0x4f, 0xc7, 0x63, 0x68, 0xbf, 0xe2, 0xea,
	0xcd, 0xe4, 0x23, 0xf7, 0xd5, 0xa8, 0xb8, 0x78, 0x69, 0xda, 0x12, 0x68, 0xaa, 0x3c, 0xe1, 0xe3,
	0x54, 0x84, 0xcb, 0x8d, 0x1b, 0x73, 0xe0, 0x9d, 0x08, 0xc9, 0x27, 0xb0, 0xaf, 0xcb, 0x75, 0x93,
	0xce, 0xb2, 0xfe, 0x4a, 0x86, 0x5f, 0x40, 0xa3, 0x9c, 0x32, 0x93, 0xd5, 0xa1, 0xd7, 0x9f, 0x06,
	0x2d, 0x0f, 0xf4, 0x0c, 0x75, 0xf0, 0x63, 0x03, 0xb6, 0xf4, 0x48, 0x88, 0x93, 0xb7, 0xaf, 0xf1,
	0x77, 0x04, 0x7b, 0xcb, 0xee, 0xf8, 0xe9, 0xaa, 0x93, 0xd6, 0x54, 0x74, 0x9e, 0xdd, 0x8e, 0x5c,
	0x16, 0x22, 0xf4, 0xd7, 0x9f, 0x7d, 0xab, 0x69, 0x7d, 0xfd, 0xfd, 0xf7, 0x9b, 0x45, 0xc8, 0x41,
	0xf5, 0x91, 0xbb, 0xd3, 0x25, 0xd1, 0x11, 0x3a, 0xc4, 0x29, 0x6c, 0x96, 0x97, 0x8b, 0x1f, 0xae,
	0xac, 0x57, 0x9d, 0x2d, 0x87, 0xdc, 0x44, 0xd1, 0x01, 0x1e, 0x57, 0x02, 0xd8, 0xe4, 0xde, 0x42,
	0x00, 0x59, 0x50, 0xe7, 0xb6, 0x39, 0x34, 0xf4, 0x5f, 0xc2, 0x2b, 0x0f, 0x5d, 0x7c, 0xcf, 0xce,
	0xa3, 0x1b, 0x39, 0xda, 0xf9, 0x49, 0xc5, 0xb9, 0x43, 0x5a, 0x0b, 0xce, 0x7a, 0x71, 0x84, 0x0e,
	0x87, 0x77, 0xdf, 0x83, 0xd9, 0xce, 0xfa, 0x93, 0xcd, 0x62, 0x62, 0x9f, 0xff, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x72, 0x1b, 0xa5, 0x17, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResolverAPIClient is the client API for ResolverAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResolverAPIClient interface {
	GetObjectSchemas(ctx context.Context, in *GetObjectSchemasRequest, opts ...grpc.CallOption) (*GetObjectSchemasResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type resolverAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewResolverAPIClient(cc grpc.ClientConnInterface) ResolverAPIClient {
	return &resolverAPIClient{cc}
}

func (c *resolverAPIClient) GetObjectSchemas(ctx context.Context, in *GetObjectSchemasRequest, opts ...grpc.CallOption) (*GetObjectSchemasResponse, error) {
	out := new(GetObjectSchemasResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/GetObjectSchemas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolverAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolverAPIClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolverAPIServer is the server API for ResolverAPI service.
type ResolverAPIServer interface {
	GetObjectSchemas(context.Context, *GetObjectSchemasRequest) (*GetObjectSchemasResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
}

// UnimplementedResolverAPIServer can be embedded to have forward compatible implementations.
type UnimplementedResolverAPIServer struct {
}

func (*UnimplementedResolverAPIServer) GetObjectSchemas(ctx context.Context, req *GetObjectSchemasRequest) (*GetObjectSchemasResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetObjectSchemas not implemented")
}
func (*UnimplementedResolverAPIServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedResolverAPIServer) Resolve(ctx context.Context, req *ResolveRequest) (*ResolveResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Resolve not implemented")
}

func RegisterResolverAPIServer(s *grpc.Server, srv ResolverAPIServer) {
	s.RegisterService(&_ResolverAPI_serviceDesc, srv)
}

func _ResolverAPI_GetObjectSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).GetObjectSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/GetObjectSchemas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).GetObjectSchemas(ctx, req.(*GetObjectSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolverAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolverAPI_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResolverAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.resolver.v1.ResolverAPI",
	HandlerType: (*ResolverAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObjectSchemas",
			Handler:    _ResolverAPI_GetObjectSchemas_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ResolverAPI_Search_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _ResolverAPI_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resolver/v1/resolver_api.proto",
}
