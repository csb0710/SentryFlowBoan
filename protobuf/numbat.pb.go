// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: numbat.proto

package protobuf

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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp    string            `protobuf:"bytes,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Id           uint64            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	SrcNamespace string            `protobuf:"bytes,3,opt,name=srcNamespace,proto3" json:"srcNamespace,omitempty"`
	SrcName      string            `protobuf:"bytes,4,opt,name=srcName,proto3" json:"srcName,omitempty"`
	SrcLabel     map[string]string `protobuf:"bytes,5,rep,name=srcLabel,proto3" json:"srcLabel,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SrcIP        string            `protobuf:"bytes,6,opt,name=srcIP,proto3" json:"srcIP,omitempty"`
	SrcPort      string            `protobuf:"bytes,7,opt,name=srcPort,proto3" json:"srcPort,omitempty"`
	SrcType      string            `protobuf:"bytes,8,opt,name=srcType,proto3" json:"srcType,omitempty"`
	DstNamespace string            `protobuf:"bytes,9,opt,name=dstNamespace,proto3" json:"dstNamespace,omitempty"`
	DstName      string            `protobuf:"bytes,10,opt,name=dstName,proto3" json:"dstName,omitempty"`
	DstLabel     map[string]string `protobuf:"bytes,11,rep,name=dstLabel,proto3" json:"dstLabel,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DstIP        string            `protobuf:"bytes,12,opt,name=dstIP,proto3" json:"dstIP,omitempty"`
	DstPort      string            `protobuf:"bytes,13,opt,name=dstPort,proto3" json:"dstPort,omitempty"`
	DstType      string            `protobuf:"bytes,14,opt,name=dstType,proto3" json:"dstType,omitempty"`
	Protocol     string            `protobuf:"bytes,15,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Method       string            `protobuf:"bytes,16,opt,name=method,proto3" json:"method,omitempty"`
	Path         string            `protobuf:"bytes,17,opt,name=path,proto3" json:"path,omitempty"`
	ResponseCode int64             `protobuf:"varint,18,opt,name=responseCode,proto3" json:"responseCode,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_numbat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_numbat_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *Log) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Log) GetSrcNamespace() string {
	if x != nil {
		return x.SrcNamespace
	}
	return ""
}

func (x *Log) GetSrcName() string {
	if x != nil {
		return x.SrcName
	}
	return ""
}

func (x *Log) GetSrcLabel() map[string]string {
	if x != nil {
		return x.SrcLabel
	}
	return nil
}

func (x *Log) GetSrcIP() string {
	if x != nil {
		return x.SrcIP
	}
	return ""
}

func (x *Log) GetSrcPort() string {
	if x != nil {
		return x.SrcPort
	}
	return ""
}

func (x *Log) GetSrcType() string {
	if x != nil {
		return x.SrcType
	}
	return ""
}

func (x *Log) GetDstNamespace() string {
	if x != nil {
		return x.DstNamespace
	}
	return ""
}

func (x *Log) GetDstName() string {
	if x != nil {
		return x.DstName
	}
	return ""
}

func (x *Log) GetDstLabel() map[string]string {
	if x != nil {
		return x.DstLabel
	}
	return nil
}

func (x *Log) GetDstIP() string {
	if x != nil {
		return x.DstIP
	}
	return ""
}

func (x *Log) GetDstPort() string {
	if x != nil {
		return x.DstPort
	}
	return ""
}

func (x *Log) GetDstType() string {
	if x != nil {
		return x.DstType
	}
	return ""
}

func (x *Log) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Log) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Log) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Log) GetResponseCode() int64 {
	if x != nil {
		return x.ResponseCode
	}
	return 0
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname  string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	IpAddress string `protobuf:"bytes,2,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_numbat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_numbat_proto_rawDescGZIP(), []int{1}
}

