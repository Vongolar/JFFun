// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: game.proto

package Dgame

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

type GameType int32

const (
	GameType__          GameType = 0
	GameType_LostCities GameType = 1
)

// Enum value maps for GameType.
var (
	GameType_name = map[int32]string{
		0: "_",
		1: "LostCities",
	}
	GameType_value = map[string]int32{
		"_":          0,
		"LostCities": 1,
	}
)

func (x GameType) Enum() *GameType {
	p := new(GameType)
	*p = x
	return p
}

func (x GameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameType) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (GameType) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x GameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameType.Descriptor instead.
func (GameType) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

type MatchState int32

const (
	MatchState_noMatch  MatchState = 0 //不在匹配中
	MatchState_matching MatchState = 1 //匹配中
	MatchState_waitSure MatchState = 2 //等待确认
	MatchState_sure     MatchState = 3 //确认状态
)

// Enum value maps for MatchState.
var (
	MatchState_name = map[int32]string{
		0: "noMatch",
		1: "matching",
		2: "waitSure",
		3: "sure",
	}
	MatchState_value = map[string]int32{
		"noMatch":  0,
		"matching": 1,
		"waitSure": 2,
		"sure":     3,
	}
)

func (x MatchState) Enum() *MatchState {
	p := new(MatchState)
	*p = x
	return p
}

func (x MatchState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchState) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[1].Descriptor()
}

func (MatchState) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[1]
}

func (x MatchState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchState.Descriptor instead.
func (MatchState) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Act  int32  `protobuf:"varint,1,opt,name=act,proto3" json:"act,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetAct() int32 {
	if x != nil {
		return x.Act
	}
	return 0
}

func (x *Action) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type StateEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int32 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *StateEnd) Reset() {
	*x = StateEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateEnd) ProtoMessage() {}

func (x *StateEnd) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateEnd.ProtoReflect.Descriptor instead.
func (*StateEnd) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *StateEnd) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type SyncGameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int32  `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SyncGameState) Reset() {
	*x = SyncGameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncGameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncGameState) ProtoMessage() {}

func (x *SyncGameState) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncGameState.ProtoReflect.Descriptor instead.
func (*SyncGameState) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *SyncGameState) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SyncGameState) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SyncGamePlayerAct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Act    int32  `protobuf:"varint,1,opt,name=act,proto3" json:"act,omitempty"`
	Player string `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SyncGamePlayerAct) Reset() {
	*x = SyncGamePlayerAct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncGamePlayerAct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncGamePlayerAct) ProtoMessage() {}

func (x *SyncGamePlayerAct) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncGamePlayerAct.ProtoReflect.Descriptor instead.
func (*SyncGamePlayerAct) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *SyncGamePlayerAct) GetAct() int32 {
	if x != nil {
		return x.Act
	}
	return 0
}

func (x *SyncGamePlayerAct) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *SyncGamePlayerAct) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

//等待
type WaitDeadLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     int64  `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Player string `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *WaitDeadLine) Reset() {
	*x = WaitDeadLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitDeadLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitDeadLine) ProtoMessage() {}

func (x *WaitDeadLine) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitDeadLine.ProtoReflect.Descriptor instead.
func (*WaitDeadLine) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *WaitDeadLine) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *WaitDeadLine) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

//单人匹配
type MatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game GameType `protobuf:"varint,1,opt,name=game,proto3,enum=data.Dgame.GameType" json:"game,omitempty"`
}

func (x *MatchReq) Reset() {
	*x = MatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchReq) ProtoMessage() {}

func (x *MatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchReq.ProtoReflect.Descriptor instead.
func (*MatchReq) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *MatchReq) GetGame() GameType {
	if x != nil {
		return x.Game
	}
	return GameType__
}

//匹配状态
type MatchStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State          MatchState `protobuf:"varint,1,opt,name=state,proto3,enum=data.Dgame.MatchState" json:"state,omitempty"`
	Game           GameType   `protobuf:"varint,2,opt,name=game,proto3,enum=data.Dgame.GameType" json:"game,omitempty"`
	StartTS        int64      `protobuf:"varint,3,opt,name=startTS,proto3" json:"startTS,omitempty"`
	SureDeadLineTS int64      `protobuf:"varint,4,opt,name=sureDeadLineTS,proto3" json:"sureDeadLineTS,omitempty"`
}

func (x *MatchStateInfo) Reset() {
	*x = MatchStateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchStateInfo) ProtoMessage() {}

func (x *MatchStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchStateInfo.ProtoReflect.Descriptor instead.
func (*MatchStateInfo) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *MatchStateInfo) GetState() MatchState {
	if x != nil {
		return x.State
	}
	return MatchState_noMatch
}

func (x *MatchStateInfo) GetGame() GameType {
	if x != nil {
		return x.Game
	}
	return GameType__
}

func (x *MatchStateInfo) GetStartTS() int64 {
	if x != nil {
		return x.StartTS
	}
	return 0
}

func (x *MatchStateInfo) GetSureDeadLineTS() int64 {
	if x != nil {
		return x.SureDeadLineTS
	}
	return 0
}

//匹配成功信息
type MatchSucInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *MatchStateInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Players []string        `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"` //匹配成功的玩家
}

