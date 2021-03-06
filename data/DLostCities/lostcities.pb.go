// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: lostcities.proto

package DLostCities

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

type State int32

const (
	State_deal    State = 0 //发牌
	State_play    State = 1 //出牌
	State_getCard State = 2 //获得牌
	State_balance State = 3 //结算
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "deal",
		1: "play",
		2: "getCard",
		3: "balance",
	}
	State_value = map[string]int32{
		"deal":    0,
		"play":    1,
		"getCard": 2,
		"balance": 3,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_lostcities_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_lostcities_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{0}
}

type Act int32

const (
	Act__      Act = 0
	Act_invest Act = 1 //投资
	Act_drop   Act = 2 //弃牌
	Act_choose Act = 3 //选择牌
	Act_draw   Act = 4 //抽牌
)

// Enum value maps for Act.
var (
	Act_name = map[int32]string{
		0: "_",
		1: "invest",
		2: "drop",
		3: "choose",
		4: "draw",
	}
	Act_value = map[string]int32{
		"_":      0,
		"invest": 1,
		"drop":   2,
		"choose": 3,
		"draw":   4,
	}
)

func (x Act) Enum() *Act {
	p := new(Act)
	*p = x
	return p
}

func (x Act) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Act) Descriptor() protoreflect.EnumDescriptor {
	return file_lostcities_proto_enumTypes[1].Descriptor()
}

func (Act) Type() protoreflect.EnumType {
	return &file_lostcities_proto_enumTypes[1]
}

func (x Act) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Act.Descriptor instead.
func (Act) EnumDescriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{1}
}

type GameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State     State         `protobuf:"varint,1,opt,name=state,proto3,enum=data.DLostCities.State" json:"state,omitempty"`
	Players   []*PlayerInfo `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	DropArea  []*DropArea   `protobuf:"bytes,3,rep,name=dropArea,proto3" json:"dropArea,omitempty"`
	Deck      int32         `protobuf:"varint,4,opt,name=deck,proto3" json:"deck,omitempty"`
	CurPlayer string        `protobuf:"bytes,5,opt,name=curPlayer,proto3" json:"curPlayer,omitempty"`
}

func (x *GameInfo) Reset() {
	*x = GameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameInfo) ProtoMessage() {}

func (x *GameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameInfo.ProtoReflect.Descriptor instead.
func (*GameInfo) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{0}
}

func (x *GameInfo) GetState() State {
	if x != nil {
		return x.State
	}
	return State_deal
}

func (x *GameInfo) GetPlayers() []*PlayerInfo {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameInfo) GetDropArea() []*DropArea {
	if x != nil {
		return x.DropArea
	}
	return nil
}

func (x *GameInfo) GetDeck() int32 {
	if x != nil {
		return x.Deck
	}
	return 0
}

func (x *GameInfo) GetCurPlayer() string {
	if x != nil {
		return x.CurPlayer
	}
	return ""
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string  `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Hands    []*Card `protobuf:"bytes,2,rep,name=hands,proto3" json:"hands,omitempty"` //手牌
	Table    []*Card `protobuf:"bytes,3,rep,name=table,proto3" json:"table,omitempty"` //出牌区
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerInfo) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerInfo) GetHands() []*Card {
	if x != nil {
		return x.Hands
	}
	return nil
}

func (x *PlayerInfo) GetTable() []*Card {
	if x != nil {
		return x.Table
	}
	return nil
}

type DropArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City  int32   `protobuf:"varint,1,opt,name=city,proto3" json:"city,omitempty"`
	Cards []*Card `protobuf:"bytes,2,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *DropArea) Reset() {
	*x = DropArea{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropArea) ProtoMessage() {}

func (x *DropArea) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropArea.ProtoReflect.Descriptor instead.
func (*DropArea) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{2}
}

func (x *DropArea) GetCity() int32 {
	if x != nil {
		return x.City
	}
	return 0
}

func (x *DropArea) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City  int32 `protobuf:"varint,1,opt,name=city,proto3" json:"city,omitempty"`
	Point int32 `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{3}
}

func (x *Card) GetCity() int32 {
	if x != nil {
		return x.City
	}
	return 0
}

func (x *Card) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type Cards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *Cards) Reset() {
	*x = Cards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cards) ProtoMessage() {}

func (x *Cards) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cards.ProtoReflect.Descriptor instead.
func (*Cards) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{4}
}

