// Code generated by protoc-gen-gogo.
// source: client.proto
// DO NOT EDIT!

/*
	Package messages is a generated protocol buffer package.

	It is generated from these files:
		client.proto
		server.proto

	It has these top-level messages:
		Client
		FetchFile
		WriteToFile
		ReadFromFile
		SeekIndex
		Impeach
*/
package messages

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Client_MessageType int32

const (
	Client_FetchFile    Client_MessageType = 1
	Client_WriteToFile  Client_MessageType = 2
	Client_ReadFromFile Client_MessageType = 3
	Client_SeekIndex    Client_MessageType = 4
	Client_Impeach      Client_MessageType = 5
)

var Client_MessageType_name = map[int32]string{
	1: "FetchFile",
	2: "WriteToFile",
	3: "ReadFromFile",
	4: "SeekIndex",
	5: "Impeach",
}
var Client_MessageType_value = map[string]int32{
	"FetchFile":    1,
	"WriteToFile":  2,
	"ReadFromFile": 3,
	"SeekIndex":    4,
	"Impeach":      5,
}

func (x Client_MessageType) Enum() *Client_MessageType {
	p := new(Client_MessageType)
	*p = x
	return p
}
func (x Client_MessageType) String() string {
	return proto.EnumName(Client_MessageType_name, int32(x))
}
func (x *Client_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Client_MessageType_value, data, "Client_MessageType")
	if err != nil {
		return err
	}
	*x = Client_MessageType(value)
	return nil
}

type Client struct {
	MessageId        *uint64             `protobuf:"varint,1,req,name=messageId" json:"messageId,omitempty"`
	MessageType      *Client_MessageType `protobuf:"varint,2,req,name=messageType,enum=messages.Client_MessageType" json:"messageType,omitempty"`
	FetchFile        *FetchFile          `protobuf:"bytes,3,opt,name=fetchFile" json:"fetchFile,omitempty"`
	WriteToFile      *WriteToFile        `protobuf:"bytes,4,opt,name=writeToFile" json:"writeToFile,omitempty"`
	ReadFromFile     *ReadFromFile       `protobuf:"bytes,5,opt,name=readFromFile" json:"readFromFile,omitempty"`
	SeekIndex        *SeekIndex          `protobuf:"bytes,6,opt,name=seekIndex" json:"seekIndex,omitempty"`
	Impeach          *Impeach            `protobuf:"bytes,7,opt,name=impeach" json:"impeach,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}

func (m *Client) GetMessageId() uint64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *Client) GetMessageType() Client_MessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return Client_FetchFile
}

func (m *Client) GetFetchFile() *FetchFile {
	if m != nil {
		return m.FetchFile
	}
	return nil
}

func (m *Client) GetWriteToFile() *WriteToFile {
	if m != nil {
		return m.WriteToFile
	}
	return nil
}

func (m *Client) GetReadFromFile() *ReadFromFile {
	if m != nil {
		return m.ReadFromFile
	}
	return nil
}

func (m *Client) GetSeekIndex() *SeekIndex {
	if m != nil {
		return m.SeekIndex
	}
	return nil
}

func (m *Client) GetImpeach() *Impeach {
	if m != nil {
		return m.Impeach
	}
	return nil
}

type FetchFile struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	FileId           *uint64 `protobuf:"varint,2,req,name=fileId" json:"fileId,omitempty"`
	Create           *bool   `protobuf:"varint,3,req,name=create" json:"create,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FetchFile) Reset()         { *m = FetchFile{} }
func (m *FetchFile) String() string { return proto.CompactTextString(m) }
func (*FetchFile) ProtoMessage()    {}

func (m *FetchFile) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FetchFile) GetFileId() uint64 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *FetchFile) GetCreate() bool {
	if m != nil && m.Create != nil {
		return *m.Create
	}
	return false
}

type WriteToFile struct {
	FileId           *uint64 `protobuf:"varint,1,req,name=fileId" json:"fileId,omitempty"`
	Data             []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WriteToFile) Reset()         { *m = WriteToFile{} }
func (m *WriteToFile) String() string { return proto.CompactTextString(m) }
func (*WriteToFile) ProtoMessage()    {}

func (m *WriteToFile) GetFileId() uint64 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *WriteToFile) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ReadFromFile struct {
	FileId           *uint64 `protobuf:"varint,1,req,name=fileId" json:"fileId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReadFromFile) Reset()         { *m = ReadFromFile{} }
func (m *ReadFromFile) String() string { return proto.CompactTextString(m) }
func (*ReadFromFile) ProtoMessage()    {}

func (m *ReadFromFile) GetFileId() uint64 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

type SeekIndex struct {
	FileId           *uint64 `protobuf:"varint,1,req,name=fileId" json:"fileId,omitempty"`
	Index            *uint64 `protobuf:"varint,2,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SeekIndex) Reset()         { *m = SeekIndex{} }
func (m *SeekIndex) String() string { return proto.CompactTextString(m) }
func (*SeekIndex) ProtoMessage()    {}

func (m *SeekIndex) GetFileId() uint64 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *SeekIndex) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type Impeach struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Impeach) Reset()         { *m = Impeach{} }
func (m *Impeach) String() string { return proto.CompactTextString(m) }
func (*Impeach) ProtoMessage()    {}

