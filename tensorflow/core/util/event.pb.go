// Code generated by protoc-gen-go.
// source: tensorflow/core/util/event.proto
// DO NOT EDIT!

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/util/event.proto

It has these top-level messages:
	Event
	LogMessage
	SessionLog
	TaggedRunMetadata
*/
package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow5 "github.com/helinwang/tfsum/tensorflow/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogMessage_Level int32

const (
	LogMessage_UNKNOWN LogMessage_Level = 0
	LogMessage_DEBUG   LogMessage_Level = 10
	LogMessage_INFO    LogMessage_Level = 20
	LogMessage_WARN    LogMessage_Level = 30
	LogMessage_ERROR   LogMessage_Level = 40
	LogMessage_FATAL   LogMessage_Level = 50
)

var LogMessage_Level_name = map[int32]string{
	0:  "UNKNOWN",
	10: "DEBUG",
	20: "INFO",
	30: "WARN",
	40: "ERROR",
	50: "FATAL",
}
var LogMessage_Level_value = map[string]int32{
	"UNKNOWN": 0,
	"DEBUG":   10,
	"INFO":    20,
	"WARN":    30,
	"ERROR":   40,
	"FATAL":   50,
}

func (x LogMessage_Level) String() string {
	return proto.EnumName(LogMessage_Level_name, int32(x))
}
func (LogMessage_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type SessionLog_SessionStatus int32

const (
	SessionLog_STATUS_UNSPECIFIED SessionLog_SessionStatus = 0
	SessionLog_START              SessionLog_SessionStatus = 1
	SessionLog_STOP               SessionLog_SessionStatus = 2
	SessionLog_CHECKPOINT         SessionLog_SessionStatus = 3
)

var SessionLog_SessionStatus_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "START",
	2: "STOP",
	3: "CHECKPOINT",
}
var SessionLog_SessionStatus_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"START":              1,
	"STOP":               2,
	"CHECKPOINT":         3,
}