func (x *Cards) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type BalanceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*Score `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *BalanceResult) Reset() {
	*x = BalanceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResult) ProtoMessage() {}

func (x *BalanceResult) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResult.ProtoReflect.Descriptor instead.
func (*BalanceResult) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{5}
}

func (x *BalanceResult) GetDetails() []*Score {
	if x != nil {
		return x.Details
	}
	return nil
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player  string  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Score   int32   `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`            //总分
	Details []int32 `protobuf:"varint,3,rep,packed,name=details,proto3" json:"details,omitempty"` //各堆得分
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lostcities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_lostcities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_lostcities_proto_rawDescGZIP(), []int{6}
}

func (x *Score) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *Score) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score) GetDetails() []int32 {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_lostcities_proto protoreflect.FileDescriptor

var file_lostcities_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x73, 0x74, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x36, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x72, 0x6f, 0x70,
	0x41, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x64, 0x72,
	0x6f, 0x70, 0x41, 0x72, 0x65, 0x61, 0x52, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x41, 0x72, 0x65, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x64, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x75, 0x72, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x05, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x52, 0x05, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x4c, 0x0a, 0x08, 0x64, 0x72, 0x6f,
	0x70, 0x41, 0x72, 0x65, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x05, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x22, 0x42, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x4f, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x79,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x03, 0x2a, 0x38, 0x0a, 0x03,
	0x61, 0x63, 0x74, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x72, 0x6f, 0x70, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04,
	0x64, 0x72, 0x61, 0x77, 0x10, 0x04, 0x42, 0x18, 0x5a, 0x16, 0x4a, 0x46, 0x46, 0x75, 0x6e, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x44, 0x4c, 0x6f, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lostcities_proto_rawDescOnce sync.Once
	file_lostcities_proto_rawDescData = file_lostcities_proto_rawDesc
)

func file_lostcities_proto_rawDescGZIP() []byte {
	file_lostcities_proto_rawDescOnce.Do(func() {
		file_lostcities_proto_rawDescData = protoimpl.X.CompressGZIP(file_lostcities_proto_rawDescData)
	})
	return file_lostcities_proto_rawDescData
}

var file_lostcities_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_lostcities_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lostcities_proto_goTypes = []interface{}{
	(State)(0),            // 0: data.DLostCities.state
	(Act)(0),              // 1: data.DLostCities.act
	(*GameInfo)(nil),      // 2: data.DLostCities.gameInfo
	(*PlayerInfo)(nil),    // 3: data.DLostCities.playerInfo
	(*DropArea)(nil),      // 4: data.DLostCities.dropArea
	(*Card)(nil),          // 5: data.DLostCities.card
	(*Cards)(nil),         // 6: data.DLostCities.cards
	(*BalanceResult)(nil), // 7: data.DLostCities.balanceResult
	(*Score)(nil),         // 8: data.DLostCities.score
}
var file_lostcities_proto_depIdxs = []int32{
	0, // 0: data.DLostCities.gameInfo.state:type_name -> data.DLostCities.state
	3, // 1: data.DLostCities.gameInfo.players:type_name -> data.DLostCities.playerInfo
	4, // 2: data.DLostCities.gameInfo.dropArea:type_name -> data.DLostCities.dropArea
	5, // 3: data.DLostCities.playerInfo.hands:type_name -> data.DLostCities.card
	5, // 4: data.DLostCities.playerInfo.table:type_name -> data.DLostCities.card
	5, // 5: data.DLostCities.dropArea.cards:type_name -> data.DLostCities.card
	5, // 6: data.DLostCities.cards.cards:type_name -> data.DLostCities.card
	8, // 7: data.DLostCities.balanceResult.details:type_name -> data.DLostCities.score
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_lostcities_proto_init() }
func file_lostcities_proto_init() {
	if File_lostcities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lostcities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameInfo); i {
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
		file_lostcities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_lostcities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropArea); i {
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
		file_lostcities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_lostcities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cards); i {
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
		file_lostcities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResult); i {
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
		file_lostcities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
			RawDescriptor: file_lostcities_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lostcities_proto_goTypes,
		DependencyIndexes: file_lostcities_proto_depIdxs,
		EnumInfos:         file_lostcities_proto_enumTypes,
		MessageInfos:      file_lostcities_proto_msgTypes,
	}.Build()
	File_lostcities_proto = out.File
	file_lostcities_proto_rawDesc = nil
	file_lostcities_proto_goTypes = nil
	file_lostcities_proto_depIdxs = nil
}
