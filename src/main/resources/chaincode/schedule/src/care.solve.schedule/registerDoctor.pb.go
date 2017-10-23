// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registerDoctor.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	registerDoctor.proto

It has these top-level messages:
	PatientPublic
	DoctorPublic
	DoctorCollection
	Schedule
	Slot
	RegistrationInfo
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Slot_Type int32

const (
	Slot_FREE Slot_Type = 0
	Slot_BUSY Slot_Type = 1
)

var Slot_Type_name = map[int32]string{
	0: "FREE",
	1: "BUSY",
}
var Slot_Type_value = map[string]int32{
	"FREE": 0,
	"BUSY": 1,
}

func (x Slot_Type) String() string {
	return proto.EnumName(Slot_Type_name, int32(x))
}
func (Slot_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type PatientPublic struct {
	PatientId string `protobuf:"bytes,1,opt,name=PatientId" json:"PatientId,omitempty"`
}

func (m *PatientPublic) Reset()                    { *m = PatientPublic{} }
func (m *PatientPublic) String() string            { return proto.CompactTextString(m) }
func (*PatientPublic) ProtoMessage()               {}
func (*PatientPublic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PatientPublic) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

type DoctorPublic struct {
	DoctorId  string `protobuf:"bytes,1,opt,name=DoctorId" json:"DoctorId,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
}

func (m *DoctorPublic) Reset()                    { *m = DoctorPublic{} }
func (m *DoctorPublic) String() string            { return proto.CompactTextString(m) }
func (*DoctorPublic) ProtoMessage()               {}
func (*DoctorPublic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DoctorPublic) GetDoctorId() string {
	if m != nil {
		return m.DoctorId
	}
	return ""
}

func (m *DoctorPublic) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *DoctorPublic) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type DoctorCollection struct {
	Doctors []*DoctorPublic `protobuf:"bytes,1,rep,name=doctors" json:"doctors,omitempty"`
}

func (m *DoctorCollection) Reset()                    { *m = DoctorCollection{} }
func (m *DoctorCollection) String() string            { return proto.CompactTextString(m) }
func (*DoctorCollection) ProtoMessage()               {}
func (*DoctorCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DoctorCollection) GetDoctors() []*DoctorPublic {
	if m != nil {
		return m.Doctors
	}
	return nil
}

type Schedule struct {
	ScheduleId string  `protobuf:"bytes,1,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
	DoctorId   string  `protobuf:"bytes,2,opt,name=DoctorId" json:"DoctorId,omitempty"`
	Slots      []*Slot `protobuf:"bytes,3,rep,name=slots" json:"slots,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Schedule) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *Schedule) GetDoctorId() string {
	if m != nil {
		return m.DoctorId
	}
	return ""
}

func (m *Schedule) GetSlots() []*Slot {
	if m != nil {
		return m.Slots
	}
	return nil
}

type Slot struct {
	SlotId           string            `protobuf:"bytes,1,opt,name=SlotId" json:"SlotId,omitempty"`
	TimeStart        uint64            `protobuf:"varint,2,opt,name=TimeStart" json:"TimeStart,omitempty"`
	TimeFinish       uint64            `protobuf:"varint,3,opt,name=TimeFinish" json:"TimeFinish,omitempty"`
	Avaliable        Slot_Type         `protobuf:"varint,4,opt,name=avaliable,enum=main.Slot_Type" json:"avaliable,omitempty"`
	RegistrationInfo *RegistrationInfo `protobuf:"bytes,5,opt,name=registrationInfo" json:"registrationInfo,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Slot) GetSlotId() string {
	if m != nil {
		return m.SlotId
	}
	return ""
}

func (m *Slot) GetTimeStart() uint64 {
	if m != nil {
		return m.TimeStart
	}
	return 0
}

func (m *Slot) GetTimeFinish() uint64 {
	if m != nil {
		return m.TimeFinish
	}
	return 0
}

func (m *Slot) GetAvaliable() Slot_Type {
	if m != nil {
		return m.Avaliable
	}
	return Slot_FREE
}

func (m *Slot) GetRegistrationInfo() *RegistrationInfo {
	if m != nil {
		return m.RegistrationInfo
	}
	return nil
}

type RegistrationInfo struct {
	Description string `protobuf:"bytes,1,opt,name=Description" json:"Description,omitempty"`
	PatientId   string `protobuf:"bytes,2,opt,name=PatientId" json:"PatientId,omitempty"`
}

func (m *RegistrationInfo) Reset()                    { *m = RegistrationInfo{} }
func (m *RegistrationInfo) String() string            { return proto.CompactTextString(m) }
func (*RegistrationInfo) ProtoMessage()               {}
func (*RegistrationInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegistrationInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegistrationInfo) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func init() {
	proto.RegisterType((*PatientPublic)(nil), "main.PatientPublic")
	proto.RegisterType((*DoctorPublic)(nil), "main.DoctorPublic")
	proto.RegisterType((*DoctorCollection)(nil), "main.DoctorCollection")
	proto.RegisterType((*Schedule)(nil), "main.Schedule")
	proto.RegisterType((*Slot)(nil), "main.Slot")
	proto.RegisterType((*RegistrationInfo)(nil), "main.RegistrationInfo")
	proto.RegisterEnum("main.Slot_Type", Slot_Type_name, Slot_Type_value)
}

func init() { proto.RegisterFile("registerDoctor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x5d, 0xcb, 0xd3, 0x30,
	0x14, 0xb6, 0x5b, 0xdf, 0xd7, 0xed, 0x4c, 0x5f, 0x4b, 0x90, 0x51, 0x87, 0x48, 0xe9, 0x55, 0x2f,
	0x5c, 0xc1, 0xf9, 0x07, 0xa4, 0x6e, 0x83, 0x81, 0xc8, 0x48, 0xe7, 0x85, 0xde, 0xa5, 0x6d, 0x74,
	0xc1, 0xac, 0x19, 0x49, 0x36, 0xd8, 0x6f, 0xf6, 0x4f, 0x48, 0x9a, 0x7e, 0xad, 0xef, 0x55, 0xfb,
	0x3c, 0xe7, 0x9c, 0xe7, 0x7c, 0xe4, 0x81, 0xb7, 0x92, 0xfe, 0x61, 0x4a, 0x53, 0xb9, 0x16, 0xb9,
	0x16, 0x32, 0x3e, 0x4b, 0xa1, 0x05, 0x72, 0x4f, 0x84, 0x95, 0xe1, 0x12, 0x5e, 0xef, 0x89, 0x66,
	0xb4, 0xd4, 0xfb, 0x4b, 0xc6, 0x59, 0x8e, 0xde, 0xc3, 0xb4, 0x26, 0x76, 0x85, 0xef, 0x04, 0x4e,
	0x34, 0xc5, 0x1d, 0x11, 0x16, 0xf0, 0xca, 0x8a, 0xd4, 0xd9, 0x0b, 0x98, 0x58, 0xdc, 0x26, 0xb7,
	0xd8, 0x28, 0x6d, 0x99, 0x54, 0xfa, 0x3b, 0x39, 0x51, 0x7f, 0x64, 0x95, 0x5a, 0xc2, 0x54, 0x7e,
	0x23, 0x75, 0x70, 0x6c, 0x2b, 0x1b, 0x1c, 0x7e, 0x01, 0xcf, 0xaa, 0x7c, 0x15, 0x9c, 0xd3, 0x5c,
	0x33, 0x51, 0xa2, 0x8f, 0xf0, 0xb2, 0xa8, 0x38, 0xe5, 0x3b, 0xc1, 0x38, 0x9a, 0xad, 0x50, 0x6c,
	0x16, 0x88, 0xfb, 0xe3, 0xe0, 0x26, 0x25, 0x3c, 0xc2, 0x24, 0xcd, 0x8f, 0xb4, 0xb8, 0x70, 0x8a,
	0x3e, 0x00, 0x34, 0xff, 0xed, 0x94, 0x3d, 0xe6, 0x6e, 0x87, 0xd1, 0x60, 0x87, 0x00, 0x1e, 0x14,
	0x17, 0x5a, 0xf9, 0xe3, 0xaa, 0x27, 0xd8, 0x9e, 0x29, 0x17, 0x1a, 0xdb, 0x40, 0xf8, 0xcf, 0x01,
	0xd7, 0x60, 0x34, 0x87, 0x47, 0xf3, 0x6d, 0x5b, 0xd4, 0xc8, 0x9c, 0xe1, 0xc0, 0x4e, 0x34, 0xd5,
	0x44, 0xea, 0x4a, 0xdf, 0xc5, 0x1d, 0x61, 0x86, 0x33, 0x60, 0xcb, 0x4a, 0xa6, 0x8e, 0xd5, 0x21,
	0x5c, 0xdc, 0x63, 0xd0, 0x12, 0xa6, 0xe4, 0x4a, 0x38, 0x23, 0x19, 0xa7, 0xbe, 0x1b, 0x38, 0xd1,
	0xd3, 0xea, 0x4d, 0x37, 0x44, 0x7c, 0xb8, 0x9d, 0x29, 0xee, 0x32, 0x50, 0x02, 0x9e, 0x7d, 0x6c,
	0x49, 0xcc, 0xd5, 0x76, 0xe5, 0x6f, 0xe1, 0x3f, 0x04, 0x4e, 0x34, 0x5b, 0xcd, 0x6d, 0x15, 0x1e,
	0x44, 0xf1, 0xb3, 0xfc, 0x70, 0x01, 0xae, 0x91, 0x45, 0x13, 0x70, 0xb7, 0x78, 0xb3, 0xf1, 0x5e,
	0x98, 0xbf, 0xe4, 0x47, 0xfa, 0xd3, 0x73, 0x42, 0x0c, 0xde, 0x50, 0x01, 0x05, 0x30, 0x5b, 0x53,
	0x95, 0x4b, 0x76, 0x36, 0x54, 0xbd, 0x7d, 0x9f, 0xba, 0xf7, 0xd4, 0x68, 0xe0, 0xa9, 0xe4, 0x13,
	0xbc, 0xcb, 0x89, 0xa4, 0xb1, 0x12, 0xfc, 0x4a, 0xe3, 0x8c, 0xe4, 0x7f, 0x69, 0x59, 0xc4, 0xb4,
	0xd4, 0x4c, 0xdf, 0x92, 0xa7, 0xe6, 0xa1, 0xf6, 0xc6, 0xb4, 0xea, 0x57, 0xe5, 0xda, 0xec, 0xb1,
	0xb2, 0xf0, 0xe7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x2c, 0x6d, 0xd4, 0xda, 0x02, 0x00,
	0x00,
}