func (x SessionLog_SessionStatus) String() string {
	return proto.EnumName(SessionLog_SessionStatus_name, int32(x))
}
func (SessionLog_SessionStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// Protocol buffer representing an event that happened during
// the execution of a Brain model.
type Event struct {
	// Timestamp of the event.
	WallTime float64 `protobuf:"fixed64,1,opt,name=wall_time,json=wallTime" json:"wall_time,omitempty"`
	// Global step of the event.
	Step int64 `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	// Types that are valid to be assigned to What:
	//	*Event_FileVersion
	//	*Event_GraphDef
	//	*Event_Summary
	//	*Event_LogMessage
	//	*Event_SessionLog
	//	*Event_TaggedRunMetadata
	//	*Event_MetaGraphDef
	What isEvent_What `protobuf_oneof:"what"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isEvent_What interface {
	isEvent_What()
}

type Event_FileVersion struct {
	FileVersion string `protobuf:"bytes,3,opt,name=file_version,json=fileVersion,oneof"`
}
type Event_GraphDef struct {
	GraphDef []byte `protobuf:"bytes,4,opt,name=graph_def,json=graphDef,proto3,oneof"`
}
type Event_Summary struct {
	Summary *tensorflow5.Summary `protobuf:"bytes,5,opt,name=summary,oneof"`
}
type Event_LogMessage struct {
	LogMessage *LogMessage `protobuf:"bytes,6,opt,name=log_message,json=logMessage,oneof"`
}
type Event_SessionLog struct {
	SessionLog *SessionLog `protobuf:"bytes,7,opt,name=session_log,json=sessionLog,oneof"`
}
type Event_TaggedRunMetadata struct {
	TaggedRunMetadata *TaggedRunMetadata `protobuf:"bytes,8,opt,name=tagged_run_metadata,json=taggedRunMetadata,oneof"`
}
type Event_MetaGraphDef struct {
	MetaGraphDef []byte `protobuf:"bytes,9,opt,name=meta_graph_def,json=metaGraphDef,proto3,oneof"`
}

func (*Event_FileVersion) isEvent_What()       {}
func (*Event_GraphDef) isEvent_What()          {}
func (*Event_Summary) isEvent_What()           {}
func (*Event_LogMessage) isEvent_What()        {}
func (*Event_SessionLog) isEvent_What()        {}
func (*Event_TaggedRunMetadata) isEvent_What() {}
func (*Event_MetaGraphDef) isEvent_What()      {}

func (m *Event) GetWhat() isEvent_What {
	if m != nil {
		return m.What
	}
	return nil
}

func (m *Event) GetWallTime() float64 {
	if m != nil {
		return m.WallTime
	}
	return 0
}

func (m *Event) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *Event) GetFileVersion() string {
	if x, ok := m.GetWhat().(*Event_FileVersion); ok {
		return x.FileVersion
	}
	return ""
}

func (m *Event) GetGraphDef() []byte {
	if x, ok := m.GetWhat().(*Event_GraphDef); ok {
		return x.GraphDef
	}
	return nil
}

func (m *Event) GetSummary() *tensorflow5.Summary {
	if x, ok := m.GetWhat().(*Event_Summary); ok {
		return x.Summary
	}
	return nil
}

func (m *Event) GetLogMessage() *LogMessage {
	if x, ok := m.GetWhat().(*Event_LogMessage); ok {
		return x.LogMessage
	}
	return nil
}

func (m *Event) GetSessionLog() *SessionLog {
	if x, ok := m.GetWhat().(*Event_SessionLog); ok {
		return x.SessionLog
	}
	return nil
}

func (m *Event) GetTaggedRunMetadata() *TaggedRunMetadata {
	if x, ok := m.GetWhat().(*Event_TaggedRunMetadata); ok {
		return x.TaggedRunMetadata
	}
	return nil
}

func (m *Event) GetMetaGraphDef() []byte {
	if x, ok := m.GetWhat().(*Event_MetaGraphDef); ok {
		return x.MetaGraphDef
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_FileVersion)(nil),
		(*Event_GraphDef)(nil),
		(*Event_Summary)(nil),
		(*Event_LogMessage)(nil),
		(*Event_SessionLog)(nil),
		(*Event_TaggedRunMetadata)(nil),
		(*Event_MetaGraphDef)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// what
	switch x := m.What.(type) {
	case *Event_FileVersion:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FileVersion)
	case *Event_GraphDef:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.GraphDef)
	case *Event_Summary:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Summary); err != nil {
			return err
		}
	case *Event_LogMessage:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogMessage); err != nil {
			return err
		}
	case *Event_SessionLog:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SessionLog); err != nil {
			return err
		}
	case *Event_TaggedRunMetadata:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TaggedRunMetadata); err != nil {
			return err
		}
	case *Event_MetaGraphDef:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.MetaGraphDef)
	case nil:
	default:
		return fmt.Errorf("Event.What has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 3: // what.file_version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.What = &Event_FileVersion{x}
		return true, err
	case 4: // what.graph_def
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.What = &Event_GraphDef{x}
		return true, err
	case 5: // what.summary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(tensorflow5.Summary)
		err := b.DecodeMessage(msg)
		m.What = &Event_Summary{msg}
		return true, err
	case 6: // what.log_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogMessage)
		err := b.DecodeMessage(msg)
		m.What = &Event_LogMessage{msg}
		return true, err
	case 7: // what.session_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SessionLog)
		err := b.DecodeMessage(msg)
		m.What = &Event_SessionLog{msg}
		return true, err
	case 8: // what.tagged_run_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaggedRunMetadata)
		err := b.DecodeMessage(msg)
		m.What = &Event_TaggedRunMetadata{msg}
		return true, err
	case 9: // what.meta_graph_def
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.What = &Event_MetaGraphDef{x}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// what
	switch x := m.What.(type) {
	case *Event_FileVersion:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FileVersion)))
		n += len(x.FileVersion)
	case *Event_GraphDef:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.GraphDef)))
		n += len(x.GraphDef)
	case *Event_Summary:
		s := proto.Size(x.Summary)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_LogMessage:
		s := proto.Size(x.LogMessage)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_SessionLog:
		s := proto.Size(x.SessionLog)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_TaggedRunMetadata:
		s := proto.Size(x.TaggedRunMetadata)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_MetaGraphDef:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.MetaGraphDef)))
		n += len(x.MetaGraphDef)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Protocol buffer used for logging messages to the events file.
