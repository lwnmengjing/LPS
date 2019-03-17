// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/domain/domain.proto

package go_micro_srv_LPS_Domain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Service struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{1}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Metadata struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{2}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metadata) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateModel struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Service              []string    `protobuf:"bytes,2,rep,name=service,proto3" json:"service,omitempty"`
	Metadata             []*Metadata `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	ExpiredAt            string      `protobuf:"bytes,4,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateModel) Reset()         { *m = CreateModel{} }
func (m *CreateModel) String() string { return proto.CompactTextString(m) }
func (*CreateModel) ProtoMessage()    {}
func (*CreateModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{3}
}

func (m *CreateModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModel.Unmarshal(m, b)
}
func (m *CreateModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModel.Marshal(b, m, deterministic)
}
func (m *CreateModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModel.Merge(m, src)
}
func (m *CreateModel) XXX_Size() int {
	return xxx_messageInfo_CreateModel.Size(m)
}
func (m *CreateModel) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModel.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModel proto.InternalMessageInfo

func (m *CreateModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateModel) GetService() []string {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *CreateModel) GetMetadata() []*Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateModel) GetExpiredAt() string {
	if m != nil {
		return m.ExpiredAt
	}
	return ""
}

type Model struct {
	Id                   uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Service              []string    `protobuf:"bytes,3,rep,name=service,proto3" json:"service,omitempty"`
	Metadata             []*Metadata `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty"`
	ExpiredAt            string      `protobuf:"bytes,5,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{4}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetService() []string {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Model) GetMetadata() []*Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Model) GetExpiredAt() string {
	if m != nil {
		return m.ExpiredAt
	}
	return ""
}

type ModelId struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelId) Reset()         { *m = ModelId{} }
func (m *ModelId) String() string { return proto.CompactTextString(m) }
func (*ModelId) ProtoMessage()    {}
func (*ModelId) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{5}
}

func (m *ModelId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelId.Unmarshal(m, b)
}
func (m *ModelId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelId.Marshal(b, m, deterministic)
}
func (m *ModelId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelId.Merge(m, src)
}
func (m *ModelId) XXX_Size() int {
	return xxx_messageInfo_ModelId.Size(m)
}
func (m *ModelId) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelId.DiscardUnknown(m)
}

var xxx_messageInfo_ModelId proto.InternalMessageInfo

func (m *ModelId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Page struct {
	PageSize             int64    `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	TotalItem            int64    `protobuf:"varint,2,opt,name=totalItem,proto3" json:"totalItem,omitempty"`
	Page                 int64    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Pages                int64    `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{6}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Page) GetTotalItem() int64 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *Page) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Page) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

type Models struct {
	Data                 []*Model `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page                 *Page    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Models) Reset()         { *m = Models{} }
func (m *Models) String() string { return proto.CompactTextString(m) }
func (*Models) ProtoMessage()    {}
func (*Models) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{7}
}

func (m *Models) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Models.Unmarshal(m, b)
}
func (m *Models) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Models.Marshal(b, m, deterministic)
}
func (m *Models) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Models.Merge(m, src)
}
func (m *Models) XXX_Size() int {
	return xxx_messageInfo_Models.Size(m)
}
func (m *Models) XXX_DiscardUnknown() {
	xxx_messageInfo_Models.DiscardUnknown(m)
}

var xxx_messageInfo_Models proto.InternalMessageInfo

func (m *Models) GetData() []*Model {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Models) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

