// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/example.proto

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "zap"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             string               `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	EmployeeNumber       int64                `protobuf:"varint,4,opt,name=employee_number,json=employeeNumber" json:"employee_number,omitempty"`
	PhysicalDesk         string               `protobuf:"bytes,6,opt,name=physical_desk,json=physicalDesk" json:"physical_desk,omitempty"`
	Service              *ServiceMsg          `protobuf:"bytes,8,opt,name=service" json:"service,omitempty"`
	Blocked              []string             `protobuf:"bytes,9,rep,name=blocked" json:"blocked,omitempty"`
	Extra                map[string]string    `protobuf:"bytes,10,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HireDate             *timestamp.Timestamp `protobuf:"bytes,11,opt,name=hire_date,json=hireDate" json:"hire_date,omitempty"`
	List                 []*ServiceMsg        `protobuf:"bytes,12,rep,name=list" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_94596b6a55643ede, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmployeeNumber() int64 {
	if m != nil {
		return m.EmployeeNumber
	}
	return 0
}

func (m *User) GetPhysicalDesk() string {
	if m != nil {
		return m.PhysicalDesk
	}
	return ""
}

func (m *User) GetService() *ServiceMsg {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *User) GetBlocked() []string {
	if m != nil {
		return m.Blocked
	}
	return nil
}

func (m *User) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *User) GetHireDate() *timestamp.Timestamp {
	if m != nil {
		return m.HireDate
	}
	return nil
}

func (m *User) GetList() []*ServiceMsg {
	if m != nil {
		return m.List
	}
	return nil
}

type ServiceMsg struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceMsg) Reset()         { *m = ServiceMsg{} }
func (m *ServiceMsg) String() string { return proto.CompactTextString(m) }
func (*ServiceMsg) ProtoMessage()    {}
func (*ServiceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_94596b6a55643ede, []int{1}
}
func (m *ServiceMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceMsg.Unmarshal(m, b)
}
func (m *ServiceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceMsg.Marshal(b, m, deterministic)
}
func (dst *ServiceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMsg.Merge(dst, src)
}
func (m *ServiceMsg) XXX_Size() int {
	return xxx_messageInfo_ServiceMsg.Size(m)
}
func (m *ServiceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMsg proto.InternalMessageInfo

func (m *ServiceMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "example.User")
	proto.RegisterMapType((map[string]string)(nil), "example.User.ExtraEntry")
	proto.RegisterType((*ServiceMsg)(nil), "example.ServiceMsg")
}

func init() { proto.RegisterFile("example/example.proto", fileDescriptor_example_94596b6a55643ede) }

var fileDescriptor_example_94596b6a55643ede = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x3f, 0xdb, 0xdd, 0xcc, 0xb6, 0x05, 0x19, 0x90, 0xac, 0x5c, 0x88, 0x16, 0x24,
	0x72, 0xc1, 0x41, 0xe5, 0x40, 0xc5, 0x11, 0xb5, 0x47, 0x7a, 0x30, 0x70, 0x8e, 0x9c, 0xcd, 0x74,
	0x6b, 0xc5, 0xd9, 0x44, 0xb6, 0xb7, 0x6a, 0xfa, 0x2e, 0xbc, 0x07, 0xcf, 0xc0, 0x43, 0x21, 0x14,
	0x27, 0x66, 0x91, 0xe0, 0x64, 0xcf, 0x37, 0x3f, 0xcd, 0x8c, 0xbe, 0x0f, 0x5e, 0xe0, 0x83, 0x68,
	0x7b, 0x85, 0xc5, 0xfc, 0xb2, 0x5e, 0x77, 0xb6, 0x23, 0xcb, 0xb9, 0x4c, 0x5f, 0xee, 0xba, 0x6e,
	0xa7, 0xb0, 0x70, 0x72, 0x75, 0xb8, 0x2d, 0xac, 0x6c, 0xd1, 0x58, 0xd1, 0xf6, 0x13, 0x99, 0x9e,
	0x3d, 0x8a, 0xbe, 0x78, 0x14, 0x73, 0xb9, 0xf9, 0x19, 0x41, 0xfc, 0xcd, 0xa0, 0x26, 0xe7, 0x10,
	0xca, 0x9a, 0x06, 0x59, 0x90, 0x27, 0x3c, 0x94, 0x35, 0x79, 0x0d, 0x70, 0x2b, 0xb5, 0xb1, 0xe5,
	0x5e, 0xb4, 0x48, 0xc3, 0x51, 0xff, 0xb4, 0xf8, 0xf1, 0xeb, 0x7b, 0x14, 0xf0, 0xc4, 0x35, 0x6e,
	0x44, 0x8b, 0x64, 0x03, 0x89, 0x12, 0x1e, 0x8a, 0xfe, 0x86, 0x56, 0xa3, 0xee, 0x18, 0x06, 0x4f,
	0xb0, 0xed, 0x55, 0x37, 0x20, 0x96, 0xfb, 0x43, 0x5b, 0xa1, 0xa6, 0x71, 0x16, 0xe4, 0x91, 0x27,
	0xcf, 0x7d, 0xf7, 0xc6, 0x35, 0xc9, 0x2b, 0x38, 0xeb, 0xef, 0x06, 0x23, 0xb7, 0x42, 0x95, 0x35,
	0x9a, 0x86, 0x9e, 0xb8, 0xa3, 0x4e, 0xbd, 0x78, 0x85, 0xa6, 0x21, 0x6f, 0x61, 0x69, 0x50, 0xdf,
	0xcb, 0x2d, 0xd2, 0x55, 0x16, 0xe4, 0xeb, 0x8b, 0x67, 0xcc, 0x3b, 0xf2, 0x65, 0xd2, 0x3f, 0x9b,
	0x1d, 0xf7, 0x0c, 0xa1, 0xb0, 0xac, 0x54, 0xb7, 0x6d, 0xb0, 0xa6, 0x49, 0x16, 0xe5, 0x09, 0xf7,
	0x25, 0x61, 0xb0, 0xc0, 0x07, 0xab, 0x05, 0x85, 0x2c, 0xca, 0xd7, 0x17, 0xf4, 0xcf, 0x98, 0xd1,
	0x15, 0x76, 0x3d, 0xb6, 0xae, 0xf7, 0x56, 0x0f, 0x7c, 0xc2, 0xc8, 0x07, 0x48, 0xee, 0xa4, 0xc6,
	0xb2, 0x16, 0x16, 0xe9, 0xda, 0xad, 0x4e, 0xd9, 0x64, 0x3a, 0xf3, 0xa6, 0xb3, 0xaf, 0xde, 0x74,
	0xbe, 0x1a, 0xe1, 0x2b, 0x61, 0x91, 0xbc, 0x81, 0x58, 0x49, 0x63, 0xe9, 0xa9, 0xdb, 0xf3, 0xdf,
	0x73, 0x1d, 0x90, 0x5e, 0x02, 0x1c, 0xd7, 0x92, 0xa7, 0x10, 0x35, 0x38, 0xcc, 0xc1, 0x8c, 0x5f,
	0xf2, 0x1c, 0x16, 0xf7, 0x42, 0x1d, 0xe6, 0x50, 0xf8, 0x54, 0x7c, 0x0c, 0x2f, 0x83, 0xcd, 0x3b,
	0x80, 0xe3, 0xb4, 0x7f, 0x12, 0x25, 0x10, 0x1f, 0xb3, 0xe4, 0xee, 0x5f, 0x9d, 0xb8, 0x93, 0xdf,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x5a, 0xcf, 0xe1, 0x57, 0x02, 0x00, 0x00,
}