func (x *ClientInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ClientInfo) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type APIMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerAPICounts map[string]uint64 `protobuf:"bytes,1,rep,name=perAPICounts,proto3" json:"perAPICounts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // @todo: add some more metrics here
}

func (x *APIMetric) Reset() {
	*x = APIMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIMetric) ProtoMessage() {}

func (x *APIMetric) ProtoReflect() protoreflect.Message {
	mi := &file_numbat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIMetric.ProtoReflect.Descriptor instead.
func (*APIMetric) Descriptor() ([]byte, []int) {
	return file_numbat_proto_rawDescGZIP(), []int{2}
}

func (x *APIMetric) GetPerAPICounts() map[string]uint64 {
	if x != nil {
		return x.PerAPICounts
	}
	return nil
}

var File_numbat_proto protoreflect.FileDescriptor

var file_numbat_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x9b, 0x05, 0x0a, 0x03, 0x4c, 0x6f, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x72, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x72, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08,
	0x73, 0x72, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x72,
	0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x72, 0x63,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x72, 0x63, 0x49, 0x50, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72,
	0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a,
	0x08, 0x64, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x44,
	0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x64, 0x73,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x73, 0x74, 0x49, 0x50, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x73, 0x74, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x3b, 0x0a, 0x0d,
	0x53, 0x72, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x44, 0x73, 0x74,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x97,
	0x01, 0x0a, 0x09, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x49, 0x0a, 0x0c,
	0x70, 0x65, 0x72, 0x41, 0x50, 0x49, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x50,
	0x49, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x41, 0x50, 0x49, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x75, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62,
	0x61, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x6f,
	0x67, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x42,
	0x11, 0x5a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_numbat_proto_rawDescOnce sync.Once
	file_numbat_proto_rawDescData = file_numbat_proto_rawDesc
)

func file_numbat_proto_rawDescGZIP() []byte {
	file_numbat_proto_rawDescOnce.Do(func() {
		file_numbat_proto_rawDescData = protoimpl.X.CompressGZIP(file_numbat_proto_rawDescData)
	})
	return file_numbat_proto_rawDescData
}

var file_numbat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_numbat_proto_goTypes = []interface{}{
	(*Log)(nil),        // 0: protobuf.Log
	(*ClientInfo)(nil), // 1: protobuf.ClientInfo
	(*APIMetric)(nil),  // 2: protobuf.APIMetric
	nil,                // 3: protobuf.Log.SrcLabelEntry
	nil,                // 4: protobuf.Log.DstLabelEntry
	nil,                // 5: protobuf.APIMetric.PerAPICountsEntry
}
var file_numbat_proto_depIdxs = []int32{
	3, // 0: protobuf.Log.srcLabel:type_name -> protobuf.Log.SrcLabelEntry
	4, // 1: protobuf.Log.dstLabel:type_name -> protobuf.Log.DstLabelEntry
	5, // 2: protobuf.APIMetric.perAPICounts:type_name -> protobuf.APIMetric.PerAPICountsEntry
	1, // 3: protobuf.Numbat.GetLog:input_type -> protobuf.ClientInfo
	1, // 4: protobuf.Numbat.GetAPIMetrics:input_type -> protobuf.ClientInfo
	0, // 5: protobuf.Numbat.GetLog:output_type -> protobuf.Log
	2, // 6: protobuf.Numbat.GetAPIMetrics:output_type -> protobuf.APIMetric
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_numbat_proto_init() }
func file_numbat_proto_init() {
	if File_numbat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_numbat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_numbat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
		file_numbat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIMetric); i {
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
			RawDescriptor: file_numbat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_numbat_proto_goTypes,
		DependencyIndexes: file_numbat_proto_depIdxs,
		MessageInfos:      file_numbat_proto_msgTypes,
	}.Build()
	File_numbat_proto = out.File
	file_numbat_proto_rawDesc = nil
	file_numbat_proto_goTypes = nil
	file_numbat_proto_depIdxs = nil
}
