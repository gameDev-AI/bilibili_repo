// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ErrorHandling.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// Protobuf version of a java.lang.StackTraceElement
// so we can serialize exceptions.
type StackTraceElementMessage struct {
	DeclaringClass   *string `protobuf:"bytes,1,opt,name=declaring_class,json=declaringClass" json:"declaring_class,omitempty"`
	MethodName       *string `protobuf:"bytes,2,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	FileName         *string `protobuf:"bytes,3,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	LineNumber       *int32  `protobuf:"varint,4,opt,name=line_number,json=lineNumber" json:"line_number,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StackTraceElementMessage) Reset()                    { *m = StackTraceElementMessage{} }
func (m *StackTraceElementMessage) String() string            { return proto.CompactTextString(m) }
func (*StackTraceElementMessage) ProtoMessage()               {}
func (*StackTraceElementMessage) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *StackTraceElementMessage) GetDeclaringClass() string {
	if m != nil && m.DeclaringClass != nil {
		return *m.DeclaringClass
	}
	return ""
}

func (m *StackTraceElementMessage) GetMethodName() string {
	if m != nil && m.MethodName != nil {
		return *m.MethodName
	}
	return ""
}

func (m *StackTraceElementMessage) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *StackTraceElementMessage) GetLineNumber() int32 {
	if m != nil && m.LineNumber != nil {
		return *m.LineNumber
	}
	return 0
}