type LogMessage struct {
	Level   LogMessage_Level `protobuf:"varint,1,opt,name=level,enum=tensorflow.LogMessage_Level" json:"level,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *LogMessage) Reset()                    { *m = LogMessage{} }
func (m *LogMessage) String() string            { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()               {}
func (*LogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogMessage) GetLevel() LogMessage_Level {
	if m != nil {
		return m.Level
	}
	return LogMessage_UNKNOWN
}

func (m *LogMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Protocol buffer used for logging session state.
type SessionLog struct {
	Status SessionLog_SessionStatus `protobuf:"varint,1,opt,name=status,enum=tensorflow.SessionLog_SessionStatus" json:"status,omitempty"`
	// This checkpoint_path contains both the path and filename.
	CheckpointPath string `protobuf:"bytes,2,opt,name=checkpoint_path,json=checkpointPath" json:"checkpoint_path,omitempty"`
	Msg            string `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *SessionLog) Reset()                    { *m = SessionLog{} }
func (m *SessionLog) String() string            { return proto.CompactTextString(m) }
func (*SessionLog) ProtoMessage()               {}
func (*SessionLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SessionLog) GetStatus() SessionLog_SessionStatus {
	if m != nil {
		return m.Status
	}
	return SessionLog_STATUS_UNSPECIFIED
}

func (m *SessionLog) GetCheckpointPath() string {
	if m != nil {
		return m.CheckpointPath
	}
	return ""
}

func (m *SessionLog) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// For logging the metadata output for a single session.run() call.
type TaggedRunMetadata struct {
	// Tag name associated with this metadata.
	Tag string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	// Byte-encoded version of the `RunMetadata` proto in order to allow lazy
	// deserialization.
	RunMetadata []byte `protobuf:"bytes,2,opt,name=run_metadata,json=runMetadata,proto3" json:"run_metadata,omitempty"`
}

func (m *TaggedRunMetadata) Reset()                    { *m = TaggedRunMetadata{} }
func (m *TaggedRunMetadata) String() string            { return proto.CompactTextString(m) }
func (*TaggedRunMetadata) ProtoMessage()               {}
func (*TaggedRunMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TaggedRunMetadata) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *TaggedRunMetadata) GetRunMetadata() []byte {
	if m != nil {
		return m.RunMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "tensorflow.Event")
	proto.RegisterType((*LogMessage)(nil), "tensorflow.LogMessage")
	proto.RegisterType((*SessionLog)(nil), "tensorflow.SessionLog")
	proto.RegisterType((*TaggedRunMetadata)(nil), "tensorflow.TaggedRunMetadata")
	proto.RegisterEnum("tensorflow.LogMessage_Level", LogMessage_Level_name, LogMessage_Level_value)
	proto.RegisterEnum("tensorflow.SessionLog_SessionStatus", SessionLog_SessionStatus_name, SessionLog_SessionStatus_value)
}