type Condition struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ExpiredAt            string   `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	PageSize             int64    `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Offset               int64    `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa37030272e0e28, []int{8}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Condition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Condition) GetExpiredAt() string {
	if m != nil {
		return m.ExpiredAt
	}
	return ""
}

func (m *Condition) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Condition) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "go.micro.srv.LPS.Domain.Empty")
	proto.RegisterType((*Service)(nil), "go.micro.srv.LPS.Domain.Service")
	proto.RegisterType((*Metadata)(nil), "go.micro.srv.LPS.Domain.Metadata")
	proto.RegisterType((*CreateModel)(nil), "go.micro.srv.LPS.Domain.CreateModel")
	proto.RegisterType((*Model)(nil), "go.micro.srv.LPS.Domain.Model")
	proto.RegisterType((*ModelId)(nil), "go.micro.srv.LPS.Domain.ModelId")
	proto.RegisterType((*Page)(nil), "go.micro.srv.LPS.Domain.Page")
	proto.RegisterType((*Models)(nil), "go.micro.srv.LPS.Domain.Models")
	proto.RegisterType((*Condition)(nil), "go.micro.srv.LPS.Domain.Condition")
}

func init() { proto.RegisterFile("proto/domain/domain.proto", fileDescriptor_bfa37030272e0e28) }

var fileDescriptor_bfa37030272e0e28 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x38, 0x4d, 0x9b, 0x57, 0x81, 0xd0, 0x13, 0x82, 0xac, 0xa2, 0x50, 0x2c, 0x0e,
	0x3d, 0x05, 0x11, 0xce, 0x1c, 0xd0, 0xc6, 0xa1, 0x68, 0x45, 0x55, 0x2a, 0xb8, 0x22, 0x33, 0xbf,
	0x55, 0x86, 0xa6, 0x8e, 0x12, 0x53, 0x18, 0x47, 0xbe, 0x01, 0x47, 0xce, 0x7c, 0x51, 0x14, 0x3b,
	0x2d, 0xac, 0x5a, 0xb6, 0x4a, 0x3b, 0xd9, 0xef, 0xd9, 0x7e, 0xef, 0xf7, 0xff, 0xdb, 0x09, 0x1c,
	0x15, 0xa5, 0x36, 0xfa, 0xb9, 0xd4, 0xb9, 0x50, 0xeb, 0x66, 0x48, 0x6c, 0x0e, 0x1f, 0x2e, 0x75,
	0x92, 0xab, 0xb3, 0x52, 0x27, 0x55, 0xb9, 0x49, 0x4e, 0xe7, 0x8b, 0xe4, 0xc4, 0x2e, 0xf3, 0x1e,
	0x74, 0xdf, 0xe4, 0x85, 0xb9, 0xe0, 0x23, 0xe8, 0x2d, 0xa8, 0xdc, 0xa8, 0x33, 0x42, 0x84, 0x60,
	0x2d, 0x72, 0x8a, 0xbd, 0xb1, 0x37, 0x89, 0x32, 0x3b, 0xe7, 0x29, 0xf4, 0x67, 0x64, 0x84, 0x14,
	0x46, 0xe0, 0x3d, 0x60, 0x5f, 0xe8, 0xa2, 0x59, 0xae, 0xa7, 0x78, 0x1f, 0xba, 0x1b, 0xb1, 0xfa,
	0x4a, 0xb1, 0x6f, 0x73, 0x2e, 0xe0, 0xbf, 0x3d, 0x18, 0x1c, 0x97, 0x24, 0x0c, 0xcd, 0xb4, 0xa4,
	0xd5, 0x55, 0x75, 0x31, 0x86, 0x5e, 0xe5, 0xda, 0xc6, 0xfe, 0x98, 0x4d, 0xa2, 0x6c, 0x1b, 0xe2,
	0x2b, 0xe8, 0xe7, 0x4d, 0xc7, 0x98, 0x8d, 0xd9, 0x64, 0x90, 0x3e, 0x4d, 0x5a, 0x54, 0x24, 0x5b,
	0xb4, 0x6c, 0x77, 0x04, 0x47, 0x00, 0xf4, 0xbd, 0x50, 0x25, 0xc9, 0x8f, 0xc2, 0xc4, 0x81, 0x6d,
	0x19, 0x35, 0x99, 0xd7, 0x86, 0xff, 0xf1, 0xa0, 0xeb, 0xa8, 0xee, 0x82, 0xaf, 0xa4, 0x65, 0xba,
	0x93, 0xf9, 0x4a, 0xee, 0x28, 0xfd, 0xab, 0x29, 0x59, 0x3b, 0x65, 0x70, 0x5b, 0xca, 0xee, 0x3e,
	0xe5, 0x11, 0xf4, 0x2c, 0xe4, 0x54, 0xee, 0x63, 0xf2, 0xcf, 0x10, 0xcc, 0xc5, 0x92, 0x70, 0x08,
	0xfd, 0x42, 0x2c, 0x69, 0xa1, 0x7e, 0x38, 0x63, 0x59, 0xb6, 0x8b, 0xf1, 0x11, 0x44, 0x46, 0x1b,
	0xb1, 0x9a, 0x1a, 0xca, 0xad, 0x1e, 0x96, 0xfd, 0x4b, 0xd4, 0x42, 0xeb, 0x9d, 0x31, 0xb3, 0x0b,
	0x76, 0x5e, 0x5f, 0x64, 0x3d, 0x56, 0xd6, 0x30, 0x96, 0xb9, 0x80, 0x6b, 0x08, 0x2d, 0x46, 0x85,
	0x29, 0x04, 0x56, 0xaa, 0x67, 0xa5, 0x3e, 0x6e, 0x97, 0x5a, 0x6f, 0xcf, 0xec, 0x5e, 0x7c, 0xd1,
	0xf4, 0xa9, 0x01, 0x06, 0xe9, 0xa8, 0xf5, 0x4c, 0x2d, 0xc7, 0x61, 0xf0, 0x9f, 0x1e, 0x44, 0xc7,
	0x7a, 0x2d, 0x95, 0x51, 0x7a, 0x7d, 0xd0, 0x0d, 0x5d, 0x36, 0x92, 0xed, 0x19, 0x79, 0xc9, 0xa5,
	0x60, 0xcf, 0xa5, 0x07, 0x10, 0xea, 0xf3, 0xf3, 0x8a, 0x9c, 0xff, 0x2c, 0x6b, 0xa2, 0xf4, 0x17,
	0x83, 0xd0, 0xa1, 0xe1, 0x1c, 0x42, 0xf7, 0x90, 0xf1, 0x59, 0x2b, 0xfe, 0x7f, 0x2f, 0x7d, 0xd8,
	0x6e, 0x8c, 0xfb, 0xd8, 0x3a, 0xf8, 0x16, 0xc2, 0xf7, 0x85, 0xac, 0x2b, 0xde, 0x60, 0xe2, 0x41,
	0xb5, 0x82, 0x0f, 0x8a, 0xbe, 0xe1, 0xf8, 0xfa, 0x4a, 0x53, 0x39, 0xbc, 0xa1, 0x17, 0xef, 0xe0,
	0x3b, 0x08, 0x4f, 0x68, 0x45, 0x86, 0x90, 0xb7, 0x2b, 0xdd, 0xde, 0xcc, 0x01, 0x6c, 0x33, 0x08,
	0x4e, 0x55, 0x65, 0x0e, 0xaa, 0xf6, 0xe4, 0x7a, 0xba, 0x8a, 0x77, 0x3e, 0x85, 0xf6, 0x77, 0xf6,
	0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xf6, 0x90, 0xfa, 0xeb, 0x04, 0x00, 0x00,
}
