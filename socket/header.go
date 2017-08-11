// Copyright 2017 HenryLee. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-gogo.
// source: header.proto
// DO NOT EDIT!

/*
	Package socket is a generated protocol buffer package.

	It is generated from these files:
		header.proto

	It has these top-level messages:
		Header
*/

package socket

import (
	"fmt"
	"io"

	"github.com/gogo/protobuf/proto"
)

// Header packet header
type Header struct {
	// Packet id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Service type
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// As reply, it indicates the service status code
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// As reply, it indicates the service status text
	Codec string `protobuf:"bytes,4,opt,name=codec,proto3" json:"codec,omitempty"`
	// Service URI
	Gzip int32 `protobuf:"varint,5,opt,name=gzip,proto3" json:"gzip,omitempty"`
	// Body encoding type
	StatusCode int32 `protobuf:"varint,6,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// Body gizp compression level
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptorHeader, []int{0} }

func (m *Header) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Header) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Header) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Header) GetCodec() string {
	if m != nil {
		return m.Codec
	}
	return ""
}

func (m *Header) GetGzip() int32 {
	if m != nil {
		return m.Gzip
	}
	return 0
}

func (m *Header) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Header) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Header)(nil), "socket.header")
}
func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.Type))
	}
	if len(m.Uri) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Uri)))
		i += copy(dAtA[i:], m.Uri)
	}
	if len(m.Codec) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Codec)))
		i += copy(dAtA[i:], m.Codec)
	}
	if m.Gzip != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.Gzip))
	}
	if m.StatusCode != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.StatusCode))
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	return i, nil
}

func encodeFixed64Header(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Header(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHeader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Header) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovHeader(uint64(m.Type))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	l = len(m.Codec)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	if m.Gzip != 0 {
		n += 1 + sovHeader(uint64(m.Gzip))
	}
	if m.StatusCode != 0 {
		n += 1 + sovHeader(uint64(m.StatusCode))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	return n
}

func sovHeader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeader(x uint64) (n int) {
	return sovHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeader
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gzip", wireType)
			}
			m.Gzip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gzip |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			m.StatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeader
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
func skipHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHeader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeader
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
				next, err := skipHeader(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthHeader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeader   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("header.proto", fileDescriptorHeader) }

var fileDescriptorHeader = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0xce, 0x4f, 0xce, 0x4e, 0x2d,
	0x51, 0x5a, 0xc8, 0xc8, 0xc5, 0x06, 0x91, 0x10, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x4b, 0x8b, 0x32, 0x25,
	0x98, 0xc1, 0x8a, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xe4, 0xfc, 0x94, 0xd4, 0x64, 0x09, 0x16,
	0xb0, 0x18, 0x84, 0x03, 0xd2, 0x9b, 0x5e, 0x95, 0x59, 0x20, 0xc1, 0x0a, 0xd1, 0x0b, 0x62, 0x0b,
	0xc9, 0x73, 0x71, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16, 0xc7, 0x83, 0xd4, 0x48, 0xb0, 0x81, 0xa5,
	0xb8, 0x20, 0x42, 0xce, 0xf9, 0x29, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x10, 0x9e, 0x04, 0x3b, 0xd8,
	0x2c, 0x28, 0xcf, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xaf, 0x9c, 0x39, 0xe2, 0xd4, 0x00, 0x00, 0x00,
}