func init() { proto.RegisterFile("tensorflow/core/util/event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x65, 0xc3, 0xa7, 0x07, 0x4a, 0x9d, 0x4d, 0x15, 0x59, 0x6d, 0x53, 0x51, 0x5a, 0x35, 0x9c,
	0x40, 0xa2, 0xa7, 0x4a, 0xbd, 0x40, 0x42, 0x62, 0x14, 0x62, 0xac, 0xb5, 0x69, 0x8e, 0xd6, 0x36,
	0x59, 0x8c, 0x15, 0xdb, 0x8b, 0xbc, 0x4b, 0x50, 0xff, 0x4f, 0xfb, 0x8b, 0xfa, 0x67, 0x7a, 0xac,
	0xd6, 0x36, 0x71, 0xbe, 0x6e, 0xb3, 0x6f, 0xde, 0x1b, 0xcf, 0xbc, 0xf1, 0x40, 0x47, 0xb2, 0x58,
	0xf0, 0x64, 0x19, 0xf2, 0xed, 0xe0, 0x9a, 0x27, 0x6c, 0xb0, 0x91, 0x41, 0x38, 0x60, 0x77, 0x2c,
	0x96, 0xfd, 0x75, 0xc2, 0x25, 0xc7, 0x50, 0x30, 0xde, 0x1e, 0x3f, 0x65, 0x2f, 0x13, 0x1a, 0xb1,
	0x2d, 0x4f, 0x6e, 0x07, 0x62, 0x13, 0x45, 0x34, 0xf9, 0x95, 0x89, 0xba, 0xbf, 0xcb, 0x50, 0x9d,
	0xa8, 0x22, 0xf8, 0x1d, 0x68, 0x5b, 0x1a, 0x86, 0x9e, 0x0c, 0x22, 0x66, 0xa0, 0x0e, 0xea, 0x21,
	0xd2, 0x50, 0x80, 0x1b, 0x44, 0x0c, 0x63, 0xa8, 0x08, 0xc9, 0xd6, 0xc6, 0x5e, 0x07, 0xf5, 0xca,
	0x24, 0x8d, 0xf1, 0x27, 0x68, 0x2d, 0x83, 0x90, 0x79, 0x77, 0x2c, 0x11, 0x01, 0x8f, 0x8d, 0x72,
	0x07, 0xf5, 0x34, 0xb3, 0x44, 0x9a, 0x0a, 0xfd, 0x91, 0x81, 0xf8, 0x08, 0x34, 0x3f, 0xa1, 0xeb,
	0x95, 0x77, 0xc3, 0x96, 0x46, 0xa5, 0x83, 0x7a, 0x2d, 0xb3, 0x44, 0x1a, 0x29, 0x74, 0xca, 0x96,
	0x78, 0x00, 0xf5, 0xbc, 0x1f, 0xa3, 0xda, 0x41, 0xbd, 0xe6, 0xf0, 0xa0, 0x5f, 0x74, 0xde, 0x77,
	0xb2, 0x94, 0x59, 0x22, 0x3b, 0x16, 0xfe, 0x06, 0xcd, 0x90, 0xfb, 0x5e, 0xc4, 0x84, 0xa0, 0x3e,
	0x33, 0x6a, 0xa9, 0xe8, 0xf0, 0xa1, 0x68, 0xc6, 0xfd, 0xcb, 0x2c, 0x6b, 0x96, 0x08, 0x84, 0xf7,
	0x2f, 0x25, 0x15, 0x4c, 0xa8, 0xae, 0xbc, 0x90, 0xfb, 0x46, 0xfd, 0xb9, 0xd4, 0xc9, 0xd2, 0x33,
	0xee, 0x2b, 0xa9, 0xb8, 0x7f, 0xe1, 0x39, 0x1c, 0x48, 0xea, 0xfb, 0xec, 0xc6, 0x4b, 0x36, 0xb1,
	0x17, 0x31, 0x49, 0x6f, 0xa8, 0xa4, 0x46, 0x23, 0x2d, 0x71, 0xf4, 0xb0, 0x84, 0x9b, 0xd2, 0xc8,
	0x26, 0xbe, 0xcc, 0x49, 0x66, 0x89, 0xec, 0xcb, 0xa7, 0x20, 0xfe, 0x02, 0x6d, 0x55, 0xc5, 0x2b,
	0xbc, 0xd1, 0x72, 0x6f, 0x5a, 0x0a, 0x3f, 0xcf, 0xfd, 0x19, 0xd7, 0xa0, 0xb2, 0x5d, 0x51, 0xd9,
	0xfd, 0x83, 0x00, 0x8a, 0xc1, 0xf0, 0x10, 0xaa, 0x21, 0xbb, 0x63, 0x61, 0xba, 0xa7, 0xf6, 0xf0,
	0xfd, 0xcb, 0xf3, 0xf7, 0x67, 0x8a, 0x43, 0x32, 0x2a, 0x36, 0xa0, 0xbe, 0x73, 0x4d, 0x6d, 0x51,
	0x23, 0xbb, 0x67, 0x77, 0x0a, 0xd5, 0x94, 0x89, 0x9b, 0x50, 0x5f, 0x58, 0x17, 0xd6, 0xfc, 0xca,
	0xd2, 0x4b, 0x58, 0x83, 0xea, 0xe9, 0x64, 0xbc, 0x38, 0xd7, 0x01, 0x37, 0xa0, 0x32, 0xb5, 0xce,
	0xe6, 0xfa, 0x1b, 0x15, 0x5d, 0x8d, 0x88, 0xa5, 0x7f, 0x50, 0xe9, 0x09, 0x21, 0x73, 0xa2, 0xf7,
	0x54, 0x78, 0x36, 0x72, 0x47, 0x33, 0x7d, 0xd8, 0xfd, 0x8b, 0x00, 0x0a, 0x17, 0xf1, 0x77, 0xa8,
	0x09, 0x49, 0xe5, 0x46, 0xe4, 0x8d, 0x7e, 0x7e, 0xd9, 0xed, 0x5d, 0xe8, 0xa4, 0x5c, 0x92, 0x6b,
	0xf0, 0x31, 0xbc, 0xbe, 0x5e, 0xb1, 0xeb, 0xdb, 0x35, 0x0f, 0x62, 0xe9, 0xad, 0xa9, 0x5c, 0xe5,
	0x9d, 0xb7, 0x0b, 0xd8, 0xa6, 0x72, 0x85, 0x75, 0x28, 0x47, 0xc2, 0xcf, 0x7e, 0x40, 0xa2, 0xc2,
	0xee, 0x0c, 0x5e, 0x3d, 0xaa, 0x89, 0x0f, 0x01, 0x3b, 0xee, 0xc8, 0x5d, 0x38, 0xde, 0xc2, 0x72,
	0xec, 0xc9, 0xc9, 0xf4, 0x6c, 0x3a, 0x39, 0xcd, 0xa6, 0x74, 0xdc, 0x11, 0x71, 0x75, 0xa4, 0x66,
	0x73, 0xdc, 0xb9, 0xad, 0xef, 0xe1, 0x36, 0xc0, 0x89, 0x39, 0x39, 0xb9, 0xb0, 0xe7, 0x53, 0xcb,
	0xd5, 0xcb, 0x5d, 0x13, 0xf6, 0x9f, 0xed, 0x55, 0x7d, 0x54, 0x52, 0x3f, 0x1d, 0x4c, 0x23, 0x2a,
	0xc4, 0x1f, 0xa1, 0xf5, 0xe8, 0xf7, 0x50, 0xcd, 0xb6, 0x48, 0x33, 0x29, 0x44, 0xe3, 0x63, 0x38,
	0xe0, 0x89, 0xff, 0xd0, 0x05, 0x75, 0xc6, 0xe3, 0x66, 0x7a, 0x82, 0xb6, 0xba, 0x48, 0x61, 0xa3,
	0x7f, 0x08, 0xfd, 0xac, 0xa5, 0xe7, 0xf9, 0xf5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x78,
	0xba, 0x51, 0xf7, 0x03, 0x00, 0x00,
}
