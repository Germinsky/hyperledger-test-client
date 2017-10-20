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
	ScheduleRecord
	ScheduleRequest
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
	ScheduleId string            `protobuf:"bytes,1,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
	DoctorId   string            `protobuf:"bytes,2,opt,name=DoctorId" json:"DoctorId,omitempty"`
	Records    []*ScheduleRecord `protobuf:"bytes,3,rep,name=Records" json:"Records,omitempty"`
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

func (m *Schedule) GetRecords() []*ScheduleRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type Slot struct {
	TimeStart  uint64 `protobuf:"varint,1,opt,name=TimeStart" json:"TimeStart,omitempty"`
	TimeFinish uint64 `protobuf:"varint,2,opt,name=TimeFinish" json:"TimeFinish,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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

type ScheduleRecord struct {
	RecordId    string `protobuf:"bytes,1,opt,name=RecordId" json:"RecordId,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
	PatientId   string `protobuf:"bytes,3,opt,name=PatientId" json:"PatientId,omitempty"`
	Slot        *Slot  `protobuf:"bytes,4,opt,name=Slot" json:"Slot,omitempty"`
}

func (m *ScheduleRecord) Reset()                    { *m = ScheduleRecord{} }
func (m *ScheduleRecord) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRecord) ProtoMessage()               {}
func (*ScheduleRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ScheduleRecord) GetRecordId() string {
	if m != nil {
		return m.RecordId
	}
	return ""
}

func (m *ScheduleRecord) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ScheduleRecord) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func (m *ScheduleRecord) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

type ScheduleRequest struct {
	PatientId   string `protobuf:"bytes,1,opt,name=PatientId" json:"PatientId,omitempty"`
	DoctorId    string `protobuf:"bytes,2,opt,name=DoctorId" json:"DoctorId,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	Slot        *Slot  `protobuf:"bytes,4,opt,name=Slot" json:"Slot,omitempty"`
}

