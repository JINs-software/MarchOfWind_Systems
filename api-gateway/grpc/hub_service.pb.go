// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: proto/hub_service.proto

package grpc

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

type CreateMatchRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomName   string `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	MaxPlayers int32  `protobuf:"varint,2,opt,name=maxPlayers,proto3" json:"maxPlayers,omitempty"`
}

func (x *CreateMatchRoomRequest) Reset() {
	*x = CreateMatchRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hub_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchRoomRequest) ProtoMessage() {}

func (x *CreateMatchRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hub_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateMatchRoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_hub_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMatchRoomRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *CreateMatchRoomRequest) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

type CreateMatchRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int32 `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
}

func (x *CreateMatchRoomResponse) Reset() {
	*x = CreateMatchRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hub_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchRoomResponse) ProtoMessage() {}

func (x *CreateMatchRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hub_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateMatchRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_hub_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMatchRoomResponse) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type GetMatchRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMatchRoomsRequest) Reset() {
	*x = GetMatchRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hub_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchRoomsRequest) ProtoMessage() {}

func (x *GetMatchRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hub_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchRoomsRequest.ProtoReflect.Descriptor instead.
func (*GetMatchRoomsRequest) Descriptor() ([]byte, []int) {
	return file_proto_hub_service_proto_rawDescGZIP(), []int{2}
}

type GetMatchRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*MatchRoomInfo `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetMatchRoomsResponse) Reset() {
	*x = GetMatchRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hub_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchRoomsResponse) ProtoMessage() {}

func (x *GetMatchRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hub_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetMatchRoomsResponse) Descriptor() ([]byte, []int) {
	return file_proto_hub_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMatchRoomsResponse) GetRooms() []*MatchRoomInfo {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type MatchRoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId         int32  `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	RoomName       string `protobuf:"bytes,2,opt,name=roomName,proto3" json:"roomName,omitempty"`
	MaxPlayers     int32  `protobuf:"varint,3,opt,name=maxPlayers,proto3" json:"maxPlayers,omitempty"`
	CurrentPlayers int32  `protobuf:"varint,4,opt,name=currentPlayers,proto3" json:"currentPlayers,omitempty"`
}

func (x *MatchRoomInfo) Reset() {
	*x = MatchRoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hub_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRoomInfo) ProtoMessage() {}

func (x *MatchRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hub_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRoomInfo.ProtoReflect.Descriptor instead.
func (*MatchRoomInfo) Descriptor() ([]byte, []int) {
	return file_proto_hub_service_proto_rawDescGZIP(), []int{4}
}

func (x *MatchRoomInfo) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *MatchRoomInfo) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *MatchRoomInfo) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

func (x *MatchRoomInfo) GetCurrentPlayers() int32 {
	if x != nil {
		return x.CurrentPlayers
	}
	return 0
}

var File_proto_hub_service_proto protoreflect.FileDescriptor

var file_proto_hub_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x75, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x54, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x31, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x43, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x32, 0xaa, 0x01, 0x0a, 0x0a, 0x48, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x13, 0x5a, 0x11, 0x68, 0x75, 0x62, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hub_service_proto_rawDescOnce sync.Once
	file_proto_hub_service_proto_rawDescData = file_proto_hub_service_proto_rawDesc
)

func file_proto_hub_service_proto_rawDescGZIP() []byte {
	file_proto_hub_service_proto_rawDescOnce.Do(func() {
		file_proto_hub_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hub_service_proto_rawDescData)
	})
	return file_proto_hub_service_proto_rawDescData
}

var file_proto_hub_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_hub_service_proto_goTypes = []any{
	(*CreateMatchRoomRequest)(nil),  // 0: proto.CreateMatchRoomRequest
	(*CreateMatchRoomResponse)(nil), // 1: proto.CreateMatchRoomResponse
	(*GetMatchRoomsRequest)(nil),    // 2: proto.GetMatchRoomsRequest
	(*GetMatchRoomsResponse)(nil),   // 3: proto.GetMatchRoomsResponse
	(*MatchRoomInfo)(nil),           // 4: proto.MatchRoomInfo
}
var file_proto_hub_service_proto_depIdxs = []int32{
	4, // 0: proto.GetMatchRoomsResponse.rooms:type_name -> proto.MatchRoomInfo
	0, // 1: proto.HubService.CreateMatchRoom:input_type -> proto.CreateMatchRoomRequest
	2, // 2: proto.HubService.GetMatchRooms:input_type -> proto.GetMatchRoomsRequest
	1, // 3: proto.HubService.CreateMatchRoom:output_type -> proto.CreateMatchRoomResponse
	3, // 4: proto.HubService.GetMatchRooms:output_type -> proto.GetMatchRoomsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_hub_service_proto_init() }
func file_proto_hub_service_proto_init() {
	if File_proto_hub_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hub_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMatchRoomRequest); i {
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
		file_proto_hub_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMatchRoomResponse); i {
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
		file_proto_hub_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetMatchRoomsRequest); i {
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
		file_proto_hub_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetMatchRoomsResponse); i {
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
		file_proto_hub_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MatchRoomInfo); i {
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
			RawDescriptor: file_proto_hub_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hub_service_proto_goTypes,
		DependencyIndexes: file_proto_hub_service_proto_depIdxs,
		MessageInfos:      file_proto_hub_service_proto_msgTypes,
	}.Build()
	File_proto_hub_service_proto = out.File
	file_proto_hub_service_proto_rawDesc = nil
	file_proto_hub_service_proto_goTypes = nil
	file_proto_hub_service_proto_depIdxs = nil
}