func (m *Impeach) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("messages.Client_MessageType", Client_MessageType_name, Client_MessageType_value)
}
func (m *Client) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MessageId = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageType", wireType)
			}
			var v Client_MessageType
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (Client_MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MessageType = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FetchFile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FetchFile == nil {
				m.FetchFile = &FetchFile{}
			}
			if err := m.FetchFile.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WriteToFile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WriteToFile == nil {
				m.WriteToFile = &WriteToFile{}
			}
			if err := m.WriteToFile.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadFromFile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReadFromFile == nil {
				m.ReadFromFile = &ReadFromFile{}
			}
			if err := m.ReadFromFile.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeekIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SeekIndex == nil {
				m.SeekIndex = &SeekIndex{}
			}
			if err := m.SeekIndex.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Impeach", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Impeach == nil {
				m.Impeach = &Impeach{}
			}
			if err := m.Impeach.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *FetchFile) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Name = &s
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileId = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Create", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Create = &b
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *WriteToFile) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileId = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append([]byte{}, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ReadFromFile) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileId = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *SeekIndex) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileId = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Index = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Impeach) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Name = &s
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Client) Size() (n int) {
	var l int
	_ = l
	if m.MessageId != nil {
		n += 1 + sovClient(uint64(*m.MessageId))
	}
	if m.MessageType != nil {
		n += 1 + sovClient(uint64(*m.MessageType))
	}
	if m.FetchFile != nil {
		l = m.FetchFile.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	if m.WriteToFile != nil {
		l = m.WriteToFile.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	if m.ReadFromFile != nil {
		l = m.ReadFromFile.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	if m.SeekIndex != nil {
		l = m.SeekIndex.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	if m.Impeach != nil {
		l = m.Impeach.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FetchFile) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovClient(uint64(l))
	}
	if m.FileId != nil {
		n += 1 + sovClient(uint64(*m.FileId))
	}
	if m.Create != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WriteToFile) Size() (n int) {
	var l int
	_ = l
	if m.FileId != nil {
		n += 1 + sovClient(uint64(*m.FileId))
	}
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovClient(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReadFromFile) Size() (n int) {
	var l int
	_ = l
	if m.FileId != nil {
		n += 1 + sovClient(uint64(*m.FileId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SeekIndex) Size() (n int) {
	var l int
	_ = l
	if m.FileId != nil {
		n += 1 + sovClient(uint64(*m.FileId))
	}
	if m.Index != nil {
		n += 1 + sovClient(uint64(*m.Index))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Impeach) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovClient(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Client) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Client) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MessageId != nil {
		data[i] = 0x8
		i++
		i = encodeVarintClient(data, i, uint64(*m.MessageId))
	}
	if m.MessageType != nil {
		data[i] = 0x10
		i++
		i = encodeVarintClient(data, i, uint64(*m.MessageType))
	}
	if m.FetchFile != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintClient(data, i, uint64(m.FetchFile.Size()))
		n1, err := m.FetchFile.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.WriteToFile != nil {
		data[i] = 0x22
		i++
		i = encodeVarintClient(data, i, uint64(m.WriteToFile.Size()))
		n2, err := m.WriteToFile.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ReadFromFile != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintClient(data, i, uint64(m.ReadFromFile.Size()))
		n3, err := m.ReadFromFile.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.SeekIndex != nil {
		data[i] = 0x32
		i++
		i = encodeVarintClient(data, i, uint64(m.SeekIndex.Size()))
		n4, err := m.SeekIndex.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Impeach != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintClient(data, i, uint64(m.Impeach.Size()))
		n5, err := m.Impeach.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FetchFile) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *FetchFile) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintClient(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.FileId != nil {
		data[i] = 0x10
		i++
		i = encodeVarintClient(data, i, uint64(*m.FileId))
	}
	if m.Create != nil {
		data[i] = 0x18
		i++
		if *m.Create {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WriteToFile) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WriteToFile) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FileId != nil {
		data[i] = 0x8
		i++
		i = encodeVarintClient(data, i, uint64(*m.FileId))
	}
	if m.Data != nil {
		data[i] = 0x12
		i++
		i = encodeVarintClient(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ReadFromFile) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReadFromFile) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FileId != nil {
		data[i] = 0x8
		i++
		i = encodeVarintClient(data, i, uint64(*m.FileId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SeekIndex) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SeekIndex) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FileId != nil {
		data[i] = 0x8
		i++
		i = encodeVarintClient(data, i, uint64(*m.FileId))
	}
	if m.Index != nil {
		data[i] = 0x10
		i++
		i = encodeVarintClient(data, i, uint64(*m.Index))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Impeach) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Impeach) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintClient(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Client(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Client(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintClient(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