func (m *ScheduleRequest) Reset()                    { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()               {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ScheduleRequest) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func (m *ScheduleRequest) GetDoctorId() string {
	if m != nil {
		return m.DoctorId
	}
	return ""
}

func (m *ScheduleRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ScheduleRequest) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func init() {
	proto.RegisterType((*PatientPublic)(nil), "main.PatientPublic")
	proto.RegisterType((*DoctorPublic)(nil), "main.DoctorPublic")
	proto.RegisterType((*DoctorCollection)(nil), "main.DoctorCollection")
	proto.RegisterType((*Schedule)(nil), "main.Schedule")
	proto.RegisterType((*Slot)(nil), "main.Slot")
	proto.RegisterType((*ScheduleRecord)(nil), "main.ScheduleRecord")
	proto.RegisterType((*ScheduleRequest)(nil), "main.ScheduleRequest")
}

func init() { proto.RegisterFile("registerDoctor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0x25, 0x4d, 0xf8, 0xda, 0x4e, 0xbf, 0x3f, 0x42, 0x2f, 0xa2, 0x48, 0x09, 0xb9, 0xea, 0x85,
	0x06, 0xac, 0x2f, 0x20, 0xb5, 0x14, 0x04, 0x91, 0x92, 0x7a, 0xe5, 0x5d, 0xba, 0x19, 0xec, 0x62,
	0x9a, 0xad, 0xbb, 0xd3, 0x82, 0xef, 0xa0, 0xef, 0x2c, 0xbb, 0x9b, 0xbf, 0x56, 0xa8, 0x77, 0x7b,
	0xce, 0xce, 0xce, 0x39, 0x67, 0x76, 0x60, 0x28, 0xf1, 0x85, 0x2b, 0x42, 0x39, 0x13, 0x8c, 0x84,
	0x8c, 0xb7, 0x52, 0x90, 0xf0, 0xbd, 0x4d, 0xca, 0x8b, 0xe8, 0x0a, 0xfe, 0x2c, 0x52, 0xe2, 0x58,
	0xd0, 0x62, 0xb7, 0xca, 0x39, 0xf3, 0x2f, 0xa0, 0x5f, 0x12, 0xf7, 0x59, 0xe0, 0x84, 0xce, 0xb8,
	0x9f, 0x34, 0x44, 0x94, 0xc1, 0x6f, 0xdb, 0xa4, 0xac, 0x3e, 0x87, 0x9e, 0xc5, 0x75, 0x71, 0x8d,
	0x75, 0xa7, 0x39, 0x97, 0x8a, 0x1e, 0xd3, 0x0d, 0x06, 0x1d, 0xdb, 0xa9, 0x26, 0xf4, 0xcb, 0x87,
	0xb4, 0xbc, 0x74, 0xed, 0xcb, 0x0a, 0x47, 0xb7, 0xf0, 0xdf, 0x76, 0xb9, 0x13, 0x79, 0x8e, 0x8c,
	0xb8, 0x28, 0xfc, 0x4b, 0xe8, 0x66, 0x86, 0x53, 0x81, 0x13, 0xba, 0xe3, 0xc1, 0xc4, 0x8f, 0x75,
	0x80, 0xb8, 0x6d, 0x27, 0xa9, 0x4a, 0xa2, 0x3d, 0xf4, 0x96, 0x6c, 0x8d, 0xd9, 0x2e, 0x47, 0x7f,
	0x04, 0x50, 0x9d, 0x6b, 0x97, 0x2d, 0xe6, 0x20, 0x43, 0xe7, 0x28, 0x43, 0x0c, 0xdd, 0x04, 0x99,
	0x90, 0x99, 0x0a, 0x5c, 0xa3, 0x3a, 0xb4, 0xaa, 0xd5, 0x73, 0x7b, 0x99, 0x54, 0x45, 0xd1, 0x0c,
	0xbc, 0x65, 0x2e, 0x48, 0x67, 0x7f, 0xe2, 0x1b, 0x5c, 0x52, 0x2a, 0xc9, 0x48, 0x7a, 0x49, 0x43,
	0x68, 0x47, 0x1a, 0xcc, 0x79, 0xc1, 0xd5, 0xda, 0x68, 0x7a, 0x49, 0x8b, 0x89, 0x3e, 0x1c, 0xf8,
	0x7b, 0xa8, 0xa0, 0x4d, 0xda, 0x53, 0x33, 0xe8, 0x0a, 0xfb, 0x21, 0x0c, 0x66, 0xa8, 0x98, 0xe4,
	0x5b, 0x3d, 0xa9, 0x32, 0x43, 0x9b, 0x3a, 0xfc, 0x54, 0xf7, 0xe8, 0x53, 0xfd, 0x91, 0x35, 0x1d,
	0x78, 0xa1, 0x33, 0x1e, 0x4c, 0xa0, 0x4c, 0x98, 0x0b, 0x4a, 0x0c, 0x1f, 0x7d, 0x3a, 0xf0, 0xaf,
	0xb1, 0xf3, 0xb6, 0x43, 0x45, 0xa7, 0xd7, 0xe4, 0xe4, 0x48, 0x8f, 0xdc, 0xba, 0xdf, 0xdd, 0xfe,
	0xe0, 0x67, 0x7a, 0x0d, 0x67, 0x2c, 0x95, 0x18, 0x2b, 0x91, 0xef, 0x31, 0x5e, 0xa5, 0xec, 0x15,
	0x8b, 0x2c, 0xc6, 0x82, 0x38, 0xbd, 0x4f, 0xeb, 0xc1, 0x2d, 0xf4, 0x96, 0xab, 0x67, 0xb3, 0xe6,
	0xab, 0x5f, 0x66, 0xe7, 0x6f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x79, 0x32, 0x2b, 0xdf, 0x0b,
	0x03, 0x00, 0x00,
}
