// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package ChatRPC

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAccountRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type CreateAccountReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountReply) Reset()         { *m = CreateAccountReply{} }
func (m *CreateAccountReply) String() string { return proto.CompactTextString(m) }
func (*CreateAccountReply) ProtoMessage()    {}
func (*CreateAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *CreateAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountReply.Unmarshal(m, b)
}
func (m *CreateAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountReply.Marshal(b, m, deterministic)
}
func (m *CreateAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountReply.Merge(m, src)
}
func (m *CreateAccountReply) XXX_Size() int {
	return xxx_messageInfo_CreateAccountReply.Size(m)
}
func (m *CreateAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountReply proto.InternalMessageInfo

func (m *CreateAccountReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogInAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInAccountRequest) Reset()         { *m = LogInAccountRequest{} }
func (m *LogInAccountRequest) String() string { return proto.CompactTextString(m) }
func (*LogInAccountRequest) ProtoMessage()    {}
func (*LogInAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *LogInAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInAccountRequest.Unmarshal(m, b)
}
func (m *LogInAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInAccountRequest.Marshal(b, m, deterministic)
}
func (m *LogInAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInAccountRequest.Merge(m, src)
}
func (m *LogInAccountRequest) XXX_Size() int {
	return xxx_messageInfo_LogInAccountRequest.Size(m)
}
func (m *LogInAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogInAccountRequest proto.InternalMessageInfo

func (m *LogInAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogInAccountRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type LogInAccountReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInAccountReply) Reset()         { *m = LogInAccountReply{} }
func (m *LogInAccountReply) String() string { return proto.CompactTextString(m) }
func (*LogInAccountReply) ProtoMessage()    {}
func (*LogInAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *LogInAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInAccountReply.Unmarshal(m, b)
}
func (m *LogInAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInAccountReply.Marshal(b, m, deterministic)
}
func (m *LogInAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInAccountReply.Merge(m, src)
}
func (m *LogInAccountReply) XXX_Size() int {
	return xxx_messageInfo_LogInAccountReply.Size(m)
}
func (m *LogInAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_LogInAccountReply proto.InternalMessageInfo

func (m *LogInAccountReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EnterChatRoomRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RoomName             string   `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	Passwd               string   `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterChatRoomRequest) Reset()         { *m = EnterChatRoomRequest{} }
func (m *EnterChatRoomRequest) String() string { return proto.CompactTextString(m) }
func (*EnterChatRoomRequest) ProtoMessage()    {}
func (*EnterChatRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
}

func (m *EnterChatRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterChatRoomRequest.Unmarshal(m, b)
}
func (m *EnterChatRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterChatRoomRequest.Marshal(b, m, deterministic)
}
func (m *EnterChatRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterChatRoomRequest.Merge(m, src)
}
func (m *EnterChatRoomRequest) XXX_Size() int {
	return xxx_messageInfo_EnterChatRoomRequest.Size(m)
}
func (m *EnterChatRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterChatRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnterChatRoomRequest proto.InternalMessageInfo

func (m *EnterChatRoomRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *EnterChatRoomRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *EnterChatRoomRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type EnterChatRoomReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterChatRoomReply) Reset()         { *m = EnterChatRoomReply{} }
func (m *EnterChatRoomReply) String() string { return proto.CompactTextString(m) }
func (*EnterChatRoomReply) ProtoMessage()    {}
func (*EnterChatRoomReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{5}
}

func (m *EnterChatRoomReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterChatRoomReply.Unmarshal(m, b)
}
func (m *EnterChatRoomReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterChatRoomReply.Marshal(b, m, deterministic)
}
func (m *EnterChatRoomReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterChatRoomReply.Merge(m, src)
}
func (m *EnterChatRoomReply) XXX_Size() int {
	return xxx_messageInfo_EnterChatRoomReply.Size(m)
}
func (m *EnterChatRoomReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterChatRoomReply.DiscardUnknown(m)
}

var xxx_messageInfo_EnterChatRoomReply proto.InternalMessageInfo

func (m *EnterChatRoomReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LeaveChatRoomRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveChatRoomRequest) Reset()         { *m = LeaveChatRoomRequest{} }
func (m *LeaveChatRoomRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveChatRoomRequest) ProtoMessage()    {}
func (*LeaveChatRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{6}
}

func (m *LeaveChatRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveChatRoomRequest.Unmarshal(m, b)
}
func (m *LeaveChatRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveChatRoomRequest.Marshal(b, m, deterministic)
}
func (m *LeaveChatRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveChatRoomRequest.Merge(m, src)
}
func (m *LeaveChatRoomRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveChatRoomRequest.Size(m)
}
func (m *LeaveChatRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveChatRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveChatRoomRequest proto.InternalMessageInfo

func (m *LeaveChatRoomRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type LeaveChatRoomReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveChatRoomReply) Reset()         { *m = LeaveChatRoomReply{} }
func (m *LeaveChatRoomReply) String() string { return proto.CompactTextString(m) }
func (*LeaveChatRoomReply) ProtoMessage()    {}
func (*LeaveChatRoomReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{7}
}

func (m *LeaveChatRoomReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveChatRoomReply.Unmarshal(m, b)
}
func (m *LeaveChatRoomReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveChatRoomReply.Marshal(b, m, deterministic)
}
func (m *LeaveChatRoomReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveChatRoomReply.Merge(m, src)
}
func (m *LeaveChatRoomReply) XXX_Size() int {
	return xxx_messageInfo_LeaveChatRoomReply.Size(m)
}
func (m *LeaveChatRoomReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveChatRoomReply.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveChatRoomReply proto.InternalMessageInfo

func (m *LeaveChatRoomReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetListOfRoomRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListOfRoomRequest) Reset()         { *m = GetListOfRoomRequest{} }
func (m *GetListOfRoomRequest) String() string { return proto.CompactTextString(m) }
func (*GetListOfRoomRequest) ProtoMessage()    {}
func (*GetListOfRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{8}
}

func (m *GetListOfRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListOfRoomRequest.Unmarshal(m, b)
}
func (m *GetListOfRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListOfRoomRequest.Marshal(b, m, deterministic)
}
func (m *GetListOfRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListOfRoomRequest.Merge(m, src)
}
func (m *GetListOfRoomRequest) XXX_Size() int {
	return xxx_messageInfo_GetListOfRoomRequest.Size(m)
}
func (m *GetListOfRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListOfRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListOfRoomRequest proto.InternalMessageInfo

func (m *GetListOfRoomRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type GetListOfRoomReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListOfRoomReply) Reset()         { *m = GetListOfRoomReply{} }
func (m *GetListOfRoomReply) String() string { return proto.CompactTextString(m) }
func (*GetListOfRoomReply) ProtoMessage()    {}
func (*GetListOfRoomReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{9}
}

func (m *GetListOfRoomReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListOfRoomReply.Unmarshal(m, b)
}
func (m *GetListOfRoomReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListOfRoomReply.Marshal(b, m, deterministic)
}
func (m *GetListOfRoomReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListOfRoomReply.Merge(m, src)
}
func (m *GetListOfRoomReply) XXX_Size() int {
	return xxx_messageInfo_GetListOfRoomReply.Size(m)
}
func (m *GetListOfRoomReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListOfRoomReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetListOfRoomReply proto.InternalMessageInfo

func (m *GetListOfRoomReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMessageRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RoomName             string   `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{10}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SendMessageRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *SendMessageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SendMessageRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type SendMessageReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageReply) Reset()         { *m = SendMessageReply{} }
func (m *SendMessageReply) String() string { return proto.CompactTextString(m) }
func (*SendMessageReply) ProtoMessage()    {}
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{11}
}

func (m *SendMessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReply.Unmarshal(m, b)
}
func (m *SendMessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReply.Marshal(b, m, deterministic)
}
func (m *SendMessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReply.Merge(m, src)
}
func (m *SendMessageReply) XXX_Size() int {
	return xxx_messageInfo_SendMessageReply.Size(m)
}
func (m *SendMessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReply proto.InternalMessageInfo

func (m *SendMessageReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "ChatRPC.CreateAccountRequest")
	proto.RegisterType((*CreateAccountReply)(nil), "ChatRPC.CreateAccountReply")
	proto.RegisterType((*LogInAccountRequest)(nil), "ChatRPC.LogInAccountRequest")
	proto.RegisterType((*LogInAccountReply)(nil), "ChatRPC.LogInAccountReply")
	proto.RegisterType((*EnterChatRoomRequest)(nil), "ChatRPC.EnterChatRoomRequest")
	proto.RegisterType((*EnterChatRoomReply)(nil), "ChatRPC.EnterChatRoomReply")
	proto.RegisterType((*LeaveChatRoomRequest)(nil), "ChatRPC.LeaveChatRoomRequest")
	proto.RegisterType((*LeaveChatRoomReply)(nil), "ChatRPC.LeaveChatRoomReply")
	proto.RegisterType((*GetListOfRoomRequest)(nil), "ChatRPC.GetListOfRoomRequest")
	proto.RegisterType((*GetListOfRoomReply)(nil), "ChatRPC.GetListOfRoomReply")
	proto.RegisterType((*SendMessageRequest)(nil), "ChatRPC.SendMessageRequest")
	proto.RegisterType((*SendMessageReply)(nil), "ChatRPC.SendMessageReply")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdb, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xe9, 0x86, 0x00, 0x5d, 0x6c, 0x92, 0xbd, 0x67, 0x13, 0xd2, 0x5d, 0x30, 0x31, 0xbd,
	0xf2, 0x42, 0xab, 0x91, 0x27, 0x00, 0x62, 0x88, 0x49, 0x39, 0x04, 0xef, 0xbc, 0x31, 0x63, 0x19,
	0x81, 0x84, 0xce, 0xd4, 0xce, 0xa0, 0xe1, 0x05, 0x7c, 0x50, 0x9f, 0xc4, 0xf4, 0x24, 0xd3, 0x03,
	0x35, 0x84, 0x3b, 0xd6, 0x5a, 0xcc, 0xff, 0xff, 0x2c, 0xbe, 0x19, 0x00, 0x7b, 0x85, 0x85, 0xe9,
	0x7a, 0x4c, 0x30, 0x54, 0x1b, 0xae, 0xb0, 0x98, 0xcf, 0x86, 0xc6, 0x00, 0x5a, 0x43, 0x8f, 0x60,
	0x41, 0xfa, 0xb6, 0xcd, 0xb6, 0x54, 0xcc, 0xc9, 0xeb, 0x96, 0x70, 0x81, 0x10, 0x54, 0x28, 0x76,
	0x88, 0xa6, 0x9c, 0x2b, 0x17, 0xea, 0x3c, 0xf8, 0x8c, 0xda, 0x50, 0x75, 0x31, 0xe7, 0xef, 0x0b,
	0xed, 0x57, 0xd0, 0x8d, 0x2a, 0xc3, 0x04, 0x94, 0xd2, 0x70, 0x37, 0x3b, 0xa4, 0x41, 0xcd, 0x21,
	0x9c, 0xe3, 0x65, 0x2c, 0x12, 0x97, 0x46, 0x1f, 0xfe, 0x59, 0x6c, 0x79, 0x4f, 0x4f, 0xb0, 0xbc,
	0x82, 0xbf, 0x49, 0x89, 0x62, 0xc7, 0x15, 0xb4, 0xee, 0xa8, 0x20, 0x5e, 0xf0, 0xab, 0x19, 0x73,
	0x62, 0xcb, 0x0e, 0xa8, 0x5b, 0x4e, 0xbc, 0x27, 0xc9, 0xb7, 0xee, 0x37, 0x26, 0xbe, 0x77, 0x07,
	0x54, 0x8f, 0x31, 0x27, 0x1c, 0x86, 0xf6, 0x75, 0xbf, 0x31, 0x49, 0x06, 0x2b, 0xa7, 0x77, 0x91,
	0x72, 0x2a, 0x4e, 0xd6, 0x83, 0x96, 0x45, 0xf0, 0x1b, 0x39, 0x26, 0x99, 0x6f, 0x92, 0x3a, 0xf4,
	0xa3, 0xc9, 0x88, 0x08, 0x6b, 0xcd, 0xc5, 0xf4, 0xe5, 0x18, 0x93, 0xd4, 0xa1, 0x62, 0x93, 0x0f,
	0x05, 0xd0, 0x03, 0xa1, 0x8b, 0x71, 0x58, 0x9f, 0xbe, 0x62, 0x0d, 0x6a, 0x36, 0xa3, 0x82, 0x50,
	0x11, 0xed, 0x38, 0x2e, 0x51, 0x17, 0x54, 0xb1, 0x76, 0x08, 0x17, 0xd8, 0x71, 0xb5, 0x4a, 0x30,
	0xdb, 0x37, 0x8c, 0x4b, 0xf8, 0x93, 0xc8, 0x51, 0x18, 0xfb, 0xf6, 0xb3, 0x0c, 0x0d, 0x7f, 0x8f,
	0x63, 0x4c, 0xf1, 0x92, 0x78, 0x68, 0x0a, 0xcd, 0x04, 0xcc, 0xe8, 0xcc, 0x8c, 0xee, 0x8a, 0x99,
	0x77, 0x51, 0xf4, 0xce, 0xa1, 0xb1, 0xbb, 0xd9, 0x19, 0xa5, 0x1b, 0x05, 0x59, 0xf0, 0x5b, 0x46,
	0x15, 0x75, 0xbf, 0x0f, 0xe4, 0x5c, 0x02, 0x5d, 0x3f, 0x30, 0x8d, 0xd5, 0xc6, 0xd0, 0x4c, 0xf0,
	0x25, 0xc5, 0xcb, 0x23, 0x5c, 0x8a, 0x97, 0xc5, 0xd2, 0x28, 0xf9, 0x72, 0x09, 0x92, 0x24, 0xb9,
	0x3c, 0x2c, 0x25, 0xb9, 0x2c, 0x80, 0xa1, 0x5c, 0x82, 0x19, 0x49, 0x2e, 0x0f, 0x40, 0x49, 0x2e,
	0x8b, 0x9a, 0x51, 0x42, 0x23, 0x68, 0x48, 0xff, 0x24, 0xda, 0x7f, 0x3b, 0xcb, 0x99, 0xfe, 0x3f,
	0x7f, 0x18, 0x08, 0x0d, 0xda, 0x83, 0xf8, 0xc1, 0x9b, 0x29, 0x8f, 0xaa, 0x79, 0x1d, 0x15, 0xcf,
	0xd5, 0xe0, 0x35, 0xec, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xd9, 0xec, 0xd0, 0x1b, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatManagerClient is the client API for ChatManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatManagerClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (ChatManager_CreateAccountClient, error)
	LogInAccount(ctx context.Context, in *LogInAccountRequest, opts ...grpc.CallOption) (ChatManager_LogInAccountClient, error)
	EnterChatRoom(ctx context.Context, in *EnterChatRoomRequest, opts ...grpc.CallOption) (*EnterChatRoomReply, error)
	LeaveChatRoom(ctx context.Context, in *LeaveChatRoomRequest, opts ...grpc.CallOption) (*LeaveChatRoomReply, error)
	GetListOfRoom(ctx context.Context, in *GetListOfRoomRequest, opts ...grpc.CallOption) (*GetListOfRoomReply, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageReply, error)
}

type chatManagerClient struct {
	cc *grpc.ClientConn
}

func NewChatManagerClient(cc *grpc.ClientConn) ChatManagerClient {
	return &chatManagerClient{cc}
}

func (c *chatManagerClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (ChatManager_CreateAccountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatManager_serviceDesc.Streams[0], "/ChatRPC.ChatManager/CreateAccount", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatManagerCreateAccountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatManager_CreateAccountClient interface {
	Recv() (*CreateAccountReply, error)
	grpc.ClientStream
}

type chatManagerCreateAccountClient struct {
	grpc.ClientStream
}

func (x *chatManagerCreateAccountClient) Recv() (*CreateAccountReply, error) {
	m := new(CreateAccountReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatManagerClient) LogInAccount(ctx context.Context, in *LogInAccountRequest, opts ...grpc.CallOption) (ChatManager_LogInAccountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatManager_serviceDesc.Streams[1], "/ChatRPC.ChatManager/LogInAccount", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatManagerLogInAccountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatManager_LogInAccountClient interface {
	Recv() (*LogInAccountReply, error)
	grpc.ClientStream
}

type chatManagerLogInAccountClient struct {
	grpc.ClientStream
}

func (x *chatManagerLogInAccountClient) Recv() (*LogInAccountReply, error) {
	m := new(LogInAccountReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatManagerClient) EnterChatRoom(ctx context.Context, in *EnterChatRoomRequest, opts ...grpc.CallOption) (*EnterChatRoomReply, error) {
	out := new(EnterChatRoomReply)
	err := c.cc.Invoke(ctx, "/ChatRPC.ChatManager/EnterChatRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatManagerClient) LeaveChatRoom(ctx context.Context, in *LeaveChatRoomRequest, opts ...grpc.CallOption) (*LeaveChatRoomReply, error) {
	out := new(LeaveChatRoomReply)
	err := c.cc.Invoke(ctx, "/ChatRPC.ChatManager/LeaveChatRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatManagerClient) GetListOfRoom(ctx context.Context, in *GetListOfRoomRequest, opts ...grpc.CallOption) (*GetListOfRoomReply, error) {
	out := new(GetListOfRoomReply)
	err := c.cc.Invoke(ctx, "/ChatRPC.ChatManager/GetListOfRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatManagerClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageReply, error) {
	out := new(SendMessageReply)
	err := c.cc.Invoke(ctx, "/ChatRPC.ChatManager/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatManagerServer is the server API for ChatManager service.
type ChatManagerServer interface {
	CreateAccount(*CreateAccountRequest, ChatManager_CreateAccountServer) error
	LogInAccount(*LogInAccountRequest, ChatManager_LogInAccountServer) error
	EnterChatRoom(context.Context, *EnterChatRoomRequest) (*EnterChatRoomReply, error)
	LeaveChatRoom(context.Context, *LeaveChatRoomRequest) (*LeaveChatRoomReply, error)
	GetListOfRoom(context.Context, *GetListOfRoomRequest) (*GetListOfRoomReply, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageReply, error)
}

// UnimplementedChatManagerServer can be embedded to have forward compatible implementations.
type UnimplementedChatManagerServer struct {
}

func (*UnimplementedChatManagerServer) CreateAccount(req *CreateAccountRequest, srv ChatManager_CreateAccountServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedChatManagerServer) LogInAccount(req *LogInAccountRequest, srv ChatManager_LogInAccountServer) error {
	return status.Errorf(codes.Unimplemented, "method LogInAccount not implemented")
}
func (*UnimplementedChatManagerServer) EnterChatRoom(ctx context.Context, req *EnterChatRoomRequest) (*EnterChatRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnterChatRoom not implemented")
}
func (*UnimplementedChatManagerServer) LeaveChatRoom(ctx context.Context, req *LeaveChatRoomRequest) (*LeaveChatRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChatRoom not implemented")
}
func (*UnimplementedChatManagerServer) GetListOfRoom(ctx context.Context, req *GetListOfRoomRequest) (*GetListOfRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfRoom not implemented")
}
func (*UnimplementedChatManagerServer) SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterChatManagerServer(s *grpc.Server, srv ChatManagerServer) {
	s.RegisterService(&_ChatManager_serviceDesc, srv)
}

func _ChatManager_CreateAccount_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateAccountRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatManagerServer).CreateAccount(m, &chatManagerCreateAccountServer{stream})
}

type ChatManager_CreateAccountServer interface {
	Send(*CreateAccountReply) error
	grpc.ServerStream
}

type chatManagerCreateAccountServer struct {
	grpc.ServerStream
}

func (x *chatManagerCreateAccountServer) Send(m *CreateAccountReply) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatManager_LogInAccount_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogInAccountRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatManagerServer).LogInAccount(m, &chatManagerLogInAccountServer{stream})
}

type ChatManager_LogInAccountServer interface {
	Send(*LogInAccountReply) error
	grpc.ServerStream
}

type chatManagerLogInAccountServer struct {
	grpc.ServerStream
}

func (x *chatManagerLogInAccountServer) Send(m *LogInAccountReply) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatManager_EnterChatRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterChatRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatManagerServer).EnterChatRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatRPC.ChatManager/EnterChatRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatManagerServer).EnterChatRoom(ctx, req.(*EnterChatRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatManager_LeaveChatRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveChatRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatManagerServer).LeaveChatRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatRPC.ChatManager/LeaveChatRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatManagerServer).LeaveChatRoom(ctx, req.(*LeaveChatRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatManager_GetListOfRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOfRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatManagerServer).GetListOfRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatRPC.ChatManager/GetListOfRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatManagerServer).GetListOfRoom(ctx, req.(*GetListOfRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatManager_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatManagerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatRPC.ChatManager/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatManagerServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChatRPC.ChatManager",
	HandlerType: (*ChatManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnterChatRoom",
			Handler:    _ChatManager_EnterChatRoom_Handler,
		},
		{
			MethodName: "LeaveChatRoom",
			Handler:    _ChatManager_LeaveChatRoom_Handler,
		},
		{
			MethodName: "GetListOfRoom",
			Handler:    _ChatManager_GetListOfRoom_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatManager_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateAccount",
			Handler:       _ChatManager_CreateAccount_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogInAccount",
			Handler:       _ChatManager_LogInAccount_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
