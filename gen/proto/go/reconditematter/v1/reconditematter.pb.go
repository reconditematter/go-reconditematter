// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: reconditematter/v1/reconditematter.proto

package reconditematterv1

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

type RandomNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RandomNamesRequest) Reset() {
	*x = RandomNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNamesRequest) ProtoMessage() {}

func (x *RandomNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNamesRequest.ProtoReflect.Descriptor instead.
func (*RandomNamesRequest) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{0}
}

func (x *RandomNamesRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type HumanName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family string `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Given  string `protobuf:"bytes,2,opt,name=given,proto3" json:"given,omitempty"`
	Gender string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *HumanName) Reset() {
	*x = HumanName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HumanName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HumanName) ProtoMessage() {}

func (x *HumanName) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HumanName.ProtoReflect.Descriptor instead.
func (*HumanName) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{1}
}

func (x *HumanName) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *HumanName) GetGiven() string {
	if x != nil {
		return x.Given
	}
	return ""
}

func (x *HumanName) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type RandomNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Names []*HumanName `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *RandomNamesResponse) Reset() {
	*x = RandomNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNamesResponse) ProtoMessage() {}

func (x *RandomNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNamesResponse.ProtoReflect.Descriptor instead.
func (*RandomNamesResponse) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{2}
}

func (x *RandomNamesResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RandomNamesResponse) GetNames() []*HumanName {
	if x != nil {
		return x.Names
	}
	return nil
}

type FibonacciPointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FibonacciPointsRequest) Reset() {
	*x = FibonacciPointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciPointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciPointsRequest) ProtoMessage() {}

func (x *FibonacciPointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciPointsRequest.ProtoReflect.Descriptor instead.
func (*FibonacciPointsRequest) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{3}
}

func (x *FibonacciPointsRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *GeoPoint) Reset() {
	*x = GeoPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPoint) ProtoMessage() {}

func (x *GeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPoint.ProtoReflect.Descriptor instead.
func (*GeoPoint) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{4}
}

func (x *GeoPoint) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GeoPoint) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type FibonacciPointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  uint32      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Points []*GeoPoint `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *FibonacciPointsResponse) Reset() {
	*x = FibonacciPointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciPointsResponse) ProtoMessage() {}

func (x *FibonacciPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciPointsResponse.ProtoReflect.Descriptor instead.
func (*FibonacciPointsResponse) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{5}
}

func (x *FibonacciPointsResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FibonacciPointsResponse) GetPoints() []*GeoPoint {
	if x != nil {
		return x.Points
	}
	return nil
}

type FibonacciCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  uint32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Origin *GeoPoint `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *FibonacciCellRequest) Reset() {
	*x = FibonacciCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciCellRequest) ProtoMessage() {}

func (x *FibonacciCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciCellRequest.ProtoReflect.Descriptor instead.
func (*FibonacciCellRequest) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{6}
}

func (x *FibonacciCellRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FibonacciCellRequest) GetOrigin() *GeoPoint {
	if x != nil {
		return x.Origin
	}
	return nil
}

type FibonacciCellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  uint32      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Origin *GeoPoint   `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Points []*GeoPoint `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *FibonacciCellResponse) Reset() {
	*x = FibonacciCellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciCellResponse) ProtoMessage() {}

func (x *FibonacciCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciCellResponse.ProtoReflect.Descriptor instead.
func (*FibonacciCellResponse) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{7}
}

func (x *FibonacciCellResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FibonacciCellResponse) GetOrigin() *GeoPoint {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *FibonacciCellResponse) GetPoints() []*GeoPoint {
	if x != nil {
		return x.Points
	}
	return nil
}

type GeoHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point  *GeoPoint `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
	Length uint32    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *GeoHashRequest) Reset() {
	*x = GeoHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoHashRequest) ProtoMessage() {}

func (x *GeoHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoHashRequest.ProtoReflect.Descriptor instead.
func (*GeoHashRequest) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{8}
}

func (x *GeoHashRequest) GetPoint() *GeoPoint {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *GeoHashRequest) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type GeoHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Geohash string `protobuf:"bytes,1,opt,name=geohash,proto3" json:"geohash,omitempty"`
}

func (x *GeoHashResponse) Reset() {
	*x = GeoHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoHashResponse) ProtoMessage() {}

func (x *GeoHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reconditematter_v1_reconditematter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoHashResponse.ProtoReflect.Descriptor instead.
func (*GeoHashResponse) Descriptor() ([]byte, []int) {
	return file_reconditematter_v1_reconditematter_proto_rawDescGZIP(), []int{9}
}

func (x *GeoHashResponse) GetGeohash() string {
	if x != nil {
		return x.Geohash
	}
	return ""
}

var File_reconditematter_v1_reconditematter_proto protoreflect.FileDescriptor

var file_reconditematter_v1_reconditematter_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x2a,
	0x0a, 0x12, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x09, 0x48, 0x75,
	0x6d, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x69, 0x76, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x60, 0x0a,
	0x13, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x75, 0x6d, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x2e, 0x0a, 0x16, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x2e, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22,
	0x65, 0x0a, 0x17, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x34, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x62, 0x0a, 0x14, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61,
	0x63, 0x63, 0x69, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65,
	0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x46,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x34, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x0e, 0x47, 0x65, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x6f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x6f, 0x68, 0x61, 0x73,
	0x68, 0x32, 0x9e, 0x03, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x4d,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0b,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0f,
	0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x2a, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x46, 0x69, 0x62, 0x6f,
	0x6e, 0x61, 0x63, 0x63, 0x69, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63,
	0x63, 0x69, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x07, 0x47, 0x65, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x61, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reconditematter_v1_reconditematter_proto_rawDescOnce sync.Once
	file_reconditematter_v1_reconditematter_proto_rawDescData = file_reconditematter_v1_reconditematter_proto_rawDesc
)

func file_reconditematter_v1_reconditematter_proto_rawDescGZIP() []byte {
	file_reconditematter_v1_reconditematter_proto_rawDescOnce.Do(func() {
		file_reconditematter_v1_reconditematter_proto_rawDescData = protoimpl.X.CompressGZIP(file_reconditematter_v1_reconditematter_proto_rawDescData)
	})
	return file_reconditematter_v1_reconditematter_proto_rawDescData
}

var file_reconditematter_v1_reconditematter_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_reconditematter_v1_reconditematter_proto_goTypes = []interface{}{
	(*RandomNamesRequest)(nil),      // 0: reconditematter.v1.RandomNamesRequest
	(*HumanName)(nil),               // 1: reconditematter.v1.HumanName
	(*RandomNamesResponse)(nil),     // 2: reconditematter.v1.RandomNamesResponse
	(*FibonacciPointsRequest)(nil),  // 3: reconditematter.v1.FibonacciPointsRequest
	(*GeoPoint)(nil),                // 4: reconditematter.v1.GeoPoint
	(*FibonacciPointsResponse)(nil), // 5: reconditematter.v1.FibonacciPointsResponse
	(*FibonacciCellRequest)(nil),    // 6: reconditematter.v1.FibonacciCellRequest
	(*FibonacciCellResponse)(nil),   // 7: reconditematter.v1.FibonacciCellResponse
	(*GeoHashRequest)(nil),          // 8: reconditematter.v1.GeoHashRequest
	(*GeoHashResponse)(nil),         // 9: reconditematter.v1.GeoHashResponse
}
var file_reconditematter_v1_reconditematter_proto_depIdxs = []int32{
	1,  // 0: reconditematter.v1.RandomNamesResponse.names:type_name -> reconditematter.v1.HumanName
	4,  // 1: reconditematter.v1.FibonacciPointsResponse.points:type_name -> reconditematter.v1.GeoPoint
	4,  // 2: reconditematter.v1.FibonacciCellRequest.origin:type_name -> reconditematter.v1.GeoPoint
	4,  // 3: reconditematter.v1.FibonacciCellResponse.origin:type_name -> reconditematter.v1.GeoPoint
	4,  // 4: reconditematter.v1.FibonacciCellResponse.points:type_name -> reconditematter.v1.GeoPoint
	4,  // 5: reconditematter.v1.GeoHashRequest.point:type_name -> reconditematter.v1.GeoPoint
	0,  // 6: reconditematter.v1.ReconditeMatterService.RandomNames:input_type -> reconditematter.v1.RandomNamesRequest
	3,  // 7: reconditematter.v1.ReconditeMatterService.FibonacciPoints:input_type -> reconditematter.v1.FibonacciPointsRequest
	6,  // 8: reconditematter.v1.ReconditeMatterService.FibonacciCell:input_type -> reconditematter.v1.FibonacciCellRequest
	8,  // 9: reconditematter.v1.ReconditeMatterService.GeoHash:input_type -> reconditematter.v1.GeoHashRequest
	2,  // 10: reconditematter.v1.ReconditeMatterService.RandomNames:output_type -> reconditematter.v1.RandomNamesResponse
	5,  // 11: reconditematter.v1.ReconditeMatterService.FibonacciPoints:output_type -> reconditematter.v1.FibonacciPointsResponse
	7,  // 12: reconditematter.v1.ReconditeMatterService.FibonacciCell:output_type -> reconditematter.v1.FibonacciCellResponse
	9,  // 13: reconditematter.v1.ReconditeMatterService.GeoHash:output_type -> reconditematter.v1.GeoHashResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_reconditematter_v1_reconditematter_proto_init() }
func file_reconditematter_v1_reconditematter_proto_init() {
	if File_reconditematter_v1_reconditematter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reconditematter_v1_reconditematter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNamesRequest); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HumanName); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNamesResponse); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciPointsRequest); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoPoint); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciPointsResponse); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciCellRequest); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciCellResponse); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoHashRequest); i {
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
		file_reconditematter_v1_reconditematter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoHashResponse); i {
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
			RawDescriptor: file_reconditematter_v1_reconditematter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reconditematter_v1_reconditematter_proto_goTypes,
		DependencyIndexes: file_reconditematter_v1_reconditematter_proto_depIdxs,
		MessageInfos:      file_reconditematter_v1_reconditematter_proto_msgTypes,
	}.Build()
	File_reconditematter_v1_reconditematter_proto = out.File
	file_reconditematter_v1_reconditematter_proto_rawDesc = nil
	file_reconditematter_v1_reconditematter_proto_goTypes = nil
	file_reconditematter_v1_reconditematter_proto_depIdxs = nil
}
