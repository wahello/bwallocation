// Copyright 2020 Anapaya Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: proto/bw_allocation/v1/bw_allocation.proto

package bw_allocation

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReserveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The destination ISD-AS.
	DstIsdAs uint64 `protobuf:"varint,1,opt,name=dst_isd_as,json=dstIsdAs,proto3" json:"dst_isd_as,omitempty"`
	// The source host of the requested bandwidth reservation.
	SrcHost []byte `protobuf:"bytes,2,opt,name=src_host,json=srcHost,proto3" json:"src_host,omitempty"`
	// The destination host of the requested bandwidth reservation.
	DstHost []byte `protobuf:"bytes,3,opt,name=dst_host,json=dstHost,proto3" json:"dst_host,omitempty"`
	// The reservation ID of the requested bandwidth reservation.
	Id uint32 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	// The expiration time of the reservation. The timestamp is encoded as
	// number of seconds elapsed since January 1, 1970 UTC.
	ExpirationTime int64 `protobuf:"varint,5,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	// The bandwidth of the reservation in bps.
	Bandwidth uint64 `protobuf:"varint,6,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// The raw SCION path that is used with the bandwidth reservation.
	Path []byte `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ReserveRequest) Reset() {
	*x = ReserveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveRequest) ProtoMessage() {}

func (x *ReserveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveRequest.ProtoReflect.Descriptor instead.
func (*ReserveRequest) Descriptor() ([]byte, []int) {
	return file_proto_bw_allocation_v1_bw_allocation_proto_rawDescGZIP(), []int{0}
}

func (x *ReserveRequest) GetDstIsdAs() uint64 {
	if x != nil {
		return x.DstIsdAs
	}
	return 0
}

func (x *ReserveRequest) GetSrcHost() []byte {
	if x != nil {
		return x.SrcHost
	}
	return nil
}

func (x *ReserveRequest) GetDstHost() []byte {
	if x != nil {
		return x.DstHost
	}
	return nil
}

func (x *ReserveRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReserveRequest) GetExpirationTime() int64 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

func (x *ReserveRequest) GetBandwidth() uint64 {
	if x != nil {
		return x.Bandwidth
	}
	return 0
}

func (x *ReserveRequest) GetPath() []byte {
	if x != nil {
		return x.Path
	}
	return nil
}

type ReserveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message authentication code that authenticates the reservation.
	Mac []byte `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *ReserveResponse) Reset() {
	*x = ReserveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveResponse) ProtoMessage() {}

func (x *ReserveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveResponse.ProtoReflect.Descriptor instead.
func (*ReserveResponse) Descriptor() ([]byte, []int) {
	return file_proto_bw_allocation_v1_bw_allocation_proto_rawDescGZIP(), []int{1}
}

func (x *ReserveResponse) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

type BandwidthHint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maxium bandwidth that the services is able to allocate at this point
	// in time.
	MaxBandwidth uint64 `protobuf:"varint,1,opt,name=max_bandwidth,json=maxBandwidth,proto3" json:"max_bandwidth,omitempty"`
}

func (x *BandwidthHint) Reset() {
	*x = BandwidthHint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BandwidthHint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BandwidthHint) ProtoMessage() {}

func (x *BandwidthHint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BandwidthHint.ProtoReflect.Descriptor instead.
func (*BandwidthHint) Descriptor() ([]byte, []int) {
	return file_proto_bw_allocation_v1_bw_allocation_proto_rawDescGZIP(), []int{2}
}

func (x *BandwidthHint) GetMaxBandwidth() uint64 {
	if x != nil {
		return x.MaxBandwidth
	}
	return 0
}

type ExpirationHint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maxmimum acceptable expiration time.
	MaxExpirationTime int64 `protobuf:"varint,1,opt,name=max_expiration_time,json=maxExpirationTime,proto3" json:"max_expiration_time,omitempty"`
}

func (x *ExpirationHint) Reset() {
	*x = ExpirationHint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpirationHint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirationHint) ProtoMessage() {}