// *
// Cause of a remote failure for a generic exception. Contains
// all the information for a generic exception as well as
// optional info about the error for generic info passing
// (which should be another protobuffed class).
type GenericExceptionMessage struct {
	ClassName        *string                     `protobuf:"bytes,1,opt,name=class_name,json=className" json:"class_name,omitempty"`
	Message          *string                     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	ErrorInfo        []byte                      `protobuf:"bytes,3,opt,name=error_info,json=errorInfo" json:"error_info,omitempty"`
	Trace            []*StackTraceElementMessage `protobuf:"bytes,4,rep,name=trace" json:"trace,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *GenericExceptionMessage) Reset()                    { *m = GenericExceptionMessage{} }
func (m *GenericExceptionMessage) String() string            { return proto.CompactTextString(m) }
func (*GenericExceptionMessage) ProtoMessage()               {}
func (*GenericExceptionMessage) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *GenericExceptionMessage) GetClassName() string {
	if m != nil && m.ClassName != nil {
		return *m.ClassName
	}
	return ""
}

func (m *GenericExceptionMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *GenericExceptionMessage) GetErrorInfo() []byte {
	if m != nil {
		return m.ErrorInfo
	}
	return nil
}

func (m *GenericExceptionMessage) GetTrace() []*StackTraceElementMessage {
	if m != nil {
		return m.Trace
	}
	return nil
}

// *
// Exception sent across the wire when a remote task needs
// to notify other tasks that it failed and why
type ForeignExceptionMessage struct {
	Source           *string                  `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	GenericException *GenericExceptionMessage `protobuf:"bytes,2,opt,name=generic_exception,json=genericException" json:"generic_exception,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ForeignExceptionMessage) Reset()                    { *m = ForeignExceptionMessage{} }
func (m *ForeignExceptionMessage) String() string            { return proto.CompactTextString(m) }
func (*ForeignExceptionMessage) ProtoMessage()               {}
func (*ForeignExceptionMessage) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *ForeignExceptionMessage) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *ForeignExceptionMessage) GetGenericException() *GenericExceptionMessage {
	if m != nil {
		return m.GenericException
	}
	return nil
}

func init() {
	proto.RegisterType((*StackTraceElementMessage)(nil), "pb.StackTraceElementMessage")
	proto.RegisterType((*GenericExceptionMessage)(nil), "pb.GenericExceptionMessage")
	proto.RegisterType((*ForeignExceptionMessage)(nil), "pb.ForeignExceptionMessage")
}

func init() { proto.RegisterFile("ErrorHandling.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5d, 0x4b, 0xf3, 0x30,
	0x14, 0xc7, 0xc9, 0x5e, 0x9e, 0xe7, 0xe9, 0xe9, 0x83, 0x2f, 0x15, 0x5c, 0x61, 0xca, 0xc6, 0x6e,
	0x1c, 0x5e, 0xf4, 0x62, 0x1f, 0x61, 0xb2, 0x39, 0x2f, 0x1c, 0x52, 0xbd, 0x2f, 0x69, 0x7a, 0xda,
	0x06, 0xdb, 0xa4, 0xa4, 0x19, 0x08, 0x7e, 0x19, 0x11, 0xfc, 0x9e, 0x92, 0x64, 0x15, 0x54, 0x76,
	0x99, 0xdf, 0xf9, 0x9f, 0xf4, 0x77, 0x4e, 0x03, 0x67, 0x2b, 0xa5, 0xa4, 0xda, 0x50, 0x91, 0x55,
	0x5c, 0x14, 0x51, 0xa3, 0xa4, 0x96, 0x41, 0xaf, 0x49, 0x67, 0xef, 0x04, 0xc2, 0x47, 0x4d, 0xd9,
	0xf3, 0x93, 0xa2, 0x0c, 0x57, 0x15, 0xd6, 0x28, 0xf4, 0x3d, 0xb6, 0x2d, 0x2d, 0x30, 0xb8, 0x82,
	0xe3, 0x0c, 0x59, 0x45, 0x15, 0x17, 0x45, 0xc2, 0x2a, 0xda, 0xb6, 0x21, 0x99, 0x92, 0xb9, 0x17,
	0x1f, 0x7d, 0xe1, 0x1b, 0x43, 0x83, 0x09, 0xf8, 0x35, 0xea, 0x52, 0x66, 0x89, 0xa0, 0x35, 0x86,
	0x3d, 0x1b, 0x02, 0x87, 0xb6, 0xb4, 0xc6, 0x60, 0x0c, 0x5e, 0xce, 0x2b, 0x74, 0xe5, 0xbe, 0x2d,
	0xff, 0x33, 0xc0, 0x16, 0x27, 0xe0, 0x57, 0x5c, 0x60, 0x22, 0x76, 0x75, 0x8a, 0x2a, 0x1c, 0x4c,
	0xc9, 0x7c, 0x18, 0x83, 0x41, 0x5b, 0x4b, 0x66, 0x1f, 0x04, 0x46, 0xb7, 0x28, 0x50, 0x71, 0xb6,
	0x7a, 0x61, 0xd8, 0x68, 0x2e, 0x45, 0xe7, 0x78, 0x09, 0x60, 0xcd, 0xdc, 0xd5, 0x4e, 0xcf, 0xb3,
	0xc4, 0xde, 0x1d, 0xc2, 0xdf, 0xda, 0x25, 0xf7, 0x56, 0xdd, 0xd1, 0x34, 0xa2, 0x59, 0x4a, 0xc2,
	0x45, 0x2e, 0xad, 0xd3, 0xff, 0xd8, 0xb3, 0xe4, 0x4e, 0xe4, 0x32, 0x58, 0xc0, 0x50, 0x9b, 0x95,
	0x84, 0x83, 0x69, 0x7f, 0xee, 0x2f, 0x2e, 0xa2, 0x26, 0x8d, 0x0e, 0x2d, 0x2a, 0x76, 0xd1, 0xd9,
	0x2b, 0x8c, 0xd6, 0x52, 0x21, 0x2f, 0xc4, 0x2f, 0xcd, 0x73, 0xf8, 0xd3, 0xca, 0x9d, 0x62, 0x9d,
	0xe2, 0xfe, 0x14, 0x6c, 0xe0, 0xb4, 0x70, 0x93, 0x25, 0xd8, 0xf5, 0x58, 0x53, 0x7f, 0x31, 0x36,
	0x9f, 0x3c, 0x30, 0x76, 0x7c, 0x52, 0xfc, 0x28, 0x2c, 0xd7, 0x70, 0x2d, 0x55, 0x11, 0xd1, 0x86,
	0xb2, 0x12, 0xa3, 0x92, 0x66, 0x52, 0x36, 0x51, 0x99, 0xd2, 0x16, 0xdd, 0xef, 0x4e, 0x77, 0x79,
	0x64, 0x9b, 0xa8, 0xc6, 0x6c, 0xf9, 0xfd, 0x41, 0x3c, 0x98, 0x40, 0xbb, 0x21, 0x6f, 0x84, 0x7c,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x59, 0x62, 0x7e, 0xbe, 0x2b, 0x02, 0x00, 0x00,
}
