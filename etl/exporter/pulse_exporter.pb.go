// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etl/exporter/pulse_exporter.proto

package exporter

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetNextPulseResponse struct {
	PulseNumber     uint32 `protobuf:"varint,1,opt,name=PulseNumber,proto3" json:"PulseNumber,omitempty"`
	PrevPulseNumber uint32 `protobuf:"varint,2,opt,name=PrevPulseNumber,proto3" json:"PrevPulseNumber,omitempty"`
	RecordAmount    uint32 `protobuf:"varint,3,opt,name=RecordAmount,proto3" json:"RecordAmount,omitempty"`
}

func (m *GetNextPulseResponse) Reset()      { *m = GetNextPulseResponse{} }
func (*GetNextPulseResponse) ProtoMessage() {}
func (*GetNextPulseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6d0e77f1c7203d, []int{0}
}
func (m *GetNextPulseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNextPulseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNextPulseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNextPulseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextPulseResponse.Merge(m, src)
}
func (m *GetNextPulseResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetNextPulseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextPulseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextPulseResponse proto.InternalMessageInfo

func (m *GetNextPulseResponse) GetPulseNumber() uint32 {
	if m != nil {
		return m.PulseNumber
	}
	return 0
}

func (m *GetNextPulseResponse) GetPrevPulseNumber() uint32 {
	if m != nil {
		return m.PrevPulseNumber
	}
	return 0
}

func (m *GetNextPulseResponse) GetRecordAmount() uint32 {
	if m != nil {
		return m.RecordAmount
	}
	return 0
}

type GetNextPulseRequest struct {
	PulseNumberFrom uint32   `protobuf:"varint,1,opt,name=PulseNumberFrom,proto3" json:"PulseNumberFrom,omitempty"`
	Prototypes      [][]byte `protobuf:"bytes,2,rep,name=Prototypes,proto3" json:"Prototypes,omitempty"`
}

func (m *GetNextPulseRequest) Reset()      { *m = GetNextPulseRequest{} }
func (*GetNextPulseRequest) ProtoMessage() {}
func (*GetNextPulseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6d0e77f1c7203d, []int{1}
}
func (m *GetNextPulseRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNextPulseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNextPulseRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNextPulseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextPulseRequest.Merge(m, src)
}
func (m *GetNextPulseRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetNextPulseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextPulseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextPulseRequest proto.InternalMessageInfo

func (m *GetNextPulseRequest) GetPulseNumberFrom() uint32 {
	if m != nil {
		return m.PulseNumberFrom
	}
	return 0
}

func (m *GetNextPulseRequest) GetPrototypes() [][]byte {
	if m != nil {
		return m.Prototypes
	}
	return nil
}

func init() {
	proto.RegisterType((*GetNextPulseResponse)(nil), "exporter.GetNextPulseResponse")
	proto.RegisterType((*GetNextPulseRequest)(nil), "exporter.GetNextPulseRequest")
}

func init() { proto.RegisterFile("etl/exporter/pulse_exporter.proto", fileDescriptor_1a6d0e77f1c7203d) }

var fileDescriptor_1a6d0e77f1c7203d = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x7d, 0xad, 0x84, 0xd0, 0x91, 0x0a, 0xc9, 0x30, 0x44, 0x48, 0x9c, 0x42, 0xa6, 0x4c,
	0x2d, 0x02, 0x5e, 0x00, 0x24, 0x60, 0xab, 0x42, 0x5e, 0xa0, 0x52, 0xe0, 0xb6, 0xa6, 0x0e, 0x8e,
	0x83, 0xca, 0xc6, 0xc4, 0xcc, 0x63, 0xf0, 0x28, 0x8c, 0x19, 0x3b, 0x12, 0x67, 0x61, 0xec, 0x23,
	0xa0, 0x9a, 0x16, 0xd2, 0x8a, 0x8e, 0xf7, 0xf9, 0xb3, 0xef, 0xbf, 0x33, 0x9e, 0xb0, 0x19, 0x0f,
	0x78, 0x9a, 0x2b, 0x6d, 0x58, 0x0f, 0xf2, 0x72, 0x5c, 0xf0, 0x68, 0x55, 0xf6, 0x73, 0xad, 0x8c,
	0x92, 0xbb, 0xab, 0x3a, 0x7c, 0x05, 0x3c, 0xbc, 0x65, 0x33, 0xe4, 0xa9, 0x89, 0x17, 0x66, 0xc2,
	0x45, 0xae, 0x26, 0x05, 0xcb, 0x00, 0xf7, 0x1c, 0x18, 0x96, 0x59, 0xca, 0xda, 0x87, 0x00, 0xa2,
	0x5e, 0xd2, 0x46, 0x32, 0xc2, 0xfd, 0x58, 0xf3, 0x53, 0xdb, 0xea, 0x38, 0x6b, 0x13, 0xcb, 0x10,
	0xbd, 0x84, 0xef, 0x95, 0x7e, 0xb8, 0xcc, 0x54, 0x39, 0x31, 0x7e, 0xd7, 0x69, 0x6b, 0x2c, 0x1c,
	0xe1, 0xc1, 0x7a, 0x8e, 0xc7, 0x92, 0x0b, 0xe3, 0x9a, 0xfc, 0xbd, 0x74, 0xa3, 0x55, 0xb6, 0x8c,
	0xb2, 0x89, 0x25, 0x21, 0xc6, 0x8b, 0xe1, 0xcc, 0x73, 0xce, 0x85, 0xdf, 0x09, 0xba, 0x91, 0x97,
	0xb4, 0xc8, 0x59, 0x8a, 0x3d, 0x77, 0xe5, 0x7a, 0x39, 0xba, 0xbc, 0x43, 0xaf, 0xdd, 0x51, 0x1e,
	0xf7, 0x7f, 0xb7, 0xf4, 0x4f, 0x92, 0x23, 0xda, 0x76, 0xfc, 0xb3, 0xb0, 0x50, 0x9c, 0xc2, 0xd5,
	0x45, 0x55, 0x93, 0x98, 0xd5, 0x24, 0xe6, 0x35, 0xc1, 0x8b, 0x25, 0x78, 0xb7, 0x04, 0x1f, 0x96,
	0xa0, 0xb2, 0x04, 0x9f, 0x96, 0xe0, 0xcb, 0x92, 0x98, 0x5b, 0x82, 0xb7, 0x86, 0x44, 0xd5, 0x90,
	0x98, 0x35, 0x24, 0xd2, 0x1d, 0xf7, 0x29, 0xe7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xc1,
	0x73, 0xcf, 0xb9, 0x01, 0x00, 0x00,
}

func (this *GetNextPulseResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetNextPulseResponse)
	if !ok {
		that2, ok := that.(GetNextPulseResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PulseNumber != that1.PulseNumber {
		return false
	}
	if this.PrevPulseNumber != that1.PrevPulseNumber {
		return false
	}
	if this.RecordAmount != that1.RecordAmount {
		return false
	}
	return true
}
func (this *GetNextPulseRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetNextPulseRequest)
	if !ok {
		that2, ok := that.(GetNextPulseRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PulseNumberFrom != that1.PulseNumberFrom {
		return false
	}
	if len(this.Prototypes) != len(that1.Prototypes) {
		return false
	}
	for i := range this.Prototypes {
		if !bytes.Equal(this.Prototypes[i], that1.Prototypes[i]) {
			return false
		}
	}
	return true
}
func (this *GetNextPulseResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&exporter.GetNextPulseResponse{")
	s = append(s, "PulseNumber: "+fmt.Sprintf("%#v", this.PulseNumber)+",\n")
	s = append(s, "PrevPulseNumber: "+fmt.Sprintf("%#v", this.PrevPulseNumber)+",\n")
	s = append(s, "RecordAmount: "+fmt.Sprintf("%#v", this.RecordAmount)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetNextPulseRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&exporter.GetNextPulseRequest{")
	s = append(s, "PulseNumberFrom: "+fmt.Sprintf("%#v", this.PulseNumberFrom)+",\n")
	s = append(s, "Prototypes: "+fmt.Sprintf("%#v", this.Prototypes)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPulseExporter(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PulseExporterClient is the client API for PulseExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PulseExporterClient interface {
	GetNextPulse(ctx context.Context, in *GetNextPulseRequest, opts ...grpc.CallOption) (PulseExporter_GetNextPulseClient, error)
}

type pulseExporterClient struct {
	cc *grpc.ClientConn
}

func NewPulseExporterClient(cc *grpc.ClientConn) PulseExporterClient {
	return &pulseExporterClient{cc}
}

func (c *pulseExporterClient) GetNextPulse(ctx context.Context, in *GetNextPulseRequest, opts ...grpc.CallOption) (PulseExporter_GetNextPulseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PulseExporter_serviceDesc.Streams[0], "/exporter.PulseExporter/GetNextPulse", opts...)
	if err != nil {
		return nil, err
	}
	x := &pulseExporterGetNextPulseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PulseExporter_GetNextPulseClient interface {
	Recv() (*GetNextPulseResponse, error)
	grpc.ClientStream
}

type pulseExporterGetNextPulseClient struct {
	grpc.ClientStream
}

func (x *pulseExporterGetNextPulseClient) Recv() (*GetNextPulseResponse, error) {
	m := new(GetNextPulseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PulseExporterServer is the server API for PulseExporter service.
type PulseExporterServer interface {
	GetNextPulse(*GetNextPulseRequest, PulseExporter_GetNextPulseServer) error
}

// UnimplementedPulseExporterServer can be embedded to have forward compatible implementations.
type UnimplementedPulseExporterServer struct {
}

func (*UnimplementedPulseExporterServer) GetNextPulse(req *GetNextPulseRequest, srv PulseExporter_GetNextPulseServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNextPulse not implemented")
}

func RegisterPulseExporterServer(s *grpc.Server, srv PulseExporterServer) {
	s.RegisterService(&_PulseExporter_serviceDesc, srv)
}

func _PulseExporter_GetNextPulse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNextPulseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PulseExporterServer).GetNextPulse(m, &pulseExporterGetNextPulseServer{stream})
}

type PulseExporter_GetNextPulseServer interface {
	Send(*GetNextPulseResponse) error
	grpc.ServerStream
}

type pulseExporterGetNextPulseServer struct {
	grpc.ServerStream
}

func (x *pulseExporterGetNextPulseServer) Send(m *GetNextPulseResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PulseExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exporter.PulseExporter",
	HandlerType: (*PulseExporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNextPulse",
			Handler:       _PulseExporter_GetNextPulse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "etl/exporter/pulse_exporter.proto",
}

func (m *GetNextPulseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNextPulseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNextPulseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RecordAmount != 0 {
		i = encodeVarintPulseExporter(dAtA, i, uint64(m.RecordAmount))
		i--
		dAtA[i] = 0x18
	}
	if m.PrevPulseNumber != 0 {
		i = encodeVarintPulseExporter(dAtA, i, uint64(m.PrevPulseNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.PulseNumber != 0 {
		i = encodeVarintPulseExporter(dAtA, i, uint64(m.PulseNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetNextPulseRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNextPulseRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNextPulseRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prototypes) > 0 {
		for iNdEx := len(m.Prototypes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Prototypes[iNdEx])
			copy(dAtA[i:], m.Prototypes[iNdEx])
			i = encodeVarintPulseExporter(dAtA, i, uint64(len(m.Prototypes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PulseNumberFrom != 0 {
		i = encodeVarintPulseExporter(dAtA, i, uint64(m.PulseNumberFrom))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPulseExporter(dAtA []byte, offset int, v uint64) int {
	offset -= sovPulseExporter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetNextPulseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PulseNumber != 0 {
		n += 1 + sovPulseExporter(uint64(m.PulseNumber))
	}
	if m.PrevPulseNumber != 0 {
		n += 1 + sovPulseExporter(uint64(m.PrevPulseNumber))
	}
	if m.RecordAmount != 0 {
		n += 1 + sovPulseExporter(uint64(m.RecordAmount))
	}
	return n
}

func (m *GetNextPulseRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PulseNumberFrom != 0 {
		n += 1 + sovPulseExporter(uint64(m.PulseNumberFrom))
	}
	if len(m.Prototypes) > 0 {
		for _, b := range m.Prototypes {
			l = len(b)
			n += 1 + l + sovPulseExporter(uint64(l))
		}
	}
	return n
}

func sovPulseExporter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPulseExporter(x uint64) (n int) {
	return sovPulseExporter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetNextPulseResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetNextPulseResponse{`,
		`PulseNumber:` + fmt.Sprintf("%v", this.PulseNumber) + `,`,
		`PrevPulseNumber:` + fmt.Sprintf("%v", this.PrevPulseNumber) + `,`,
		`RecordAmount:` + fmt.Sprintf("%v", this.RecordAmount) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetNextPulseRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetNextPulseRequest{`,
		`PulseNumberFrom:` + fmt.Sprintf("%v", this.PulseNumberFrom) + `,`,
		`Prototypes:` + fmt.Sprintf("%v", this.Prototypes) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPulseExporter(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetNextPulseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPulseExporter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNextPulseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNextPulseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PulseNumber", wireType)
			}
			m.PulseNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PulseNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevPulseNumber", wireType)
			}
			m.PrevPulseNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevPulseNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordAmount", wireType)
			}
			m.RecordAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordAmount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPulseExporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPulseExporter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPulseExporter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetNextPulseRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPulseExporter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNextPulseRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNextPulseRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PulseNumberFrom", wireType)
			}
			m.PulseNumberFrom = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PulseNumberFrom |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prototypes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPulseExporter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPulseExporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prototypes = append(m.Prototypes, make([]byte, postIndex-iNdEx))
			copy(m.Prototypes[len(m.Prototypes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPulseExporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPulseExporter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPulseExporter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPulseExporter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPulseExporter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPulseExporter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPulseExporter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPulseExporter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPulseExporter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPulseExporter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPulseExporter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPulseExporter = fmt.Errorf("proto: unexpected end of group")
)