func (x *ExpirationHint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirationHint.ProtoReflect.Descriptor instead.
func (*ExpirationHint) Descriptor() ([]byte, []int) {
	return file_proto_bw_allocation_v1_bw_allocation_proto_rawDescGZIP(), []int{3}
}

func (x *ExpirationHint) GetMaxExpirationTime() int64 {
	if x != nil {
		return x.MaxExpirationTime
	}
	return 0
}

var File_proto_bw_allocation_v1_bw_allocation_proto protoreflect.FileDescriptor

var file_proto_bw_allocation_v1_bw_allocation_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x61, 0x6e,
	0x61, 0x70, 0x61, 0x79, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x77, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xcf, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x69, 0x73, 0x64, 0x5f, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x73, 0x74, 0x49, 0x73, 0x64, 0x41, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x72, 0x63, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x73, 0x72, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x73, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x23,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6d, 0x61, 0x63, 0x22, 0x34, 0x0a, 0x0d, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x48, 0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x40, 0x0a, 0x0e, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d,
	0x61, 0x78, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x19,
	0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x74, 0x68, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x07, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x12, 0x2e, 0x2e, 0x61, 0x6e, 0x61, 0x70, 0x61, 0x79, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x6e, 0x61, 0x70, 0x61, 0x79, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x61, 0x70, 0x61, 0x79, 0x61, 0x2f, 0x62, 0x77,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bw_allocation_v1_bw_allocation_proto_rawDescOnce sync.Once
	file_proto_bw_allocation_v1_bw_allocation_proto_rawDescData = file_proto_bw_allocation_v1_bw_allocation_proto_rawDesc
)

func file_proto_bw_allocation_v1_bw_allocation_proto_rawDescGZIP() []byte {
	file_proto_bw_allocation_v1_bw_allocation_proto_rawDescOnce.Do(func() {
		file_proto_bw_allocation_v1_bw_allocation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bw_allocation_v1_bw_allocation_proto_rawDescData)
	})
	return file_proto_bw_allocation_v1_bw_allocation_proto_rawDescData
}

var file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_bw_allocation_v1_bw_allocation_proto_goTypes = []interface{}{
	(*ReserveRequest)(nil),  // 0: anapaya.proto.bw_allocation.v1.ReserveRequest
	(*ReserveResponse)(nil), // 1: anapaya.proto.bw_allocation.v1.ReserveResponse
	(*BandwidthHint)(nil),   // 2: anapaya.proto.bw_allocation.v1.BandwidthHint
	(*ExpirationHint)(nil),  // 3: anapaya.proto.bw_allocation.v1.ExpirationHint
}
var file_proto_bw_allocation_v1_bw_allocation_proto_depIdxs = []int32{
	0, // 0: anapaya.proto.bw_allocation.v1.BandwithAllocationService.Reserve:input_type -> anapaya.proto.bw_allocation.v1.ReserveRequest
	1, // 1: anapaya.proto.bw_allocation.v1.BandwithAllocationService.Reserve:output_type -> anapaya.proto.bw_allocation.v1.ReserveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_bw_allocation_v1_bw_allocation_proto_init() }
func file_proto_bw_allocation_v1_bw_allocation_proto_init() {
	if File_proto_bw_allocation_v1_bw_allocation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveRequest); i {
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
		file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveResponse); i {
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
		file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BandwidthHint); i {
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
		file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpirationHint); i {
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
			RawDescriptor: file_proto_bw_allocation_v1_bw_allocation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bw_allocation_v1_bw_allocation_proto_goTypes,
		DependencyIndexes: file_proto_bw_allocation_v1_bw_allocation_proto_depIdxs,
		MessageInfos:      file_proto_bw_allocation_v1_bw_allocation_proto_msgTypes,
	}.Build()
	File_proto_bw_allocation_v1_bw_allocation_proto = out.File
	file_proto_bw_allocation_v1_bw_allocation_proto_rawDesc = nil
	file_proto_bw_allocation_v1_bw_allocation_proto_goTypes = nil
	file_proto_bw_allocation_v1_bw_allocation_proto_depIdxs = nil
}
