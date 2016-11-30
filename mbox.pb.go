// Code generated by protoc-gen-go.
// source: mbox.proto
// DO NOT EDIT!

/*
Package mbox is a generated protocol buffer package.

It is generated from these files:
	mbox.proto

It has these top-level messages:
	Message
	Result
	Folder
*/
package mbox

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	From    string   `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To      []string `protobuf:"bytes,2,rep,name=to" json:"to,omitempty"`
	Subject string   `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	// Types that are valid to be assigned to Body:
	//	*Message_Text
	//	*Message_Html
	Body isMessage_Body `protobuf_oneof:"body"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isMessage_Body interface {
	isMessage_Body()
}

type Message_Text struct {
	Text string `protobuf:"bytes,4,opt,name=text,oneof"`
}
type Message_Html struct {
	Html string `protobuf:"bytes,5,opt,name=html,oneof"`
}

func (*Message_Text) isMessage_Body() {}
func (*Message_Html) isMessage_Body() {}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetText() string {
	if x, ok := m.GetBody().(*Message_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Message) GetHtml() string {
	if x, ok := m.GetBody().(*Message_Html); ok {
		return x.Html
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Text)(nil),
		(*Message_Html)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// body
	switch x := m.Body.(type) {
	case *Message_Text:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Message_Html:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Html)
	case nil:
	default:
		return fmt.Errorf("Message.Body has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 4: // body.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Body = &Message_Text{x}
		return true, err
	case 5: // body.html
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Body = &Message_Html{x}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// body
	switch x := m.Body.(type) {
	case *Message_Text:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Message_Html:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Html)))
		n += len(x.Html)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Result struct {
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Folder struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Folder) Reset()                    { *m = Folder{} }
func (m *Folder) String() string            { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()               {}
func (*Folder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Folder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "mbox.Message")
	proto.RegisterType((*Result)(nil), "mbox.Result")
	proto.RegisterType((*Folder)(nil), "mbox.Folder")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mail service

type MailClient interface {
	// Send mail
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Result, error)
	// Receive new mail
	Receive(ctx context.Context, in *Folder, opts ...grpc.CallOption) (Mail_ReceiveClient, error)
}

type mailClient struct {
	cc *grpc.ClientConn
}

func NewMailClient(cc *grpc.ClientConn) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/mbox.Mail/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailClient) Receive(ctx context.Context, in *Folder, opts ...grpc.CallOption) (Mail_ReceiveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Mail_serviceDesc.Streams[0], c.cc, "/mbox.Mail/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &mailReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mail_ReceiveClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type mailReceiveClient struct {
	grpc.ClientStream
}

func (x *mailReceiveClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Mail service

type MailServer interface {
	// Send mail
	Send(context.Context, *Message) (*Result, error)
	// Receive new mail
	Receive(*Folder, Mail_ReceiveServer) error
}

func RegisterMailServer(s *grpc.Server, srv MailServer) {
	s.RegisterService(&_Mail_serviceDesc, srv)
}

func _Mail_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbox.Mail/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).Send(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mail_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Folder)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MailServer).Receive(m, &mailReceiveServer{stream})
}

type Mail_ReceiveServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type mailReceiveServer struct {
	grpc.ServerStream
}

func (x *mailReceiveServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Mail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mbox.Mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Mail_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _Mail_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mbox.proto",
}

func init() { proto.RegisterFile("mbox.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0x86, 0x30,
	0x10, 0x86, 0xe5, 0xfb, 0x6a, 0xd1, 0x8b, 0x3a, 0x5c, 0x1c, 0x1a, 0xe2, 0x40, 0x70, 0x61, 0x22,
	0x46, 0xff, 0x81, 0x83, 0x71, 0x61, 0xa9, 0xf1, 0x07, 0x50, 0x38, 0x15, 0xd3, 0x52, 0x43, 0x8b,
	0xc1, 0xf8, 0xe7, 0xa5, 0x05, 0x06, 0xb7, 0xf7, 0x79, 0xda, 0x5e, 0xde, 0x1e, 0x80, 0x51, 0x76,
	0xae, 0xbe, 0x46, 0xeb, 0x2d, 0xb2, 0x90, 0x8b, 0x5f, 0x48, 0x6b, 0x72, 0xae, 0x79, 0x27, 0x44,
	0x60, 0x6f, 0xa3, 0x35, 0x22, 0xc9, 0x93, 0xf2, 0x5c, 0xc6, 0x8c, 0x57, 0x70, 0xf0, 0x56, 0x1c,
	0xf2, 0xe3, 0x62, 0x96, 0x84, 0x02, 0x52, 0x37, 0xa9, 0x4f, 0x6a, 0xbd, 0x38, 0xc6, 0x6b, 0x3b,
	0xe2, 0x35, 0x30, 0x4f, 0xb3, 0x17, 0x2c, 0xe8, 0xe7, 0x13, 0x19, 0x29, 0xd8, 0x0f, 0x6f, 0xb4,
	0x38, 0xdd, 0x6d, 0xa0, 0x47, 0x0e, 0x4c, 0xd9, 0xee, 0xa7, 0x38, 0x03, 0x2e, 0xc9, 0x4d, 0xda,
	0x17, 0x37, 0xc0, 0x9f, 0xac, 0xee, 0x68, 0x0c, 0x2d, 0x86, 0xc6, 0xd0, 0xde, 0x22, 0xe4, 0xfb,
	0x57, 0x60, 0x75, 0xd3, 0x6b, 0xbc, 0x05, 0xf6, 0x42, 0x43, 0x87, 0x97, 0x55, 0xfc, 0xc7, 0x56,
	0x3c, 0xbb, 0x58, 0x71, 0x1d, 0x85, 0x25, 0xa4, 0x92, 0x5a, 0xea, 0xbf, 0x09, 0xb7, 0x83, 0x75,
	0x72, 0xf6, 0xff, 0xd5, 0x5d, 0xa2, 0x78, 0x5c, 0xc4, 0xc3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x72, 0x91, 0xdf, 0x54, 0x16, 0x01, 0x00, 0x00,
}