func (x *MatchSucInfo) Reset() {
	*x = MatchSucInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchSucInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchSucInfo) ProtoMessage() {}

func (x *MatchSucInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchSucInfo.ProtoReflect.Descriptor instead.
func (*MatchSucInfo) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

func (x *MatchSucInfo) GetInfo() *MatchStateInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *MatchSucInfo) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x44, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x73, 0x79,
	0x6e, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x11, 0x73, 0x79, 0x6e, 0x63, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x0c, 0x77, 0x61, 0x69, 0x74,
	0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x34, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x44, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x44, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x53, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x53, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x75, 0x72, 0x65, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x54, 0x53, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x72, 0x65, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e,
	0x65, 0x54, 0x53, 0x22, 0x58, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2a, 0x21, 0x0a,
	0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x10, 0x01,
	0x2a, 0x3f, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x6e, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x77, 0x61, 0x69,
	0x74, 0x53, 0x75, 0x72, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x75, 0x72, 0x65, 0x10,
	0x03, 0x42, 0x12, 0x5a, 0x10, 0x4a, 0x46, 0x46, 0x75, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x44, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_game_proto_goTypes = []interface{}{
	(GameType)(0),             // 0: data.Dgame.GameType
	(MatchState)(0),           // 1: data.Dgame.matchState
	(*Action)(nil),            // 2: data.Dgame.action
	(*StateEnd)(nil),          // 3: data.Dgame.stateEnd
	(*SyncGameState)(nil),     // 4: data.Dgame.syncGameState
	(*SyncGamePlayerAct)(nil), // 5: data.Dgame.syncGamePlayerAct
	(*WaitDeadLine)(nil),      // 6: data.Dgame.waitDeadLine
	(*MatchReq)(nil),          // 7: data.Dgame.matchReq
	(*MatchStateInfo)(nil),    // 8: data.Dgame.matchStateInfo
	(*MatchSucInfo)(nil),      // 9: data.Dgame.matchSucInfo
}
var file_game_proto_depIdxs = []int32{
	0, // 0: data.Dgame.matchReq.game:type_name -> data.Dgame.GameType
	1, // 1: data.Dgame.matchStateInfo.state:type_name -> data.Dgame.matchState
	0, // 2: data.Dgame.matchStateInfo.game:type_name -> data.Dgame.GameType
	8, // 3: data.Dgame.matchSucInfo.info:type_name -> data.Dgame.matchStateInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateEnd); i {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncGameState); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncGamePlayerAct); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitDeadLine); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchReq); i {
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
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchStateInfo); i {
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
		file_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchSucInfo); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
