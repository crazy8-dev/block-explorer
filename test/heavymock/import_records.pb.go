// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/heavymock/import_records.proto

package heavymock

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	exporter "github.com/insolar/insolar/ledger/heavy/exporter"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Ok struct {
	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (m *Ok) Reset()      { *m = Ok{} }
func (*Ok) ProtoMessage() {}
func (*Ok) Descriptor() ([]byte, []int) {
	return fileDescriptor_c565bd77abcb08b8, []int{0}
}
func (m *Ok) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ok) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ok.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ok) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ok.Merge(m, src)
}
func (m *Ok) XXX_Size() int {
	return m.Size()
}
func (m *Ok) XXX_DiscardUnknown() {
	xxx_messageInfo_Ok.DiscardUnknown(m)
}

var xxx_messageInfo_Ok proto.InternalMessageInfo

func (m *Ok) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*Ok)(nil), "heavymock.Ok")
}

func init() {
	proto.RegisterFile("test/heavymock/import_records.proto", fileDescriptor_c565bd77abcb08b8)
}

var fileDescriptor_c565bd77abcb08b8 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x49, 0x2d, 0x2e,
	0xd1, 0xcf, 0x48, 0x4d, 0x2c, 0xab, 0xcc, 0xcd, 0x4f, 0xce, 0xd6, 0xcf, 0xcc, 0x2d, 0xc8, 0x2f,
	0x2a, 0x89, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x84, 0xcb, 0x4b, 0xb9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x67, 0xe6, 0x15, 0xe7, 0xe7, 0x24, 0x16, 0xc1, 0xe9, 0x9c, 0xd4, 0x94, 0xf4, 0xd4, 0x22, 0x88,
	0x61, 0xfa, 0xa9, 0x15, 0x20, 0x83, 0x52, 0x8b, 0xf4, 0x21, 0x46, 0xc5, 0xc3, 0xf8, 0x10, 0x23,
	0x95, 0x44, 0xb8, 0x98, 0xfc, 0xb3, 0x85, 0xf8, 0xb8, 0x98, 0xf2, 0xb3, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x82, 0x98, 0xf2, 0xb3, 0x8d, 0x9c, 0xb8, 0x04, 0x3d, 0x60, 0x56, 0x79, 0xe6, 0x42,
	0x34, 0x08, 0xe9, 0x72, 0xb1, 0x41, 0xd8, 0x42, 0x02, 0x7a, 0x70, 0x53, 0x82, 0xc0, 0xa6, 0x4a,
	0xf1, 0xea, 0xc1, 0x9d, 0xa6, 0xe7, 0x9f, 0xad, 0xc4, 0xa0, 0xc1, 0xe8, 0x64, 0x72, 0xe1, 0xa1,
	0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1,
	0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c,
	0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0x32, 0x83, 0x69, 0x10, 0x01, 0x00, 0x00,
}

func (this *Ok) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ok)
	if !ok {
		that2, ok := that.(Ok)
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
	if this.Ok != that1.Ok {
		return false
	}
	return true
}
func (this *Ok) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&heavymock.Ok{")
	s = append(s, "Ok: "+fmt.Sprintf("%#v", this.Ok)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringImportRecords(v interface{}, typ string) string {
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

// HeavymockImporterClient is the client API for HeavymockImporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeavymockImporterClient interface {
	Import(ctx context.Context, opts ...grpc.CallOption) (HeavymockImporter_ImportClient, error)
}

type heavymockImporterClient struct {
	cc *grpc.ClientConn
}

func NewHeavymockImporterClient(cc *grpc.ClientConn) HeavymockImporterClient {
	return &heavymockImporterClient{cc}
}

func (c *heavymockImporterClient) Import(ctx context.Context, opts ...grpc.CallOption) (HeavymockImporter_ImportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HeavymockImporter_serviceDesc.Streams[0], "/heavymock.HeavymockImporter/Import", opts...)
	if err != nil {
		return nil, err
	}
	x := &heavymockImporterImportClient{stream}
	return x, nil
}

type HeavymockImporter_ImportClient interface {
	Send(*exporter.Record) error
	CloseAndRecv() (*Ok, error)
	grpc.ClientStream
}

type heavymockImporterImportClient struct {
	grpc.ClientStream
}

func (x *heavymockImporterImportClient) Send(m *exporter.Record) error {
	return x.ClientStream.SendMsg(m)
}

func (x *heavymockImporterImportClient) CloseAndRecv() (*Ok, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Ok)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeavymockImporterServer is the server API for HeavymockImporter service.
type HeavymockImporterServer interface {
	Import(HeavymockImporter_ImportServer) error
}

func RegisterHeavymockImporterServer(s *grpc.Server, srv HeavymockImporterServer) {
	s.RegisterService(&_HeavymockImporter_serviceDesc, srv)
}

func _HeavymockImporter_Import_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HeavymockImporterServer).Import(&heavymockImporterImportServer{stream})
}

type HeavymockImporter_ImportServer interface {
	SendAndClose(*Ok) error
	Recv() (*exporter.Record, error)
	grpc.ServerStream
}

type heavymockImporterImportServer struct {
	grpc.ServerStream
}

func (x *heavymockImporterImportServer) SendAndClose(m *Ok) error {
	return x.ServerStream.SendMsg(m)
}

func (x *heavymockImporterImportServer) Recv() (*exporter.Record, error) {
	m := new(exporter.Record)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HeavymockImporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heavymock.HeavymockImporter",
	HandlerType: (*HeavymockImporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Import",
			Handler:       _HeavymockImporter_Import_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test/heavymock/import_records.proto",
}

func (m *Ok) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ok) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ok {
		dAtA[i] = 0x8
		i++
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintImportRecords(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Ok) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	return n
}

func sovImportRecords(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImportRecords(x uint64) (n int) {
	return sovImportRecords(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Ok) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Ok{`,
		`Ok:` + fmt.Sprintf("%v", this.Ok) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringImportRecords(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Ok) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImportRecords
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
			return fmt.Errorf("proto: Ok: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ok: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImportRecords
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipImportRecords(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImportRecords
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthImportRecords
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
func skipImportRecords(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImportRecords
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
					return 0, ErrIntOverflowImportRecords
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowImportRecords
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
				return 0, ErrInvalidLengthImportRecords
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthImportRecords
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImportRecords
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipImportRecords(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthImportRecords
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthImportRecords = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImportRecords   = fmt.Errorf("proto: integer overflow")
)