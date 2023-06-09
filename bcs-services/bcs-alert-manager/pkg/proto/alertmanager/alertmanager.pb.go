// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/alertmanager/alertmanager.proto

package alertmanager

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateRawAlertInfoReq struct {
	Starttime            int64             `protobuf:"varint,1,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Endtime              int64             `protobuf:"varint,2,opt,name=endtime,proto3" json:"endtime,omitempty"`
	Generatorurl         string            `protobuf:"bytes,3,opt,name=generatorurl,proto3" json:"generatorurl,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateRawAlertInfoReq) Reset()         { *m = CreateRawAlertInfoReq{} }
func (m *CreateRawAlertInfoReq) String() string { return proto.CompactTextString(m) }
func (*CreateRawAlertInfoReq) ProtoMessage()    {}
func (*CreateRawAlertInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{0}
}

func (m *CreateRawAlertInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRawAlertInfoReq.Unmarshal(m, b)
}
func (m *CreateRawAlertInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRawAlertInfoReq.Marshal(b, m, deterministic)
}
func (m *CreateRawAlertInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRawAlertInfoReq.Merge(m, src)
}
func (m *CreateRawAlertInfoReq) XXX_Size() int {
	return xxx_messageInfo_CreateRawAlertInfoReq.Size(m)
}
func (m *CreateRawAlertInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRawAlertInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRawAlertInfoReq proto.InternalMessageInfo

func (m *CreateRawAlertInfoReq) GetStarttime() int64 {
	if m != nil {
		return m.Starttime
	}
	return 0
}

func (m *CreateRawAlertInfoReq) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

func (m *CreateRawAlertInfoReq) GetGeneratorurl() string {
	if m != nil {
		return m.Generatorurl
	}
	return ""
}

func (m *CreateRawAlertInfoReq) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *CreateRawAlertInfoReq) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type CreateRawAlertInfoResp struct {
	ErrCode              uint64   `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRawAlertInfoResp) Reset()         { *m = CreateRawAlertInfoResp{} }
func (m *CreateRawAlertInfoResp) String() string { return proto.CompactTextString(m) }
func (*CreateRawAlertInfoResp) ProtoMessage()    {}
func (*CreateRawAlertInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{1}
}

func (m *CreateRawAlertInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRawAlertInfoResp.Unmarshal(m, b)
}
func (m *CreateRawAlertInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRawAlertInfoResp.Marshal(b, m, deterministic)
}
func (m *CreateRawAlertInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRawAlertInfoResp.Merge(m, src)
}
func (m *CreateRawAlertInfoResp) XXX_Size() int {
	return xxx_messageInfo_CreateRawAlertInfoResp.Size(m)
}
func (m *CreateRawAlertInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRawAlertInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRawAlertInfoResp proto.InternalMessageInfo

func (m *CreateRawAlertInfoResp) GetErrCode() uint64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CreateRawAlertInfoResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type CreateBusinessAlertInfoReq struct {
	Starttime    int64  `protobuf:"varint,1,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Endtime      int64  `protobuf:"varint,2,opt,name=endtime,proto3" json:"endtime,omitempty"`
	Generatorurl string `protobuf:"bytes,3,opt,name=generatorurl,proto3" json:"generatorurl,omitempty"`
	AlarmType    string `protobuf:"bytes,4,opt,name=alarmType,proto3" json:"alarmType,omitempty"`
	ClusterID    string `protobuf:"bytes,5,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	// 报警annotation
	AlertAnnotation *AlertAnnotation `protobuf:"bytes,6,opt,name=alertAnnotation,proto3" json:"alertAnnotation,omitempty"`
	// 模块类型报警label
	ModuleAlertLabel *ModuleAlertLabel `protobuf:"bytes,7,opt,name=moduleAlertLabel,proto3" json:"moduleAlertLabel,omitempty"`
	// 资源类型报警label
	ResourceAlertLabel   *ResourceAlertLabel `protobuf:"bytes,8,opt,name=resourceAlertLabel,proto3" json:"resourceAlertLabel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateBusinessAlertInfoReq) Reset()         { *m = CreateBusinessAlertInfoReq{} }
func (m *CreateBusinessAlertInfoReq) String() string { return proto.CompactTextString(m) }
func (*CreateBusinessAlertInfoReq) ProtoMessage()    {}
func (*CreateBusinessAlertInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{2}
}

func (m *CreateBusinessAlertInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBusinessAlertInfoReq.Unmarshal(m, b)
}
func (m *CreateBusinessAlertInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBusinessAlertInfoReq.Marshal(b, m, deterministic)
}
func (m *CreateBusinessAlertInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBusinessAlertInfoReq.Merge(m, src)
}
func (m *CreateBusinessAlertInfoReq) XXX_Size() int {
	return xxx_messageInfo_CreateBusinessAlertInfoReq.Size(m)
}
func (m *CreateBusinessAlertInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBusinessAlertInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBusinessAlertInfoReq proto.InternalMessageInfo

func (m *CreateBusinessAlertInfoReq) GetStarttime() int64 {
	if m != nil {
		return m.Starttime
	}
	return 0
}

func (m *CreateBusinessAlertInfoReq) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

func (m *CreateBusinessAlertInfoReq) GetGeneratorurl() string {
	if m != nil {
		return m.Generatorurl
	}
	return ""
}

func (m *CreateBusinessAlertInfoReq) GetAlarmType() string {
	if m != nil {
		return m.AlarmType
	}
	return ""
}

func (m *CreateBusinessAlertInfoReq) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

func (m *CreateBusinessAlertInfoReq) GetAlertAnnotation() *AlertAnnotation {
	if m != nil {
		return m.AlertAnnotation
	}
	return nil
}

func (m *CreateBusinessAlertInfoReq) GetModuleAlertLabel() *ModuleAlertLabel {
	if m != nil {
		return m.ModuleAlertLabel
	}
	return nil
}

func (m *CreateBusinessAlertInfoReq) GetResourceAlertLabel() *ResourceAlertLabel {
	if m != nil {
		return m.ResourceAlertLabel
	}
	return nil
}

type CreateBusinessAlertInfoResp struct {
	ErrCode              uint64   `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBusinessAlertInfoResp) Reset()         { *m = CreateBusinessAlertInfoResp{} }
func (m *CreateBusinessAlertInfoResp) String() string { return proto.CompactTextString(m) }
func (*CreateBusinessAlertInfoResp) ProtoMessage()    {}
func (*CreateBusinessAlertInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{3}
}

func (m *CreateBusinessAlertInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBusinessAlertInfoResp.Unmarshal(m, b)
}
func (m *CreateBusinessAlertInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBusinessAlertInfoResp.Marshal(b, m, deterministic)
}
func (m *CreateBusinessAlertInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBusinessAlertInfoResp.Merge(m, src)
}
func (m *CreateBusinessAlertInfoResp) XXX_Size() int {
	return xxx_messageInfo_CreateBusinessAlertInfoResp.Size(m)
}
func (m *CreateBusinessAlertInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBusinessAlertInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBusinessAlertInfoResp proto.InternalMessageInfo

func (m *CreateBusinessAlertInfoResp) GetErrCode() uint64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CreateBusinessAlertInfoResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type AlertAnnotation struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertAnnotation) Reset()         { *m = AlertAnnotation{} }
func (m *AlertAnnotation) String() string { return proto.CompactTextString(m) }
func (*AlertAnnotation) ProtoMessage()    {}
func (*AlertAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{4}
}

func (m *AlertAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertAnnotation.Unmarshal(m, b)
}
func (m *AlertAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertAnnotation.Marshal(b, m, deterministic)
}
func (m *AlertAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertAnnotation.Merge(m, src)
}
func (m *AlertAnnotation) XXX_Size() int {
	return xxx_messageInfo_AlertAnnotation.Size(m)
}
func (m *AlertAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_AlertAnnotation proto.InternalMessageInfo

func (m *AlertAnnotation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AlertAnnotation) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type CommonAlertLabel struct {
	AlarmType            string   `protobuf:"bytes,1,opt,name=alarmType,proto3" json:"alarmType,omitempty"`
	ClusterID            string   `protobuf:"bytes,2,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonAlertLabel) Reset()         { *m = CommonAlertLabel{} }
func (m *CommonAlertLabel) String() string { return proto.CompactTextString(m) }
func (*CommonAlertLabel) ProtoMessage()    {}
func (*CommonAlertLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{5}
}

func (m *CommonAlertLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonAlertLabel.Unmarshal(m, b)
}
func (m *CommonAlertLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonAlertLabel.Marshal(b, m, deterministic)
}
func (m *CommonAlertLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonAlertLabel.Merge(m, src)
}
func (m *CommonAlertLabel) XXX_Size() int {
	return xxx_messageInfo_CommonAlertLabel.Size(m)
}
func (m *CommonAlertLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonAlertLabel.DiscardUnknown(m)
}

var xxx_messageInfo_CommonAlertLabel proto.InternalMessageInfo

func (m *CommonAlertLabel) GetAlarmType() string {
	if m != nil {
		return m.AlarmType
	}
	return ""
}

func (m *CommonAlertLabel) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

type ModuleAlertLabel struct {
	ModuleName           string   `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	ModuleIP             string   `protobuf:"bytes,2,opt,name=moduleIP,proto3" json:"moduleIP,omitempty"`
	AlarmName            string   `protobuf:"bytes,3,opt,name=alarmName,proto3" json:"alarmName,omitempty"`
	AlarmLevel           string   `protobuf:"bytes,4,opt,name=alarmLevel,proto3" json:"alarmLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModuleAlertLabel) Reset()         { *m = ModuleAlertLabel{} }
func (m *ModuleAlertLabel) String() string { return proto.CompactTextString(m) }
func (*ModuleAlertLabel) ProtoMessage()    {}
func (*ModuleAlertLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{6}
}

func (m *ModuleAlertLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModuleAlertLabel.Unmarshal(m, b)
}
func (m *ModuleAlertLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModuleAlertLabel.Marshal(b, m, deterministic)
}
func (m *ModuleAlertLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleAlertLabel.Merge(m, src)
}
func (m *ModuleAlertLabel) XXX_Size() int {
	return xxx_messageInfo_ModuleAlertLabel.Size(m)
}
func (m *ModuleAlertLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleAlertLabel.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleAlertLabel proto.InternalMessageInfo

func (m *ModuleAlertLabel) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *ModuleAlertLabel) GetModuleIP() string {
	if m != nil {
		return m.ModuleIP
	}
	return ""
}

func (m *ModuleAlertLabel) GetAlarmName() string {
	if m != nil {
		return m.AlarmName
	}
	return ""
}

func (m *ModuleAlertLabel) GetAlarmLevel() string {
	if m != nil {
		return m.AlarmLevel
	}
	return ""
}

type ResourceAlertLabel struct {
	AlarmName            string   `protobuf:"bytes,1,opt,name=alarmName,proto3" json:"alarmName,omitempty"`
	NameSpace            string   `protobuf:"bytes,2,opt,name=nameSpace,proto3" json:"nameSpace,omitempty"`
	AlarmResourceType    string   `protobuf:"bytes,3,opt,name=alarmResourceType,proto3" json:"alarmResourceType,omitempty"`
	AlarmResourceName    string   `protobuf:"bytes,4,opt,name=alarmResourceName,proto3" json:"alarmResourceName,omitempty"`
	AlarmID              string   `protobuf:"bytes,5,opt,name=alarmID,proto3" json:"alarmID,omitempty"`
	AlarmLevel           string   `protobuf:"bytes,6,opt,name=alarmLevel,proto3" json:"alarmLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceAlertLabel) Reset()         { *m = ResourceAlertLabel{} }
func (m *ResourceAlertLabel) String() string { return proto.CompactTextString(m) }
func (*ResourceAlertLabel) ProtoMessage()    {}
func (*ResourceAlertLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaad32c28dd2f644, []int{7}
}

func (m *ResourceAlertLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAlertLabel.Unmarshal(m, b)
}
func (m *ResourceAlertLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAlertLabel.Marshal(b, m, deterministic)
}
func (m *ResourceAlertLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAlertLabel.Merge(m, src)
}
func (m *ResourceAlertLabel) XXX_Size() int {
	return xxx_messageInfo_ResourceAlertLabel.Size(m)
}
func (m *ResourceAlertLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAlertLabel.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAlertLabel proto.InternalMessageInfo

func (m *ResourceAlertLabel) GetAlarmName() string {
	if m != nil {
		return m.AlarmName
	}
	return ""
}

func (m *ResourceAlertLabel) GetNameSpace() string {
	if m != nil {
		return m.NameSpace
	}
	return ""
}

func (m *ResourceAlertLabel) GetAlarmResourceType() string {
	if m != nil {
		return m.AlarmResourceType
	}
	return ""
}

func (m *ResourceAlertLabel) GetAlarmResourceName() string {
	if m != nil {
		return m.AlarmResourceName
	}
	return ""
}

func (m *ResourceAlertLabel) GetAlarmID() string {
	if m != nil {
		return m.AlarmID
	}
	return ""
}

func (m *ResourceAlertLabel) GetAlarmLevel() string {
	if m != nil {
		return m.AlarmLevel
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRawAlertInfoReq)(nil), "alertmanager.CreateRawAlertInfoReq")
	proto.RegisterMapType((map[string]string)(nil), "alertmanager.CreateRawAlertInfoReq.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "alertmanager.CreateRawAlertInfoReq.LabelsEntry")
	proto.RegisterType((*CreateRawAlertInfoResp)(nil), "alertmanager.CreateRawAlertInfoResp")
	proto.RegisterType((*CreateBusinessAlertInfoReq)(nil), "alertmanager.CreateBusinessAlertInfoReq")
	proto.RegisterType((*CreateBusinessAlertInfoResp)(nil), "alertmanager.CreateBusinessAlertInfoResp")
	proto.RegisterType((*AlertAnnotation)(nil), "alertmanager.AlertAnnotation")
	proto.RegisterType((*CommonAlertLabel)(nil), "alertmanager.CommonAlertLabel")
	proto.RegisterType((*ModuleAlertLabel)(nil), "alertmanager.ModuleAlertLabel")
	proto.RegisterType((*ResourceAlertLabel)(nil), "alertmanager.ResourceAlertLabel")
}

func init() {
	proto.RegisterFile("proto/alertmanager/alertmanager.proto", fileDescriptor_aaad32c28dd2f644)
}

var fileDescriptor_aaad32c28dd2f644 = []byte{
	// 1795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x58, 0xdb, 0x6f, 0xdb, 0xd6,
	0x19, 0xdf, 0xf1, 0xdd, 0x27, 0x06, 0xea, 0x1d, 0x24, 0x8d, 0xa6, 0x25, 0x0b, 0xc3, 0xa4, 0x9d,
	0xc4, 0x3a, 0x56, 0xc2, 0x06, 0x6b, 0xac, 0x22, 0x43, 0x0f, 0xe3, 0x76, 0x33, 0x96, 0xb4, 0x19,
	0xdb, 0x6e, 0xdd, 0xd0, 0x22, 0xa0, 0xe5, 0x53, 0x41, 0xab, 0x44, 0xaa, 0xa4, 0x6c, 0xc3, 0xe8,
	0x02, 0x38, 0x9d, 0x9d, 0xba, 0x4a, 0xb2, 0xa6, 0x6c, 0x96, 0xb4, 0xce, 0xc5, 0x0d, 0xd6, 0xe5,
	0x02, 0xcc, 0xd1, 0x02, 0x77, 0x9e, 0x2b, 0xb5, 0xeb, 0xff, 0xa0, 0x3e, 0x0c, 0xd8, 0xb3, 0x48,
	0x4a, 0xc0, 0xb0, 0x3f, 0xc0, 0x0f, 0xc3, 0xa0, 0x73, 0x48, 0x89, 0x12, 0xe5, 0xce, 0xc3, 0x1e,
	0xf6, 0xd0, 0x17, 0x5b, 0xfc, 0x9d, 0xef, 0xfe, 0x7d, 0xe7, 0xf7, 0x51, 0x82, 0x8f, 0x65, 0x75,
	0x2d, 0xa7, 0xc5, 0x94, 0x34, 0xd1, 0x73, 0x19, 0x45, 0x55, 0x92, 0x44, 0x6f, 0x79, 0x18, 0xa5,
	0xe7, 0x68, 0xc8, 0x8f, 0x85, 0xf7, 0x24, 0x35, 0x2d, 0x99, 0x26, 0x31, 0x25, 0x9b, 0x8a, 0x29,
	0xaa, 0xaa, 0xe5, 0x94, 0x5c, 0x4a, 0x53, 0x0d, 0x26, 0x1b, 0x1e, 0xa1, 0xff, 0x12, 0x87, 0x92,
	0x44, 0x3d, 0x64, 0xcc, 0x2a, 0xc9, 0xba, 0x49, 0x2d, 0x4b, 0x25, 0x3a, 0x48, 0xef, 0x9e, 0x51,
	0xd2, 0xa9, 0x29, 0x25, 0x47, 0x62, 0xde, 0x07, 0x76, 0xc0, 0x2f, 0xf6, 0xc3, 0x5d, 0x27, 0x74,
	0xa2, 0xe4, 0x88, 0xac, 0xcc, 0xe2, 0xba, 0xfb, 0x09, 0xf5, 0x75, 0x4d, 0x26, 0x6f, 0x22, 0x19,
	0x0e, 0x1a, 0x39, 0x45, 0xcf, 0xe5, 0x52, 0x19, 0x12, 0x02, 0x1c, 0x88, 0x74, 0x4b, 0x47, 0x4d,
	0x2c, 0x08, 0x4d, 0x54, 0xdc, 0x6b, 0xbd, 0x9f, 0xaf, 0xae, 0xde, 0xb1, 0x3e, 0x9f, 0xb7, 0xee,
	0x5e, 0xb0, 0xaf, 0xad, 0xd7, 0xae, 0x7d, 0x1a, 0x61, 0xff, 0xec, 0xa5, 0xb5, 0xe8, 0xa6, 0xd4,
	0xc3, 0x77, 0x71, 0xdf, 0x92, 0x9b, 0x0a, 0xe8, 0x59, 0xd8, 0x4f, 0xd4, 0x29, 0x6a, 0xb1, 0x8b,
	0x5a, 0x7c, 0xc2, 0xc4, 0x11, 0xc1, 0xc3, 0x3c, 0x7b, 0x4e, 0xe9, 0x03, 0xfb, 0xe6, 0x4a, 0xc0,
	0x9e, 0xec, 0xc9, 0xa1, 0x97, 0xe1, 0x50, 0x92, 0xa8, 0x44, 0x57, 0x72, 0x9a, 0x3e, 0xad, 0xa7,
	0x43, 0xdd, 0x1c, 0x88, 0x0c, 0x4a, 0x47, 0x4c, 0xfc, 0xb8, 0xd0, 0x72, 0x20, 0x3e, 0x6a, 0xfd,
	0x76, 0xc5, 0x59, 0x2b, 0x39, 0xa5, 0x95, 0x97, 0xe5, 0x93, 0x23, 0xd6, 0x7b, 0x85, 0xca, 0x46,
	0xd1, 0xb9, 0x5f, 0xdc, 0x94, 0x7a, 0xf5, 0xee, 0xd0, 0xfc, 0x80, 0xdc, 0x22, 0x8d, 0x7e, 0x07,
	0xe0, 0x0e, 0x5f, 0xe9, 0x42, 0x3d, 0x5c, 0x77, 0x64, 0x87, 0x78, 0x74, 0xb4, 0xa5, 0x53, 0x1d,
	0x8b, 0x35, 0x8a, 0x9b, 0x6a, 0xcf, 0xaa, 0x39, 0x7d, 0x4e, 0x3a, 0x6e, 0x62, 0x51, 0xf0, 0x5b,
	0x13, 0x0f, 0xd8, 0x6b, 0xf7, 0xaa, 0x77, 0x7f, 0x3f, 0x52, 0x9b, 0x5f, 0xb6, 0x36, 0x36, 0xec,
	0x0f, 0x0b, 0x19, 0x62, 0x18, 0x4a, 0x92, 0x38, 0xab, 0xef, 0xd6, 0x6e, 0xde, 0xb2, 0x96, 0x3e,
	0xb1, 0xaf, 0xae, 0x6e, 0x4a, 0xbd, 0xcb, 0xa0, 0x6b, 0x00, 0xc8, 0x7e, 0x4d, 0x74, 0x11, 0xc0,
	0xbe, 0xb4, 0x32, 0x49, 0xd2, 0x46, 0xa8, 0x97, 0x86, 0x14, 0xdb, 0x4e, 0x48, 0x27, 0xa9, 0x06,
	0x8b, 0xe6, 0x84, 0x89, 0x8f, 0x09, 0xae, 0x0d, 0x71, 0xf4, 0x0d, 0x32, 0x57, 0x1f, 0x83, 0x69,
	0x52, 0xbb, 0xf2, 0xc0, 0x9a, 0xff, 0xdc, 0x2a, 0x7c, 0x36, 0xc2, 0xdc, 0xdb, 0x1f, 0x16, 0xac,
	0xcb, 0x77, 0xe8, 0xdf, 0x8b, 0x95, 0x8d, 0x79, 0xd6, 0x8d, 0x46, 0x4c, 0xae, 0x7e, 0xf8, 0x87,
	0x70, 0xb8, 0x3d, 0x5d, 0x34, 0x0c, 0xbb, 0xdf, 0x20, 0x73, 0x74, 0x4c, 0x06, 0xe5, 0xfa, 0x47,
	0xb4, 0x13, 0xf6, 0x52, 0x1f, 0xb4, 0xd1, 0x83, 0x32, 0x7b, 0x88, 0x77, 0x1d, 0x03, 0xe1, 0x31,
	0xb8, 0xc3, 0x17, 0xdb, 0x7f, 0xa3, 0x1a, 0xff, 0x1b, 0x30, 0xf1, 0x57, 0x00, 0xbe, 0x2a, 0x74,
	0x9e, 0x59, 0x71, 0xc4, 0x5a, 0xba, 0x6e, 0x95, 0x8a, 0x4e, 0x69, 0xa1, 0x52, 0x5a, 0xb7, 0x97,
	0xae, 0xda, 0x37, 0x2e, 0x59, 0xf9, 0xdb, 0xce, 0xf2, 0x82, 0x9d, 0xff, 0x43, 0x75, 0xf5, 0x4e,
	0xa5, 0x78, 0xa1, 0x52, 0x5a, 0x8f, 0xd4, 0x87, 0xe2, 0xca, 0x4a, 0xb4, 0x0c, 0xfc, 0x55, 0x2e,
	0x03, 0xaf, 0x3e, 0x06, 0x7a, 0xf3, 0x2d, 0x8e, 0xf7, 0x1d, 0xf1, 0x71, 0xee, 0x2d, 0xde, 0xed,
	0x16, 0x1f, 0xe7, 0x78, 0x32, 0x43, 0xd4, 0x1c, 0x37, 0xa9, 0x4d, 0xcd, 0xf1, 0x67, 0x47, 0x38,
	0x9e, 0xa9, 0x31, 0x21, 0x6d, 0x6a, 0x3a, 0x4d, 0xce, 0xa8, 0x4a, 0x86, 0x0a, 0x4e, 0x26, 0x8c,
	0x43, 0xb4, 0x53, 0x87, 0xdc, 0x56, 0xf1, 0x23, 0x1c, 0xaf, 0xa4, 0x15, 0x3d, 0xd3, 0x10, 0xd1,
	0xb4, 0x0c, 0xcf, 0x9d, 0x3d, 0xcb, 0xff, 0x13, 0xc0, 0x47, 0x3b, 0xe5, 0x64, 0x64, 0xd1, 0x71,
	0xd8, 0x4f, 0x74, 0xfd, 0x84, 0x36, 0xc5, 0xae, 0x61, 0x8f, 0x74, 0xc0, 0xc4, 0x9c, 0xe0, 0x61,
	0xe2, 0xae, 0x6a, 0xe1, 0x2f, 0xf6, 0xc3, 0xb7, 0xab, 0x5f, 0x5e, 0xb1, 0xae, 0xdf, 0x72, 0xf2,
	0xeb, 0xf6, 0xfc, 0x39, 0xe7, 0xe3, 0x73, 0xb2, 0x77, 0x8e, 0x9e, 0x82, 0x7d, 0x44, 0xd7, 0x4f,
	0x19, 0x49, 0x56, 0x4e, 0x69, 0x9f, 0x89, 0xf7, 0x08, 0x2e, 0x24, 0x22, 0xa6, 0x5c, 0xbb, 0xf2,
	0x51, 0xb5, 0x50, 0xa8, 0x7c, 0x79, 0xdb, 0x3e, 0x57, 0x90, 0xdd, 0xb3, 0xf8, 0xab, 0x26, 0xfe,
	0x05, 0xfc, 0xb9, 0xb0, 0x45, 0x58, 0xa2, 0xb0, 0x9d, 0x5a, 0xb3, 0xa8, 0xca, 0xc0, 0x8b, 0xa6,
	0x0c, 0x5c, 0xeb, 0xfc, 0xfd, 0x7e, 0x18, 0x66, 0x96, 0xa5, 0x69, 0x23, 0xa5, 0x12, 0xc3, 0xf8,
	0x86, 0xb3, 0xcf, 0x24, 0x1c, 0xa4, 0xa3, 0xf1, 0xd2, 0x5c, 0x96, 0x84, 0x7a, 0xa8, 0xcd, 0x71,
	0x9a, 0x71, 0x03, 0x15, 0xf7, 0xb2, 0xd2, 0x3a, 0x0f, 0x4b, 0xd6, 0xad, 0x0b, 0x11, 0x9d, 0x18,
	0xda, 0xb4, 0x9e, 0x20, 0x31, 0x36, 0x75, 0xd1, 0x4d, 0x69, 0xa7, 0x8e, 0xe4, 0x01, 0x0f, 0x96,
	0xfb, 0x18, 0x2e, 0x37, 0x0d, 0xa0, 0xe3, 0x70, 0x30, 0x91, 0x9e, 0x36, 0x72, 0x44, 0x9f, 0x18,
	0x0f, 0xf5, 0x7a, 0xe3, 0xb0, 0x4b, 0x68, 0xa2, 0xe2, 0x40, 0xed, 0xfa, 0x79, 0xe7, 0xaf, 0x9f,
	0x4c, 0x8c, 0x6f, 0x4a, 0x3d, 0x7a, 0xd7, 0x30, 0x90, 0x9b, 0x67, 0xc8, 0x80, 0x8f, 0xd0, 0x71,
	0x6e, 0x5e, 0xff, 0x50, 0x1f, 0x07, 0x22, 0x3b, 0xc4, 0xbd, 0xad, 0x84, 0x84, 0x5b, 0x85, 0x24,
	0xc1, 0xc4, 0xfb, 0x84, 0x76, 0x55, 0x71, 0x88, 0x65, 0x63, 0x9d, 0x5f, 0xb4, 0x1e, 0x7c, 0xb6,
	0x29, 0xf5, 0xe6, 0x41, 0xdd, 0x5d, 0xbb, 0x18, 0x9a, 0x86, 0xc3, 0x2c, 0x11, 0x6a, 0x95, 0x32,
	0x47, 0xa8, 0x9f, 0x7a, 0xfd, 0x5e, 0xab, 0xd7, 0x53, 0x6d, 0x52, 0x52, 0xb4, 0xde, 0x92, 0x80,
	0xb2, 0x88, 0xec, 0x7b, 0xb7, 0xad, 0x9b, 0xd7, 0x98, 0x77, 0xfb, 0xe3, 0x77, 0xaa, 0x85, 0xf3,
	0x72, 0x40, 0x0a, 0x9d, 0x85, 0xc8, 0xab, 0xa4, 0xcf, 0xf1, 0x00, 0x75, 0xcc, 0xb5, 0x3a, 0x96,
	0x03, 0x72, 0x6c, 0xb2, 0x3a, 0x18, 0x10, 0x51, 0xf5, 0xcf, 0x0b, 0x76, 0xf1, 0x72, 0x8b, 0xf3,
	0x0e, 0x72, 0xf1, 0x3c, 0x30, 0xf1, 0x79, 0x00, 0x7f, 0x03, 0x84, 0xaf, 0xb9, 0x25, 0x62, 0x6c,
	0x7b, 0x7c, 0xf7, 0x60, 0xd9, 0x5a, 0x5a, 0xb7, 0x2e, 0x5e, 0x8d, 0x96, 0x41, 0xf3, 0x3a, 0x94,
	0x41, 0x73, 0x30, 0xca, 0xe0, 0x91, 0xb6, 0x9e, 0x95, 0x41, 0xb3, 0xef, 0xfc, 0xbf, 0x00, 0xfc,
	0xee, 0x96, 0x51, 0xfc, 0x1f, 0x19, 0x2a, 0x69, 0xe2, 0x29, 0x38, 0x29, 0x7c, 0x5d, 0x6c, 0xe2,
	0xe1, 0xad, 0x4a, 0xc4, 0xaa, 0xb2, 0x3d, 0xb2, 0xfa, 0x3b, 0x80, 0xed, 0x05, 0x42, 0x4f, 0xc3,
	0x7e, 0x77, 0x27, 0xb0, 0xdd, 0x25, 0xed, 0x37, 0xf1, 0x6e, 0xc1, 0xc3, 0xc4, 0x21, 0x97, 0x9d,
	0xdc, 0xe9, 0x66, 0x77, 0xc9, 0x3b, 0x45, 0x47, 0x61, 0x7f, 0x42, 0xcb, 0x64, 0x88, 0x9a, 0x73,
	0x73, 0x0e, 0x53, 0x65, 0x17, 0xf3, 0x94, 0xed, 0xb5, 0x7b, 0xb5, 0x77, 0xf2, 0xb2, 0x07, 0xc7,
	0x5f, 0x32, 0xf1, 0x4f, 0xe1, 0x0b, 0xc2, 0x70, 0xfb, 0xa4, 0x8b, 0x47, 0x58, 0x1a, 0xcd, 0x9d,
	0xc5, 0x0a, 0x15, 0xf1, 0xdf, 0xb0, 0x98, 0x3b, 0x73, 0xd4, 0x66, 0xb4, 0x0c, 0xbc, 0x58, 0xf8,
	0xa5, 0x2e, 0x38, 0x7c, 0x42, 0xcb, 0x64, 0x34, 0xd5, 0x37, 0xfe, 0xcf, 0xf9, 0xd9, 0x88, 0xe5,
	0x17, 0x31, 0xf1, 0x63, 0x7e, 0x36, 0x0a, 0xb5, 0xb0, 0x51, 0x22, 0x3b, 0x1d, 0xcb, 0x90, 0x4c,
	0x2c, 0xa5, 0x45, 0xfd, 0x8c, 0x73, 0xc6, 0xcf, 0x38, 0x2c, 0x55, 0x6c, 0xe2, 0xa7, 0xfc, 0x8c,
	0x23, 0x78, 0x8c, 0x13, 0x71, 0x9b, 0x44, 0xcd, 0x7a, 0xa0, 0x53, 0x5a, 0xa9, 0x6c, 0xcc, 0xc7,
	0xb9, 0xb1, 0xb1, 0xb1, 0xb1, 0x68, 0x90, 0x93, 0xe2, 0x3f, 0x31, 0xf1, 0x8f, 0xe1, 0x73, 0x42,
	0x20, 0x03, 0x31, 0xec, 0x86, 0xb7, 0xbc, 0x60, 0x2d, 0xfe, 0xd1, 0x5a, 0x7c, 0x48, 0xd7, 0x34,
	0xab, 0x4b, 0xeb, 0xe8, 0xfb, 0x06, 0x7d, 0xb1, 0x1b, 0x06, 0x8a, 0x8b, 0xc6, 0x21, 0x64, 0xec,
	0xf0, 0xbc, 0x92, 0xf1, 0x6a, 0x71, 0xd0, 0xc4, 0x61, 0xc1, 0x07, 0x8b, 0x43, 0x8c, 0x54, 0xac,
	0xcb, 0x97, 0x9c, 0xbb, 0x7f, 0xf2, 0xc2, 0xf4, 0x09, 0xa0, 0x27, 0xe1, 0x00, 0x7b, 0x9a, 0x38,
	0xed, 0xd6, 0x61, 0xb7, 0x89, 0x77, 0x0a, 0x0d, 0x50, 0x1c, 0x60, 0x16, 0x26, 0x4e, 0xcb, 0x0d,
	0x0c, 0xfd, 0xc8, 0xed, 0x02, 0xf5, 0xcc, 0xf6, 0x4c, 0xb4, 0xce, 0xa5, 0x4d, 0x54, 0x44, 0x2c,
	0xcd, 0x0c, 0xc9, 0xe9, 0xa9, 0x44, 0xab, 0xfb, 0xa6, 0x14, 0x7a, 0x1e, 0x42, 0xfa, 0x70, 0x92,
	0xcc, 0x90, 0xb4, 0xbb, 0x5d, 0x46, 0x4d, 0xfc, 0x84, 0xe0, 0x83, 0x1b, 0xeb, 0x65, 0xf5, 0x5d,
	0xa7, 0x78, 0x37, 0x92, 0x52, 0x5f, 0xd7, 0x62, 0x44, 0xd7, 0x35, 0x3d, 0x36, 0xab, 0xe8, 0x6a,
	0x54, 0xf6, 0x89, 0xc6, 0x5f, 0x33, 0xf1, 0x2f, 0xe1, 0x2b, 0x1d, 0x26, 0x91, 0xf3, 0x93, 0x6b,
	0x9d, 0x88, 0x28, 0xc5, 0xb5, 0xd4, 0xde, 0x57, 0x95, 0x32, 0x68, 0xe4, 0xea, 0xf5, 0xa4, 0x0e,
	0xf3, 0x37, 0x7a, 0x20, 0x0a, 0xb2, 0x2a, 0x8a, 0xfb, 0xcb, 0xc1, 0x1a, 0xb1, 0xc7, 0xc4, 0xdf,
	0xf1, 0x97, 0xc3, 0x5b, 0x2a, 0xb4, 0x10, 0xfe, 0x0a, 0x3c, 0x03, 0x07, 0xeb, 0xef, 0x5c, 0x2f,
	0x66, 0x95, 0x84, 0xfb, 0x62, 0x29, 0xf1, 0xb4, 0x94, 0x0d, 0xd4, 0xe3, 0x66, 0xeb, 0xfd, 0x2f,
	0xea, 0xea, 0xf7, 0x8b, 0xb5, 0x6b, 0x9f, 0xca, 0xcd, 0x63, 0x74, 0x06, 0x7e, 0x9b, 0x9a, 0xf3,
	0x02, 0xa3, 0x57, 0xa3, 0xb1, 0xfc, 0xbf, 0x2f, 0x04, 0x4f, 0xbd, 0xe6, 0x30, 0xbb, 0xec, 0xa2,
	0x78, 0xcd, 0x09, 0x4a, 0xa3, 0x57, 0xda, 0x1c, 0xd0, 0x34, 0x59, 0xaf, 0x84, 0xa0, 0x03, 0x7f,
	0xf7, 0xdd, 0xc0, 0x59, 0xd2, 0x41, 0x31, 0xf4, 0x0c, 0xec, 0xa7, 0x60, 0x63, 0xeb, 0x3f, 0x6e,
	0xe2, 0x03, 0x82, 0x87, 0x89, 0x21, 0xfa, 0x46, 0xeb, 0xdf, 0x4c, 0xda, 0xac, 0x5a, 0xbf, 0x10,
	0xb2, 0x27, 0xd2, 0x36, 0x40, 0x7d, 0xff, 0xf3, 0x00, 0x3d, 0x6d, 0xe2, 0x63, 0xf0, 0x07, 0x42,
	0x87, 0x2e, 0x8b, 0x9c, 0x3f, 0x90, 0x4e, 0x23, 0x24, 0x3a, 0xdd, 0x70, 0x88, 0x2a, 0x9c, 0x62,
	0x1b, 0x18, 0x39, 0x00, 0xa2, 0xe0, 0x6b, 0x2a, 0x3a, 0xb0, 0x8d, 0xef, 0x49, 0xe1, 0x83, 0xff,
	0x59, 0xc8, 0xc8, 0xf2, 0x8b, 0xc0, 0xc4, 0xaf, 0x21, 0xdc, 0xb6, 0x4c, 0xdc, 0xc6, 0x2e, 0x2f,
	0xb0, 0xaf, 0x14, 0xee, 0x32, 0x61, 0xdc, 0x5b, 0xd9, 0x58, 0xb6, 0xf2, 0xb7, 0xad, 0xf7, 0x0a,
	0xb5, 0xc5, 0x4b, 0xce, 0x17, 0x0f, 0xac, 0x1b, 0x45, 0xfb, 0xa3, 0x42, 0x34, 0x1c, 0x62, 0x26,
	0x82, 0x0a, 0x6f, 0x7f, 0x55, 0xf9, 0xa0, 0x6b, 0x1f, 0x1f, 0x6e, 0xfd, 0x95, 0x60, 0xe6, 0x48,
	0x4c, 0x57, 0x66, 0x29, 0x64, 0xc4, 0x81, 0x80, 0xfe, 0x01, 0xe0, 0xee, 0x2d, 0xd6, 0x1d, 0x8a,
	0x74, 0xca, 0xa4, 0xd3, 0x7b, 0x43, 0x38, 0xba, 0x4d, 0x49, 0x23, 0xcb, 0xff, 0xda, 0xc4, 0x2f,
	0x20, 0x61, 0xcb, 0xbc, 0xfd, 0x4b, 0x94, 0xa6, 0x11, 0xde, 0xcf, 0x64, 0xdd, 0xf4, 0xbd, 0x57,
	0x8f, 0x40, 0xa6, 0x07, 0xf9, 0x7d, 0x81, 0x4c, 0x27, 0xdd, 0x08, 0x1a, 0xe9, 0x4a, 0x3f, 0x33,
	0xf1, 0x8b, 0x48, 0x80, 0xfb, 0x03, 0xdf, 0xa4, 0x38, 0x83, 0xe8, 0x33, 0xa9, 0x04, 0xe1, 0xf0,
	0xe9, 0x09, 0x6e, 0x5c, 0x4b, 0x88, 0xbd, 0x87, 0x47, 0x0f, 0x8f, 0x1e, 0x11, 0x00, 0x10, 0x87,
	0x95, 0x6c, 0x36, 0x9d, 0x4a, 0xd0, 0x3d, 0x18, 0xfb, 0x95, 0xa1, 0xa9, 0xf1, 0x00, 0x32, 0xd9,
	0x47, 0x7f, 0xfd, 0x78, 0xf2, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x74, 0x3a, 0xce, 0x99,
	0x11, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertManagerClient is the client API for AlertManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertManagerClient interface {
	CreateRawAlertInfo(ctx context.Context, in *CreateRawAlertInfoReq, opts ...grpc.CallOption) (*CreateRawAlertInfoResp, error)
	CreateBusinessAlertInfo(ctx context.Context, in *CreateBusinessAlertInfoReq, opts ...grpc.CallOption) (*CreateBusinessAlertInfoResp, error)
}

type alertManagerClient struct {
	cc *grpc.ClientConn
}

func NewAlertManagerClient(cc *grpc.ClientConn) AlertManagerClient {
	return &alertManagerClient{cc}
}

func (c *alertManagerClient) CreateRawAlertInfo(ctx context.Context, in *CreateRawAlertInfoReq, opts ...grpc.CallOption) (*CreateRawAlertInfoResp, error) {
	out := new(CreateRawAlertInfoResp)
	err := c.cc.Invoke(ctx, "/alertmanager.AlertManager/CreateRawAlertInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertManagerClient) CreateBusinessAlertInfo(ctx context.Context, in *CreateBusinessAlertInfoReq, opts ...grpc.CallOption) (*CreateBusinessAlertInfoResp, error) {
	out := new(CreateBusinessAlertInfoResp)
	err := c.cc.Invoke(ctx, "/alertmanager.AlertManager/CreateBusinessAlertInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertManagerServer is the server API for AlertManager service.
type AlertManagerServer interface {
	CreateRawAlertInfo(context.Context, *CreateRawAlertInfoReq) (*CreateRawAlertInfoResp, error)
	CreateBusinessAlertInfo(context.Context, *CreateBusinessAlertInfoReq) (*CreateBusinessAlertInfoResp, error)
}

// UnimplementedAlertManagerServer can be embedded to have forward compatible implementations.
type UnimplementedAlertManagerServer struct {
}

func (*UnimplementedAlertManagerServer) CreateRawAlertInfo(ctx context.Context, req *CreateRawAlertInfoReq) (*CreateRawAlertInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRawAlertInfo not implemented")
}
func (*UnimplementedAlertManagerServer) CreateBusinessAlertInfo(ctx context.Context, req *CreateBusinessAlertInfoReq) (*CreateBusinessAlertInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusinessAlertInfo not implemented")
}

func RegisterAlertManagerServer(s *grpc.Server, srv AlertManagerServer) {
	s.RegisterService(&_AlertManager_serviceDesc, srv)
}

func _AlertManager_CreateRawAlertInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRawAlertInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertManagerServer).CreateRawAlertInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanager.AlertManager/CreateRawAlertInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertManagerServer).CreateRawAlertInfo(ctx, req.(*CreateRawAlertInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertManager_CreateBusinessAlertInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessAlertInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertManagerServer).CreateBusinessAlertInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanager.AlertManager/CreateBusinessAlertInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertManagerServer).CreateBusinessAlertInfo(ctx, req.(*CreateBusinessAlertInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alertmanager.AlertManager",
	HandlerType: (*AlertManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRawAlertInfo",
			Handler:    _AlertManager_CreateRawAlertInfo_Handler,
		},
		{
			MethodName: "CreateBusinessAlertInfo",
			Handler:    _AlertManager_CreateBusinessAlertInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/alertmanager/alertmanager.proto",
}
